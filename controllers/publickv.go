package controllers

import (
	"fmt"
	"github.com/trustkeys/trustkeyskv/models"
	"github.com/astaxie/beego"
	"github.com/trustkeys/trustkeyskv/common/util"
)

var (
	tkPublicModel models.TrustKeysKVModelIf
)


// Operations about Public Key-value store 
type PublicKVController struct {
	beego.Controller
	
}

// Set model for controller
func SetPublicModel(aModel models.TrustKeysKVModelIf) {
	tkPublicModel = aModel
}

var GetMessage = util.GetMessage

// @Title PutItem
// @Description Put key-value to cloud
// @Param	pubKey		path 	string	true		"Public key in hex"
// @Param	appID		path 	string	true		"App ID"
// @Param	key		query  	string	true		"The Key"
// @Param	val		query  	string	true		"The Value"
// @Param 	sig		query 	string 	true 	"signature of a message consisted tk_key_value "
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /putitem/:appID/:pubKey [post]
func (o *PublicKVController) PutItem(pubKey, appID string) {

	// pubKey := o.GetString("pubKey")
	sig := o.GetString("sig")
	

	fmt.Println("PutItem ",o)

	//Todo: check sessionID 

	appID = o.GetString(":appID")
	key := o.GetString("key") 
	val := o.GetString("val") 
	fmt.Printf("pubKey: %s, appID: %s, object: %s : %s, sig: %s \n",pubKey , appID, key, val, sig)
	

	if util.CheckSignature(pubKey, GetMessage(pubKey, appID, key, val ) , sig ) {
		var aMap = map[string]string{"Key":key , "errCode": "0", "desc": "success"}

		if tkPublicModel != nil{
			ok, oldVal, transID := tkPublicModel.Put(appID, pubKey, key, val);
			if ok {
				aMap["oldValue"] = oldVal;
				aMap["transactionID"] = transID;
			} else {
				aMap["errCode"] = "-2"; // not ok
				aMap["desc"] = "backend error"
			}
		}
		o.Data["json"] = aMap;

	} else {
		o.Data["json"] = map[string]string{"Key":key , "errCode": "-1", "desc": "Invalid signature "}
	}

	o.ServeJSON()
}


// @Title Get
// @Description find key-value by key
// @Param	pubKey		query 	string	true		"Public Key of a user"
// @Param	appID		query 	string	true		"appID"
// @Param	key		query 	string	true		"the key of kv you want to get"
// @Success 200 {object} models.KVObject
// @Failure 403 : empty object
// @router /get [get]
func (o *PublicKVController) Get() {
	pubKey := o.GetString("pubKey")
	appID := o.GetString("appID")
	key := o.GetString("key") //o.Ctx.Input.Param(":key")

	fmt.Printf("pubKey: %s, appID: %s, key: %s",pubKey , appID , key)
	var aMap = map[string]string{"Key":key }

	if tkPublicModel != nil {
		ok, value, lastestTrans := tkPublicModel.Get(appID, pubKey, key)
		if ok {
			// aMap["errCode"] = "0";
			// aMap["value"] = value;
			// aMap["lastestTransactionID"] = lastestTrans
			o.Data["json"] = &models.KVObject{
				Key:key, 
				Value: value,
				TransactionID : lastestTrans,
			}
			o.ServeJSON();
			return;
		} else {
			aMap["errCode"] = "-1";
			aMap["desc"] = "Model error or empty data"
		}
		
	}
	o.Data["json"] = aMap;

	o.ServeJSON()
}

