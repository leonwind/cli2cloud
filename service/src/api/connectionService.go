package api

import (
	"crypto/md5"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// Create valid and unique id for a client
func CreateNewID(w http.ResponseWriter, request *http.Request) {
	// Create unique id for the client by hashing the ip address and the current
	// timestamp and encoding the hash using base62 ([0-9][A-Z][a-z])
	// Use MD5 since since it is significantly faster than SHA2 and not security relevant
	ipAddr := request.RemoteAddr
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	hash := md5.Sum([]byte(ipAddr + timestamp))

	// encode hash into base62 and use the first 5 characters as the unique id
	// 62^5 = 916132832 different unique ids
	uniqueID := encodeBase62(hash)[:5]

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(uniqueID))
}

func encodeBase62(toEncode [16]byte) string {
	encoded := big.NewInt(0)
	encoded.SetBytes(toEncode[:])
	return encoded.Text(62)
}