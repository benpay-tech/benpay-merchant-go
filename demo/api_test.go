package demo

import (
	"fmt"
	"testing"
	"time"

	"github.com/benpay-tech/benpay-merchant-go/client"
)

func TestCreatePayment(t *testing.T) {
	cli := client.NewClient(ApiKey, MerchantPrivateKey, PlatformPublicKey)
	resp, err := cli.CreatePayment(&client.CreatePaymentRequest{
		OutTradeNo:   time.Now().Format("20060102150405"),
		Coin:         "BUSD",
		CoinAmount:   "0.05",
		MerchantNote: "merchant note",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
}

func TestGetPaymentInfo(t *testing.T) {
	cli := client.NewClient(ApiKey, MerchantPrivateKey, PlatformPublicKey)
	resp, err := cli.GetPaymentInfo(&client.GetPaymentInfoRequest{
		PaymentId: "eb7dbae231fd492e8928d6423777125b",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
}

func TestGetPaymentList(t *testing.T) {
	cli := client.NewClient(ApiKey, MerchantPrivateKey, PlatformPublicKey)
	resp, err := cli.GetPaymentList(&client.GetPaymentListRequest{
		PaymentId: "eb7dbae231fd492e8928d6423777125b",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
	for _, v := range resp.Data.Payments {
		fmt.Printf("Payments:%+v\n", v)
	}
}
