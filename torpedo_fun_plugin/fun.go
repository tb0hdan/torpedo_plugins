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
	torpedo_registry.Config.RegisterHandler("sudo", FunProcessMessage)
	torpedo_registry.Config.RegisterHelp("sudo", "Run sudo on this machine")
	torpedo_registry.Config.RegisterHandler("rm", FunProcessMessage)
	torpedo_registry.Config.RegisterHelp("rm", "Remove files on this machine")
	torpedo_registry.Config.RegisterHandler("shutdown", FunProcessMessage)
	torpedo_registry.Config.RegisterHelp("shutdown", "Shutdown this machine for good")
	torpedo_registry.Config.RegisterHandler("halt", FunProcessMessage)
	torpedo_registry.Config.RegisterHelp("halt", "Halt this machine for good")
	torpedo_registry.Config.RegisterHandler("reboot", FunProcessMessage)
	torpedo_registry.Config.RegisterHelp("reboot", "Reboot this machine")
	torpedo_registry.Config.RegisterHandler("poweroff", FunProcessMessage)
	torpedo_registry.Config.RegisterHelp("poweroff", "Power off this machine")
	torpedo_registry.Config.RegisterHandler("kill", FunProcessMessage)
	torpedo_registry.Config.RegisterHelp("kill", "Terminate any process running on this machine")
}
