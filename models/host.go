package models

type Host struct {
	IP       string `json:"ip"`
	Password string `json:"password"`
	Port     int    `json:"port,default=22"`
	User     string `json:"user,default=root"`
}
