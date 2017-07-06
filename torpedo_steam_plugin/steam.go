package torpedo_steam_plugin

import (
	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
	"github.com/tb0hdan/torpedo_plugins/torpedo_steam_plugin/games"
	"fmt"
	"strings"
)

const HelpMessage = "Usage: %ssteam [user|id|deals]. Query SteamPowered.com for user info/current game deals"

func SteamProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string
	_, command, _ := common.GetRequestedFeature(incoming_message)
	client := games.NewClient("steam", games.SteamStoreURL)
	switch strings.TrimSpace(strings.Split(command, " ")[0]) {
	case "user":
		username := strings.TrimSpace(strings.Split(command, "user")[1])
		if username == "" {
			message = "Please provide username"
		} else {
			user, _ := client.SearchSteamUser(username, "")
			if user != nil {
				message = fmt.Sprintf("Nickname: %s\n", user.PersonaName)
				message += fmt.Sprintf("Steam ID: %s\n", user.SteamID)
			} else {
				message ="User not found"
			}
		}
	case "id":
		steamid := strings.TrimSpace(strings.Split(command, "id")[1])
		if steamid == "" {
			message = "Please provide steam id"
		} else {
			user, _ := client.SearchSteamUser("", steamid)
			if user != nil {
				message = fmt.Sprintf("Nickname: %s\n", user.PersonaName)
				message += fmt.Sprintf("Steam ID: %s\n", user.SteamID)
			} else {
				message = "User not found"
			}
		}
	case "deals":
		for _, item := range client.SteamShowNew() {
			platforms := ""
			for _, platform := range item.Platforms {
				platforms += fmt.Sprintf("%s ", platform)
			}
			message += fmt.Sprintf("%s [%s] - Regular Price: %v, Current Price: %v, Difference: %v%%\n", item.GameURL, platforms, item.RegularPrice, item.CurrentPrice, item.DiscountPercentage)
		}
	default:
		message = fmt.Sprintf(HelpMessage, api.CommandPrefix)
	}
	api.Bot.PostMessage(channel, message, api)
}

func init() {
	torpedo_registry.Config.RegisterHandler("steam", SteamProcessMessage)
	torpedo_registry.Config.RegisterHelp("steam", HelpMessage)
}
