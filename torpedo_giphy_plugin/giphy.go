package torpedo_giphy_plugin

import (
	"github.com/tb0hdan/torpedo_registry"
	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_plugins/torpedo_giphy_plugin/giphy"
)

func GiphyProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string
	var richmsg torpedo_registry.RichMessage
	cu := &common.Utils{}
	logger := cu.NewLog("giphy-process-message")

	client := giphy.NewClient()

	_, command, message := common.GetRequestedFeature(incoming_message)
	if command != "" {
		logger.Printf("Got command %s\n", command)
		giphyResponse := client.GiphySearch(command)
		if giphyResponse.Meta.Status == 200 {
			richmsg = torpedo_registry.RichMessage{BarColor: "#36a64f",
				Text:      command,
				TitleLink: giphyResponse.Data[0].URL,
				ImageURL:  giphyResponse.Data[0].Images.OriginalImage.URL}
			message = ""
		} else {
			message = "Your request to Giphy could not be processed"
		}
	}
	api.Bot.PostMessage(channel, message, api, richmsg)
}

func init() {
	torpedo_registry.Config.RegisterHelpAndHandler("giphy", "Get Giphy.com image", GiphyProcessMessage)
}

