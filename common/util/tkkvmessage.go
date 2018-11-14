package util


func GetMessage(pubKey, appID, Key,Value string ) string {
	return "TrustKeys:" + pubKey + appID + Key + Value;
}
