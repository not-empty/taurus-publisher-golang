# Go Bull publisher

Go bull publisher is a lib that provides a publisher for bull queues

## install

    go get github.com/kiwfy/golang-bull-publisher

## usage

Below is an example which shows some common use cases for publisher.


```go
package main

import (
	"context"
	"time"

	bullpublisher "github.com/kiwfy/golang-bull-publisher"
	"github.com/redis/go-redis/v9"
)

type testeando struct {
	Name string `json:"name"`
	Data int64  `json:"data"`
}

func main() {
	redis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	context := context.Background()
	publisher := &bullpublisher.Publisher{
		Redis:   redis,
		Context: context,
	}

	publisher.AddJob(
		"teste",
		&testando{
			Name: "levy",
			Data: time.Now().Unix(),
		},
		bullpublisher.Options{
			Attempts:           1,
			Backoff:            0,
			Delay:              0,
			Lifo:               false,
			PreventParsingData: false,
			Priority:           0,
			RemoveOnComplete:   10,
			RemoveOnFail:       1,
			Timeout:            0,
			JobId:              "123t87236482",
			Timestamp:          time.Now().Unix(),
		},
		"process",
	)
}

```
