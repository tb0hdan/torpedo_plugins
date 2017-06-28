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
	torpedo_registry.Config.RegisterHandler("youtube", YoutubeProcessMessage)
	torpedo_registry.Config.RegisterHelp("youtube", "Get Youtube.com URL for specified track")
	torpedo_registry.Config.RegisterHandler("distance", DistanceProcessMessage)
	torpedo_registry.Config.RegisterHelp("distance", "Get driving distance between cities")
	torpedo_registry.Config.RegisterPreParser("googlewebapp", ConfigureGooglePlugin)
	torpedo_registry.Config.RegisterPostParser("googlewebapp", ParseGooglePlugin)
}
