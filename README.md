# Golang SDK for Benpay merchant

# API Documentation
- [Official documentation](https://docs.benpay.org/zh-Hans/BenPay/Doc/api_doc/)

# Installation

```bash
go get github.com/benpay-tech/benpay-merchant-go
```

# Usage
> Consider `/v1/payment/create` as a reference, the full code can be located in the demo directory.

* Get api key and secret

https://www.benpay.org/paymvp/business

* create payment

```golang
package main

import (
    "fmt"
    "github.com/benpay-tech/benpay-merchant-go/client"
    "time"
)

func main() {
    cli := client.NewClient(ApiKey, MerchantPrivateKey, PlatformPublicKey)
    resp, err := cli.CreatePayment(&client.CreatePaymentRequest{
        OutTradeNo:   time.Now().Format("20060102150405"),
        Coin:         "BUSD",
        CoinAmount:   "0.05",
        MerchantNote: "merchant note",
    })
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
    }
}
```
