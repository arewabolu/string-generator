package randomid

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

//GenerateRandomBytes makes a byte slice of len(n)
//then randomly selects any number of bytes to be
//read into f
func GenerateRandomBytes(n int) ([]byte, error) {
	buff := make([]byte, n)
	_, err := rand.Read(buff)
	// Note that err == nil only if we read len(b) bytes.
	if err == io.EOF {
		println("finished reading file!s")
	}
	if err != nil {
		log.Println(err)
	}

	return buff, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string in combination with
// buff from GenerateRandomBytes
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func Generator() string {
	token, err := GenerateRandomString(6)
	if err != nil {
		// Return some error to the user,
		// but log the details internally.
	}
	return token
}
