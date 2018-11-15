// @APIVersion 1.0.0
// @Title Public Key-value store for trustkeys.network
// @Description An awesome key-value store for mobile application that verify ECDSA digital signature with secp256k1
// @Contact thanhnt@123xe.vn
// @TermsOfServiceUrl https://kvpublic.trustkeys.network/swagger/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/trustkeys/trustkeyskv/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/publickv",
			beego.NSInclude(
				&controllers.PublicKVController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

