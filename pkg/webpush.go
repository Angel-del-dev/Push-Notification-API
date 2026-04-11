package pkg

import (
	"encoding/json"
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

func SendNotification(subscription StoredSubscription, publicKey string, privateKey string, payload map[string]string) (int, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 500, err
	}

	resp := &webpush.Subscription{
		Endpoint: subscription.Endpoint,
		Keys: webpush.Keys{
			P256dh: subscription.P256dh,
			Auth:   subscription.Auth,
		},
	}

	options := &webpush.Options{
		Subscriber:      "mailto:admin@tu-dominio.com",
		VAPIDPublicKey:  publicKey,
		VAPIDPrivateKey: privateKey,
		TTL:             60 * 60,
	}

	//ctx, cancel := time.WithTimeout(time.Background(), 10*time.Second)
	//defer cancel()

	response, err := webpush.SendNotification(jsonPayload, resp, options)

	return response.StatusCode, err
}

type StoredSubscription struct {
	Endpoint string
	P256dh   string
	Auth     string
}
