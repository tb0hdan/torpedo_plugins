package torpedo_pinterest_plugin

import (
	"fmt"
	"strings"

	"github.com/tb0hdan/torpedo_plugins/torpedo_pinterest_plugin/pinterest"
	"github.com/tb0hdan/torpedo_registry"

	common "github.com/tb0hdan/torpedo_common"

	"flag"
)

var PinterestToken *string

func  ConfigurePinterestPlugin(cfg *torpedo_registry.ConfigStruct) {
	PinterestToken = flag.String("pinterest_token", "", "Pinterest Client Token")

}

func  ParsePinterestPlugin(cfg *torpedo_registry.ConfigStruct) {
	cfg.SetConfig("pinteresttoken", *PinterestToken)
	if cfg.GetConfig()["pinteresttoken"] == "" {
		cfg.SetConfig("pinteresttoken", common.GetStripEnv("PINTEREST"))
	}
}


func PinterestProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var richmsg torpedo_registry.RichMessage

	requestedFeature, command, message := common.GetRequestedFeature(incoming_message, "board")
	command = strings.Split(command, " ")[0]

	switch command {
	case "board":
		board := strings.TrimSpace(strings.TrimPrefix(incoming_message, fmt.Sprintf("%s %s", requestedFeature, command)))
		if board != "" {
			api := pinterest.New(torpedo_registry.Config.GetConfig()["pinteresttoken"])
			images, err := api.GetImagesForBoard(board)
			if err != nil {
				return
			}
			richmsg = torpedo_registry.RichMessage{BarColor: "#36a64f",
				Text:      board,
				Title:     board,
				TitleLink: pinterest.PINTEREST_API_BASE + board,
				ImageURL:  images[0]}
		}
	default:
		if command != "" {
			message = fmt.Sprintf("Command %s not available yet", command)
		}
	}

	api.Bot.PostMessage(channel, message, api, richmsg)
}

func init() {
	torpedo_registry.Config.RegisterHelpAndHandler("pinterest", "Get pinterest picture from specific board", PinterestProcessMessage)
	torpedo_registry.Config.RegisterParser("pinterest", ConfigurePinterestPlugin, ParsePinterestPlugin)
}
