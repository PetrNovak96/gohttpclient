package model

type Config struct {
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}