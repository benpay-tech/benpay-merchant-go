# SDK Usage Guide

## Prepare Your Configuration

### Copy the example config file

``` bash
cd demo
cp config.example.yaml config.yaml
```

### Edit config.yaml

Replace the following fields with your own credentials and keys:

* api_key – Your API key.

* merchant_private_key – Your merchant RSA private key.

* platform_public_key – The platform RSA public key.

``` bash
api_key: "your_api_key_here"

merchant_private_key: |
  -----BEGIN PRIVATE KEY-----
  YOUR_PRIVATE_KEY_CONTENT
  -----END PRIVATE KEY-----

platform_public_key: |
  -----BEGIN RSA PUBLIC KEY-----
  YOUR_PUBLIC_KEY_CONTENT
  -----END RSA PUBLIC KEY-----

```

## Run Demo Tests

``` bash
go test -run CreatePayment          # create payment
go test -run TestGetPaymentInfo     # get payment info
go test -run TestGetPaymentList     # get payment list
go test -run TestWebhook            # webhook
```
