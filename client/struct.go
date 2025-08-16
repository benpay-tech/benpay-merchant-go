package client

import "encoding/json"

type CommonReply struct {
	Code int
	Msg  string
}

type CreatePaymentRequest struct {
	MerchantOrderNo string `json:"merchant_order_no"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	MerchantNote    string `json:"merchant_note"`
}

type CreatePaymentData struct {
	PaymentID string `json:"payment_id"`
}

type CreatePaymentReply struct {
	CommonReply
	Data *CreatePaymentData `json:"data"`
}

type GetPaymentInfoRequest struct {
	PaymentID string `json:"payment_id"`
}

type Payment struct {
	PaymentID       string `json:"payment_id"`
	MerchantOrderNo string `json:"merchant_order_no"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	Status          string `json:"status"`
	CreatedAt       string `json:"created_at"`
	MerchantNote    string `json:"merchant_note"`
	PayAt           string `json:"pay_at"`
	SettleAt        string `json:"settle_at"`
	SettleAmount    string `json:"settle_amount"`
	SettleTxHash    string `json:"settle_tx_hash"`
	RefundAmount    string `json:"refund_amount"`
	RefundAt        string `json:"refund_at"`
	RefundTxHash    string `json:"refund_tx_hash"`
	RefundAddress   string `json:"refund_address"`
	BillDueTo       string `json:"bill_due_to"`
	ExpireAt        string `json:"expire_at"`
}

type GetPaymentInfoReply struct {
	CommonReply
	Data *Payment `json:"data"`
}

type GetPaymentListRequest struct {
	Page            int32  `json:"page"`
	Limit           int32  `json:"limit"`
	CreatedAtBegin  int64  `json:"created_at_begin"`
	CreatedAtEnd    int64  `json:"created_at_end"`
	PaymentID       string `json:"payment_id"`
	MerchantOrderNo string `json:"merchant_order_no"`
}

type PaymentListData struct {
	Payments []*Payment `json:"list"`
	Total    string     `json:"total"`
}

type GetPaymentListReply struct {
	CommonReply
	Data *PaymentListData `json:"data"`
}

type Notification struct {
	NotifyID   string          `json:"id"`
	NotifyType string          `json:"type"`
	NotifyData json.RawMessage `json:"data"`
}

type NotifyBaseData struct {
	PaymentID       string `json:"payment_id"`
	Chain           string `json:"chain"`
	MerchantID      string `json:"mechant_id"`
	MerchantOrderNo string `json:"merchant_order_no"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	Status          string `json:"status"`
}

type PayNotifyData struct {
	NotifyBaseData
	PayAt     string `json:"pay_at"`
	BillDueTo string `json:"bill_due_to"`
}

type RefundNotifyData struct {
	NotifyBaseData
	RefundAmount  string `json:"refund_amount"`
	RefundTime    string `json:"refund_time"`
	RefundTxHash  string `json:"refund_tx_hash"`
	RefundAddress string `json:"refund_address"`
}

type SettleNotifyData struct {
	NotifyBaseData
	SettleAt     string `json:"settle_at"`
	SettleAmount string `json:"settle_amount"`
	SettleTxHash string `json:"settle_tx_hash"`
}
