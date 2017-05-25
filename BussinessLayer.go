package main

import (
	"encoding/json"
	"errors"

	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func initializeChaincode(stub shim.ChaincodeStubInterface, args []string) error {
	var ok bool
	var err error
	ok, err = createDatabase(stub, args)
	if !ok {
		return err
	}
	ok, err = initializeUsers(stub)
	if !ok {
		return err
	}
	return nil
}

func saveContractDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var contractDetails contract
	var err error
	var ok bool

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Need 1 arguments")
	}

	json.Unmarshal([]byte(args[0]), &contractDetails)

	contractDetails = addContractInformation(contractDetails)

	ok, err = insertContractDetails(stub, contractDetails)
	if !ok && err == nil {
		return nil, errors.New("Error in adding OrderDetails record")
	}

	ok, err = updateUsersContractList(stub, contractDetails)
	if !ok {
		return nil, err
	}

	return nil, nil
}

func getContractDetailsByContractId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	contractId := args[0]
	contractDetails, _ := getContractDetails(stub, contractId)

	jsonAsBytes, _ := json.Marshal(contractDetails)
	return jsonAsBytes, nil

}

func saveAttachmentDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var ok bool

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Need 3 arguments")
	}

	contractId := args[0]
	attachmentName := args[1]
	documentBlob := args[2]

	ok, err = insertAttachmentDetails(stub, contractId, attachmentName, documentBlob)
	if !ok && err == nil {
		return nil, errors.New("Error in inserting attachment")
	}

	return nil, err
}

func getAttachment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Need 2 arguments")
	}

	contractId := args[0]
	attachmentName := args[1]

	jsonAsBytes, err := getAttachmentDetails(stub, contractId, attachmentName)
	if err != nil {
		return nil, errors.New("Error in downloading the attachment")
	}

	return jsonAsBytes, nil
}

func addContractInformation(contractDetails contract) contract {
	contractDetails.ContractId = strconv.Itoa(rand.Int() + 1)
	current_time := time.Now().Local()
	contractDetails.ContractCreateDate = current_time.Format("2006-01-02")
	contractDetails.IsLCAttached = false
	contractDetails.IsPOAttached = true
	contractDetails.IsBillOfLedingAttached = false
	contractDetails.IsInvoiceListAttached = false
	contractDetails.ActionPendingOn = "Buyer"
	contractDetails.ContractStatus = "Contract Created"

	return contractDetails
}

func updateUsersContractList(stub shim.ChaincodeStubInterface, contractDetails contract) (bool, error) {
	var ok bool
	var userContractList []string

	//Update Seller's Contract List
	userContractList, ok = getUserContractList(stub, contractDetails.SellerDetails.Seller.UserId)
	if !ok {
		return ok, errors.New("Error in geting Seller's contract list")
	}
	userContractList = append(userContractList, contractDetails.ContractId)
	ok = updateUserContractList(stub, contractDetails.SellerDetails.Seller.UserId, userContractList)
	if !ok {
		return ok, errors.New("Error in updating Seller's contract list")
	}

	//Update SellerBank's Contract List
	userContractList, ok = getUserContractList(stub, contractDetails.SellerDetails.SellerBank.UserId)
	if !ok {
		return ok, errors.New("Error in geting SellerBank's contract list")
	}
	userContractList = append(userContractList, contractDetails.ContractId)
	ok = updateUserContractList(stub, contractDetails.SellerDetails.SellerBank.UserId, userContractList)
	if !ok {
		return ok, errors.New("Error in updating SellerBank's contract list")
	}

	//Update Buyer's Contract List
	userContractList, ok = getUserContractList(stub, contractDetails.BuyerDetails.Buyer.UserId)
	if !ok {
		return ok, errors.New("Error in geting Buyer's contract list")
	}
	userContractList = append(userContractList, contractDetails.ContractId)
	ok = updateUserContractList(stub, contractDetails.BuyerDetails.Buyer.UserId, userContractList)
	if !ok {
		return ok, errors.New("Error in updating Buyer's contract list")
	}

	//Update BuyerBank's Contract List
	userContractList, ok = getUserContractList(stub, contractDetails.BuyerDetails.BuyerBank.UserId)
	if !ok {
		return ok, errors.New("Error in geting BuyerBank's contract list")
	}
	userContractList = append(userContractList, contractDetails.ContractId)
	ok = updateUserContractList(stub, contractDetails.BuyerDetails.BuyerBank.UserId, userContractList)
	if !ok {
		return ok, errors.New("Error in updating BuyerBank's contract list")
	}

	//Update Transporter's Contract List
	userContractList, ok = getUserContractList(stub, contractDetails.DeliveryDetails.TransporterDetails.UserId)
	if !ok {
		return ok, errors.New("Error in geting Transporter's contract list")
	}
	userContractList = append(userContractList, contractDetails.ContractId)
	ok = updateUserContractList(stub, contractDetails.DeliveryDetails.TransporterDetails.UserId, userContractList)
	if !ok {
		return ok, errors.New("Error in updating Transporter's contract list")
	}

	return true, nil
}

func getContractDetailsByUserId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var contractDetails []contract
	var contract contract

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Need 1 argument")
	}
	userId := args[0]

	contractIdList, ok := getUserContractList(stub, userId)
	if !ok {
		return nil, errors.New("Error in geting user specific contract list")
	}

	for _, element := range contractIdList {
		contractId := element
		contract, _ = getContractDetails(stub, contractId)
		contractDetails = append(contractDetails, contract)
	}
	contractAsBytes, _ := json.Marshal(contractDetails)
	return contractAsBytes, nil

}
func UpdateContractStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var ok bool
	var err error
	//var status statusMaintained
	//var contractLists contract

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Need 3 arguments")
	}

	userID := args[0]
	contractID := args[1]
	current_time := time.Now().Local()
	contractList, _ := getContractDetails(stub, contractID)

	contractStatus := contractList.ContractStatus
	//for seller
	if contractList.SellerDetails.Seller.UserId == userID {
		if contractStatus == LC_Approved {
			contractList.ContractStatus = Ready_For_Shipment
			contractList.ActionPendingOn = "Transporter"
			contractList.ReadyForShipmentBySellerDate = current_time.Format("2006-01-02")
		} else if contractStatus == Shipment_Delivered {
			contractList.ContractStatus = Invoice_Created
			contractList.ActionPendingOn = "Seller Bank"
			contractList.InvoiceCreatedBySellerDate = current_time.Format("2006-01-02")
		}
	}

	//for buyer
	if contractList.BuyerDetails.Buyer.UserId == userID {
		if contractStatus == Contract_Created {
			contractList.ContractStatus = Contract_Accepted
			contractList.ActionPendingOn = "Buyer Bank"
			contractList.ApprovedContractByBuyerDate = current_time.Format("2006-01-02")
		} else if contractStatus == Payment_Completed_to_Seller_Bank {
			contractList.ContractStatus = Contract_Completed
			contractList.ActionPendingOn = Contract_Completed
			contractList.ContractCompletedByBuyerDate = current_time.Format("2006-01-02")
		} else if contractStatus == Shipment_Inprogress {
			contractList.ContractStatus = Shipment_Delivered
			contractList.ActionPendingOn = "Seller"
			contractList.ShipmentDeliveredByBuyerDate = current_time.Format("2006-01-02")
		}
	}

	//for sellerBank
	if contractList.SellerDetails.SellerBank.UserId == userID {
		if contractStatus == LC_Created {
			contractList.ContractStatus = LC_Approved
			contractList.ActionPendingOn = "Seller"
			contractList.LCApprovedBySellerBankDate = current_time.Format("2006-01-02")
		} else if contractStatus == Invoice_Created {
			contractList.ContractStatus = Payment_Completed_to_Seller
			contractList.ActionPendingOn = "Buyer Bank"
			contractList.PaymentCompletedToSellerBySellerBankDate = current_time.Format("2006-01-02")
		}
	}

	//for buyerBank
	if contractList.BuyerDetails.BuyerBank.UserId == userID {
		if contractStatus == Contract_Accepted {
			contractList.ContractStatus = LC_Created
			contractList.ActionPendingOn = "Seller Bank"
			contractList.LCCreatedByBuyerBankDate = current_time.Format("2006-01-02")
		} else if contractStatus == Payment_Completed_to_Seller {
			contractList.ContractStatus = Payment_Completed_to_Seller_Bank
			contractList.ActionPendingOn = "Buyer"
			contractList.PaymentCompletedToSellerBankByBuyerBankDate = current_time.Format("2006-01-02")
		}
	}

	//for transporter
	if contractList.DeliveryDetails.TransporterDetails.UserId == userID {
		if contractStatus == Ready_For_Shipment {
			contractList.ContractStatus = Shipment_Inprogress
			contractList.ActionPendingOn = "Buyer"
			contractList.ShipmentInProgressByTransDate = current_time.Format("2006-01-02")
		}
	}

	contractList.LastUpdatedDate = current_time.Format("2006-01-02")
	//status = setStructStatus(stub, status, userID, contractStatus)
	ok = updateContractListByContractID(stub, contractID, contractList)
	if !ok {
		return nil, errors.New("Error in updating contract list")
	}

	return nil, err

}
