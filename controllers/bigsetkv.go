package controllers

import (
	"strings"
	"fmt"
	"context"
	"github.com/trustkeys/trustkeyskv/models"
	"encoding/json"
	"github.com/astaxie/beego"
	thriftpool "github.com/OpenStars/thriftpool"
	bs "github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"
	   "github.com/OpenStars/backendclients/go/bigset/transports"
	"github.com/trustkeys/trustkeyskv/common/util"
)

var (	
	BSHost string
	BSPort string
)

func ConfigSetBSHostPort(host string, port string){
	BSHost = host;
	BSPort = port;
}

func getBSClient() *thriftpool.ThriftSocketClient{
	return transports.GetStringBsListI64Client(BSHost, BSPort)
}

// Operations about simplekv
type BigsetKVController struct {
	beego.Controller
	tkPublicModel models.TrustKeysKVModelIf
}

func GetMessage(pubKey, appID, Key,Value string ) string {
	return "TrustKeys:" + pubKey + appID + Key + Value;
}

// @Title PutItem
// @Description Put key-value to cloud
// @Param	pubKey		path 	string	true		"Public key in hex"
// @Param	appID		path 	string	true		"App ID"
// @Param	key		query  	string	true		"The Key"
// @Param	val		query  	string	true		"The Value"
// @Param 	sig		query 	string 	true 	"signature of a message consisted tk_key_value "
// @Success 200 {string} models.KVObject.Key
// @Failure 403 body is empty
// @router /putitem/:appID/:pubKey [post]
func (o *BigsetKVController) PutItem(pubKey, appID string) {

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

		if o.tkPublicModel != nil{
			ok, oldVal, transID := o.tkPublicModel.Put(appID, pubKey, key, val);
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



	// client := getBSClient(); 
	
	// if client != nil {
	// 	defer client.BackToPool();
	// 	fmt.Print(  client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(),  
	// 			bs.TStringKey(appID)  , &bs.TItem{[]byte( key ), []byte(val)}) )
	// }
	// //models.KV[ob.Key] = ob.Value;
	o.ServeJSON()
}


// // @Title Post
// // @Description Put key-value to cloud
// // @Param	sessionID		query 	string	true		"Session ID"
// // @Param	appID		query 	string	true		"App ID"
// // @Param	body		body 	models.KVObject	true		"The KVObject content"
// // @Success 200 {string} models.KVObject.Key
// // @Failure 403 body is empty
// // @router / [post]
// func (o *BigsetKVController) Post() {

// 	sessionID := o.GetString("sessionID")

// 	//Todo: check sessionID 

// 	appID := o.GetString("appID")
// 	fmt.Printf("sessionID: %s, appID: %s, object: ",sessionID , appID)

// 	var ob models.KVObject
// 	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

// 	fmt.Println(ob);

// 	//simplekvid := models.AddOne(ob)
// 	o.Data["json"] = map[string]string{"Key":ob.Key}

// 	client := getBSClient();
	
// 	if client != nil {
// 		defer client.BackToPool();
// 		fmt.Print(  client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(),  
// 				bs.TStringKey(appID)  , &bs.TItem{[]byte( ob.Key ), []byte(ob.Value)  }) )
// 	}
// 	//models.KV[ob.Key] = ob.Value;
// 	o.ServeJSON()
// }


// @Title Get
// @Description find key-value by key
// @Param	pubKey		query 	string	true		"Public Key of a user"
// @Param	appID		query 	string	true		"appID"
// @Param	key		query 	string	true		"the key of kv you want to get"
// @Success 200 {simplekv} models.KVObject
// @Failure 403 : empty object
// @router /get [get]
func (o *BigsetKVController) Get() {
	pubKey := o.GetString("pubKey")
	appID := o.GetString("appID")
	key := o.GetString("key") //o.Ctx.Input.Param(":key")

	fmt.Printf("pubKey: %s, appID: %s, key: %s",sessionID , appID , key)
	
	// if key != "" {
	// 	client := getBSClient();
	// 	if client != nil {
	// 		defer client.BackToPool();
	// 		res, _ := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(context.Background(),  bs.TStringKey(appID), []byte(key) )
	// 		fmt.Print("result: ")
	// 		fmt.Print(res)
	// 		fmt.Print("\n")
	// 		if (res !=nil && res.Item != nil && res.Item.Value != nil) {
				
	// 			fmt.Println((string)(res.Item.Value[:]) )

	// 			o.Data["json"] = &models.KVObject{
	// 				Key:key,
	// 				Value:  (string)(res.Item.Value[:] )} 
	// 		} else {
	// 			o.Data["json"] = &models.KVObject{
	// 				Key:key,
	// 				Value: "notfound",
	// 			}
	// 		}
	// 	}
	// }
	o.ServeJSON()
}


// @Title MultiGet
// @Description find key-value by key
// @Param	sessionID		query 	string	true		"Session ID of user/app"
// @Param	appID		query 	string	true		"appID"
// @Param	keys		query 	[]string	true		"the keys of kv you want to get, saparated by comma, "
// @Success 200 {bigsetkv} []models.KVObject
// @Failure 403 : empty object
// @router /multiget [get]
func (o *BigsetKVController) MultiGet() {
	sessionID := o.GetString("sessionID")
	appID := o.GetString("appID")
	
	skeys := o.GetString("keys")

	keys:=strings.Split(skeys, ",")

	fmt.Printf("Multiget  sessionID: %s, appID: %s , keys: ",sessionID , appID, keys, len(keys) )
	
	
	if len(keys) != 0 {
		client := getBSClient();
		if client != nil {
			defer client.BackToPool();
		}

		kvs := []*models.KVObject{}
		for _, key := range keys {
			fmt.Println("get item ...", key)
			res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(context.Background(),  bs.TStringKey(appID), []byte(key) )
			fmt.Println("get item ", key, res, err);
			if (res !=nil && res.Item != nil && res.Item.Value != nil) {
				fmt.Println("get item ok ", res);
				kvs = append(kvs,&models.KVObject{
					Key: key,
					Value:  (string)(res.Item.Value[:] )} )				
			}
		}
		fmt.Println("kvs: ",kvs)
		o.Data["json"] = kvs
	}
	o.ServeJSON()
}


// @Title Remove
// @Description find key-value by key
// @Param	sessionID		query 	string	true		"Session ID of user/app"
// @Param	appID		query 	string	true		"appID"
// @Param	keys		query 	[]string	true		"the keys of kv you want to get, saparated by comma, "
// @Success 200 {deletedKeys} []string
// @Failure 403 : empty object
// @router /remove [get]
func (o *BigsetKVController) Remove() {
	sessionID := o.GetString("sessionID")
	appID := o.GetString("appID")
	skeys := o.GetString("keys")

	keys:=strings.Split(skeys, ",")

	fmt.Printf("Remove sessionID: %s, appID: %s ",sessionID , appID )
	fmt.Println(keys)
	
	if len(keys) != 0 {
		client := getBSClient();
		if client != nil {
			defer client.BackToPool();
		}
		for _, key := range keys {
			fmt.Println("get item ...", key)
			client.Client.(*bs.TStringBigSetKVServiceClient).BsRemoveItem(context.Background(),  bs.TStringKey(appID), []byte(key) )

		}

	}
	o.Data["json"] = keys
	o.ServeJSON()
}
