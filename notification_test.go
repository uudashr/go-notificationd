package notificationd_test

import (
	"fmt"
	"time"

	notificationd "github.com/uudashr/go-notificationd"
)

func ExampleReadNotification() {
	type ProductCreated struct {
		Name string `json:"name"`
	}

	b := []byte(`{
		"id": 1,
		"name": "ProductCreated",
		"event": {
			"name": "Giant Spray"
		},
		"version": 1,
		"occuredTime": "2017-11-13T10:42:04.754Z"
	}`)

	reader, err := notificationd.ReadNotification(b)
	if err != nil {
		panic(err)
	}

	var event ProductCreated
	if err = reader.ReadEvent(&event); err != nil {
		panic(err)
	}

	fmt.Println("ID:", reader.ID())
	fmt.Println("Name:", reader.Name())
	fmt.Println("Event:", event)
	fmt.Println("Version:", reader.Version())
	fmt.Println("OccuredTime:", reader.OccuredTime().Format(time.RFC3339Nano))

	// Output:
	// ID: 1
	// Name: ProductCreated
	// Event: {Giant Spray}
	// Version: 1
	// OccuredTime: 2017-11-13T10:42:04.754Z
}
