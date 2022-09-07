package web

type Config struct {
	Address   string `yaml:"address"`
	Api       Api    `yaml:"api"`
	ShortHost string `yaml:"shortHost"`
}

type Api struct {
	Host     string `yaml:"address"`
	Insecure bool   `yaml:"insecure"`
}
