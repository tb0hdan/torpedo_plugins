package torpedo_history_plugin

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_common/database"
	"github.com/tb0hdan/torpedo_registry"
)

func RunHistorySearch(channel, pattern string) (reply string) {
	var query bson.M

	result := []torpedo_registry.MessageHistoryItem{}

	db_uri := torpedo_registry.Config.GetConfig()["mongo"]
	db := database.New(db_uri, "torpedobot")
	session, collection, err := db.GetCollection("chatHistory")
	if err != nil {
		return
	}
	defer session.Close()
	query = bson.M{"message": bson.RegEx{pattern, "i"}, "channel": channel}
	err = collection.Find(query).All(&result)
	if len(result) == 0 {
		query = bson.M{"nick": bson.RegEx{pattern, "i"}, "channel": channel}
		err = collection.Find(query).All(&result)
	}

	for _, msgitem := range result {
		reply += fmt.Sprintf("%s - %s: %s\n", time.Unix(msgitem.Timestamp, 0).String(), msgitem.Nick, msgitem.Message)
	}

	if reply == "" {
		reply = "No results found\n"
	}
	return
}

func HistoryProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string

	history_help_msg := fmt.Sprintf("Usage: %shistory [search] <pattern>", api.CommandPrefix)

	_, command, _ := common.GetRequestedFeature(incoming_message)
	if command != "" {
		switch strings.Split(command, " ")[0] {
		case "search":
			pattern := strings.TrimSpace(strings.TrimPrefix(command, "search"))
			if pattern == "" {
				message = history_help_msg
			} else {
				message = RunHistorySearch(fmt.Sprintf("%v%s", channel, api.UserProfile.Server), pattern)
			}
		default:
			message = history_help_msg
		}
	} else {
		message = history_help_msg
	}
	api.Bot.PostMessage(channel, message, api)
	return
}

func init() {
	help_msg := "Show chat history"
	torpedo_registry.Config.RegisterHelpAndHandler("history", help_msg, HistoryProcessMessage)
}
