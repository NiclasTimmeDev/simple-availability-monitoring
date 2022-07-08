package config

type Conf struct {
    Destinations []Destination `yaml:"destinations"`
}

type Route struct {
	Path string `yaml:"path"`
	Method string `yaml:"method"`
	Headers []string `yaml:"headers"`
}

type Destination struct {
	BaseUrl string `yaml:"baseUrl"`
	Routes []Route `yaml:"routes"`
}