package torpedo_fun_plugin

import (
	"fmt"
	"strings"
	"time"
	"github.com/tb0hdan/torpedo_registry"
	common "github.com/tb0hdan/torpedo_common"
)

func ParseFunCommand(feature, command string) (result string) {
	switch feature {
	case "rm":
		if command == "" {
			result = `rm: missing operand
Try 'rm --help' for more information.
`
		} else {
			result = "rm: removing files"
		}
	case "halt", "poweroff", "reboot", "shutdown":
		if command == "" {
			t := time.Now()
			hm := t.Format("15:04")
			result = fmt.Sprintf("Broadcast message from root@localhost\n\t(/dev/tty1) at %s ...\n\nThe system is going down for %s NOW!", hm, feature)
		} else {

		}
	case "kill":
		if command == "" {
			result = "kill: usage: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]"
		} else {
			result = fmt.Sprintf("-bash: kill: (%s) - No such process", command)
		}
	default:
		result = fmt.Sprintf("-bash: %s: command not found", feature)
	}
	return
}

func FunProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	cu := &common.Utils{}
	logger := cu.NewLog("fun-process-message")
	requestedFeature, command, _ := common.GetRequestedFeature(incoming_message)
	logger.Printf("Feature: %s, command: %s\n", requestedFeature, command)
	message := ParseFunCommand(strings.TrimLeft(requestedFeature, api.CommandPrefix), command)
	api.Bot.PostMessage(channel, message, api)
}


func init() {
	torpedo_registry.Config.RegisterHelpAndHandler("sudo", "Run sudo on this machine", FunProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("rm", "Remove files on this machine", FunProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("shutdown", "Shutdown this machine for good", FunProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("halt",  "Halt this machine for good", FunProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("reboot", "Reboot this machine", FunProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("poweroff",  "Power off this machine", FunProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("kill",  "Terminate any process running on this machine", FunProcessMessage)
}
