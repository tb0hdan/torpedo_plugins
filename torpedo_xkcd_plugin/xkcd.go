package torpedo_xkcd_plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"strings"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
)

type XKCDResponse struct {
	Month      string `json:"month"`
	Day        string `json:"day"`
	Num        int64  `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	PostID     string `json:"postid,omitempty"`
}

func GetXKCD(postId string, logger *log.Logger) (result XKCDResponse, err error) {
	URL := fmt.Sprintf("https://xkcd.com/%s/info.0.json", postId)
	if postId == "" || postId == "0" {
		resp, err := http.Get("https://c.xkcd.com/random/comic/")
		defer resp.Body.Close()
		if err != nil {
			logger.Printf("http.Get => %v", err.Error())
		} else {
			finalURL := resp.Request.URL.String()
			URL = fmt.Sprintf("%s/info.0.json", finalURL)
			postId = strings.TrimPrefix(finalURL, "https://xkcd.com/")
		}
	} else if _, err = strconv.ParseInt(postId, 10, 64); err != nil {
		return
	}

	resp, err := http.Get(URL)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("http.Get => %v", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll => %v", err)
		return
	}
	result = XKCDResponse{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Printf("json.Unmarshal => %v", err)
		return
	}
	result.PostID = strings.TrimSuffix(postId, "/")
	return
}

func XKCDProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string
	cu := &common.Utils{}
	logger := cu.NewLog("xkcd-process-message")
	_, command, _ := common.GetRequestedFeature(incoming_message)
	result, err := GetXKCD(command, logger)
	if err != nil {
		message = fmt.Sprintf("An error occured while processing XKCD request: %+v\n", err)
		api.Bot.PostMessage(channel, message, api)
	} else {
		richmsg := torpedo_registry.RichMessage{ImageURL: result.Img, Text: result.SafeTitle}
		api.Bot.PostMessage(channel, fmt.Sprintf("XKCD %s", result.PostID), api, richmsg)
	}
}

func init() {
	torpedo_registry.Config.RegisterHandler("xkcd", XKCDProcessMessage)
	torpedo_registry.Config.RegisterHelp("xkcd", "Get XKCD random strip. Provide integer ID to get specific one.")
}
