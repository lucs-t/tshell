package models

type SSHConfig struct {
	Ak string `json:"ak"`
	Sk string `json:"sk"`
	Region string `json:"region"`
	Bucket string `json:"bucket"`
}