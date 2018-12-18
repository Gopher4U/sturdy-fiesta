package config

import "os"

type Config struct {
	HttpServer 	HttpServer 	`json:"http_server"`
	Mqtt		Mqtt 		`json:"mqtt"`
}

type HttpServer struct {
	Port		string		`json:"port"`
}

type Mqtt struct {
	Host		string		`json:"host"`
	Port		string		`json:"port"`
	User		string		`json:"user"`
	Pass		string		`json:"pass"`
	ClientName	string		`json:"client_name"`
}

func CreateConfig() *Config {
	config := &Config{}

	//Http Config
	config.HttpServer.Port 		=	os.Getenv("HTTP_PORT")

	//MQTT Server Config
	config.Mqtt.Host 			=	os.Getenv("MQTT_HOST")
	config.Mqtt.Port 			=	os.Getenv("MQTT_PORT")

	//MQTT User
	config.Mqtt.User 			=	os.Getenv("MQTT_USER")
	config.Mqtt.Pass 			=	os.Getenv("MQTT_PASS")

	//Client App Name
	config.Mqtt.ClientName		=	os.Getenv("MQTT_CLIENT")

	return config
}
