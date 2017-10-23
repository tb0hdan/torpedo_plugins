package torpedo_history_plugin

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_common/database"
	"github.com/tb0hdan/torpedo_registry"
	"regexp"
)

type SearchOption struct {
	SearchOptions []string
	Limit int
	Page int
	Multi int
	Before int
	After int
	Query string
}


func (so *SearchOption) Parse(str string) (err error){
	so.SearchOptions = []string{"limit", "page", "multi", "before", "after"}
	err_msg := ""
	// Set default values
	so.Limit = 10
	so.Page = 1
	so.Before = 0
	so.After = 0
	so.Multi = 0
	//
	matcher := regexp.MustCompile(`.+\:[\d+]`)
	for _, str := range strings.Split(str, " ") {
		param := strings.Split(str, ":")[0]
		value, _ := strconv.Atoi(strings.TrimLeft(str, param + ":"))
		if matcher.MatchString(str) && common.IsInArray(param, so.SearchOptions){
			switch param {
			case "limit":
				if 500 >= value && value >= 0 {
					so.Limit = value
				} else {
					err_msg = "Limit has to be within 0-500 range"
				}
			case "page":
				so.Page = value
			case "multi":
				if value == 1 || value == 0 {
					so.Multi = value
				} else {
					err_msg = "Multi has to be within 0-1 range"
				}
			case "before":
				if 500 >= value && value >= 0 {
					so.Before = value
				} else {
					err_msg = "Before has to be within 0-500 range"
				}
			case "after":
				if 500 >= value && value >= 0 {
					so.After = value
				} else {
					err_msg = "After has to be within 0-500 range"
				}
			}
		} else {
			if so.Query == "" {
				so.Query = str
			} else {
				so.Query += " " + str
			}
		}
	}
	if err_msg != "" {
		err = errors.New(err_msg)
	}
	return
}


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
	so := &SearchOption{}
	err = so.Parse(pattern)
	if err != nil {
		reply = fmt.Sprintf("%+v\n", err)
		return
	}
	query = bson.M{"message": bson.RegEx{so.Query, "i"}, "channel": channel}
	err = collection.Find(query).Limit(so.Limit).All(&result)
	if len(result) == 0 {
		query = bson.M{"nick": bson.RegEx{so.Query, "i"}, "channel": channel}
		err = collection.Find(query).Limit(so.Limit).All(&result)
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
