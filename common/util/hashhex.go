package util
import (
	"encoding/hex"
	"crypto/sha256"
)


func Hash256Hex(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	var out []byte;
	out = h.Sum(nil);
	return hex.EncodeToString(out);
}

func Hash256(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	var out []byte;
	out = h.Sum(nil);
	return out
	// return hex.EncodeToString(out);
}


func HashBytes256Hex(s []byte) string {
	h := sha256.New()
	h.Write(s)
	var out []byte;
	out = h.Sum(nil);
	return hex.EncodeToString(out);
}
