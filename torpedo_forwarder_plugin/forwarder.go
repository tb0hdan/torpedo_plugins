package torpedo_forwarder_plugin

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/tb0hdan/torpedo_registry"
)

func ForwarderProcessTextMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	apiKey := "xoxb-"
	destination := "D-"
	s_channel := fmt.Sprintf("%s", channel)
	if destination == s_channel {
		return
	}
	account := torpedo_registry.Accounts.GetAccountByAPIKey(apiKey)
	if account == nil {
		return
	}
	switch capi := account.API.(type) {
	case *slack.Client:
		params := slack.PostMessageParameters{}
		capi.PostMessage(destination, fmt.Sprintf("%s", channel)+" "+api.UserProfile.Nick+" "+incoming_message, params)
	}
	return
}

func init() {
	torpedo_registry.Config.RegisterTextMessageHandler("forwarder", ForwarderProcessTextMessage)
}
