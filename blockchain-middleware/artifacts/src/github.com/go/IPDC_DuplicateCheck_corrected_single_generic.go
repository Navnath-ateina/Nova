package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *IPDCChaincode) query_customer_invoice_duplicate_passed(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	var arguments []string

	var record_specification map[string]interface{}

	err := json.Unmarshal([]byte(args[0]), &record_specification)

	if err != nil {

		fmt.Println("Error in decoding Record JSON")

		return shim.Error("JSON format for Record is not correct")
	}

	if len(record_specification) != 3 {

		fmt.Println("Input should have three fields")

		return shim.Error("Input should have three fields")
	}

	var additional_json interface{}

	var specs map[string]interface{}

	var ok bool

	additional_json, ok = map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {

				record_specification[spec] = additional_json_data[spec]
			}

		} else {
			fmt.Println("Invalid additional JSON fields in specification")

			return shim.Error("Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Invalid keys map specification.")

		return shim.Error("Invalid keys map specification.")
	}

	var composite_key = make(map[string]interface{})

	composite_key["Duplicate_check_passed"], ok = specs["Duplicate_check_passed"]

	if !ok {
		fmt.Println("Duplicate_check_passed does not exist in config")

		return shim.Error("Duplicate_check_passed does not exist in config")
	}

	var compositekeyJsonString []byte

	compositekeyJsonString, err = json.Marshal(composite_key)

	if err != nil {

		fmt.Println("Unable to Marshal composite key specification to JSON")

		return shim.Error("Unable to Marshal composite key specification to JSON")
	}

	var concatenated_record_json []byte

	concatenated_record_json, err = json.Marshal(record_specification)

	if err != nil {

		fmt.Println("Unable to Marshal Concatenated Record to JSON")

		return shim.Error("Unable to Marshal Concatenated Record to JSON")
	}

	arguments = append(arguments, string(concatenated_record_json))

	arguments = append(arguments, string(compositekeyJsonString))

	err_iterator, tableRowsIterator := t.query_by_composite_key_primitive_string_args(stub, arguments)

	if err_iterator != nil {

		fmt.Println("Error in query_by_composite_key_primitive_string_args " + err_iterator.Error())

		//return shim.Success(nil)

		return shim.Error("Error in query_by_composite_key_primitive_string_args " + err_iterator.Error())

	}

	fmt.Println("***********Entering fetchDisbursedRecordsFromCompositeKeys***********")

	records, err4 := t.fetchDuplicateCheckPassedRecordFromCompositeKey(stub, tableRowsIterator)

	fmt.Println("***********Exiting fetchDisbursedRecordsFromCompositeKeys***********")

	if len(records) != 0 {

		return shim.Success([]byte(records[0]))

	} else if err4 != nil {

		fmt.Println("Error: " + err4.Error())

		return shim.Error("Error: " + err4.Error())

	} else {

		return shim.Success(nil)
	}

}

func (t *IPDCChaincode) fetchDuplicateCheckPassedRecordFromCompositeKey(stub shim.ChaincodeStubInterface, tableRowsIterator shim.StateQueryIteratorInterface) ([]string, error) {

	var i int

	var outputstringarray []string

	errstring := ""

	defer tableRowsIterator.Close()

	for i = 0; tableRowsIterator.HasNext(); i++ {

		indexKey, err := tableRowsIterator.Next()

		if err != nil {

			return nil, err
		}

		objectType, compositeKeyParts, err := stub.SplitCompositeKey(indexKey.Key)

		if err != nil {

			return nil, err

		}
		var completeKey string

		length_of_compositekey := len(compositeKeyParts)

		for j := 0; j < length_of_compositekey; j++ {

			if j == 0 {

				completeKey += compositeKeyParts[j]

			} else {

				completeKey += "~" + compositeKeyParts[j]

			}
		}

		fmt.Println("ObjectType is " + objectType)

		fmt.Println("Key is " + completeKey)

		returnedKey := compositeKeyParts[length_of_compositekey-1]

		valAsBytes, err2 := stub.GetState(returnedKey)

		if err2 != nil {

			fmt.Println("Failed to get state for key: " + returnedKey)

			errstring = errstring + fmt.Sprintf("Failed to get state for key: %s. ", returnedKey)

			continue

		} else if valAsBytes == nil {

			fmt.Println("No value for key : " + returnedKey)

			errstring = errstring + fmt.Sprintf("No value for key : %s. ", returnedKey)

			continue
		}

		outputstringarray = append(outputstringarray, string(valAsBytes))

		break

	}

	var error_return error

	if errstring == "" {

		error_return = nil

	} else {

		error_return = errors.New(errstring)
	}

	return outputstringarray, error_return
}

func (t *IPDCChaincode) invoke_cross_channel_duplicate_check(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Get map specification for update duplicate-status=Failed

	key_for_func_invoice_duplicate_failed := "FunctionName*" + "invoke_invoice_duplicate_failed"

	valAsBytes_invoice_duplicate_failed, err_invoice_duplicate_failed := stub.GetState(key_for_func_invoice_duplicate_failed)

	if err_invoice_duplicate_failed != nil {

		return shim.Error("Failed to get specification for invoke_invoice_duplicate_failed: " + err_invoice_duplicate_failed.Error())

	} else if valAsBytes_invoice_duplicate_failed == nil {
		fmt.Println("No value for key : " + key_for_func_invoice_duplicate_failed)

		return shim.Error("No value for key : " + key_for_func_invoice_duplicate_failed)
	}

	var json_specification_invoice_duplicate_failed interface{}

	err_invoice_duplicate_failed = json.Unmarshal(valAsBytes_invoice_duplicate_failed, &json_specification_invoice_duplicate_failed)

	if err_invoice_duplicate_failed != nil {
		fmt.Println("Error in decoding Specification JSON for invoke_invoice_duplicate_failed")

		return shim.Error("JSON format for Specification is not correct for invoke_invoice_duplicate_failed")
	}

	map_specification_invoice_duplicate_failed, ok_invoice_duplicate_failed := json_specification_invoice_duplicate_failed.(map[string]interface{})

	if !ok_invoice_duplicate_failed {
		fmt.Println("Error Parsing map_specification for invoke_invoice_duplicate_failed")

		return shim.Error("Error Parsing map_specification for invoke_invoice_duplicate_failed")
	}

	//Get map specification for update duplicate-status=Passed

	key_for_func_invoice_duplicate_passed := "FunctionName*" + "invoke_invoice_duplicate_passed"

	valAsBytes_invoice_duplicate_passed, err_invoice_duplicate_passed := stub.GetState(key_for_func_invoice_duplicate_passed)

	if err_invoice_duplicate_passed != nil {

		return shim.Error("Failed to get specification for invoke_invoice_duplicate_passed: " + err_invoice_duplicate_passed.Error())

	} else if valAsBytes_invoice_duplicate_passed == nil {
		fmt.Println("No value for key : " + key_for_func_invoice_duplicate_passed)

		return shim.Error("No value for key : " + key_for_func_invoice_duplicate_passed)
	}

	var json_specification_invoice_duplicate_passed interface{}

	err_invoice_duplicate_passed = json.Unmarshal(valAsBytes_invoice_duplicate_passed, &json_specification_invoice_duplicate_passed)

	if err_invoice_duplicate_passed != nil {
		fmt.Println("Error in decoding Specification JSON for invoke_invoice_duplicate_passed")

		return shim.Error("JSON format for Specification is not correct for invoke_invoice_duplicate_passed")
	}

	map_specification_invoice_duplicate_passed, ok_invoice_duplicate_passed := json_specification_invoice_duplicate_passed.(map[string]interface{})

	if !ok_invoice_duplicate_passed {
		fmt.Println("Error Parsing map_specification for invoke_invoice_invoice_duplicate_passed")

		return shim.Error("Error Parsing map_specification for invoke_invoice_invoice_duplicate_passed")
	}

	//Query in same channel

	key_for_func_same_channel := "FunctionName*query_customer_invoice_duplicate_passed"

	valAsBytes_same_channel, err_same_channel := stub.GetState(key_for_func_same_channel)

	if err_same_channel != nil {

		return shim.Error("Failed to get state: " + err_same_channel.Error())

	} else if valAsBytes_same_channel == nil {

		fmt.Println("No value for key : " + key_for_func_same_channel)

		return shim.Error("No value for key : " + key_for_func_same_channel)
	}

	var json_specification_same_channel interface{}

	err_same_channel = json.Unmarshal(valAsBytes_same_channel, &json_specification_same_channel)

	if err_same_channel != nil {

		fmt.Println("Error in decoding Specification JSON for same channel query")

		return shim.Error("JSON format for Specification is not correct for same channel query")
	}

	map_specification_same_channel, ok_same_channel1 := json_specification_same_channel.(map[string]interface{})

	if !ok_same_channel1 {

		fmt.Println("Error Parsing map_specification for same channel query")

		return shim.Error("Error Parsing map_specification for same channel query")
	}

	var successstring string

	//Get list of channels and chaincodes

	var list_unmarshal interface{}

	err := json.Unmarshal([]byte(args[0]), &list_unmarshal)

	if err != nil {

		return shim.Error("Incorrect Specification of Channels and Chaincodes")
	}

	var list_interface []interface{}

	var ok1 bool

	list_interface, ok1 = list_unmarshal.([]interface{})

	if !ok1 {
		return shim.Error("Incorrect list specification of Channels and Chaincodes")
	}

	var list_channel_chaincode_maps [](map[string]string)

	for i, channel_chaincode_interface := range list_interface {

		channel_chaincode_map, ok := channel_chaincode_interface.(map[string]interface{})

		if !ok {
			fmt.Println(fmt.Sprintf("Error: Cannot parse element %d of list of maps. ", i))

			return shim.Error(fmt.Sprintf("Error: Cannot parse element %d of list of maps. ", i))

		}

		var channel_string, chaincode_string string

		channel_string, ok = channel_chaincode_map["channel-name"].(string)

		if !ok {
			fmt.Println(fmt.Sprintf("Error: Cannot parse channel-name in element %d of list of maps. ", i))

			return shim.Error(fmt.Sprintf("Error: Cannot parse channel-name in element %d of list of maps. ", i))

		}

		chaincode_string, ok = channel_chaincode_map["chaincode-name"].(string)

		if !ok {
			fmt.Println(fmt.Sprintf("Error: Cannot parse chaincode-name in element %d of list of maps. ", i))

			return shim.Error(fmt.Sprintf("Error: Cannot parse chaincode-name in element %d of list of maps. ", i))

		}

		var channel_chaincode_map_string map[string]string

		channel_chaincode_map_string = make(map[string]string)

		channel_chaincode_map_string["channel-name"] = channel_string

		channel_chaincode_map_string["chaincode-name"] = chaincode_string

		list_channel_chaincode_maps = append(list_channel_chaincode_maps, channel_chaincode_map_string)
	}

	// fetch all <duplicate-status=False, disbursement-status=False, clearance-status=True> records

	var nested_map map[string](map[string](map[string]string))

	nested_map = make(map[string](map[string](map[string]string)))

	var anchor_vendor_processed_map map[string](map[string]string)

	anchor_vendor_processed_map = make(map[string](map[string]string))

	// First get all the onboarded vendors with financing enabled

	compositekeyJsonString_for_vendors := "{\"To_iterate_vendors\":[\"record-type\",\"financing-enabled\",\"BDType\"]}"

	concatenated_record_json_for_vendors := "{\"record-type\":\"onboarded_vendor\",\"financing-enabled\":\"True\"}"

	var arguments_for_vendors []string

	arguments_for_vendors = append(arguments_for_vendors, concatenated_record_json_for_vendors)

	arguments_for_vendors = append(arguments_for_vendors, compositekeyJsonString_for_vendors)

	err_iterator_for_vendors, tableRowsIterator_for_vendors := t.query_by_composite_key_primitive_string_args(stub, arguments_for_vendors)

	if err_iterator_for_vendors != nil {

		return shim.Error("Error: Unable to Fetch Vendors")
	}

	defer tableRowsIterator_for_vendors.Close()

	var proceessed_records_counter int

	var exceeded_processing_limit bool

	exceeded_processing_limit = false

	proceessed_records_counter = 1

	for i_vendors := 0; tableRowsIterator_for_vendors.HasNext(); i_vendors++ {

		if exceeded_processing_limit {

			break
		}

		indexKey_for_vendors, err_for_vendors := tableRowsIterator_for_vendors.Next()
		if err_for_vendors != nil {

			return shim.Error("Error in iterating over vendors: " + err_for_vendors.Error())
		}

		objectType_for_vendors, compositeKeyParts_for_vendors, err_for_vendors := stub.SplitCompositeKey(indexKey_for_vendors.Key)

		if err_for_vendors != nil {

			return shim.Error("Error splitting composite key: " + err_for_vendors.Error())
		}

		var completeKey_for_vendors string

		length_of_compositekey_for_vendors := len(compositeKeyParts_for_vendors)

		for j := 0; j < length_of_compositekey_for_vendors; j++ {

			if j == 0 {
				completeKey_for_vendors += compositeKeyParts_for_vendors[j]
			} else {
				completeKey_for_vendors += "~" + compositeKeyParts_for_vendors[j]
			}
		}

		fmt.Println("ObjectType is ", objectType_for_vendors)
		fmt.Println("Key is ", completeKey_for_vendors)

		returnedKey_for_vendors := compositeKeyParts_for_vendors[length_of_compositekey_for_vendors-1]

		valAsBytes_for_vendors, err_for_vendors2 := stub.GetState(returnedKey_for_vendors)
		if err_for_vendors2 != nil {

			fmt.Println("Failed to get state for key: " + returnedKey_for_vendors)

			successstring = successstring + fmt.Sprintf("Failed to get state for vendor key: %s. ", returnedKey_for_vendors)

			continue

		} else if valAsBytes_for_vendors == nil {

			fmt.Println("No value for vendor key : " + returnedKey_for_vendors)

			successstring = successstring + fmt.Sprintf("No value for vendor key : %s. ", returnedKey_for_vendors)

			continue
		}

		var fullrecord_for_vendors map[string]interface{}

		err_for_vendors = json.Unmarshal(valAsBytes_for_vendors, &fullrecord_for_vendors)

		if err_for_vendors != nil {
			fmt.Println("Error unmarshalling record for key : " + returnedKey_for_vendors)

			successstring = successstring + fmt.Sprintf("Error unmarshalling record for key : %s. ", returnedKey_for_vendors)

			continue

		}

		var vendor_code_string_for_vendors, customer_code_string_for_vendors, anchor_code_string_for_vendors string

		var ok bool

		vendor_code_string_for_vendors, ok = fullrecord_for_vendors["vendor-code"].(string)

		if !ok {
			fmt.Println("Error parsing vendor-code for key :  " + returnedKey_for_vendors)

			successstring = successstring + fmt.Sprintf("Error parsing vendor-code for key : %s. ", returnedKey_for_vendors)

			continue
		}

		customer_code_string_for_vendors, ok = fullrecord_for_vendors["customer-code"].(string)

		if !ok {
			fmt.Println("Error parsing customer-code for key : " + returnedKey_for_vendors)

			successstring = successstring + fmt.Sprintf("Error parsing customer-code for key : %s. ", returnedKey_for_vendors)

			continue
		}

		anchor_code_string_for_vendors, ok = fullrecord_for_vendors["anchor-code"].(string)

		if !ok {
			fmt.Println("Error parsing anchor-code for key : " + returnedKey_for_vendors)

			successstring = successstring + fmt.Sprintf("Error parsing anchor-code for key : %s. ", returnedKey_for_vendors)

			continue
		}

		if _, ok = anchor_vendor_processed_map[anchor_code_string_for_vendors]; !ok {

			fmt.Println("Anchor-code: " + anchor_code_string_for_vendors + " not processed earlier." + "Vendor " + vendor_code_string_for_vendors)

			anchor_vendor_processed_map[anchor_code_string_for_vendors] = make(map[string]string)

			anchor_vendor_processed_map[anchor_code_string_for_vendors][vendor_code_string_for_vendors] = "Processed"

		} else if _, ok = anchor_vendor_processed_map[anchor_code_string_for_vendors][vendor_code_string_for_vendors]; !ok {

			fmt.Println("Anchor-code: " + anchor_code_string_for_vendors + " processed earlier. " + "Vendor " + vendor_code_string_for_vendors + " not processed earlier.")

			anchor_vendor_processed_map[anchor_code_string_for_vendors][vendor_code_string_for_vendors] = "Processed"

		} else {

			fmt.Println("Anchor-code: " + anchor_code_string_for_vendors + " processed earlier. " + "Vendor " + vendor_code_string_for_vendors + " processed earlier.")

			continue

		}

		// fetch all <duplicate-status=False, disbursement-status=False, clearance-status=True> records

		compositekeyJsonString_for_dup_check := "{\"Duplicate_check\":[\"record-type\",\"duplicate-status\",\"vendor-code\",\"anchor-code\"]}"

		concatenated_record_json_for_dup_check := "{\"record-type\":\"invoice\",\"duplicate-status\":\"False\",\"vendor-code\":\"" + vendor_code_string_for_vendors + "\",\"anchor-code\":\"" + anchor_code_string_for_vendors + "\"}"

		var arguments_for_dup_check []string

		arguments_for_dup_check = append(arguments_for_dup_check, concatenated_record_json_for_dup_check)

		arguments_for_dup_check = append(arguments_for_dup_check, compositekeyJsonString_for_dup_check)

		err_iterator, tableRowsIterator := t.query_by_composite_key_primitive_string_args(stub, arguments_for_dup_check)

		if err_iterator != nil {

			return shim.Error("Error: Unable to Fetch Invoices for Duplicate Check")
		}

		//loop over fetched records and do duplicate check

		defer tableRowsIterator.Close()

		for i_invoices := 0; tableRowsIterator.HasNext(); i_invoices++ {

			indexKey, err := tableRowsIterator.Next()
			if err != nil {

				return shim.Error("Error in iterating over invoices: " + err.Error())
			}

			objectType, compositeKeyParts, err := stub.SplitCompositeKey(indexKey.Key)
			if err != nil {

				return shim.Error("Error splitting composite key: " + err.Error())
			}
			var completeKey string
			length_of_compositekey := len(compositeKeyParts)
			for j := 0; j < length_of_compositekey; j++ {
				if j == 0 {
					completeKey += compositeKeyParts[j]
				} else {
					completeKey += "~" + compositeKeyParts[j]
				}
			}

			fmt.Println("ObjectType is ", objectType)
			fmt.Println("Key is ", completeKey)

			returnedKey := compositeKeyParts[length_of_compositekey-1]

			valAsBytes, err2 := stub.GetState(returnedKey)
			if err2 != nil {

				fmt.Println("Failed to get state for key: " + returnedKey)

				successstring = successstring + fmt.Sprintf("Failed to get state for key: %s. ", returnedKey)

				continue

			} else if valAsBytes == nil {

				fmt.Println("No value for key : " + returnedKey)

				successstring = successstring + fmt.Sprintf("No value for key : %s. ", returnedKey)

				continue
			}

			var fullrecord map[string]interface{}

			err = json.Unmarshal(valAsBytes, &fullrecord)

			if err != nil {
				fmt.Println("Error unmarshalling record for key : " + returnedKey)

				successstring = successstring + fmt.Sprintf("Error unmarshalling record for key : %s. ", returnedKey)

				continue

			}

			var vendor_code_string, anchor_code_string string

			var ok bool

			vendor_code_string, ok = fullrecord["vendor-code"].(string)

			if !ok {

				fmt.Println(fmt.Sprintf("Serious Error: Unable to parse Vendor code in record %s. Terminating.", returnedKey))

				return shim.Error(fmt.Sprintf("Serious Error: Unable to parse Vendor code in record %s. Terminating.", returnedKey))

			}

			if vendor_code_string != vendor_code_string_for_vendors {

				fmt.Println(fmt.Sprintf("Serious Error: Vendor code in composite key different from vendor code value in record %s. Terminating.", returnedKey))

				return shim.Error(fmt.Sprintf("Serious Error: Vendor code in composite key different from vendor code value in record %s. Terminating.", returnedKey))
			}

			anchor_code_string, ok = fullrecord["anchor-code"].(string)

			if !ok {

				fmt.Println(fmt.Sprintf("Serious Error: Unable to parse Anchor code in record %s. Terminating.", returnedKey))

				return shim.Error(fmt.Sprintf("Serious Error: Unable to parse Anchor code in record %s. Terminating.", returnedKey))

			}

			if anchor_code_string != anchor_code_string_for_vendors {

				fmt.Println(fmt.Sprintf("Serious Error: Anchor code in composite key different from vendor code value in record %s. Terminating.", returnedKey))

				return shim.Error(fmt.Sprintf("Serious Error: Anchor code in composite key different from vendor code value in record %s. Terminating.", returnedKey))
			}

			var customer_code_string string

			customer_code_string = customer_code_string_for_vendors

			//create argument for cross channel query

			var cross_channel_query map[string]string

			cross_channel_query = make(map[string]string)

			cross_channel_query["customer-code"] = customer_code_string

			cross_channel_query["invoice-number"], ok = fullrecord["invoice-number"].(string)

			if !ok {
				fmt.Println("Invalid invoice-number, ignoring invoice. ")

				successstring = successstring + fmt.Sprintf("Invalid invoice-number, ignoring invoice. ")

				continue
			}

			var fullrecord_invoice_date string

			fullrecord_invoice_date, ok = fullrecord["invoice-date"].(string)

			if !ok {
				fmt.Println("Invalid invoice-date, ignoring invoice. ")

				successstring = successstring + fmt.Sprintf("Invalid invoice-date, ignoring invoice. ")

				continue
			}

			cross_channel_query["two-digit-invoice-financial-year"], ok = fullrecord["two-digit-invoice-financial-year"].(string)

			if !ok {
				fmt.Println("Invalid two-digit-invoice-financial-year, ignoring invoice. ")

				successstring = successstring + fmt.Sprintf("Invalid two-digit-invoice-financial-year, ignoring invoice. ")

				continue
			}

			cross_channel_query_string, err_cross_channel_query_string := json.Marshal(cross_channel_query)

			if err_cross_channel_query_string != nil {

				fmt.Println(fmt.Sprintf("Unable to Mashal the argument for cross channel query for vendor code %s, customer code %s, invoice number %s, two digit year %s. ", vendor_code_string, customer_code_string, cross_channel_query["invoice-number"], cross_channel_query["two-digit-invoice-financial-year"]))

				successstring = successstring + fmt.Sprintf("Unable to Mashal the argument for cross channel query for vendor code %s, customer code %s, invoice number %s, two digit year %s. ", vendor_code_string, customer_code_string, cross_channel_query["invoice-number"], cross_channel_query["two-digit-invoice-financial-year"])

				continue
			}

			invoke_update_primary_string := "{\"invoice-number\":\"" + cross_channel_query["invoice-number"] + "\",\"invoice-date\":\"" + fullrecord_invoice_date + "\",\"anchor-code\":\"" + anchor_code_string + "\",\"vendor-code\":\"" + vendor_code_string + "\",\"customer-code\":\"" + customer_code_string + "\"}"

			var args_for_duplicate_found_update []string

			args_for_duplicate_found_update = append(args_for_duplicate_found_update, string(invoke_update_primary_string))

			var args_for_duplicate_not_found_update []string

			args_for_duplicate_not_found_update = append(args_for_duplicate_not_found_update, string(invoke_update_primary_string))

			// Do the cross channel queries

			cross_channel_queryArgs := make([][]byte, 2)

			cross_channel_queryArgs[0] = []byte("query_customer_invoice_duplicate_passed")

			cross_channel_queryArgs[1] = cross_channel_query_string

			fmt.Println("Initiating cross channel queries: " + string(cross_channel_query_string))

			var isduplicate bool

			isduplicate = false

			for index_channel_chaincode_map, channel_chaincode_map := range list_channel_chaincode_maps {

				var channel_name string

				var chaincode_name string

				channel_name, ok = channel_chaincode_map["channel-name"]

				if !ok {
					fmt.Println(fmt.Sprintf("Invalid channel-name for map number %d.", index_channel_chaincode_map))

					return shim.Error(fmt.Sprintf("Invalid channel-name for map number %d. Skipping query. ", index_channel_chaincode_map))

				}

				chaincode_name, ok = channel_chaincode_map["chaincode-name"]

				if !ok {

					fmt.Println(fmt.Sprintf("Invalid chaincode-name for map number %d.", index_channel_chaincode_map))

					return shim.Error(fmt.Sprintf("Invalid chaincode-name for map number %d.", index_channel_chaincode_map))

				}

				fmt.Println(fmt.Sprintf("Channel-name: %s, chaincode-name: %s", channel_name, chaincode_name))

				response := stub.InvokeChaincode(chaincode_name, cross_channel_queryArgs, channel_name)

				if response.Status != shim.OK {

					fmt.Println(fmt.Sprintf("Error in querying channel-name: %s, chaincode-name %s, with query: %s. Error message: %s.", channel_name, chaincode_name, string(cross_channel_query_string), string(response.Message)))

					return shim.Error(fmt.Sprintf("Error in querying channel-name: %s, chaincode-name %s, with query: %s. Error message: %s.", channel_name, chaincode_name, string(cross_channel_query_string), string(response.Message)))

				}

				if response.Payload == nil {

					fmt.Println("response.Payload == nil")

					continue

				} else {

					fmt.Println(fmt.Sprintf("Found Duplicate Invoice %s in channel %s, chaincode %s, for this invoice %s. ", string(response.Payload), channel_name, chaincode_name, invoke_update_primary_string))

					successstring = successstring + fmt.Sprintf("Found Duplicate Invoice %s in channel %s, chaincode %s, for this invoice %s. ", string(response.Payload), channel_name, chaincode_name, invoke_update_primary_string)

					isduplicate = true

					break
				}
			}

			if !isduplicate { //do a check in this channel itself

				var cross_channel_query_string_array []string

				cross_channel_query_string_array = append(cross_channel_query_string_array, string(cross_channel_query_string))

				response := t.query_customer_invoice_duplicate_passed(stub, cross_channel_query_string_array, map_specification_same_channel)

				if response.Status != shim.OK {

					fmt.Println(fmt.Sprintf("Error in querying same channel and chaincode with query: %s. Error message: %s.", string(cross_channel_query_string), string(response.Message)))

					return shim.Error(fmt.Sprintf("Error in querying same channel and chaincode with query: %s. Error message: %s.", string(cross_channel_query_string), string(response.Message)))

					//successstring = successstring + fmt.Sprintf("Error in querying same channel and chaincode with query: %s .", string(cross_channel_query_string))

				} else if response.Payload == nil {

				} else {

					fmt.Println("Found Duplicate Invoice %s in same channel and chaincode, for this invoice %s. ", string(response.Payload), invoke_update_primary_string)

					successstring = successstring + fmt.Sprintf("Found Duplicate Invoice %s in same channel and chaincode, for this invoice %s. ", string(response.Payload), invoke_update_primary_string)

					isduplicate = true

				}

			}

			if !isduplicate {

				if len(nested_map[cross_channel_query["invoice-number"]]) == 0 {

					nested_map[cross_channel_query["invoice-number"]] = make(map[string](map[string]string))

					nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]] = make(map[string]string)

					nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]][customer_code_string] = vendor_code_string

					fmt.Println("prev it 1")

				} else if len(nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]]) == 0 {

					nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]] = make(map[string]string)

					nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]][customer_code_string] = vendor_code_string

					fmt.Println("prev it 2")

				} else if _, exists_key := nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]][customer_code_string]; !exists_key {

					nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]][customer_code_string] = vendor_code_string

					fmt.Println("prev it 3")

				} else {
					var prev_vendor_code_string = nested_map[cross_channel_query["invoice-number"]][cross_channel_query["two-digit-invoice-financial-year"]][customer_code_string]

					fmt.Println(fmt.Sprintf("Duplicate Invoice processed earlier in this iteration with vendor code %s, for this invoice %s. ", prev_vendor_code_string, invoke_update_primary_string))

					successstring = successstring + fmt.Sprintf("Duplicate Invoice processed earlier in this iteration with vendor code %s, for this invoice %s. ", prev_vendor_code_string, invoke_update_primary_string)

					isduplicate = true

				}

			}

			if proceessed_records_counter > PROCESSING_LIMIT {

				fmt.Println(fmt.Sprintf("Exceeded %d", proceessed_records_counter))

				exceeded_processing_limit = true

				break
			}

			var response_update_status pb.Response

			if isduplicate {

				fmt.Println("Doing duplicate status = true update")

				response_update_status = t.invoke_insert_update(stub, args_for_duplicate_found_update, map_specification_invoice_duplicate_failed)
			} else {
				fmt.Println("Doing duplicate status = passed update")

				response_update_status = t.invoke_insert_update(stub, args_for_duplicate_not_found_update, map_specification_invoice_duplicate_passed)
			}

			if response_update_status.Status != shim.OK {

				fmt.Println(fmt.Sprintf("Error in updating invoice status for vendor code %s, customer code %s, invoice number %s, invoice date %s. ", vendor_code_string, customer_code_string, cross_channel_query["invoice-number"], fullrecord_invoice_date) + string(response_update_status.Message))

				return shim.Error(fmt.Sprintf("Error in updating invoice status for vendor code %s, customer code %s, invoice number %s, invoice date %s. ", vendor_code_string, customer_code_string, cross_channel_query["invoice-number"], fullrecord_invoice_date) + string(response_update_status.Message))
			}

			proceessed_records_counter++

			fmt.Println(fmt.Sprintf("incremented to %d", proceessed_records_counter))

		}

		fmt.Println("closing iterator")

		tableRowsIterator.Close()

	}

	fmt.Println(successstring)

	if exceeded_processing_limit {

		return shim.Success([]byte("1"))
	} else {

		return shim.Success([]byte("0"))
	}

}
