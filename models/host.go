package models

type Host struct {
	IP       string `json:"ip" bson:"ip"`
	User     string `json:"user" default:"root" bson:"user"`
	Port     int    `json:"port" bson:"port" default:"22"`
	Password string `json:"password" bson:"password"`
}
