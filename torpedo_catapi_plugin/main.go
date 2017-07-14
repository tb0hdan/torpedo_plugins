package torpedo_catapi_plugin

import (
	"flag"
	"fmt"
	"encoding/xml"
	"net/url"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
)

const CatAPIURL = "http://thecatapi.com/api/images/get"

var CatAPIKey *string

type CatAPIImage struct {
	ImageURL  string `xml:"url"`
	ID        string `xml:"id"`
	SourceURL string `xml:"source_url"`
}

type CatAPIResponse struct {
	XMLName xml.Name `xml:"response"`
	Data    struct {
		Images []*CatAPIImage `xml:"images>image"`
	} `xml:"data"`
}

func ConfigureCatAPIPlugin(cfg *torpedo_registry.ConfigStruct) {
	CatAPIKey = flag.String("catapikey", "", "TheCatAPI.com API Key")

}

func ParseCatAPIPlugin(cfg *torpedo_registry.ConfigStruct) {
	cfg.SetConfig("catapikey", *CatAPIKey)
	if cfg.GetConfig()["catapikey"] == "" {
		cfg.SetConfig("catapikey", common.GetStripEnv("CATAPIKEY"))
	}
}

func CatAPIProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string
	cu := common.Utils{}
	v := url.Values{}
	v.Add("format", "xml")
	v.Add("results_per_page", "1")
	v.Add("api_key", torpedo_registry.Config.GetConfig()["catapikey"])
	data, err := cu.GetURLBytes(fmt.Sprintf("%s?%s", CatAPIURL, v.Encode()))
	if err != nil {
		message = "Failed to process TheCatAPI.com request"
		api.Bot.PostMessage(channel, message, api)
		return
	}
	response := CatAPIResponse{}
	err = xml.Unmarshal(data, &response)
	if err != nil {
		message = "Failed to parse TheCatAPI.com response"
		api.Bot.PostMessage(channel, message, api)
		return
	}
	txt := fmt.Sprintf("TheCatAPI.com: %s", response.Data.Images[0].SourceURL)
	richmsg := torpedo_registry.RichMessage{ImageURL: response.Data.Images[0].ImageURL, Text: txt}
	api.Bot.PostMessage(channel, txt, api, richmsg)
}

func init() {
	torpedo_registry.Config.RegisterHelpAndHandler("catapi", "Get http://thecatapi.com random cat picture", CatAPIProcessMessage)
	torpedo_registry.Config.RegisterParser("catapi", ConfigureCatAPIPlugin, ParseCatAPIPlugin)
}
