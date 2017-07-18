package torpedo_currency_plugin

import "github.com/tb0hdan/torpedo_registry"

func init() {
	torpedo_registry.Config.RegisterHelpAndHandler("pb", "Get currency exchange rate from PB.UA", PBProcessMessage)
}
