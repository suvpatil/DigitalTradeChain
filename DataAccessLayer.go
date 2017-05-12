package main

import (
	"errors"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func createDatabase(stub shim.ChaincodeStubInterface, args []string) error {
	var err error	
	//Create table "ContractDetails"
	err = stub.CreateTable("contractDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "contractId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "contractList", Type: shim.ColumnDefinition_BYTES, Key: false},

		
	})
	if err != nil {
		return errors.New("Failed creating ContractDetails table")
	}

	err = stub.CreateTable("attachmentDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "contractId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "attachmentName", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "documentBlob", Type: shim.ColumnDefinition_BYTES, Key: false},
	})
	if err != nil {
		return errors.New("Failed creating attachmentDetails table.")
	}

	err = stub.CreateTable("userDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "userId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "contractIdList", Type: shim.ColumnDefinition_BYTES, Key: false},
	})
	if err != nil {
		return errors.New("Failed creating userDetails table.")
	}

	return nil

}

func insertContractDetails(stub shim.ChaincodeStubInterface, contId string, contractDetails contract) (bool, error) {
	var err error
	var ok bool
	jsonAsBytes, _ := json.Marshal(contractDetails)
	ok, err = stub.InsertRow("contractDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: contId}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: jsonAsBytes}},
			
		},
	})
	return ok, err
}

func getContractSpecificList(stub shim.ChaincodeStubInterface, contractId string) (contract, error) {
	
	var columns []shim.Column
	var contractList contract

	col1 := shim.Column{Value: &shim.Column_String_{String_: contractId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("contractDetails", columns)
	if err != nil {
		return contractList, errors.New("Failed to query table contractDetails")
	}

	json.Unmarshal(row.Columns[1], &contractList)
	return contractList, nil
}
/*
func updateContractDetails(stub shim.ChaincodeStubInterface, contractDetails contract) (bool, error) {
	ok, err := stub.ReplaceRow("ContractDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractId}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.OrderId}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentCommitment}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentConfirmation}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.InformationCounterparty}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.ForfeitingInvoice}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractCreateDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentDueDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.InvoiceStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.DeliveryStatus}},
		},
	})

	if !ok && err == nil {
		return false, errors.New("Error in updating Seller record.")
	}
	return true, nil
}
*/
