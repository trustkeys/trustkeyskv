package models

type TrustKeysKVModelIf interface{
	Put(appID, pubKey,  key, value string) (ok bool, oldValue, transctionID string )
	Get(appID, pubkey,  key string ) (ok bool, value, lastestTrans string)
}