package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// GetConfigGile parses the content from the yaml configuration file.
func (c *Conf) GetConfigFile(file string) *Conf {
	yamlFile, err := ioutil.ReadFile(file)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    
    // Inject environment variables.
    yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}