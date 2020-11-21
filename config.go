package firebase

// Config is firebase structure config.
type Config struct {
	Credentials string `json:"credentials" env:"FIREBASE_CREDENTIALS" env-default:"firebase.json"`
}
