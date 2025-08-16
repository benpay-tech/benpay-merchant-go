package demo

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/benpay-tech/benpay-merchant-go/client"
	"gopkg.in/yaml.v2"
)

type Config struct {
	APIKey             string `yaml:"api_key"`
	MerchantPrivateKey string `yaml:"merchant_private_key"`
	PlatformPublicKey  string `yaml:"platform_public_key"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func genClient() (*client.Client, error) {
	config, err := loadConfig("config.yaml")
	if err != nil {
		return nil, err
	}
	cli := client.NewClient(config.APIKey, config.MerchantPrivateKey, config.PlatformPublicKey)

	return cli, nil
}

func TestCreatePayment(t *testing.T) {
	cli, err := genClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.CreatePayment(&client.CreatePaymentRequest{
		MerchantOrderNo: time.Now().Format("20060102150405"),
		Coin:            "ETH",
		Amount:          "0.02",
		MerchantNote:    "merchant note",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
}

func TestGetPaymentInfo(t *testing.T) {
	cli, err := genClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.GetPaymentInfo(&client.GetPaymentInfoRequest{
		PaymentID: "b37a33632bb747d28f1ffe562342c517",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
}

func TestGetPaymentList(t *testing.T) {
	cli, err := genClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.GetPaymentList(&client.GetPaymentListRequest{
		// PaymentID: "1f7d8a94b3714e4ba371db5316871bbb",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp:%+v, data:%+v \n", resp, resp.Data)
	for _, v := range resp.Data.Payments {
		fmt.Printf("Payments:%+v\n", v)
	}
}
