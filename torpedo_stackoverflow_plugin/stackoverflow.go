package torpedo_stackoverflow_plugin

import (
	"fmt"
	"log"

	"github.com/tb0hdan/torpedo_plugins/torpedo_stackoverflow_plugin/stackoverflow"

	"github.com/tb0hdan/torpedo_registry"
	common "github.com/tb0hdan/torpedo_common"
)

func StackOverflowProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string
	_, command, message := common.GetRequestedFeature(incoming_message)

	if command != "" {
		log.Printf("Got command %s\n", command)
		client := stackoverflow.NewClient("")
		result, err := client.Search(incoming_message)
		if err != nil {
			message = fmt.Sprintf("An error occured during StackOverflow search: %+v\n", err)
		} else {
			message = result
		}
		if message == "" {
			message = "No results for your query"
		}
	}
	api.Bot.PostMessage(channel, message, api)
}


func init() {
	helpmsg := "Search for solution on StackOverflow.com"
	torpedo_registry.Config.RegisterHelpAndHandler("so", helpmsg, StackOverflowProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("stackoverflow", helpmsg, StackOverflowProcessMessage)
}
