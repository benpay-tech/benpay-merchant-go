package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/benpay-tech/benpay-merchant-go/utils"

	"github.com/google/uuid"
)

const BaseUrl = "https://api.benfenpay.com"

type Client struct {
	ApiKey             string //
	MerchantPrivateKey string // merchant private key
	PlatformPublicKey  string // platform public key
	BaseUrl            string
}

// NewClient Create a new API client
func NewClient(apiKey, merchantPrivateKey, platformPublicKey string) *Client {
	return &Client{
		ApiKey:             apiKey,
		MerchantPrivateKey: merchantPrivateKey,
		PlatformPublicKey:  platformPublicKey,
		BaseUrl:            BaseUrl,
	}
}

// DoRequest Execute the request and process the response
func (c *Client) DoRequest(method, path string, body interface{}, response interface{}) error {
	// Create request URL
	url := fmt.Sprintf("%s%s", c.BaseUrl, path)

	// Serialized request body
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body err: %v", err)
		}
	}

	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	nonce := uuid.New().String()
	signature, err := c.GenSign(method, path, timestamp, nonce, string(reqBody))
	if err != nil {
		return err
	}

	authorization := c.GenerateAuthorizationHeader(timestamp, nonce, signature)

	header := map[string]string{
		"Authorization": authorization,
	}

	resp, respBody, err := utils.SendPostRequest(url, reqBody, header)
	if err != nil {
		return err
	}

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			if err := json.Unmarshal(respBody, &response); err != nil {
				return fmt.Errorf("unmarshal response err: %v, body:%s", err, string(respBody))
			}
			return nil
		}
		return fmt.Errorf("request err: statusCode:%d, body:%s", resp.StatusCode, string(respBody))
	}

	// get nonce, timestamp, signature
	respNonce := resp.Header.Get("Benpay-Nonce")
	respTimestamp := resp.Header.Get("Benpay-Timestamp")
	respSignature := resp.Header.Get("Benpay-Signature")
	if respSignature == "" {
		return errors.New("response no signature")
	}

	// Decode signature
	signatureDecode, err := base64.StdEncoding.DecodeString(respSignature)
	if err != nil {
		return errors.New("response decode signature err")
	}

	// Create the message to validate
	message := fmt.Sprintf("%s\n%s\n%s\n", respTimestamp, respNonce, string(respBody))

	// Verify signature
	err = utils.VerifySignature(message, c.PlatformPublicKey, signatureDecode)
	if err != nil {
		return fmt.Errorf("response verify signature err:%v", err)
	}

	// 解析响应体
	if err := json.Unmarshal(respBody, &response); err != nil {
		return fmt.Errorf("response unmarshal body err: %v, body:%s", err, string(respBody))
	}

	return nil
}

func (c *Client) CreatePayment(createPaymentRequest *CreatePaymentRequest) (*CreatePaymentReply, error) {
	var reply CreatePaymentReply
	path := "/v1/payment/create"
	err := c.DoRequest("POST", path, createPaymentRequest, &reply)
	if err != nil {
		return nil, fmt.Errorf("DoRequest err: %v", err)
	}
	return &reply, nil
}

func (c *Client) GetPaymentInfo(getPaymentInfoRequest *GetPaymentInfoRequest) (*GetPaymentInfoReply, error) {
	var reply GetPaymentInfoReply
	path := "/v1/payment/info"
	err := c.DoRequest("POST", path, getPaymentInfoRequest, &reply)
	if err != nil {
		return nil, fmt.Errorf("DoRequest err: %v", err)
	}
	return &reply, nil
}

func (c *Client) GetPaymentList(getPaymentListRequest *GetPaymentListRequest) (*GetPaymentListReply, error) {
	var reply GetPaymentListReply
	path := "/v1/payment/list"
	err := c.DoRequest("POST", path, getPaymentListRequest, &reply)
	if err != nil {
		return nil, fmt.Errorf("DoRequest err: %v", err)
	}
	return &reply, nil
}

func (c *Client) GenerateAuthorizationHeader(timestamp, nonce, signature string) string {
	return fmt.Sprintf("BENPAY-SHA256-RSA2048 api_key=%s,timestamp=%s,nonce=%s,signature=%s",
		c.ApiKey, timestamp, nonce, signature)
}

func (c *Client) GenSign(method, path, timestamp, nonce, body string) (string, error) {
	publicParamString := fmt.Sprintf("api_key=%s,timestamp=%s,nonce=%s", c.ApiKey, timestamp, nonce)
	signString := fmt.Sprintf("%s\n%s\n%s\n%s\n", publicParamString, strings.ToUpper(method), path, body)
	signature, err := utils.SHA256WithRSA2048(signString, c.MerchantPrivateKey)
	if err != nil {
		return "", err
	}

	return signature, nil
}

func (c *Client) HandleWebhook(body []byte, nonce, timestamp, signatureBase64 string) (*Notification, error) {

	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return nil, fmt.Errorf("base64 decode err:%v", err)
	}

	// Construct message
	message := fmt.Sprintf("%s\n%s\n%s\n", timestamp, nonce, string(body))

	err = utils.VerifySignature(message, c.PlatformPublicKey, signature)
	if err != nil {
		return nil, err
	}

	var notification Notification
	err = json.Unmarshal(body, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
