package pkg

import (
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
)

func GenerateVAPIDKeys() (string, string, error) {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		return "", "", err
	}

	fmt.Println("PUBLIC KEY RAW:", publicKey, len(publicKey))
	fmt.Println("PRIVATE KEY RAW:", privateKey, len(privateKey))

	return publicKey, privateKey, nil
}
