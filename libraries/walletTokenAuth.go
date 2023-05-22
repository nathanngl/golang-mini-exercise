package libraries

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateWalletToken(customer_xid string) string {
	h := sha1.New()
	h.Write([]byte(customer_xid))

	return hex.EncodeToString(h.Sum(nil))
}
