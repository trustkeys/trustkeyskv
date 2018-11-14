package models
import(
	thriftpool "github.com/OpenStars/thriftpool"
	bs "github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"
	   "github.com/OpenStars/backendclients/go/bigset/transports"
	"context"
	"encoding/json"
	"fmt"
)


func NewTrustKeysKVAcceptAllModel(host, port string ) TrustKeysKVModelIf{
	return &SimpleTKKVModel{
		KeyAvailable: acceptAllPubKey,
		DataBSHost : host,
		DataBSPort : port,
	}
}

type CheckWhiteListPubKey func(string) bool ;

/*SimpleTKKVModel with bigset */
type SimpleTKKVModel struct  {
	//function to check if a public key is enabled to store data
	KeyAvailable CheckWhiteListPubKey

	// Big set host/port for storing data
	DataBSHost string
	DataBSPort string 
}

func (o* SimpleTKKVModel) getDataBSClient() *thriftpool.ThriftSocketClient{
	return transports.GetBsGenericClient(o.DataBSHost, o.DataBSPort);//transports.GetStringBsListI64Client(o.DataBSHost, o.DataBSPort)
}


func acceptAllPubKey(string) bool {
	return true;
}

func makeKey(pubKey, key string) []byte {
	return []byte (pubKey + ":" + key);
}

type internalValue struct {
	Value string `json:value`
	TransID string `json:trans`
}

func (o* internalValue) toString() string  {
	jsonByte, _ := json.Marshal(o)
	return string(jsonByte)
}

func (o* internalValue) toBytes() []byte  {
	jsonByte, _ := json.Marshal(o)
	return jsonByte
}

func (o* internalValue) fromBytes(aBin []byte) bool {
	if json.Unmarshal(aBin , o ) == nil {
		return true;
	}
	return false;
}


func (o* internalValue) fromString(str string) bool {
	if json.Unmarshal([]byte(str) , o ) == nil {
		return true;
	}
	return false;
}


func (o *SimpleTKKVModel) Put(appID, pubKey, key, value string) (ok bool, oldValue, transctionID string ) {
	bskey := makeKey(pubKey, key);
	testObj := &internalValue{Value:" this is my value ", TransID: "this is transiid"}
	fmt.Println("testOBJ json: ", testObj.toString() );
	// pub into bigset : appID , bsKey , value 
	aClient := o.getDataBSClient();
	if aClient != nil {
		defer aClient.BackToPool();
		aRes, Err := aClient.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(), 
				bs.TStringKey(appID), &bs.TItem{Key:bskey, Value: (&internalValue{Value:value}).toBytes() } )
		fmt.Println("put result: ", aRes, "Err: ",Err, " value: ", (&internalValue{Value:value}).toBytes() )
		if aRes !=nil && aRes.Error == bs.TErrorCode_EGood && Err == nil {
			ok = true;
			if aRes.IsSetOldItem() {
				oldValue = string(aRes.GetOldItem().Value)
				transctionID = "To be added later"
				return; 
			} 
			return ;
		}
	}
	ok = false;
	return
}

func (o *SimpleTKKVModel) Get(appID, pubKey, key string ) (ok bool, value, lastestTrans string) {
	bskey := makeKey(pubKey, key);
	// pub into bigset : appID , bsKey , value 
	aClient := o.getDataBSClient();
	if aClient != nil {
		defer aClient.BackToPool();
		aRes, Err := aClient.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(context.Background(), bs.TStringKey(appID), bs.TItemKey(bskey))
		fmt.Println(" Get result: ", aRes, " err: ", Err)
		if aRes != nil && Err == nil {
			if aRes.Item != nil {
				ok = true;
				var aValueObj = &internalValue{};
				aValueObj.fromBytes(aRes.Item.Value);
				value, lastestTrans = aValueObj.Value, aValueObj.TransID
				return 
			}
		}
	}
	ok = false;
	return

}
