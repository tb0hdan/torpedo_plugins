package torpedo_currency_plugin

import (
	"fmt"

	"strings"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
)

const (
	PBAPIBase  = "https://api.privatbank.ua/p24api"
	PBExchRate = PBAPIBase + "/pubinfo?exchange&json&coursid=11"
)

type PBResponse struct {
	Currency     string `json:"ccy"`
	BaseCurrency string `json:"base_ccy"`
	PurchaseRate string `json:"buy"`
	SaleRate     string `json:"sale"`
}

func PBProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var prefix string
	var reply string
	cu := &common.Utils{}
	result := make([]PBResponse, 0)
	err := cu.GetURLUnmarshal(PBExchRate, &result)
	if err != nil {
		api.Bot.PostMessage(channel, fmt.Sprint("An error occured: %+v\n", err), api)
		return
	}
	_, command, _ := common.GetRequestedFeature(incoming_message)

	if command == "" {
		reply = fmt.Sprintf("Usage: %spb [", api.CommandPrefix)
		for idx, item := range result {
			if idx == 0 {
				prefix = ""
			} else {
				prefix = "|"
			}
			reply += fmt.Sprintf("%s%s", prefix, item.Currency)
		}
		reply += "]"
	} else {
		for _, item := range result {
			if item.Currency == strings.ToUpper(command) {
				reply = fmt.Sprintf("Buy/Sale: %s/%s %s\n", item.PurchaseRate, item.SaleRate, item.BaseCurrency)
			}
		}
	}

	api.Bot.PostMessage(channel, reply, api)
	return
}
