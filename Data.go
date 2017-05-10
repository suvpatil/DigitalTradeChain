package main

type contract struct {
	contractId                   string `json:"contractId"`
	orderDetails                 order  `json:"orderDetails"`
	contractCreateDate           string `json:"createDate"`
	sellerInformation            user   `json:"sellerInformation"`
	buyerInformation             user   `json:"buyerInformation"`
	buyerBankInformation         user   `json:"buyerBankInformation"`
	sellerBankInformation        user   `json:"sellerBankInformation"`
	transporterInformation       user   `json:"transporterInformation"`
	paymentConditionInDays       int    `json:"paymentConditionInDays"`
	isLCAttached                 bool   `json:"isLCAttached"`
	isPOAttached                 bool   `json:"isPOAttached"`
	isInvoiceListAttached        bool   `json:"isInvoiceListAttached"`
	isBillOfLedingAttached       bool   `json:"isBillOfLedingAttached"`
	actionPendingOn              string `json:"actionPendingOn"`
	contractStatus               string `json:"contractStatus"`
	createDate                   string `json:"createDate"`
	buyerApprove                 string `json:"buyerApprove"`
	locPublishedByBuyerBank      string `json:"locPublishedByBuyerBank"`
	locPublishedBySellerBank     string `json:"locPublishedBySellerBank"`
	readyForShipmentDateBySeller string `json:"readyForShipmentDateBySeller"`
	deliveryOngoingByTransporter string `json:"deliveryOngoingByTransporter"`
	shipmentDoneByTransporter    string `json:"shipmentDoneByTransporter"`
	deliveryConfirmByBuyer       string `json:"deliveryConfirmByBuyer"`
	paymentInitiatedByBuyerBank  string `json:"paymentInitiatedByBuyerBank"`
	paymentSetteledBySellerBank  string `json:"paymentSetteledBySellerBank"`
	paymentConfirmedBySeller     string `json:"paymentConfirmedBySeller"`
}

type order struct {
	totalAmount     int       `json:"totalAmount"`
	currency        string    `json:"currency"`
	deliveryDate    string    `json:"deliveryDate"`
	deliveryAddress string    `json:"deliveryAddress"`
	articleDetails  []article `json:"articleDetails"`
}

type user struct {
	userId         string `json:"userId"`
	contactDetails string `json:"contactDetails"`
	address        string `json:"address"`
}

type article struct {
	articleName  string `json:"articleName"`
	articleDesc  string `json:"articleDesc"`
	articleQty   string `json:"articleQty"`
	articlePrice string `json:"articlePrice"`
}
