package torpedo_google_plugin

import (
	"fmt"
	"strings"

	common "github.com/tb0hdan/torpedo_common"

	"github.com/tb0hdan/torpedo_registry"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func DistanceProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	cu := &common.Utils{}
	logger := cu.NewLog("distance-process-message")
	message := fmt.Sprintf("Usage: `%sdistance address_A -> address_B`\n", api.CommandPrefix)
	_, command, _ := common.GetRequestedFeature(incoming_message)
	separator := "->"
	if command != "" && strings.Contains(command, "-&gt;") {
		separator = "-&gt;"
	}
	if command != "" && len(strings.Split(command, separator)) == 2 {
		c, err := maps.NewClient(maps.WithAPIKey(torpedo_registry.Config.GetConfig()["googlewebappkey"]))
		if err != nil {
			// Okay, fatal here...
			logger.Fatalf("fatal error: %+v\n", err)
		}
		r := &maps.DirectionsRequest{
			Origin:      strings.TrimSpace(strings.Split(command, separator)[0]),
			Destination: strings.TrimSpace(strings.Split(command, separator)[1]),
		}
		resp, _, err := c.Directions(context.Background(), r)
		if err != nil {
			logger.Printf("fatal error: %+v\n", err)
			message = "Start / Destination could not be processed"
			api.Bot.PostMessage(channel, message, api)
			return
		}
		for _, item := range resp {
			message = fmt.Sprintf("Roads: %s\n", item.Summary)
			for _, lg := range item.Legs {
				message += fmt.Sprintf("Duration: %s\n", lg.Duration)
				message += fmt.Sprintf("Distance: %s\n", lg.Distance.HumanReadable)
			}
		}
	}
	api.Bot.PostMessage(channel, message, api)
}
