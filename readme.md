# Go Bull publisher

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square&label=PRs%20Welcome)](http://makeapullrequest.com)

Go bull publisher is a lib that provides a publisher for bull queues

### Installation

    go get github.com/kiwfy/golang-bull-publisher

### Usage

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

**Kiwfy - Open your code, open your mind!**
