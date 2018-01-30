package torpedo_webui_plugin

import (
	"flag"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
)

var (
	Logger    *log.Logger
	WebUIPort *string
)

func WebUIPreParser(cfg *torpedo_registry.ConfigStruct) {
	WebUIPort = flag.String("webuiport", "", "Enable built-in WebUI by providing port value greater than zero")

}

func WebUIPostParser(cfg *torpedo_registry.ConfigStruct) {
	cfg.SetConfig("webuiport", *WebUIPort)
	if cfg.GetConfig()["webuiport"] == "" {
		cfg.SetConfig("webuiport", common.GetStripEnv("WEBUI_PORT"))
	}
}

type User struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name        string
	Description string
}

func WebUIBackgroundTask(cfg *torpedo_registry.ConfigStruct) {
	if *WebUIPort == "" {
		return
	}
	if val, ok := strconv.Atoi(*WebUIPort); ok != nil || val <= 0 {
		// log here
		Logger.Printf("Wrong value for port: `%s`. Web UI not started.", *WebUIPort)
		return
	}
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{})
	Admin := admin.New(&qor.Config{DB: DB})

	// Create resources from GORM-backend model
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})
	// Register route
	mux := http.NewServeMux()
	// amount to /admin, so visit `/admin` to view the admin interface
	Admin.MountTo("/admin", mux)

	Logger.Printf("Listening on: %s", *WebUIPort)
	http.ListenAndServe(fmt.Sprintf(":%s", *WebUIPort), mux)
}

func init() {
	cu := &common.Utils{}
	Logger = cu.NewLog("webui-log")
	torpedo_registry.Config.RegisterParser("webui", WebUIPreParser, WebUIPostParser)
	torpedo_registry.Config.RegisterCoroutine("webui", WebUIBackgroundTask)
}
