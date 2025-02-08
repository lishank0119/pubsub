[![Go Reference](https://pkg.go.dev/badge/github.com/lishank0119/pubsub.svg)](https://pkg.go.dev/github.com/lishank0119/pubsub)
[![go.mod](https://img.shields.io/github/go-mod/go-version/lishank0119/pubsub)](go.mod)

# PubSub System
[中文](README.zh-TW.md)

A lightweight internal **Publish/Subscribe (Pub/Sub)** system written in Go, designed for efficient message broadcasting with flexible subscription management.

## 🚀 Features

- **Topic-Based Messaging:** Subscribers can subscribe to specific topics.
- **Efficient Message Handling:** Uses a centralized `messages` for optimized message dispatch.
- **Flexible Subscription Control:** Supports subscribing, unsubscribing from specific topics, and bulk unsubscription.
- **Thread-Safe:** Built-in synchronization for concurrent operations.

## 📦 Installation

```bash
go get github.com/lishank0119/pubsub
```

## ⚡ Usage

### Import

```go
package main

import (
  "fmt"
  "github.com/lishank0119/pubsub"
  "time"
)

func main() {
  ps := pubsub.NewPubSub()
  subscriber := ps.NewSubscriber(256)

  subscriber.Subscribe("news", func(msg []byte) {
    fmt.Println("Received:", string(msg))
  })

  ps.Publish("news", []byte("Hello, PubSub World!"))

  time.Sleep(1 * time.Second)
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

- **Global Unsubscribe for All Topics:**

  ```go
  ps.UnsubscribeAll()
  ```

## ✅ Running Tests

```bash
go test -v
```

## 📄 License

This project is licensed under the MIT License.


