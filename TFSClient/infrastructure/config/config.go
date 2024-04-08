package config

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

var properties map[string]string

var (
	BaseUrl        = "base_%s_url"
	IpConfig       = "ip_%s"
	PortConfig     = "port_%s"
	PostMetadata   = "namenode_post_metadata"
	DeleteMetadata = "namenode_delete_metadata"
	GetReplicas    = "namenode_get_replicas"
	GetMetadata    = "namenode_get_metadata"
)

func LoadConfigs() {
	fileroute := "configs/application.properties"
	properties = make(map[string]string)

	file, err := os.Open(fileroute)
	if err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			properties[key] = value
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}
}

func GetStringPropetyBykey(key string) string {
	return properties[key]
}

func GetMapPropertyByKey(key string) map[string]string {
	property := make(map[string]string)

	_ = json.Unmarshal([]byte(properties[key]), &property)

	return property
}
