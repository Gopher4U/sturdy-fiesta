package mqtt

import (
	"fmt"
	"github.com/Gopher4U/sturdy-fiesta/RESTapi/config"
	"log"
	"net/url"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

func BuildUri(conf *config.Config) string {
	return fmt.Sprintf("mqtts://%s:%s@%s:%s/", conf.Mqtt.User, conf.Mqtt.Pass, conf.Mqtt.Host, conf.Mqtt.Port)
}

func Connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func Listen(uri *url.URL, topic string) {
	client := Connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}
