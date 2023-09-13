package models

type Host struct {
	IP       string `json:"ip"`
	User     string `json:"user" default:"root"`
	Port     int    `json:"port" default:"22"`
	Password string `json:"password"`
}
