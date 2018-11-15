package models

type TrustKeysKVModelIf interface{
	Put(appID, pubKey,  key, value string) (ok bool, oldValue, transctionID string )
	Get(appID, pubkey,  key string ) (ok bool, value, lastestTrans string)
	GetSlice(appID, pubkey, fromKey string, maxNum int32 ) (kv []KVObject, err error)
}