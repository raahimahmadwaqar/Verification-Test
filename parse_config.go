package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Addon struct {
	Name              string   `yaml:"name"`
	Versions          []string `yaml:"versions"`
	InstallationGuide string   `yaml:"installation_guide"`
}

type Config struct {
	Addons []Addon `yaml:"addons"`
}

func main() {
	data, err := ioutil.ReadFile("/repo/addons-config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	addon := config.Addons[0] // For simplicity, assume there's only one addon for now
	versions := addon.Versions

	versionsJSON, err := json.Marshal(versions)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("/tmp/versions.json", versionsJSON, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Versions written to /tmp/versions.json")
}
