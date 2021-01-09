package app

type HttpConfig struct {
	Port string `yaml:`
	Host string
	/// other config
}

type AppConfig struct {
	Services struct {
		HTTP *HttpConfig `json:"http,omitempty", yaml:"http,omitempty"`
	} ``
}
