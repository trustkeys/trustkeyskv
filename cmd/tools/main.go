package main
import (
	"github.com/trustkeys/trustkeyskv/common/util"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	
	"flag"
	"encoding/hex"
)

var (
	privKey, _ = crypto.GenerateKey()
	// keyReceiver , _ = crypto.GenerateKey()
	appID = "tkmessenger"
	key = "theKey"
	value = "theValue"
	kegenText = ""
)

func GetMessage(pubKey, appID, Key,Value string ) string {
	return "TrustKeys:" + pubKey + appID + Key + Value;
}

func main(){
	// if len(os.Args) > 1 {
	// 	fmt.Println("Sender Key:", os.Args[1])
	// 	privKey = util.GetKeyByText(os.Args[1])
	// }

	flag.StringVar(&kegenText, "keygen", "Private Text", "Text to gen private key")


	flag.StringVar(&appID, "appid", "tkmessenger", "app ID ")

	flag.StringVar(&key, "key", "tkmessenger", "app ID ")
	flag.StringVar(&value, "value", "tkmessenger", "app ID ")
	flag.Parse()
	
	fmt.Println("AppId: ", appID, " key: ", key, " value:", value," ");

	if len(kegenText) > 0 {
		privKey = util.GetKeyByText(kegenText)
	}

	pubKey := hex.EncodeToString(crypto.CompressPubkey( &privKey.PublicKey ))

	fmt.Println("Public key: ", pubKey);
	
	aMsg := GetMessage(pubKey, appID, key, value );
	sig, err := crypto.Sign( util.Hash256(aMsg), privKey )

	sigHex  := hex.EncodeToString(sig);
	if err == nil {
		fmt.Println("Signature Hex: ", sigHex)
	}
	if util.CheckSignature(pubKey, aMsg, sigHex) {
		fmt.Println("Check signature ok ")

	}

}