package torpedo_wiki_plugin

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/tb0hdan/torpedo_plugins/torpedo_wiki_plugin/wiki"
	"github.com/tb0hdan/torpedo_registry"
)

func WikiProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var richmsg torpedo_registry.RichMessage
	client := wiki.NewClient()
	command := strings.TrimSpace(strings.TrimLeft(incoming_message, fmt.Sprintf("%swiki", api.CommandPrefix)))
	message := fmt.Sprintf("Usage: %swiki query\n", api.CommandPrefix)
	if command != "" {
		message = "The page you've requested could not be found."
		summary := client.GetWikiPageExcerpt(command)
		if summary != "" {
			message = ""
			image_url, _ := client.GetWikiTitleImage(command)
			richmsg = torpedo_registry.RichMessage{BarColor: "#36a64f",
				Text:      summary,
				Title:     command,
				TitleLink: fmt.Sprintf("https://en.wikipedia.org/wiki/%s", url.QueryEscape(command)),
				ImageURL:  image_url}
		}
	}
	api.Bot.PostMessage(channel, message, api, richmsg)
}

func init() {
	torpedo_registry.Config.RegisterHandler("wiki", WikiProcessMessage)
	torpedo_registry.Config.RegisterHelp("wiki", "Get article excerpt from Wikipedia.org")
}
