[![Go Reference](https://pkg.go.dev/badge/github.com/lishank0119/pubsub.svg)](https://pkg.go.dev/github.com/lishank0119/pubsub)
[![go.mod](https://img.shields.io/github/go-mod/go-version/lishank0119/pubsub)](go.mod)

# PubSub System
[中文](README.zh-TW.md)

A lightweight internal **Publish/Subscribe (Pub/Sub)** system written in Go, designed for efficient message broadcasting with flexible subscription management.

## 🚀 Features

- **Topic-Based Messaging:** Subscribers can subscribe to specific topics.
- **Efficient Message Handling:** Uses a centralized `publish message channel` for optimized message dispatch.
- **Flexible Subscription Control:** Supports subscribing, unsubscribing from specific topics, and bulk unsubscription.
- **Thread-Safe:** Built-in synchronization for concurrent operations.

## 📦 Installation

```bash
go get -u github.com/lishank0119/pubsub
```

## ⚡ Usage

### Code

```go
package main

import (
  "fmt"
  "github.com/lishank0119/pubsub"
  "time"
)

func main() {
  ps := pubsub.NewPubSub(nil)
  subscriber := ps.NewSubscriber()

  go Publish(ps)

  subscriber.Subscribe("news", func(msg []byte) {
    fmt.Println("Received:", string(msg))
  })

  subscriber.Subscribe("news:2", func(msg []byte) {
    fmt.Println("Received(2):", string(msg))
  })

  time.Sleep(1500 * time.Millisecond)

  subscriber.Unsubscribe("news:2")

  time.Sleep(1 * time.Second)

  subscriber.UnsubscribeAll()

  select {}
}

func Publish(ps *pubsub.PubSub) {
  IntervalTime := 1 * time.Second
  ticker := time.NewTicker(IntervalTime)
  for {
    select {
    case <-ticker.C:
      if err := ps.Publish("news", []byte("Hello, PubSub World!")); err != nil {
        panic(err)
        return
      }

      if err := ps.Publish("news:2", []byte("Hello, PubSub World!(2)")); err != nil {
        panic(err)
        return
      }
    }
  }
}


```

### Subscription Management

- **Subscribe to a Topic:**

  ```go
  subscriber.Subscribe("topic", handlerFunc)
  ```

- **Unsubscribe from a Topic:**

  ```go
  subscriber.Unsubscribe("topic")
  ```

- **Unsubscribe from All Topics:**

  ```go
  subscriber.UnsubscribeAll()
  ```

- **Global Unsubscribe for a Topic:**

  ```go
  ps.UnsubscribeTopic("topic")
  ```

## ✅ Running Tests

```bash
go test -v
```

## 📄 License

This project is licensed under the MIT License.


