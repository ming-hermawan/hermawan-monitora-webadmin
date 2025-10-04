package hashpassword

import (
    "crypto/sha256"
    "encoding/hex"
)


func Get(val string) string {
    h := sha256.New()
    h.Write([]byte(val))
    return hex.EncodeToString(h.Sum(nil))
}
