package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"],
        beego.ControllerComments{
            Method: "MultiGet",
            Router: `/multiget`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"],
        beego.ControllerComments{
            Method: "PutItem",
            Router: `/putitem/:appID/:pubKey`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("pubKey", param.IsRequired, param.InPath),
				param.New("appID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeyskv/controllers:BigsetKVController"],
        beego.ControllerComments{
            Method: "Remove",
            Router: `/remove`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
