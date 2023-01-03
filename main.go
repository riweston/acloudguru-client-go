package main

import (
	"fmt"
	"os"
)

func main() {
	consumerId := os.Getenv("ACLOUDGURU_CONSUMER_ID")
	apiKey := os.Getenv("ACLOUDGURU_API_KEY")

	client, _ := NewClient(&apiKey, &consumerId)
	//subscription, _ := client.GetSubscription()
	subscription, err := client.GetUsersAll()
	if err != nil {

	}
	fmt.Printf("%+v\n", subscription)
}
