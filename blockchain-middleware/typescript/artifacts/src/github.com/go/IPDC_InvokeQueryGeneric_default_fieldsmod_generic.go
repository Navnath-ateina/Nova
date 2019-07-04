package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// =================================================================================================
// query_primary_key - Query a record by Primary Key
// =================================================================================================
func (t *IPDCChaincode) query_primary_key(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering query_primary_key***********")

	if len(args) < 1 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	var record_specification map[string]interface{}

	err := json.Unmarshal([]byte(args[0]), &record_specification)
	if err != nil {
		err = json.Unmarshal([]byte(args[1]), &record_specification)
	}
	if err != nil {

		fmt.Println("Error in decoding Input.")

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error("Error in decoding Input.")
	}

	additional_json, ok := map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {

				record_specification[spec] = additional_json_data[spec]
			}

		} else {
			fmt.Println("Error: Invalid additional JSON fields in specification.")

			fmt.Println("***********Exiting query_primary_key***********")

			return shim.Error("Error: Invalid additional JSON fields in specification.")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Error: Invalid keys map specification.")

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error("Error: Invalid keys map specification.")
	}

	if specs["primary_key"] == nil {

		fmt.Println("Error: There is no primary key specification.")

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error("Error : There is no primary key specification.")
	}

	var pk_spec []interface{}

	pk_spec, ok = specs["primary_key"].([]interface{})

	if !ok {

		fmt.Println("Error in Primary key specification.")

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error("Error in Primary key specification.")
	}

	primary_key, err_key := t.createInterfacePrimaryKey(record_specification, pk_spec)

	if err_key != nil {

		fmt.Println(err_key.Error())

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error(err_key.Error())

	}

	var valAsBytes []byte

	valAsBytes, err = stub.GetState(primary_key)

	if err != nil {

		fmt.Println("Failed to get state: " + err.Error())

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Error("Failed to get state: " + err.Error())

	} else {

		fmt.Println("Value: " + string(valAsBytes))

		fmt.Println("***********Exiting query_primary_key***********")

		return shim.Success(valAsBytes)
	}

}

// =================================================================================================
// query_primary_key_history - Query the history of a Primary Key
// =================================================================================================
func (t *IPDCChaincode) query_primary_key_history(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering query_primary_key_history***********")

	var pageno int

	var err error

	if len(args) < 1 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting query_primary_key_history***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	if len(args) < 2 {

		pageno = 1

	} else {

		pageno, err = strconv.Atoi(args[1])

		if err != nil {

			fmt.Println("Error parsing page number " + err.Error())

			fmt.Println("***********Exiting query_primary_key_history***********")

			return shim.Error("Error parsing page number " + err.Error())

		} else if pageno < 1 {

			fmt.Println("Invalid page number")

			fmt.Println("***********Exiting query_primary_key_history***********")

			return shim.Error("Invalid page number")
		}

	}

	var record_specification map[string]interface{}

	err = json.Unmarshal([]byte(args[0]), &record_specification)

	if err != nil {

		fmt.Println("Error in format of record.")

		fmt.Println("***********Exiting query_primary_key_history***********")

		return shim.Error("Error in format of record.")
	}

	additional_json, ok := map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {

				record_specification[spec] = additional_json_data[spec]
			}
		} else {
			fmt.Println("Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting query_primary_key_history***********")

			return shim.Error("Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting query_primary_key_history***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Invalid keys map specification.")

		fmt.Println("***********Exiting query_primary_key_history***********")

		return shim.Error("Invalid keys map specification.")
	}

	if specs["primary_key"] == nil {

		fmt.Println("There is no primary key specification.")

		fmt.Println("***********Exiting query_primary_key_history***********")

		return shim.Error("Error : There is no primary key specification.")
	}

	var pk_spec []interface{}

	pk_spec, ok = specs["primary_key"].([]interface{})

	if !ok {

		fmt.Println("Error in Primary key specification.")

		fmt.Println("***********Exiting query_primary_key_history***********")

		return shim.Error("Error in Primary key specification.")
	}

	primary_key, err_key := t.createInterfacePrimaryKey(record_specification, pk_spec)

	if err_key != nil {

		fmt.Println(err_key.Error())

		fmt.Println("***********Exiting query_primary_key_history***********")

		return shim.Error(err_key.Error())

	}

	fmt.Println("***********Exiting query_primary_key_history***********")

	return t.query_for_record_history(stub, primary_key, pageno)

}

// =================================================================================================
// query_update_status - Query a record for composite Key for checking update status
// =================================================================================================
func (t *IPDCChaincode) query_update_status(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering query_update_status***********")

	//var arguments []string

	var pageno int

	if len(args) < 1 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting query_update_status***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	var record_specification map[string]interface{}

	err := json.Unmarshal([]byte(args[0]), &record_specification)

	if err != nil {

		fmt.Println("Error in format of record")

		fmt.Println("***********Exiting query_update_status***********")

		return shim.Error("Error in format of record")
	}

	if len(record_specification) != 1 {

		fmt.Println("Input should contain only one status")

		fmt.Println("***********Exiting query_update_status***********")

		return shim.Error("Input should contain only one status")
	}

	if len(args) < 2 {

		pageno = 1

	} else {

		pageno, err = strconv.Atoi(args[1])

		if err != nil {

			fmt.Println("Error parsing page number " + err.Error())

			fmt.Println("***********Exiting query_update_status***********")

			return shim.Error("Error parsing page number " + err.Error())

		} else if pageno < 1 {

			fmt.Println("Invalid page number")

			fmt.Println("***********Exiting query_update_status***********")

			return shim.Error("Invalid page number")
		}

	}

	additional_json, ok := map_specification["additional_json"]

	var additional_json_data map[string]interface{}

	if ok {

		var ok1 bool

		additional_json_data, ok1 = additional_json.(map[string]interface{})

		if !ok1 {

			fmt.Println("Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting query_update_status***********")

			return shim.Error("Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, additional_json_data)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting query_update_status***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Invalid keys map specification.")

		fmt.Println("***********Exiting query_update_status***********")

		return shim.Error("Invalid keys map specification.")
	}

	var composite_key = make(map[string]interface{})

	for spec, _ := range record_specification {

		composite_key[spec], ok = specs[spec]

		if !ok {
			fmt.Println("Composite key specification missing")

			fmt.Println("***********Exiting query_update_status***********")

			return shim.Error("Composite key specification missing")
		}
	}

	for spec, _ := range additional_json_data {

		record_specification[spec] = additional_json_data[spec]
	}

	//compositekeyJsonString, err_marshal := json.Marshal(composite_key)

	//if (err_marshal != nil) {

	//	fmt.Println("Error in marshaling composite key")

	//	return shim.Error("Error in marshaling composite key")
	//}

	//var concatenated_record_json []byte

	//concatenated_record_json, err_marshal = json.Marshal(record_specification)

	//if(err_marshal != nil) {

	//	fmt.Println("Unable to Marshal Concatenated Record to JSON")

	//	return shim.Error("Unable to Marshal Concatenated Record to JSON")
	//}

	//arguments = append(arguments, string(concatenated_record_json))

	//arguments = append(arguments, string(compositekeyJsonString))

	//return t.query_by_composite_key(stub, arguments, pageno)

	fmt.Println("***********Exiting query_update_status***********")

	return t.query_by_composite_key(stub, record_specification, composite_key, pageno)

}

// =================================================================================================
// query_records_composite_key - Query records using a (partial) composite key named by first argument
// =================================================================================================
func (t *IPDCChaincode) query_records_composite_key(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering query_records_composite_key***********")

	//var arguments []string

	var pageno int

	if len(args) < 1 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting query_records_composite_key***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	//var composite_key_name = args[0]

	var composite_key_name_interface interface{}

	composite_key_name_interface = map_specification["composite_key_name"]

	composite_key_name, ok_composite_key_name := composite_key_name_interface.(string)

	if !ok_composite_key_name {
		fmt.Println("Error: Invalid composite key specification")

		fmt.Println("***********Exiting query_records_composite_key***********")

		return shim.Error("Error: Invalid composite key specification")
	}

	var record_specification map[string]interface{}

	err := json.Unmarshal([]byte(args[0]), &record_specification)

	if err != nil {

		fmt.Println("Error in format of record")

		fmt.Println("***********Exiting query_records_composite_key***********")

		return shim.Error("Error in format of record")
	}

	var check int

	check, err = t.Mandatoryfieldscheck(record_specification, map_specification)

	if check == 1 {

		fmt.Println(err.Error())

		fmt.Println("***********Exiting query_records_composite_key***********")

		return shim.Error(err.Error())
	}

	if len(args) < 2 {

		pageno = 1

	} else {

		pageno, err = strconv.Atoi(args[1])

		if err != nil {

			fmt.Println("Error parsing page number " + err.Error())

			fmt.Println("***********Exiting query_records_composite_key***********")

			return shim.Error("Error parsing page number " + err.Error())

		} else if pageno < 1 {

			fmt.Println("Invalid page number")

			fmt.Println("***********Exiting query_records_composite_key***********")

			return shim.Error("Invalid page number")
		}

	}

	additional_json, ok := map_specification["additional_json"]

	var additional_json_data map[string]interface{}

	if ok {

		var ok1 bool

		additional_json_data, ok1 = additional_json.(map[string]interface{})

		if !ok1 {

			fmt.Println("Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting query_records_composite_key***********")

			return shim.Error("Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, additional_json_data)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting query_records_composite_key***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Invalid keys map specification.")

		fmt.Println("***********Exiting query_records_composite_key***********")

		return shim.Error("Invalid keys map specification.")
	}

	var composite_key = make(map[string]interface{})

	/*for spec, _ := range record_specification {

		composite_key[spec], ok = specs[spec]

		if !ok {
			fmt.Println("Composite key specification missing")

			fmt.Println("***********Exiting query_update_status***********")

			return shim.Error("Composite key specification missing")
		}
	}*/

	composite_key[composite_key_name], ok = specs[composite_key_name]

	if !ok {
		fmt.Println("Error: composite key does not exist")

		fmt.Println("***********Exiting query_records_composite_key***********")

		return shim.Error("Error: composite key does not exist")
	}

	for spec, _ := range additional_json_data {

		record_specification[spec] = additional_json_data[spec]
	}

	//compositekeyJsonString, err_marshal := json.Marshal(composite_key)

	//if (err_marshal != nil) {

	//	fmt.Println("Error in marshaling composite key")

	//	return shim.Error("Error in marshaling composite key")
	//}

	//var concatenated_record_json []byte

	//concatenated_record_json, err_marshal = json.Marshal(record_specification)

	//if(err_marshal != nil) {

	//	fmt.Println("Unable to Marshal Concatenated Record to JSON")

	//	return shim.Error("Unable to Marshal Concatenated Record to JSON")
	//}

	//arguments = append(arguments, string(concatenated_record_json))

	//arguments = append(arguments, string(compositekeyJsonString))

	//return t.query_by_composite_key(stub, arguments, pageno)

	fmt.Println("***********Exiting query_records_composite_key***********")

	return t.query_by_composite_key(stub, record_specification, composite_key, pageno)

}

// =================================================================================================
// query_using_rich_query - Query records using a (partial) composite key named by first argument
// =================================================================================================
func (t *IPDCChaincode) query_using_rich_query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("***********Entering query_using_rich_query***********")
	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	docType := args[0]
	key := args[1]
	value := args[2]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":%s,%s:%s}}", docType, key, value)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =================================================================================================
// query_all_rich_query - Query records using a (partial) composite key named by first argument
// =================================================================================================
func (t *IPDCChaincode) query_all_rich_query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("***********Entering query_all_rich_query***********")
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	docType := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":%s}}", docType)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =================================================================================================
// query_by_id_and_status - Query records using a (partial) composite key named by first argument
// =================================================================================================
func (t *IPDCChaincode) query_by_id_and_status(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("***********Entering query_by_id_and_status***********")
	if len(args) < 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	docType := args[0]
	key := args[1]
	value := args[2]
	statusVal := args[3]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":%s,%s:%s,\"status\":%s}}", docType, key, value, statusVal)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =================================================================================================
// query_by_user - Query records using a (partial) composite key named by first argument
// =================================================================================================
func (t *IPDCChaincode) query_by_user(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("***********Entering query_by_user***********")
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryFrm := args[0]
	// key := args[1]
	// value := args[2]
	// statusVal := args[3]

	queryString := fmt.Sprintf(queryFrm)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// ===========================================================================================
// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
// ===========================================================================================
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		// buffer.WriteString("{\"Key\":")
		// buffer.WriteString("\"")
		// buffer.WriteString(queryResponse.Key)
		// buffer.WriteString("\"")

		// buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		// buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

// =================================================================================================
// invoke_update_status - update a status of a record using INVOKE
// =================================================================================================
func (t *IPDCChaincode) invoke_update_status(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering invoke_update_status***********")

	if len(args) < 2 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	var record_specification map[string]interface{}

	var err error

	err = json.Unmarshal([]byte(args[0]), &record_specification)

	if err != nil {

		fmt.Println("Error in format of record.")

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error in format of record.")
	}

	additional_json, ok := map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {

				record_specification[spec] = additional_json_data[spec]
			}
		} else {
			fmt.Println("Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting invoke_update_status***********")

			return shim.Error("Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Invalid keys_map specification.")

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Invalid keys_map specification.")
	}

	if specs["primary_key"] == nil {

		fmt.Println("There is no primary key specification.")

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error : There is no primary key specification.")
	}

	var pk_spec []interface{}

	pk_spec, ok = specs["primary_key"].([]interface{})

	if !ok {

		fmt.Println("Error in Primary key specification.")

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error in Primary key specification.")
	}

	key, err_key := t.createInterfacePrimaryKey(record_specification, pk_spec)

	if err_key != nil {

		fmt.Println(err_key.Error())

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error(err_key.Error())

	}

	var valAsBytes []byte

	valAsBytes, err = stub.GetState(key)

	if err != nil {

		fmt.Println("Error: Failed to get state for primary key. " + err.Error())

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error: Failed to get state for primary key. " + err.Error())

	} else if valAsBytes == nil {

		fmt.Println("Error: No value for key : " + key)

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error: No value for primary key.")

	}

	err = json.Unmarshal([]byte(valAsBytes), &record_specification)

	if err != nil {

		fmt.Println("Error in format of Blockchain record")

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error in format of Blockchain record")

	}

	err_del := t.delete_composite_keys(stub, specs, record_specification, key)

	if err_del != nil {

		fmt.Println("Error while deleting composite keys: " + err_del.Error())

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error while deleting composite keys: " + err_del.Error())

	}

	var to_be_updated_map map[string]interface{}

	err = json.Unmarshal([]byte(args[1]), &to_be_updated_map)

	if err != nil {

		fmt.Println("Error in format of update map")

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error in format of update map")

	}

	for spec, spec_val := range to_be_updated_map {

		var spec_val_string, spec_ok = spec_val.(string)

		if !spec_ok {

			fmt.Println("Unable to parse value of status update")

			fmt.Println("***********Exiting invoke_update_status***********")

			return shim.Error("Unable to parse value of status update")

		}

		var val_check, val_err = t.updatestatusvaliditycheck(spec, spec_val_string, map_specification)

		if val_check != 0 {

			fmt.Println(val_err.Error())

			fmt.Println("***********Exiting invoke_update_status***********")

			return shim.Error(val_err.Error())
		}

		record_specification[spec] = spec_val_string
	}

	var concatenated_record_json []byte

	concatenated_record_json, err = json.Marshal(record_specification)

	if err != nil {

		fmt.Println("Error: Unable to Marshal Concatenated Record to JSON " + err.Error())

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Error: Unable to Marshal Concatenated Record to JSON " + err.Error())
	}

	err = stub.PutState(key, []byte(concatenated_record_json))

	if err != nil {

		fmt.Println("Failed to put state : " + err.Error())

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Failed to put state : " + err.Error())
	}

	err = t.create_composite_keys(stub, specs, record_specification, key)

	if err != nil {

		fmt.Println("Received error while creating composite keys" + err.Error())

		fmt.Println("***********Exiting invoke_update_status***********")

		return shim.Error("Received error while creating composite keys" + err.Error())
	}

	fmt.Println("***********Exiting invoke_update_status***********")

	return shim.Success(nil)

}

// =====================================================================================================
// invoke_insert_update - insert a record using INVOKE if not inserted before, else update the record
// ======================================================================================================
func (t *IPDCChaincode) invoke_insert_update(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering invoke_insert_update***********")

	var success string

	if len(args) < 1 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	success = ""

	var record_specification_input map[string]interface{}

	var record_specification = make(map[string]interface{})

	var err error

	err = json.Unmarshal([]byte(args[0]), &record_specification_input)
	// fmt.Println("************************************************")
	// fmt.Println(args[0])
	// fmt.Println(record_specification_input)
	// fmt.Println(err)
	// fmt.Println("************************************************")

	if err != nil {
		err = json.Unmarshal([]byte(args[1]), &record_specification_input)
		fmt.Println("------------------------------------------------")
		fmt.Println(args[0])
		fmt.Println(record_specification_input)
		fmt.Println(err)
		fmt.Println("------------------------------------------------")
		if err != nil {

			fmt.Println("Error in reading input record")

			fmt.Println("***********Exiting invoke_insert_update***********")

			return shim.Error("Error in reading input record")
		}
	}

	err = t.StringValidation(record_specification_input, map_specification)

	if err != nil {

		fmt.Println(err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error(err.Error())
	}

	var check int

	check, err = t.Mandatoryfieldscheck(record_specification_input, map_specification)

	if check == 1 {

		fmt.Println(err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error(err.Error())
	}

	if err != nil {

		success = err.Error()
	}

	check, err = t.Datefieldscheck(record_specification_input, map_specification)

	if check == 1 {

		fmt.Println(err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error(err.Error())
	}

	check, err = t.Amountfieldscheck(record_specification_input, map_specification)

	if check == 1 {

		fmt.Println(err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error(err.Error())
	}

	record_specification, err = t.Mapinputfieldstotarget(record_specification_input, map_specification)

	if err != nil {

		fmt.Println("Error in decoding and/or mapping record: " + err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Error in decoding and/or mapping record: " + err.Error())
	}

	additional_json, ok := map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {
				fmt.Println(spec)
				record_specification[spec] = additional_json_data[spec]
			}
		} else {
			fmt.Println("Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting invoke_insert_update***********")

			return shim.Error("Invalid additional JSON fields in specification")
		}
	}

	fmt.Println("&&&&&&& New Record &&&&&&&& ", record_specification)

	var default_fields interface{}

	var default_fields_data map[string]interface{}

	var ok_default bool

	default_fields, ok = map_specification["default_fields"]

	if ok {

		default_fields_data, ok_default = default_fields.(map[string]interface{})

		if !ok_default {

			default_fields_data = make(map[string]interface{})
		}
	}

	fmt.Println("&&&&&&& Default Fields &&&&&&&& ", default_fields_data)

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Invalid keys_map specification.")

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Invalid keys_map specification.")
	}

	if specs["primary_key"] == nil {

		fmt.Println("Error: There is no primary key specification.")

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Error : There is no primary key specification.")
	}

	var pk_spec []interface{}

	pk_spec, ok = specs["primary_key"].([]interface{})

	if !ok {

		fmt.Println("Error in Primary key specification.")

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Error in Primary key specification.")
	}

	key, err_key := t.createInterfacePrimaryKey(record_specification, pk_spec)

	if err_key != nil {

		fmt.Println(err_key.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error(err_key.Error())

	}

	var valAsBytes []byte

	valAsBytes, err = stub.GetState(key)

	if err != nil {

		fmt.Println("Error Failed to get state for primary key: " + err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Failed to get state for primary key: " + err.Error())

	}

	var record_specification_old = make(map[string]interface{})

	if valAsBytes != nil {

		fmt.Println("Record is already present in blockchain for key " + key)

		err = json.Unmarshal([]byte(valAsBytes), &record_specification_old)

		if err != nil {

			fmt.Println("Error in format of blockchain record")

			fmt.Println("***********Exiting invoke_insert_update***********")

			return shim.Error("Error in format of blockchain record")
		}

		err_del := t.delete_composite_keys(stub, specs, record_specification_old, key)

		if err_del != nil {

			fmt.Println("Error in deleting composite keys" + err_del.Error())

			fmt.Println("***********Exiting invoke_insert_update***********")

			return shim.Error("Error in deleting composite keys" + err_del.Error())
		}

		fmt.Println("&&&&&&& Old Record Record &&&&&&&& ", record_specification_old)

		fmt.Println("&&&&&&& New Record Record again &&&&&&&& ", record_specification)

		for spec, _ := range record_specification {

			//if default_fields_data[spec]==nil {

			record_specification_old[spec] = record_specification[spec] // Add all the new record fields to the older record
			//}
		}

		for spec, _ := range default_fields_data {

			if record_specification_old[spec] == nil {

				record_specification_old[spec] = default_fields_data[spec]
			}
		}

		fmt.Println("&&&&&&& Updated Old Record Record &&&&&&&& ", record_specification_old)

		status, err_validation := t.validation_checks(stub, map_specification, record_specification_old)

		fmt.Println("Updated ---- Status is " + strconv.Itoa(status))

		//fmt.Println("Updated ---- err_validation is " + err_validation.Error())

		if status == -2 {

			fmt.Println("Error: Exiting due to Validation Config failure: " + err_validation.Error())

			fmt.Println("***********Exiting invoke_insert_update***********")

			return shim.Error("Error: Exiting due to Validation Config failure: " + err_validation.Error())

		}

		if err_validation != nil {

			if status == -1 {

				success = success + " " + err_validation.Error()
			}
		}

	} else {

		for spec, _ := range default_fields_data {

			record_specification_old[spec] = default_fields_data[spec]
		}

		for spec, _ := range record_specification {

			record_specification_old[spec] = record_specification[spec] // Add all the new record fields to the older record
		}

		status, err_validation := t.validation_checks(stub, map_specification, record_specification_old)

		//fmt.Println("Status is " + strconv.Itoa(status))

		//if err_validation!=nil {

		//	fmt.Println("err_validation is " + err_validation.Error())

		//	if((status == -1)||(status == -2)) {
		//
		//		success = success + " " + err_validation.Error()
		//	}

		//}

		fmt.Println("Updated ---- Status is " + strconv.Itoa(status))

		//fmt.Println("Updated ---- err_validation is " + err_validation.Error())

		if status == -2 {

			fmt.Println("Error: Exiting due to Validation Config failure: " + err_validation.Error())

			fmt.Println("***********Exiting invoke_insert_update***********")

			return shim.Error("Error: Exiting due to Validation Config failure: " + err_validation.Error())

		}

		if err_validation != nil {

			if status == -1 {

				success = success + " " + err_validation.Error()
			}
		}

	}

	var concatenated_record_json []byte

	fmt.Println("&&&&&&& Updated Old Record Record again &&&&&&&& ", record_specification_old)

	concatenated_record_json, err = json.Marshal(record_specification_old)

	if err != nil {

		fmt.Println("Error: Unable to Marshal Concatenated Record to JSON")

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Error: Unable to Marshal Concatenated Record to JSON")
	}

	fmt.Println("&&&&&&& Updated Old Record JSON &&&&&&&& " + string(concatenated_record_json))

	fmt.Println("&&&&&&& Key for Put &&&&&&&& " + key)

	err = stub.PutState(key, []byte(concatenated_record_json))

	if err != nil {

		fmt.Println("Error: Failed to put state: " + err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Error: Failed to put state: " + err.Error())
	}

	err = t.create_composite_keys(stub, specs, record_specification_old, key)

	if err != nil {

		fmt.Println("Error in creating composite keys: " + err.Error())

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Error("Error in creating composite keys: " + err.Error())
	}

	if success != "" {

		fmt.Println("Warnings! " + success)

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Success([]byte("Warnings! " + success))

	} else {

		fmt.Println("***********Exiting invoke_insert_update***********")

		return shim.Success(nil)
	}

}

// =================================================================================================
// invoke_update_status_with_modification_check - update a status of a record using INVOKE only if input record matches the ledger record in config specified fields
// =================================================================================================
func (t *IPDCChaincode) invoke_update_status_with_modification_check(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering invoke_update_status_with_modification_check***********")

	if len(args) < 2 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	var record_specification_input map[string]interface{}

	var err error

	err = json.Unmarshal([]byte(args[0]), &record_specification_input)

	if err != nil {

		fmt.Println("Error in format of record.")

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error in format of record.")
	}

	additional_json, ok := map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {

				record_specification_input[spec] = additional_json_data[spec]
			}
		} else {
			fmt.Println("Error: Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

			return shim.Error("Error: Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification_input)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Error: Invalid keys_map specification.")

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error: Invalid keys_map specification.")
	}

	if specs["primary_key"] == nil {

		fmt.Println("Error: There is no primary key specification.")

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error : There is no primary key specification.")
	}

	var pk_spec []interface{}

	pk_spec, ok = specs["primary_key"].([]interface{})

	if !ok {

		fmt.Println("Error in Primary key specification.")

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error in Primary key specification.")
	}

	key, err_key := t.createInterfacePrimaryKey(record_specification_input, pk_spec)

	if err_key != nil {

		fmt.Println(err_key.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error(err_key.Error())

	}

	var valAsBytes []byte

	valAsBytes, err = stub.GetState(key)

	if err != nil {

		fmt.Println("Error: Failed to get state: " + err.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error: Failed to get state: " + err.Error())

	} else if valAsBytes == nil {

		fmt.Println("Error: No value for primary key : " + key)

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error: No value for key")

	}

	var record_specification map[string]interface{}

	err = json.Unmarshal([]byte(valAsBytes), &record_specification)

	if err != nil {

		fmt.Println("Error in format of record")

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error in format of record")

	}

	var check int

	check, err = t.Isfieldsmodified(record_specification_input, record_specification, map_specification)

	if check != 0 {

		fmt.Println("Status Update Failed due to error in modification check. " + err.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Status Update Failed due to error in modification check. " + err.Error())
	}

	err_del := t.delete_composite_keys(stub, specs, record_specification, key)

	if err_del != nil {

		fmt.Println("Error in deleting composite keys" + err_del.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error in deleting composite keys" + err_del.Error())

	}

	var to_be_updated_map map[string]interface{}

	err = json.Unmarshal([]byte(args[1]), &to_be_updated_map)

	if err != nil {

		fmt.Println("Error in format of update map.")

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error in format of update map.")

	}

	for spec, spec_val := range to_be_updated_map {

		var spec_val_string, spec_ok = spec_val.(string)

		if !spec_ok {

			fmt.Println("Error: Unable to parse value of status update")

			fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

			return shim.Error("Error: Unable to parse value of status update")

		}

		var val_check, val_err = t.updatestatusvaliditycheck(spec, spec_val_string, map_specification)

		if val_check != 0 {

			fmt.Println(val_err.Error())

			fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

			return shim.Error(val_err.Error())
		}

		record_specification[spec] = spec_val_string
	}

	var concatenated_record_json []byte

	concatenated_record_json, err = json.Marshal(record_specification)

	if err != nil {

		fmt.Println("Error: Unable to Marshal Concatenated Record to JSON " + err.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error: Unable to Marshal Concatenated Record to JSON " + err.Error())
	}

	err = stub.PutState(key, []byte(concatenated_record_json))

	if err != nil {

		fmt.Println("Error: Failed to put state : " + err.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error: Failed to put state : " + err.Error())
	}

	err = t.create_composite_keys(stub, specs, record_specification, key)

	if err != nil {

		fmt.Println("Error in creating composite keys" + err.Error())

		fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

		return shim.Error("Error in creating composite keys" + err.Error())
	}

	fmt.Println("***********Exiting invoke_update_status_with_modification_check***********")

	return shim.Success(nil)

}

// =================================================================================================
// invoke_delete_record - delete a record using its primary key
// =================================================================================================
func (t *IPDCChaincode) invoke_delete_record(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering invoke_delete_record***********")

	if len(args) < 1 {

		fmt.Println("Error: Incorrect number of arguments")

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error: Incorrect number of arguments")
	}

	var record_specification map[string]interface{}

	var err error

	err = json.Unmarshal([]byte(args[0]), &record_specification)

	if err != nil {

		fmt.Println("Error in format of input record")

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error in format of input record")
	}

	additional_json, ok := map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {

				record_specification[spec] = additional_json_data[spec]
			}
		} else {
			fmt.Println("Error: Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting invoke_delete_record***********")

			return shim.Error("Error: Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Error: Invalid keys_map specification.")

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error: Invalid keys_map specification.")
	}

	if specs["primary_key"] == nil {

		fmt.Println("Error: invalid primary key specification.")

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error : invalid primary key specification.")
	}

	var pk_spec []interface{}

	pk_spec, ok = specs["primary_key"].([]interface{})

	if !ok {

		fmt.Println("Error in Primary key specification.")

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error in Primary key specification.")
	}

	key, err_key := t.createInterfacePrimaryKey(record_specification, pk_spec)

	if err_key != nil {

		fmt.Println(err_key.Error())

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error(err_key.Error())

	}

	var valAsBytes []byte

	valAsBytes, err = stub.GetState(key)

	if err != nil {

		fmt.Println("Error: Failed to get state. " + err.Error())

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error: Failed to get state. " + err.Error())

	} else if valAsBytes == nil {

		fmt.Println("Error: No value for primary key : " + key)

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Success([]byte("Error: No value for primary key."))

	}

	err = json.Unmarshal([]byte(valAsBytes), &record_specification)

	if err != nil {

		fmt.Println("Error in format of blockchain record.")

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error in format of blockchain record.")

	}

	err_del := t.delete_composite_keys(stub, specs, record_specification, key)

	if err_del != nil {

		fmt.Println("Error in deleting composite keys: " + err_del.Error())

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error in deleting composite keys: " + err_del.Error())

	}

	//Deleting primary key

	err_del = stub.DelState(key)

	if err_del != nil {

		fmt.Println("Error in deleting primary key: " + err_del.Error())

		fmt.Println("***********Exiting invoke_delete_record***********")

		return shim.Error("Error in deleting primary key: " + err_del.Error())

	}

	fmt.Println("***********Exiting invoke_delete_record***********")

	return shim.Success(nil)

}

// =================================================================================================
// invoke_delete_all_records - Bulk invoke to delete a type of record given in map spec
// =================================================================================================
func (t *IPDCChaincode) invoke_delete_all_records(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}) pb.Response {

	fmt.Println("***********Entering invoke_delete_all_records***********")

	var arguments []string

	var ok bool

	var additional_json interface{}

	var record_specification = make(map[string]interface{})

	additional_json, ok = map_specification["additional_json"]

	if ok {

		additional_json_data, ok1 := additional_json.(map[string]interface{})

		if ok1 {

			for spec, _ := range additional_json_data {

				record_specification[spec] = additional_json_data[spec]
			}
		} else {
			fmt.Println("Error: Invalid additional JSON fields in specification")

			fmt.Println("***********Exiting invoke_delete_all_records***********")

			return shim.Error("Error: Invalid additional JSON fields in specification")
		}
	}

	var keys_map interface{}

	var specs map[string]interface{}

	keys_map, error_keys_map := t.get_keys_map(stub, record_specification)

	if error_keys_map != nil {

		fmt.Println(error_keys_map.Error())

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error(error_keys_map.Error())
	}

	specs, ok = keys_map.(map[string]interface{})

	if !ok {

		fmt.Println("Error: Invalid keys_map specification.")

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error("Error: Invalid keys_map specification.")
	}

	var composite_key = make(map[string]interface{})

	//for spec, _ := range record_specification {
	//
	//	composite_key[spec] = specs[spec]
	//}

	composite_key["stagingdb-update-status"], ok = specs["stagingdb-update-status"]

	if !ok {

		fmt.Println("Error: Composite key specification missing for deletion.")

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error("Error: Composite key specification missing for deletion.")
	}

	compositekeyJsonString, err_marshal := json.Marshal(composite_key)

	if err_marshal != nil {

		fmt.Println("Error in marshaling composite key")

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error("Error in marshaling composite key")
	}

	record_specification["stagingdb-update-status"] = "True"

	var concatenated_record_json []byte

	concatenated_record_json, err_marshal = json.Marshal(record_specification)

	if err_marshal != nil {

		fmt.Println("Error: Unable to Marshal Concatenated Record to JSON")

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error("Error: Unable to Marshal Concatenated Record to JSON")
	}

	arguments = append(arguments, string(concatenated_record_json))

	arguments = append(arguments, string(compositekeyJsonString))

	err_delete, processed_records, records_remaining := t.delete_by_composite_key(stub, arguments, specs, PROCESSING_LIMIT)

	if err_delete != nil {

		fmt.Println(err_delete.Error())

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error(err_delete.Error())
	}

	if records_remaining {

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Success([]byte("1"))
	}

	record_specification["stagingdb-update-status"] = "False"

	concatenated_record_json, err_marshal = json.Marshal(record_specification)

	if err_marshal != nil {

		fmt.Println("Error: Unable to Marshal Concatenated Record to JSON")

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error("Error: Unable to Marshal Concatenated Record to JSON")
	}

	arguments = make([]string, 0)

	arguments = append(arguments, string(concatenated_record_json))

	arguments = append(arguments, string(compositekeyJsonString))

	PROCESSING_LIMIT_TEMP := PROCESSING_LIMIT - processed_records

	err_delete, _, records_remaining = t.delete_by_composite_key(stub, arguments, specs, PROCESSING_LIMIT_TEMP)

	if err_delete != nil {

		fmt.Println(err_delete.Error())

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Error(err_delete.Error())
	}

	if !records_remaining {

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Success([]byte("0"))

	} else {

		fmt.Println("***********Exiting invoke_delete_all_records***********")

		return shim.Success([]byte("1"))
	}

}

// =================================================================================================
// invoke_bulk - Bulk invoke to execute a list of individual non-bulk invokes given their payloads. Succeeds of all the individual invokes succeed.
// =================================================================================================
func (t *IPDCChaincode) invoke_bulk(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("***********Entering invoke_bulk***********")

	if len(args) > (PROCESSING_LIMIT + 10) {

		fmt.Println("Error: Too many invoke calls in bulk invoke")

		fmt.Println("***********Exiting invoke_bulk***********")

		return shim.Error("Error: Too many invoke calls in bulk invoke")

	}

	success_string := "["

	for i_number, individual_invoke_args := range args {

		i := fmt.Sprint(i_number)

		var list_args_interface []interface{}

		err_json := json.Unmarshal([]byte(individual_invoke_args), &list_args_interface)

		if err_json != nil {

			fmt.Println("Error: Unable to read the arguments for Invoke no. " + i + err_json.Error())

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: Unable to read the arguments for Invoke no. " + i + err_json.Error())
		}

		if len(list_args_interface) < 1 {

			fmt.Println("Error: empty payload for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: empty payload for Invoke no. " + i)
		}

		var list_args []string

		//list_args = make([]string, len(list_args_interface))

		for _, value_interface := range list_args_interface {

			value_string, ok_value_string := value_interface.(string)

			if !ok_value_string {

				fmt.Println("Error: Invalid format of payload for Invoke no. " + i)

				fmt.Println("***********Exiting invoke_bulk***********")

				return shim.Error("Error: Invalid format of payload for Invoke no. " + i)
			}

			list_args = append(list_args, value_string)
		}

		function_name := list_args[0]

		args_to_pass := list_args[1:]

		key_for_func := "FunctionName*" + function_name

		valAsBytes, err := stub.GetState(key_for_func)

		if err != nil {

			fmt.Println(fmt.Sprintf("Error: Failed to get state: " + err.Error() + " for Invoke no. " + i))

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: Failed to get state: " + err.Error() + " for Invoke no. " + i)

		} else if valAsBytes == nil {

			fmt.Println("Error: No value for key : " + key_for_func + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: No value for key : " + key_for_func + " for Invoke no. " + i)
		}

		var json_specification interface{}

		err = json.Unmarshal(valAsBytes, &json_specification)

		if err != nil {

			fmt.Println("Error in decoding Specification JSON" + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error in decoding Specification JSON" + " for Invoke no. " + i)
		}

		map_specification, ok1 := json_specification.(map[string]interface{})

		if !ok1 {
			fmt.Println("Error Parsing map_specification" + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error Parsing map_specification" + " for Invoke no. " + i)
		}

		operation, ok2 := map_specification["operation"]

		if !ok2 {
			fmt.Println("Error Parsing operation" + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error Parsing operation" + " for Invoke no. " + i)
		}

		primitive_list, ok3 := operation.(map[string]interface{})

		if !ok3 {
			fmt.Println("Error Parsing primitive list" + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error Parsing primitive list" + " for Invoke no. " + i)
		}

		if _, ok3 = primitive_list["primitive"]; !ok3 {

			fmt.Println("Error: no primitive operation" + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: no primitive operation" + " for Invoke no. " + i)
		}

		var primitive_operation string

		primitive_operation, ok3 = primitive_list["primitive"].(string)

		if !ok3 {

			fmt.Println("Error: Invalid primitive operation" + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: Invalid primitive operation" + " for Invoke no. " + i)
		}

		fmt.Println("Primitive operation for  Invoke no. " + i + " : " + primitive_operation)

		var invoke_response pb.Response

		if primitive_operation == "invoke_insert_update" {

			invoke_response = t.invoke_insert_update(stub, args_to_pass, map_specification)

		} else if primitive_operation == "invoke_update_status" {

			invoke_response = t.invoke_update_status(stub, args_to_pass, map_specification)

		} else if primitive_operation == "invoke_update_status_with_modification_check" {

			invoke_response = t.invoke_update_status_with_modification_check(stub, args_to_pass, map_specification)

		} else if primitive_operation == "invoke_delete_record" {

			invoke_response = t.invoke_delete_record(stub, args_to_pass, map_specification)

		} else if primitive_operation == "query_primary_key" || primitive_operation == "query_primary_key_history" || primitive_operation == "query_update_status" || primitive_operation == "query_customer_invoice_disbursed" || primitive_operation == "query_customer_invoice_asn_disbursed" {

			fmt.Println("Error: Query function received as Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: Query function received as Invoke no. " + i)

		} else if primitive_operation == "invoke_delete_all_records" {

			fmt.Println("Error: Delete all invoke call not allowed but received as Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: Delete all invoke call not allowed but received as Invoke no. " + i)

		} else {

			fmt.Println("Error: Invalid function " + function_name + " for Invoke no. " + i)

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error: Invalid function " + function_name + " for Invoke no. " + i)
		}

		if invoke_response.Status != shim.OK {

			fmt.Println("Error in executing Invoke no. " + i + " : " + string(invoke_response.Message))

			fmt.Println("***********Exiting invoke_bulk***********")

			return shim.Error("Error in executing Invoke no. " + i + " : " + string(invoke_response.Message))

		} else if i_number == 0 {

			success_string = success_string + string(invoke_response.Message)
		} else {

			success_string = success_string + "," + string(invoke_response.Message)
		}

		fmt.Println("Response of executing Invoke no. " + i + " : " + string(invoke_response.Message))

	}

	success_string = success_string + "]"

	fmt.Println("***********Exiting invoke_bulk***********")

	return shim.Success([]byte(success_string))

}
