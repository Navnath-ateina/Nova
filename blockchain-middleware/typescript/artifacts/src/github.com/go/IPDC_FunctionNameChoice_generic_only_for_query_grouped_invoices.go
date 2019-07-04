package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//type IPDCChaincode struct {
//}

func (t *IPDCChaincode) Internalfunctionname(function string, input_args []string) (string, error) {

	fmt.Println("***********Entering Internalfunctionname***********")

	fmt.Println("Function is: " + function)

	if function == "query_grouped_invoice_by_update_status" {

		if len(input_args) < 3 {

			fmt.Println("***********Exiting Internalfunctionname***********")

			fmt.Println("Error: Incorrect number of arguments")

			return "", errors.New("Error: Incorrect number of arguments")
		}

		//var record_string =  input_args[0]

		var record_specification_input map[string]interface{}

		var err = json.Unmarshal([]byte(input_args[2]), &record_specification_input)

		if err != nil {

			fmt.Println("***********Exiting Internalfunctionname***********")

			fmt.Println("Error parsing argument for BDType")

			return "", errors.New("Error parsing argument for BDType")
		}

		BDType_string, ok := record_specification_input["BDType"].(string)

		if !ok {
			fmt.Println("***********Exiting Internalfunctionname***********")

			fmt.Println("Error parsing value of BDType")

			return "", errors.New("Error parsing value of BDType")
		}

		if DISB_BDTYPE == strings.TrimSpace(BDType_string) {

			fmt.Println("Going to query_grouped_invoice_by_update_status_internal_BD")

			fmt.Println("***********Exiting Internalfunctionname***********")

			return "query_grouped_invoice_by_update_status_internal_BD", nil

		} else if ASN_BDTYPE == strings.TrimSpace(BDType_string) {

			fmt.Println("Going to query_grouped_invoice_by_update_status_internal_ASN")

			fmt.Println("***********Exiting Internalfunctionname***********")

			return "query_grouped_invoice_by_update_status_internal_ASN", nil

		} else if len(strings.TrimSpace(BDType_string)) == 0 {

			fmt.Println("Going to query_grouped_invoice_by_update_status_internal_ALL")

			fmt.Println("***********Exiting Internalfunctionname***********")

			return "query_grouped_invoice_by_update_status_internal_ALL", nil

		} else {
			fmt.Println("Error: Invalid BDType")

			fmt.Println("***********Exiting Internalfunctionname***********")

			return "", errors.New("Error: Invalid BDType")
		}

	} else {

		fmt.Println("***********Exiting Internalfunctionname***********")

		return function, nil

	}

}
