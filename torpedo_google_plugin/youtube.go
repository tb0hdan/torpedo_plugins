package torpedo_google_plugin

import (
	"fmt"
	"strings"

	"github.com/tb0hdan/torpedo_plugins/torpedo_google_plugin/youtube"
	"github.com/tb0hdan/torpedo_registry"
)

func YoutubeProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	message := fmt.Sprintf("Usage: %syoutube query\n", api.CommandPrefix)
	command := strings.TrimSpace(strings.TrimLeft(incoming_message, fmt.Sprintf("%syoutube", api.CommandPrefix)))
	if command != "" {
		searchResults := youtube.YoutubeSearch(command, torpedo_registry.Config.GetConfig()["googlewebappkey"], 25)
		message = fmt.Sprintf("https://youtu.be/%s", searchResults[0].VideoID)
	}
	api.Bot.PostMessage(channel, message, api)
}
