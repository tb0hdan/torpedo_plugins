package torpedo_cleverbot_plugin

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
	"github.com/ugjka/cleverbot-go"
)

var (
	CleverBotAPIKey *string
	Jobs            chan ChannelItem
	WP              *WorkerPool
)

type ChannelItem struct {
	API       *torpedo_registry.BotAPI
	ChannelID interface{}
	Message   string
}

type WorkerPool struct {
	APIKey  string
	Logger  *log.Logger
	Workers map[string]chan ChannelItem
}

func (wp *WorkerPool) dispatch(jobs <-chan ChannelItem) {
	wp.Workers = make(map[string]chan ChannelItem)
	for job := range jobs {
		wch, ok := wp.Workers[fmt.Sprintf("%+v", job.ChannelID)]
		if ok {
			// worker already started, pass job item
			wch <- job
		} else {
			// start new worker and process message
			ch := make(chan ChannelItem)
			wid := fmt.Sprintf("%+v", job.ChannelID)
			wp.Workers[wid] = ch
			go wp.worker(wid, ch)
			ch <- job
		}
	}
}

func (wp *WorkerPool) worker(wid string, jobs <-chan ChannelItem) {
	session := cleverbot.New(wp.APIKey)
	wp.Logger.Printf("Worker %s start\n", wid)
	for job := range jobs {
		answer, err := session.Ask(job.Message)
		if err == nil {
			wp.Logger.Printf("Req/Resp: `%s` -> `%s`", job.Message, answer)
			job.API.Bot.PostMessage(job.ChannelID, answer, job.API)
		} else {
			wp.Logger.Printf("Error in CleverBot worker: %+v\n", err)
		}
	}
	wp.Logger.Printf("Worker %s exit\n", wid)
}

func CleverBotBackgroundTask(cfg *torpedo_registry.ConfigStruct) {
	WP.APIKey = torpedo_registry.Config.GetConfig()["cleverbot"]
	//defer close(jobs)
	go WP.dispatch(Jobs)
}

func CleverBotProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	_, command, _ := common.GetRequestedFeature(incoming_message)
	channelItem := ChannelItem{api, channel, command}
	Jobs <- channelItem
	return
}

func CleverBotProcessTextMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	if torpedo_registry.Config.GetConfig()["cleverbot"] != "" {
		channelItem := ChannelItem{api, channel, incoming_message}
		Jobs <- channelItem
	}
	return
}

func CleverBotPreParser(cfg *torpedo_registry.ConfigStruct) {
	CleverBotAPIKey = flag.String("cleverbot", "", "CleverBot.com API Key")

}

func CleverBotPostParser(cfg *torpedo_registry.ConfigStruct) {
	cfg.SetConfig("cleverbot", *CleverBotAPIKey)
	if cfg.GetConfig()["cleverbot"] == "" {
		cfg.SetConfig("cleverbot", common.GetStripEnv("CLEVERBOT_API_KEY"))
	}
}

func init() {
	WP = &WorkerPool{}
	cu := &common.Utils{}
	WP.Logger = cu.NewLog("cleverbot-process-message")
	Jobs = make(chan ChannelItem)
	torpedo_registry.Config.RegisterParser("talk", CleverBotPreParser, CleverBotPostParser)
	torpedo_registry.Config.RegisterHelpAndHandler("talk", "Say something to bot.", CleverBotProcessMessage)
	torpedo_registry.Config.RegisterCoroutine("talk", CleverBotBackgroundTask)
	torpedo_registry.Config.RegisterTextMessageHandler("talk", CleverBotProcessTextMessage)
}
