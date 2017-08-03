package torpedo_shorteners_plugin

import (
	"net/url"
	"strings"

	"fmt"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
	"golang.org/x/net/html"
)

func FB2ProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var f func(*html.Node)
	message := "Could not parse FB embed code"

	requestedFeature, command, _ := common.GetRequestedFeature(incoming_message)

	if command == "" {
		message = fmt.Sprintf("Usage: %s%s iframe_code", api.CommandPrefix, requestedFeature)
		api.Bot.PostMessage(channel, message, api)
		return
	}

	result, _ := url.QueryUnescape(incoming_message)
	doc, err := html.Parse(strings.NewReader(result))
	if err != nil {
		api.Bot.PostMessage(channel, message, api)
		return
	}
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "iframe" {
			u, err := url.Parse(n.Attr[0].Val)
			if err != nil {
				api.Bot.PostMessage(channel, message, api)
				return
			}
			q := u.Query()
			message = q["href"][0]

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	api.Bot.PostMessage(channel, message, api)
	return
}
