package torpedo_google_plugin

import (
	"github.com/tb0hdan/torpedo_registry"
	common "github.com/tb0hdan/torpedo_common"
	"flag"
)

var GoogleWebAppKey *string

func ConfigureGooglePlugin(cfg *torpedo_registry.ConfigStruct) {
	GoogleWebAppKey = flag.String("google_webapp_key", "", "Google Data API Web Application Key")

}

func ParseGooglePlugin(cfg *torpedo_registry.ConfigStruct) {
	cfg.SetConfig("googlewebappkey", *GoogleWebAppKey)
	if cfg.GetConfig()["googlewebappkey"] == "" {
		cfg.SetConfig("googlewebappkey", common.GetStripEnv("GOOGLE_WEBAPP_KEY"))
	}
}

func init() {
	torpedo_registry.Config.RegisterHelpAndHandler("youtube", "Get Youtube.com URL for specified track", YoutubeProcessMessage)
	torpedo_registry.Config.RegisterHelpAndHandler("distance",  "Get driving distance between cities", DistanceProcessMessage)
	torpedo_registry.Config.RegisterParser("googlewebapp", ConfigureGooglePlugin, ParseGooglePlugin)
}
