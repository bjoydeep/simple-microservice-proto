package transport

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(client mqtt.Client, topic string) {

	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	println("Subscribed to topic successfully: ", topic)
}

func Publish(client mqtt.Client, jsonData []byte, topic string) {

	//println("Publishing messages..-----", string(jsonData))
	//helpful: https://github.com/eclipse/paho.mqtt.golang/blob/master/client.go#L767-L776
	token := client.Publish(topic, 1, true, jsonData)
	//call blocks till the message is sent to the broker
	token.Wait()
	//println("Published messages")

}
