package main

type contract struct {
	contractId                   string          `json:"contractId"`
	sellerDetails                sellerDetails   `json:"sellerDetails"`
	buyerDetails                 buyerDetails    `json:"buyerDetails"`
	tradeDetails                 []product       `json:"tradeDetails"`
	tradeConditions              tradeConditions `json:"tradeConditions"`
	deliveryDetails              deliveryDetails `json:"deliveryDetails"`
	contractCreateDate           string          `json:"createDate"`
	isLCAttached                 bool            `json:"isLCAttached"`
	isPOAttached                 bool            `json:"isPOAttached"`
	isInvoiceListAttached        bool            `json:"isInvoiceListAttached"`
	isBillOfLedingAttached       bool            `json:"isBillOfLedingAttached"`
	actionPendingOn              string          `json:"actionPendingOn"`
	contractStatus               string          `json:"contractStatus"`
	lastUpdatedDate              string          `json:"createDate"`
	buyerApprove                 string          `json:"buyerApprove"`
	locPublishedByBuyerBank      string          `json:"locPublishedByBuyerBank"`
	locPublishedBySellerBank     string          `json:"locPublishedBySellerBank"`
	readyForShipmentDateBySeller string          `json:"readyForShipmentDateBySeller"`
	deliveryOngoingByTransporter string          `json:"deliveryOngoingByTransporter"`
	shipmentDoneByTransporter    string          `json:"shipmentDoneByTransporter"`
	deliveryConfirmByBuyer       string          `json:"deliveryConfirmByBuyer"`
	paymentInitiatedByBuyerBank  string          `json:"paymentInitiatedByBuyerBank"`
	paymentSetteledBySellerBank  string          `json:"paymentSetteledBySellerBank"`
	paymentConfirmedBySeller     string          `json:"paymentConfirmedBySeller"`
}

type tradeConditions struct {
	paymentDuration   string `"json:"paymentDuration"`
	transportDuration string `"json:"transportDuration"`
	currency          string `"json:"currency"`
	paymentTerms      string `"json:"paymentTerms"`
}

type product struct {
	productName     string `json:"productName"`
	productDesc     string `json:"productDesc"`
	productPrice    string `json:"productPrice"`
	productQuantity string `json:"productQuantity"`
	totalAmount     string `json:"totalAmount"`
}
type sellerDetails struct {
	seller     user `json:"seller"`
	sellerBank user `json:"sellerBank"`
}
type buyerDetails struct {
	buyer     user `json:"buyer"`
	buyerBank user `json:"buyerBank"`
}

type deliveryDetails struct {
	address            string `json:"address"`
	deliveryDate       string `json:"deliveryDate"`
	transporterDetails user   `json:"transporterDetails"`
}

type user struct {
	userId    string `json:"userId"`
	userName  string `json:"userName"`
	contactNo string `json:"contactNo"`
	address   string `json:"address"`
}

type article struct {
	articleName  string `json:"articleName"`
	articleDesc  string `json:"articleDesc"`
	articleQty   string `json:"articleQty"`
	articlePrice string `json:"articlePrice"`
}
