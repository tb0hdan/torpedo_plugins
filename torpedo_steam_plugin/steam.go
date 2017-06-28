package main

import (
	"fmt"

	"github.com/tb0hdan/torpedo_registry"
	"github.com/tb0hdan/torpedo_plugins/torpedo_steam_plugin/games"
)

func SteamProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string
	client := games.NewClient("steam", games.SteamStoreURL)
	for _, item := range client.SteamShowNew() {
		platforms := ""
		for _, platform := range item.Platforms {
			platforms += fmt.Sprintf("%s ", platform)
		}
		message += fmt.Sprintf("%s [%s] - Regular Price: %v, Current Price: %v, Difference: %v%%\n", item.GameURL, platforms, item.RegularPrice, item.CurrentPrice, item.DiscountPercentage)
	}
	api.Bot.PostMessage(channel, message, api)
}

func init() {
	torpedo_registry.Config.RegisterHandler("steam", SteamProcessMessage)
	torpedo_registry.Config.RegisterHelp("steam", "Get http://store.steampowered.com/ deals")
}
