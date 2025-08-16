package demo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/benpay-tech/benpay-merchant-go/client"
)

var cli *client.Client

func handler(w http.ResponseWriter, r *http.Request) {
	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// 获取请求头
	nonce := r.Header.Get("Benpay-Nonce")
	timestamp := r.Header.Get("Benpay-Timestamp")
	signature := r.Header.Get("Benpay-Signature")

	// 打印调试信息
	fmt.Printf("Body: %s\n", string(body))
	fmt.Printf("Benpay-Nonce: %s\n", nonce)
	fmt.Printf("Benpay-Timestamp: %s\n", timestamp)
	fmt.Printf("Benpay-Signature: %s\n", signature)

	notification, err := cli.HandleWebhook(body, nonce, timestamp, signature)
	if err != nil {
		fmt.Println(err)
	}

	if notification.NotifyType == "pay" {
		var payNotifyData client.PayNotifyData
		err = json.Unmarshal(notification.NotifyData, &payNotifyData)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(payNotifyData)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))

}

func TestWebhook(t *testing.T) {
	var err error
	cli, err = genClient()
	if err != nil {
		log.Fatalf("gen client err: %v", err)
	}
	http.HandleFunc("/", handler)

	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	if err = http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

}
