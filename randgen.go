package randomid

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"flag"
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

func Generator(n int) string {
	token, err := GenerateRandomString(n)
	if err != nil {
		err = errors.New("Could not generate the random string")
		log.Printf("%v", err)
	}
	return token
}

//You can use a flag to indicate the length of the string 
/*
	var length int
	flag.IntVar(&length,"L",6,"lenght of generated string")

	flag.Parse()
	genToken:=Generator(length)
	log.Println(genToken)
*/
