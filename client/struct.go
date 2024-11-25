package client

import "encoding/json"

type CommonReply struct {
	Code int
	Msg  string
}

type CreatePaymentRequest struct {
	OutTradeNo   string `json:"out_trade_no"`
	Coin         string `json:"coin"`
	CoinAmount   string `json:"coin_amount"`
	MerchantNote string `json:"merchant_note"`
}

type CreatePaymentData struct {
	PaymentId string `json:"payment_id"`
}

type CreatePaymentReply struct {
	CommonReply
	Data *CreatePaymentData `json:"data"`
}

type GetPaymentInfoRequest struct {
	PaymentId string `json:"payment_id"`
}

type PaymentInfoData struct {
	PaymentId                 string `json:"payment_id"`
	OutTradeNo                string `json:"out_trade_no"`
	Coin                      string `json:"coin"`
	CoinAmount                string `json:"coin_amount"`
	Status                    string `json:"status"`
	TransactionHash           string `json:"transaction_hash"`
	RefundTransactionHash     string `json:"refund_transaction_hash"`
	SettlementTransactionHash string `json:"settlement_transaction_hash"`
	CrossTransactionHash      string `json:"cross_transaction_hash"`
	AccountPeriodTime         string `json:"account_period_time"`
	CreatedAt                 string `json:"created_at"`
	ArbitraeTransactionHash   string `json:"arbitrae_transaction_hash"`
	FreezeTransactionHash     string `json:"freeze_transaction_hash"`
	TimeoutTransactionHash    string `json:"timeout_transaction_hash"`
	SettlementAmount          string `json:"settlement_amount"`
	MerchantNote              string `json:"merchant_note"`
}

type GetPaymentInfoReply struct {
	CommonReply
	Data *PaymentInfoData `json:"data"`
}

type GetPaymentListRequest struct {
	Page           int32  `json:"page"`
	Limit          int32  `json:"limit"`
	CreatedAtBegin int64  `json:"created_at_begin"`
	CreatedAtEnd   int64  `json:"created_at_end"`
	PaymentId      string `json:"payment_id"`
	OutTradeNo     string `json:"out_trade_no"`
}

type PaymentItem struct {
	PaymentId                 string `json:"payment_id"`
	OutTradeNo                string `json:"out_trade_no"`
	Coin                      string `json:"coin"`
	CoinAmount                string `json:"coin_amount"`
	Status                    string `json:"status"`
	TransactionHash           string `json:"transaction_hash"`
	RefundTransactionHash     string `json:"refund_transaction_hash"`
	SettlementTransactionHash string `json:"settlement_transaction_hash"`
	CrossTransactionHash      string `json:"cross_transaction_hash"`
	AccountPeriodTime         string `json:"account_period_time"`
	CreatedAt                 string `json:"created_at"`
	ArbitraeTransactionHash   string `json:"arbitrae_transaction_hash"`
	FreezeTransactionHash     string `json:"freeze_transaction_hash"`
	TimeoutTransactionHash    string `json:"timeout_transaction_hash"`
	SettlementAmount          string `json:"settlement_amount"`
	MerchantNote              string `json:"merchant_note"`
}

type PaymentListData struct {
	Payments []*PaymentItem `json:"payments"`
	Total    string         `json:"total"`
}

type GetPaymentListReply struct {
	CommonReply
	Data *PaymentListData `json:"data"`
}

type Notification struct {
	NotifyId   string          `json:"notify_id"`
	NotifyType string          `json:"notify_type"`
	NotifyData json.RawMessage `json:"notify_data"`
}

type NotifyBaseData struct {
	PaymentId  string `json:"payment_id"`
	Chain      string `json:"chain"`
	MerchantId string `json:"mechant_id"`
	OutTradeNo string `json:"out_trade_no"`
	Coin       string `json:"coin"`
	CoinAmount string `json:"coin_amount"`
	Status     string `json:"status"`
}

type PayNotifyData struct {
	NotifyBaseData
	TransactionHash string `json:"transaction_hash"`
}

type RefundNotifyData struct {
	NotifyBaseData
	RefundTransactionHash string `json:"refund_transaction_hash"`
}

type SettlementNotifyData struct {
	NotifyBaseData
	SettlementTransactionHash string `json:"settlement_transaction_hash"`
}

type ArbitrateNotifyData struct {
	NotifyBaseData
	ArbitrateTransactionHash string `json:"arbitrate_transaction_hash"`
}

type FreezeNotifyData struct {
	NotifyBaseData
	FreezeTransactionHash string `json:"freeze_transaction_hash"`
}
