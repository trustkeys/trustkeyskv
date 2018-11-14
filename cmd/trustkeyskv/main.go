package main

import (
	_ "github.com/trustkeys/trustkeyskv/routers"
	"github.com/trustkeys/trustkeyskv/appconfig"
	"github.com/trustkeys/trustkeyskv/controllers"
	"os"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	// if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	appconfig.InitConfig()
	controllers.ConfigSetBSHostPort(appconfig.BIGSETKV_HOST, strconv.Itoa(appconfig.BIGSETKV_PORT) )
	
	os.Setenv("HOST", appconfig.RunningHost)
	os.Setenv("PORT", appconfig.ListenPort)
	
	beego.Run()
}
