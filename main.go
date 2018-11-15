package main

import (
	_ "github.com/trustkeys/trustkeyskv/routers"
	"github.com/trustkeys/trustkeyskv/appconfig"
	// "github.com/trustkeys/trustkeyskv/controllers"
	"os"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/trustkeys/trustkeyskv/controllers"
	"github.com/trustkeys/trustkeyskv/models"
	
)


func InitWithBSHostPort(bsHost, bsPort string) {
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/publickv",
	// 		beego.NSInclude(
	// 			controllers.NewPublicKVController( models.NewTrustKeysKVAcceptAllModel(bsHost, bsPort) ),
	// 		),
	// 	),

	// )
	// beego.AddNamespace(ns)
	controllers.SetPublicModel(models.NewTrustKeysKVAcceptAllModel(bsHost, bsPort) )
}


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	appconfig.InitConfig()
	InitWithBSHostPort(appconfig.BIGSETKV_HOST, strconv.Itoa(appconfig.BIGSETKV_PORT) )
	// controllers.ConfigSetBSHostPort(appconfig.BIGSETKV_HOST, strconv.Itoa(appconfig.BIGSETKV_PORT) )
	
	os.Setenv("HOST", appconfig.RunningHost)
	os.Setenv("PORT", appconfig.ListenPort)
	
	beego.Run()
}
