package consumers

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

func NewsConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "****:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.Subscribe("DSEX", nil)

	if err != nil {
		panic(err)
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			var news models.News
			if err := json.Unmarshal(msg.Value, &news); err != nil {
				fmt.Printf("Failed to parse message: %v\n", err)
				continue
			}

			if err := config.DB.Create(&news).Error; err != nil {
				fmt.Printf("Failed to save message to database: %v\n", err)
			} else {
				fmt.Println("Message saved to database successfully")
			}
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	err = c.Close()
	if err != nil {
		return
	}
}
