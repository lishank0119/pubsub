# PubSub 系統
[English](README.md)

一個用 Go 語言編寫的輕量級 **發佈/訂閱 (Pub/Sub)** 系統，專為高效訊息廣播與靈活的訂閱管理設計。

## 🚀 功能特色

- **基於主題的訊息傳遞：** 訂閱者可訂閱特定主題接收訊息。
- **高效訊息處理：** 使用集中式 `messages` 以優化訊息分派。
- **彈性訂閱控制：** 支援訂閱、取消特定主題與批次取消訂閱。
- **執行緒安全：** 內建同步機制，適用於多執行緒環境。

## 📦 安裝方式

```bash
go get -u github.com/lishank0119/pubsub
```

## ⚡ 使用方式

### 程式碼

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

### 訂閱管理

- **訂閱主題：**

  ```go
  subscriber.Subscribe("topic", handlerFunc)
  ```

- **取消訂閱特定主題：**

  ```go
  subscriber.Unsubscribe("topic")
  ```

- **取消所有訂閱：**

  ```go
  subscriber.UnsubscribeAll()
  ```

- **全域取消特定主題的所有訂閱者：**

  ```go
  ps.UnsubscribeTopic("topic")
  ```

- **全域取消所有主題的所有訂閱者：**

  ```go
  ps.UnsubscribeAll()
  ```

## ✅ 執行測試

```bash
go test -v
```

## 📄 授權條款

本專案採用 MIT 授權條款。


