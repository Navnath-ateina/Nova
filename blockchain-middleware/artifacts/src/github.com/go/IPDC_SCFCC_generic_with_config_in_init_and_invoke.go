package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"golang.org/x/exp/utf8string"
)

// IPDCChaincode example simple Chaincode implementation
//type IPDCChaincode struct {
//}

// ===================================================================================
// Main
// ===================================================================================
func main() {

	err := shim.Start(new(IPDCChaincode))

	if err != nil {

		fmt.Printf("Error starting chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *IPDCChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("***********Entering Init***********")

	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Init is running " + function)

	var username, error_username = t.getUserNameusingMsp(stub)

	if error_username != nil {

		fmt.Println("***********Exiting Init***********")

		fmt.Println("Error in parsing username: " + error_username.Error())

		return shim.Error("Error in parsing username: " + error_username.Error())

	} else {

		fmt.Println("username = " + username)
	}

	/*
		if username != "Admin@MMFSLPeerOrg.MMFSL.com" {

			fmt.Println("***********Exiting Invoke***********")

			return shim.Error("Error: username  is " + username + " with insufficient permission.")
		}
	*/

	response := t.cleanConfig(stub)

	if response.Status != shim.OK {

		fmt.Println("***********Exiting Init***********")

		return response
	}

	response = t.parseConfigFile(stub, args)

	fmt.Println("***********Exiting Init***********")

	return response

	return shim.Success(nil)
}

// ========================================
// Invoke - Our entry point for Invocations
// ========================================
func (t *IPDCChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	var username, error_username = t.getUserNameusingMsp(stub)

	if error_username != nil {

		fmt.Println("***********Exiting Init***********")

		fmt.Println("Error in parsing username: " + error_username.Error())

		return shim.Error("Error in parsing username: " + error_username.Error())

	} else {

		fmt.Println("username = " + username)
	}

	fmt.Println("***********Entering Invoke***********")

	function, args := stub.GetFunctionAndParameters()

	fmt.Println("invoke is running " + function)

	fmt.Println("args is ")

	fmt.Println(args)

	if function == "query_configs" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_configs(stub, args)
	}

	if function == "invoke_update_config" {

		var username, error_username = t.getUserNameusingMsp(stub)

		if error_username != nil {

			fmt.Println("***********Exiting Invoke***********")

			fmt.Println("Error in parsing username: " + error_username.Error())

			return shim.Error("Error in parsing username: " + error_username.Error())

		} else {

			fmt.Println("username = " + username)
		}

		/*
			if username != "Admin@MMFSLPeerOrg.MMFSL.com" {

				fmt.Println("***********Exiting Invoke***********")

				return shim.Error("Error: username  is " + username + " with insufficient permission.")
			}
		*/

		response := t.cleanConfig(stub)

		if response.Status != shim.OK {

			fmt.Println("***********Exiting Invoke***********")

			return response
		}

		response = t.parseConfigFileinvoke(stub, args)

		fmt.Println("***********Exiting Invoke***********")

		return response
	}

	if function == "invoke_cross_channel_duplicate_check" {

		if len(args) < 1 {

			fmt.Println("Error: Incorrect number of arguments")

			fmt.Println("***********Exiting Invoke***********")

			return shim.Error("Error: Incorrect number of arguments")
		}

		fmt.Println("***********Exiting Invoke***********")

		fmt.Println("***********Entering invoke_cross_channel_duplicate_check***********")

		to_return := t.invoke_cross_channel_duplicate_check(stub, args)

		fmt.Println("***********Exiting invoke_cross_channel_duplicate_check***********")

		return to_return

	}

	var error_func error

	function, error_func = t.Internalfunctionname(function, args)

	if error_func != nil {

		fmt.Println("Error in deciphering function name")

		fmt.Println("***********Exiting Invoke***********")

		return shim.Error("Error in deciphering function name. " + error_func.Error())
	}

	if function == "invoke_bulk" {

		fmt.Println("***********Exiting Invoke***********")

		return t.invoke_bulk(stub, args)
	}

	key_for_func := "FunctionName*" + function

	valAsBytes, err := stub.GetState(key_for_func)

	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to get state: " + err.Error()))

		fmt.Println("***********Exiting Invoke***********")

		return shim.Error("Failed to get state: " + err.Error())

	} else if valAsBytes == nil {

		fmt.Println("No value for key : " + key_for_func)

		fmt.Println("***********Exiting Invoke***********")

		return shim.Error("No value for key : " + key_for_func)
	}

	var json_specification interface{}

	err = json.Unmarshal(valAsBytes, &json_specification)

	if err != nil {

		fmt.Println("Error in decoding Specification JSON")

		fmt.Println("***********Exiting Invoke***********")

		return shim.Error("JSON format for Specification is not correct")
	}

	map_specification, ok1 := json_specification.(map[string]interface{})

	if !ok1 {
		fmt.Println("Error Parsing map_specification")

		fmt.Println("***********Exiting Invoke***********")

		return shim.Error("Error Parsing map_specification")
	}

	operation, ok2 := map_specification["operation"]

	if !ok2 {
		fmt.Println("Error Parsing operation")

		fmt.Println("***********Exiting Invoke***********")

		return shim.Error("Error Parsing operation")
	}

	primitive_list, ok3 := operation.(map[string]interface{})

	if !ok3 {
		fmt.Println("Error Parsing primitive list")

		fmt.Println("***********Exiting Invoke***********")

		return shim.Error("Error Parsing primitive list")
	}

	fmt.Println("Primitive operation: ")
	fmt.Println(primitive_list["primitive"])

	if primitive_list["primitive"] == "invoke_insert_update" {

		fmt.Println("***********Exiting Invoke***********")

		return t.invoke_insert_update(stub, args, map_specification)

	} else if primitive_list["primitive"] == "invoke_update_status" {

		fmt.Println("***********Exiting Invoke***********")

		return t.invoke_update_status(stub, args, map_specification)

	} else if primitive_list["primitive"] == "invoke_update_status_with_modification_check" {

		fmt.Println("***********Exiting Invoke***********")

		return t.invoke_update_status_with_modification_check(stub, args, map_specification)

	} else if primitive_list["primitive"] == "query_primary_key" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_primary_key(stub, args, map_specification)

	} else if primitive_list["primitive"] == "query_primary_key_history" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_primary_key_history(stub, args, map_specification)

	} else if primitive_list["primitive"] == "query_update_status" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_update_status(stub, args, map_specification)

	} else if primitive_list["primitive"] == "query_customer_invoice_duplicate_passed" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_customer_invoice_duplicate_passed(stub, args, map_specification)

		/*} else if primitive_list["primitive"] == "query_customer_invoice_asn_disbursed" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_customer_invoice_asn_disbursed(stub, args, map_specification)
		*/
	} else if primitive_list["primitive"] == "query_records_composite_key" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_records_composite_key(stub, args, map_specification)

	} else if primitive_list["primitive"] == "query_using_rich_query" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_using_rich_query(stub, args)

	} else if primitive_list["primitive"] == "query_all_rich_query" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_all_rich_query(stub, args)

	} else if primitive_list["primitive"] == "query_by_id_and_status" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_by_id_and_status(stub, args)

	} else if primitive_list["primitive"] == "query_by_user" {

		fmt.Println("***********Exiting Invoke***********")

		return t.query_by_user(stub, args)

	} else if primitive_list["primitive"] == "invoke_delete_record" {

		fmt.Println("***********Exiting Invoke***********")

		return t.invoke_delete_record(stub, args, map_specification)

	} else if primitive_list["primitive"] == "invoke_delete_all_records" {

		fmt.Println("***********Exiting Invoke***********")

		return t.invoke_delete_all_records(stub, args, map_specification)
	}

	fmt.Println("invoke did not find function: " + function)

	fmt.Println("***********Exiting Invoke***********")

	return shim.Error("Received unknown function invocation! " + "invoke did not find function: " + function)
}

// =================================================================================================
// cleanConfig - remove existing configuration
// =================================================================================================
func (t *IPDCChaincode) cleanConfig(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("***********Entering cleanConfig***********")

	//***********Removing the Function definitions ***************

	fmt.Println("******Now Removing the Function definitions*******")

	iterator, err_iterator := stub.GetStateByRange("FunctionName*", "FunctionName*~")

	if err_iterator != nil {

		fmt.Println("Error in the first range query: " + err_iterator.Error())

		fmt.Println("***********Exiting cleanConfig***********")

		return shim.Error("Error in the first range query: " + err_iterator.Error())
	}

	var i int

	for i = 0; iterator.HasNext(); i++ {

		indexKey, err := iterator.Next()

		if err != nil {

			fmt.Println(fmt.Sprintf("Error fetching indexKey for first iterator at i = %d : ", i) + err.Error())

			fmt.Println("***********Exiting cleanConfig***********")

			iterator.Close()

			return shim.Error(fmt.Sprintf("Error fetching indexKey for first iterator at i = %d : ", i) + err.Error())

		}

		key_to_delete := indexKey.Key

		err = stub.DelState(key_to_delete)

		if err != nil {

			fmt.Println(fmt.Sprintf("Error deleting key %s for first iterator at i = %d : ", key_to_delete, i) + err.Error())

			fmt.Println("***********Exiting cleanConfig***********")

			iterator.Close()

			return shim.Error(fmt.Sprintf("Error deleting key %s for first iterator at i = %d : ", key_to_delete, i) + err.Error())

		}
	}

	iterator.Close()

	fmt.Println("******Completed Removing the Function definitions*******")

	fmt.Println("******Now Removing the record type keys definitions*******")

	iterator, err_iterator = stub.GetStateByRange("RecordType*", "RecordType*~")

	if err_iterator != nil {

		fmt.Println("Error in the second range query: " + err_iterator.Error())

		fmt.Println("***********Exiting cleanConfig***********")

		return shim.Error("Error in the second range query: " + err_iterator.Error())
	}

	for i = 0; iterator.HasNext(); i++ {

		indexKey, err := iterator.Next()

		if err != nil {

			fmt.Println(fmt.Sprintf("Error fetching indexKey for second iterator at i = %d : ", i) + err.Error())

			fmt.Println("***********Exiting cleanConfig***********")

			iterator.Close()

			return shim.Error(fmt.Sprintf("Error fetching indexKey for second iterator at i = %d : ", i) + err.Error())

		}

		key_to_delete := indexKey.Key

		err = stub.DelState(key_to_delete)

		if err != nil {

			fmt.Println(fmt.Sprintf("Error deleting key %s for second iterator at i = %d : ", key_to_delete, i) + err.Error())

			fmt.Println("***********Exiting cleanConfig***********")

			iterator.Close()

			return shim.Error(fmt.Sprintf("Error deleting key %s for second iterator at i = %d : ", key_to_delete, i) + err.Error())

		}
	}

	iterator.Close()

	fmt.Println("******Completed Removing the record type keys definitions*******")

	fmt.Println("***********Exiting cleanConfig***********")

	return shim.Success(nil)

}

// =================================================================================================
// parseConfigFile - called in the init function --- parse the config file and store data in ledger
// =================================================================================================
func (t *IPDCChaincode) parseConfigFile(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("***********Entering parseConfigFile***********")

	fmt.Println("Parsing config_json_bytes")

	fmt.Println("Checking if config_json_bytes is valid string and has no * ` ~ \\ ? < > special characters")

	if !utf8string.NewString(string(config_json_bytes)).IsASCII() {

		fmt.Println(fmt.Sprintf("Error: Invalid config_json_bytes: not ASCII"))

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error(fmt.Sprintf("Error: Invalid config_json_bytes: not ASCII"))
	}

	if strings.ContainsAny(string(config_json_bytes), "~*\\^`?<>") {

		fmt.Println(fmt.Sprintf("Error: Invalid characters in config_json_bytes"))

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error(fmt.Sprintf("Error: Invalid characters in config_json_bytes"))
	}

	configbuffer := new(bytes.Buffer)

	if err_buf := json.Compact(configbuffer, config_json_bytes); err_buf != nil {

		fmt.Println("Error in decoding config_json_bytes: " + err_buf.Error())

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error("Error in decoding config_json_bytes: " + err_buf.Error())
	}

	var json_specification interface{}

	err := json.Unmarshal(configbuffer.Bytes(), &json_specification)

	if err != nil {

		fmt.Println("Error in Unmarshalling configbuffer for config_json_bytes: " + err.Error())

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error("Error in Unmarshalling configbuffer for config_json_bytes: " + err.Error())
	}

	map_specification, ok := json_specification.(map[string]interface{})

	if !ok {
		fmt.Println("Error in decoding configbuffer into map for config_json_bytes")

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error("Error in decoding configbuffer into map for config_json_bytes")
	}

	for key, value := range map_specification {

		value_byteArray, err2 := json.Marshal(value)

		if err2 != nil {

			fmt.Println("Error in marshalling JSON for function " + key)

			fmt.Println("***********Exiting parseConfigFile***********")

			return shim.Error("Error in marshalling JSON for function " + key)

		}

		operation := "FunctionName*" + key

		err2 = stub.PutState(operation, value_byteArray)

		if err2 != nil {

			fmt.Println("Error in PuState for function " + key)

			fmt.Println("***********Exiting parseConfigFile***********")

			return shim.Error("Error in PuState for function " + key)
		}

	}

	fmt.Println("config_json_bytes is in Ledger!!!")

	fmt.Println("Parsing record_types_to_keys_map")

	fmt.Println("Checking if record_types_to_keys_map is valid string and has no * ` ~ \\ ? < > special characters")

	if !utf8string.NewString(string(record_types_to_keys_map)).IsASCII() {

		fmt.Println(fmt.Sprintf("Error: Invalid record_types_to_keys_map: not ASCII"))

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error(fmt.Sprintf("Error: Invalid record_types_to_keys_map: not ASCII"))
	}

	if strings.ContainsAny(string(record_types_to_keys_map), "~*\\^`?<>") {

		fmt.Println(fmt.Sprintf("Error: Invalid characters in record_types_to_keys_map"))

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error(fmt.Sprintf("Error: Invalid characters in record_types_to_keys_map"))
	}

	configbuffer = new(bytes.Buffer)

	if err_buf := json.Compact(configbuffer, record_types_to_keys_map); err_buf != nil {

		fmt.Println("Error in decoding record_types_to_keys_map: " + err_buf.Error())

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error("Error in decoding record_types_to_keys_map: " + err_buf.Error())
	}

	err = json.Unmarshal(configbuffer.Bytes(), &json_specification)

	if err != nil {

		fmt.Println("Error in Unmarshalling configbuffer for record_types_to_keys_map: " + err.Error())

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error("Error in Unmarshalling configbuffer for record_types_to_keys_map: " + err.Error())
	}

	map_specification, ok = json_specification.(map[string]interface{})

	if !ok {
		fmt.Println("Error in decoding configbuffer into map for record_types_to_keys_map")

		fmt.Println("***********Exiting parseConfigFile***********")

		return shim.Error("Error in decoding configbuffer into map for record_types_to_keys_map")
	}

	for key, value := range map_specification {

		value_byteArray, err2 := json.Marshal(value)

		if err2 != nil {

			fmt.Println("Error in marshalling keys JSON for record type " + key)

			fmt.Println("***********Exiting parseConfigFile***********")

			return shim.Error("Error in marshalling  keys JSON for record type " + key)

		}

		record_type_key := "RecordType*" + key

		err2 = stub.PutState(record_type_key, value_byteArray)

		if err2 != nil {

			fmt.Println("Error in PuState for record type " + key)

			fmt.Println("***********Exiting parseConfigFile***********")

			return shim.Error("Error in PuState for record type " + key)
		}

	}

	fmt.Println("record_types_to_keys_map is in Ledger!!!")

	fmt.Println("***********Exiting parseConfigFile***********")

	return shim.Success([]byte("ConfigFile is in Ledger!!!"))
}

// =================================================================================================
// parseConfigFileinvoke - called in the special invoke function --- parse the config file and store data in ledger
// =================================================================================================
func (t *IPDCChaincode) parseConfigFileinvoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("***********Entering parseConfigFileinvoke***********")

	fmt.Println("Checking if argument is valid string and has no * ` ~ \\ ? < > special characters")

	if !utf8string.NewString(args[0]).IsASCII() {

		fmt.Println(fmt.Sprintf("Error: Invalid argument: not ASCII"))

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error(fmt.Sprintf("Error: Invalid argument: not ASCII"))
	}

	if strings.ContainsAny(args[0], "~*\\^`?<>") {

		fmt.Println(fmt.Sprintf("Error: Invalid characters in argument"))

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error(fmt.Sprintf("Error: Invalid characters in argument"))
	}

	fmt.Println("Parsing config_json_bytes")

	configbuffer := new(bytes.Buffer)

	if err_buf := json.Compact(configbuffer, []byte(args[0])); err_buf != nil {

		fmt.Println("Error in decoding argument: " + err_buf.Error())

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error("Error in decoding argument: " + err_buf.Error())
	}

	var json_specification_global interface{}

	err := json.Unmarshal(configbuffer.Bytes(), &json_specification_global)

	if err != nil {

		fmt.Println("Error in Unmarshalling configbuffer for argument: " + err.Error())

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error("Error in Unmarshalling configbuffer for argument: " + err.Error())
	}

	json_specification, ok_json_specification := json_specification_global.(map[string]interface{})

	if !ok_json_specification {

		fmt.Println("Error in decoding configbuffer into map")

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error("Error in decoding configbuffer into map")
	}

	_, ok_json_specification = json_specification["config_json_bytes"]

	if !ok_json_specification {

		fmt.Println("Error in finding specification in configbuffer for config_json_bytes")

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error("Error in finding specification in configbuffer for config_json_bytes")
	}

	map_specification, ok := json_specification["config_json_bytes"].(map[string]interface{})

	if !ok {
		fmt.Println("Error in decoding configbuffer into map for config_json_bytes")

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error("Error in decoding configbuffer into map for config_json_bytes")
	}

	for key, value := range map_specification {

		value_byteArray, err2 := json.Marshal(value)

		if err2 != nil {

			fmt.Println("Error in marshalling JSON for function " + key)

			fmt.Println("***********Exiting parseConfigFileinvoke***********")

			return shim.Error("Error in marshalling JSON for function " + key)

		}

		operation := "FunctionName*" + key

		err2 = stub.PutState(operation, value_byteArray)

		if err2 != nil {

			fmt.Println("Error in PuState for function " + key)

			fmt.Println("***********Exiting parseConfigFileinvoke***********")

			return shim.Error("Error in PuState for function " + key)
		}

	}

	fmt.Println("config_json_bytes is in Ledger!!!")

	fmt.Println("Parsing record_types_to_keys_map")
	/*
		configbuffer = new(bytes.Buffer)

		if err_buf := json.Compact(configbuffer, record_types_to_keys_map); err_buf != nil {

			fmt.Println("Error in decoding record_types_to_keys_map: " + err_buf.Error())

			fmt.Println("***********Exiting parseConfigFile***********")

			return shim.Error("Error in decoding record_types_to_keys_map: " + err_buf.Error())
		}


		err = json.Unmarshal(configbuffer.Bytes(), &json_specification)

		if(err != nil) {

			fmt.Println("Error in Unmarshalling configbuffer for record_types_to_keys_map: " + err.Error())

			fmt.Println("***********Exiting parseConfigFile***********")

			return shim.Error("Error in Unmarshalling configbuffer for record_types_to_keys_map: " + err.Error())
		}
	*/

	_, ok_json_specification = json_specification["record_types_to_keys_map"]

	if !ok_json_specification {

		fmt.Println("Error in finding specification in configbuffer for record_types_to_keys_map")

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error("Error in finding specification in configbuffer for record_types_to_keys_map")
	}

	map_specification, ok = json_specification["record_types_to_keys_map"].(map[string]interface{})

	if !ok {
		fmt.Println("Error in decoding configbuffer into map for record_types_to_keys_map")

		fmt.Println("***********Exiting parseConfigFileinvoke***********")

		return shim.Error("Error in decoding configbuffer into map for record_types_to_keys_map")
	}

	for key, value := range map_specification {

		value_byteArray, err2 := json.Marshal(value)

		if err2 != nil {

			fmt.Println("Error in marshalling keys JSON for record type " + key)

			fmt.Println("***********Exiting parseConfigFileinvoke***********")

			return shim.Error("Error in marshalling  keys JSON for record type " + key)

		}

		record_type_key := "RecordType*" + key

		err2 = stub.PutState(record_type_key, value_byteArray)

		if err2 != nil {

			fmt.Println("Error in PuState for record type " + key)

			fmt.Println("***********Exiting parseConfigFileinvoke***********")

			return shim.Error("Error in PuState for record type " + key)
		}

	}

	fmt.Println("record_types_to_keys_map is in Ledger!!!")

	fmt.Println("***********Exiting parseConfigFileinvoke***********")

	return shim.Success([]byte("ConfigFile is in Ledger!!!"))
}

// =================================================================================================
// query_configs - query the configurations which are stored in the ledger
// =================================================================================================
func (t *IPDCChaincode) query_configs(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// fmt.Println("came in query configs args is ")
	// fmt.Println(args)

	fmt.Println("***********Entering query_configs***********")

	//key:="FunctionName*" + args[0]
	key := args[0]

	valAsBytes, err := stub.GetState(key)

	if err != nil {
		fmt.Println("Failed to get state: " + err.Error())

		fmt.Println("***********Exiting query_configs***********")

		return shim.Error("Failed to get state: " + err.Error())

	} else if valAsBytes == nil {

		fmt.Println("No value for key : " + key)

		fmt.Println("***********Exiting query_configs***********")

		return shim.Error("No value for key : " + key)
	}

	fmt.Println("***********Exiting query_configs***********")

	return shim.Success(valAsBytes)

}
