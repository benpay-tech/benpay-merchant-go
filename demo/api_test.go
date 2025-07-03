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
		MerchantOrderNo: time.Now().Format("20060102150405"),
		Coin:            "BUSD",
		Amount:          "0.05",
		MerchantNote:    "merchant note",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
}

func TestGetPaymentInfo(t *testing.T) {
	cli := client.NewClient(ApiKey, MerchantPrivateKey, PlatformPublicKey)
	resp, err := cli.GetPaymentInfo(&client.GetPaymentInfoRequest{
		PaymentID: "1f7d8a94b3714e4ba371db5316871bbb",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
}

func TestGetPaymentList(t *testing.T) {
	cli := client.NewClient(ApiKey, MerchantPrivateKey, PlatformPublicKey)
	resp, err := cli.GetPaymentList(&client.GetPaymentListRequest{
		PaymentID: "1f7d8a94b3714e4ba371db5316871bbb",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
	for _, v := range resp.Data.Payments {
		fmt.Printf("Payments:%+v\n", v)
	}
}
