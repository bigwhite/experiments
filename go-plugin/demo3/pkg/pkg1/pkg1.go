package pkg1

import (
	"log"
	"plugin"
)

func init() {
	log.Println("pkg1 init")
}

func LoadPlugin(pluginPath string) error {
	_, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}
	return nil
}
