
package models
import "fmt"

type  KVObject struct{
	Key string 
	Value string 
	TransactionID string
}

// type  KVObject struct{
// 	Key string `json:key`
// 	Value string `json:value`
// }

var (
	// global key - value cache
)

func init() {
	fmt.Println("Hello world init KVOBject")
}