package models



func NewTrustKeysKVModel() TrustKeysKVModelIf{
	return nil;
}

type SimpleTKKVModel struct  {

}

func (o *SimpleTKKVModel) Put(appID, pubKey, key, value string) (ok bool, oldValue, transctionID string ) {
	ok = false;
	return
}

func (o *SimpleTKKVModel) Get(appID, pubkey, key string ) (ok bool, value, lastestTrans string) {
	ok = false;
	return;
}
