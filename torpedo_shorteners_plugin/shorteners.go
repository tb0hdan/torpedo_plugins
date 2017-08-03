package torpedo_shorteners_plugin

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
)

func QREncoderProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	cu := &common.Utils{}
	cu.SetLoggerPrefix("shorteners-plugin")
	command := strings.TrimSpace(strings.TrimLeft(incoming_message, fmt.Sprintf("%sqr", api.CommandPrefix)))

	if command == "" {
		api.Bot.PostMessage(channel, fmt.Sprintf("Usage: %sqr query\n", api.CommandPrefix), api)
	} else {
		command := strings.TrimSpace(strings.TrimLeft(incoming_message, fmt.Sprintf("%sqr", api.CommandPrefix)))
		richmsg := torpedo_registry.RichMessage{ImageURL: fmt.Sprintf("http://chart.apis.google.com/chart?cht=qr&chs=350x350&chld=M|2&chl=%s", command), Text: command}
		api.Bot.PostMessage(channel, "", api, richmsg)
	}
}

func TinyURLProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	cu := &common.Utils{}
	cu.SetLoggerPrefix("shorteners-plugin")
	command := strings.TrimSpace(strings.TrimLeft(incoming_message, fmt.Sprintf("%stinyurl", api.CommandPrefix)))

	if command == "" {
		api.Bot.PostMessage(channel, fmt.Sprintf("Usage: %stinyurl url\n", api.CommandPrefix), api)
	} else {
		command := strings.TrimSpace(strings.TrimLeft(incoming_message, fmt.Sprintf("%stinyurl", api.CommandPrefix)))
		query := url.QueryEscape(command)
		result, err := cu.GetURLBytes(fmt.Sprintf("https://tinyurl.com/api-create.php?url=%s", query))
		message := "An error occured during TinyURL encoding process"
		if err == nil {
			message = string(result)
		}
		api.Bot.PostMessage(channel, message, api)
	}
}

func CryptoProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	requestedFeature, command, message := common.GetRequestedFeature(incoming_message)
	if command != "" {
		switch requestedFeature {
		case fmt.Sprintf("%sb64e", api.CommandPrefix):
			message = base64.StdEncoding.EncodeToString([]byte(command))
		case fmt.Sprintf("%sb64d", api.CommandPrefix):
			decoded, err := base64.StdEncoding.DecodeString(command)
			if err != nil {
				message = fmt.Sprintf("%v", err)
			} else {
				message = string(decoded)
			}
		case fmt.Sprintf("%smd5", api.CommandPrefix):
			message = common.MD5Hash(command)
		case fmt.Sprintf("%ssha1", api.CommandPrefix):
			message = common.SHA1Hash(command)
		case fmt.Sprintf("%ssha256", api.CommandPrefix):
			message = common.SHA256Hash(command)
		case fmt.Sprintf("%ssha512", api.CommandPrefix):
			message = common.SHA512Hash(command)
		default:
			// should never get here
			message = "Unknown feature requested"
		}
	}
	api.Bot.PostMessage(channel, message, api)
}

func init() {
	fb_msg := "Convert FB embed to post URL"
	torpedo_registry.Config.RegisterHelpAndHandler("f2p", fb_msg, FB2ProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("fb2post", fb_msg, FB2ProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("qr", "Create QR Code from URL", QREncoderProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("tinyurl", "Shorten URL using TinyURL.com", TinyURLProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("b64e", "Base64 encode", CryptoProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("b64d", "Base64 decode", CryptoProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("md5", "Calculate message MD5 sum", CryptoProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("sha1", "Calculate message SHA1 sum", CryptoProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("sha256", "Calculate message SHA256 sum", CryptoProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("sha512", "Calculate message SHA512 sum", CryptoProcessMessage)
}
