package config

type Config struct {
	Application struct {
		MaxRequestsPerMinute int
		Port                 string
		VAPIDPubliKey        string
		VAPIDPrivateKey      string
		SecretJWT            string
		EncryptionKey        string
		HmacKey              string
	}

	Database struct {
		Host     string
		Port     int
		Name     string
		Username string
		Password string
		SSLMode  string
	}
}
