package models



func NewTrustKeysKVModel() TrustKeysKVModelIf{
	return nil;
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

func acceptAllPubKey(string) bool {
	return true;
}

func makeKey(pubKey, key string) string {
	return pubKey + ":" + key;
}



func (o *SimpleTKKVModel) Put(appID, pubKey, key, value string) (ok bool, oldValue, transctionID string ) {
	bskey := makeKey(pubKey, key);
	// pub into bigset : appID , bsKey , value 
	ok = false;
	return
}

func (o *SimpleTKKVModel) Get(appID, pubkey, key string ) (ok bool, value, lastestTrans string) {
	ok = false;
	return;
}
