package notificationd_test

import (
	"fmt"

	notificationd "github.com/uudashr/go-notificationd"
)

func ExampleReadNotification() {
	type ProductCreated struct {
		Name string `json:"name"`
	}

	var b []byte
	// ...

	reader, err := notificationd.ReadNotification(b)
	if err != nil {
		return
	}

	var event ProductCreated
	if err = reader.ReadEvent(&event); err != nil {
		return
	}

	fmt.Println("ID:", reader.ID())
	fmt.Println("Name:", reader.Name())
	fmt.Println("OccuredTime:", reader.OccuredTime())
	fmt.Println("Event:", event)
}
