package ipconfig

type SessionConfig struct {
	SessionSecret string
	Path          string
	MaxAge        int
	HttpOnly      bool
}

type Config struct {
	Session SessionConfig
}
