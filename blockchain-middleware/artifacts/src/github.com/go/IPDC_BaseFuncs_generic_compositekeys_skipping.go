package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// =================================================================================================
// delete_composite_keys - takes as input record and keys and deletes composite keys
// =================================================================================================
func (t *IPDCChaincode) delete_composite_keys(stub shim.ChaincodeStubInterface, map_specification map[string]interface{}, map_record map[string]interface{}, key string) error {

	// fmt.Println("MAP SPEC")
	// fmt.Println(map_specification)
	// fmt.Println("map record")
	// fmt.Println(map_record)
	// fmt.Println("key")
	// fmt.Println(key)

	fmt.Println("***********Entering delete_composite_keys***********")

	fmt.Println("The key is: " + key)

	for spec, field := range map_specification {

		if spec != "primary_key" {

			ck_spec, ok := field.([]interface{})

			if !ok {

				fmt.Println("Error in Composite key specification")

				fmt.Println("***********Exiting delete_composite_keys***********")

				return errors.New(fmt.Sprintf("Error : Composite Key specification %s cannot be fetched", spec))
			}

			compositekey_struct, arr_value, length_of_array, err2 := t.createInterfaceCompositeKeyStruct(map_record, ck_spec, key, 1)

			if err2 != nil && length_of_array == -1 {

				fmt.Println("***********Exiting delete_composite_keys***********")

				return err2

			} else if err2 != nil {

				fmt.Println(err2.Error())

				continue
			}

			compositekey, err3 := stub.CreateCompositeKey(compositekey_struct, arr_value[:length_of_array])

			if err3 != nil {

				fmt.Println("***********Exiting delete_composite_keys***********")

				return err3
			}

			fmt.Println("Deleting composite key ----" + compositekey)

			err := stub.DelState(compositekey)

			if err != nil {

				fmt.Println("Error in deleting composite key")

				fmt.Println("***********Exiting delete_composite_keys***********")

				return err
			}
		}
	}

	fmt.Println("***********Exiting delete_composite_keys***********")

	return nil
}

// =================================================================================================
// create_composite_keys - takes as input record and keys and creates composite keys
// =================================================================================================
func (t *IPDCChaincode) create_composite_keys(stub shim.ChaincodeStubInterface, map_specification map[string]interface{}, map_record map[string]interface{}, key string) error {

	fmt.Println("***********Entering create_composite_keys***********")

	fmt.Println("The key is: " + key)

	for spec, field := range map_specification {

		if spec != "primary_key" {

			ck_spec, ok := field.([]interface{})

			if !ok {

				fmt.Println("Error in Composite key specification for " + spec)

				fmt.Println("***********Exiting create_composite_keys***********")

				return errors.New(fmt.Sprintf("Error : Composite Key specification %s cannot be fetched", spec))
			}

			compositekey_struct, arr_value, length_of_array, err2 := t.createInterfaceCompositeKeyStruct(map_record, ck_spec, key, 1)

			if err2 != nil && length_of_array == -1 {

				fmt.Println("Error in creating Composite struct for " + spec)

				fmt.Println("***********Exiting create_composite_keys***********")

				return err2

			} else if err2 != nil {

				fmt.Println(err2.Error())

				continue

			}

			compositekey, err3 := stub.CreateCompositeKey(compositekey_struct, arr_value[:length_of_array])

			if err3 != nil {

				fmt.Println("Error in creating Composite for " + spec)

				fmt.Println("***********Exiting create_composite_keys***********")

				return err3
			}

			valueAsNULL := []byte{0x00}

			stub.PutState(compositekey, valueAsNULL)
		}
	}

	fmt.Println("***********Exiting create_composite_keys***********")

	return nil
}

// =========================================================================================
// query_by_composite_key_primitive - querying by composite key (primitive function)
// =========================================================================================
func (t *IPDCChaincode) query_by_composite_key_primitive_string_args(stub shim.ChaincodeStubInterface, args []string) (error, shim.StateQueryIteratorInterface) {

	fmt.Println("***********Entering query_by_composite_key_primitive_string_args***********")

	if len(args) < 2 {

		fmt.Println("Requires at least two arguments")

		fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

		return errors.New("Requires at least two arguments"), nil
	}

	record := args[0]

	specification := args[1]

	byteArray_record := []byte(record)

	var map_record map[string]interface{}

	err := json.Unmarshal(byteArray_record, &map_record)

	if err != nil {

		fmt.Println("Error in decoding record JSON")

		fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

		return errors.New("JSON format for Record is not correct"), nil
	}

	var json_specification interface{}

	err = json.Unmarshal([]byte(specification), &json_specification)

	if err != nil {

		fmt.Println("Error in decoding Specification JSON")

		fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

		return errors.New("JSON format for Specification is not correct"), nil
	}

	map_specification := json_specification.(map[string]interface{})

	if len(map_specification) != 1 {

		fmt.Println("Incorrect size of map specification for composite key")

		fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

		return errors.New("Incorrect size of map specification for composite key"), nil
	}

	var key string

	key = ""

	for spec, field := range map_specification {

		ck_spec, ok := field.([]interface{})

		if !ok {

			fmt.Println("Error in Composite key specification")

			fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

			return errors.New(fmt.Sprintf("Error : Composite Key specification %s cannot be fetched", spec)), nil
		}

		compositekey_struct, arr_value, length_of_array, err2 := t.createInterfaceCompositeKeyStruct(map_record, ck_spec, key, 0)

		if err2 != nil {

			fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

			return err2, nil
		}

		tableRowsIterator, err3 := stub.GetStateByPartialCompositeKey(compositekey_struct, arr_value[:length_of_array])

		if err3 != nil {

			fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

			return err3, nil
		}

		//defer tableRowsIterator.Close()

		return nil, tableRowsIterator

	}

	fmt.Println("***********Exiting query_by_composite_key_primitive_string_args***********")

	return errors.New("Something went wrong in fetching the iterator"), nil
}

// =========================================================================================
// query_by_composite_key_primitive - querying by composite key (primitive function)
// =========================================================================================
//func (t *IPDCChaincode) query_by_composite_key_primitive(stub shim.ChaincodeStubInterface, args []string) (error,shim.StateQueryIteratorInterface) {
func (t *IPDCChaincode) query_by_composite_key_primitive(stub shim.ChaincodeStubInterface, map_record map[string]interface{}, map_specification map[string]interface{}) (error, shim.StateQueryIteratorInterface) {

	fmt.Println("***********Entering query_by_composite_key_primitive***********")

	if len(map_specification) != 1 {

		fmt.Println("Incorrect size of map specification for composite key")

		fmt.Println("***********Exiting query_by_composite_key_primitive***********")

		return errors.New("Incorrect size of map specification for composite key"), nil
	}

	var key string

	key = ""

	for spec, field := range map_specification {

		ck_spec, ok := field.([]interface{})

		if !ok {
			fmt.Println(fmt.Sprintf("Error : Composite Key specification %s cannot be fetched", spec))

			fmt.Println("***********Exiting query_by_composite_key_primitive***********")

			return errors.New(fmt.Sprintf("Error : Composite Key specification %s cannot be fetched", spec)), nil
		}

		compositekey_struct, arr_value, length_of_array, err2 := t.createInterfaceCompositeKeyStruct(map_record, ck_spec, key, 0)

		if err2 != nil {

			fmt.Println("***********Exiting query_by_composite_key_primitive***********")

			return err2, nil
		}

		tableRowsIterator, err3 := stub.GetStateByPartialCompositeKey(compositekey_struct, arr_value[:length_of_array])

		if err3 != nil {

			fmt.Println("***********Exiting query_by_composite_key_primitive***********")

			return err3, nil
		}

		//defer tableRowsIterator.Close()

		return nil, tableRowsIterator

	}

	fmt.Println("***********Exiting query_by_composite_key_primitive***********")

	return errors.New("Something went wrong in fetching the iterator"), nil
}

// =========================================================================================
// query_by_composite_key - querying by composite key (base function)
// =========================================================================================
//func (t *IPDCChaincode) query_by_composite_key(stub shim.ChaincodeStubInterface, args []string, pageno int) pb.Response {
func (t *IPDCChaincode) query_by_composite_key(stub shim.ChaincodeStubInterface, map_record map[string]interface{}, map_specification map[string]interface{}, pageno int) pb.Response {

	//if len(args) != 2 {

	//	return shim.Error("Operation requires two arguments - record and specification list")
	//}

	fmt.Println("***********Entering query_by_composite_key***********")

	err, tableRowsIterator := t.query_by_composite_key_primitive(stub, map_record, map_specification)

	if err != nil {

		fmt.Println("***********Exiting query_by_composite_key***********")

		return shim.Error(err.Error())
	}

	var list_of_records string

	list_of_records = "["

	records, err2 := t.fetchRecordsFromCompositeKeys(stub, tableRowsIterator, pageno)

	if err2 != nil {

		fmt.Println("***********Exiting query_by_composite_key***********")

		return shim.Error(err2.Error())
	}

	list_of_records += string(records)

	list_of_records += "]"

	fmt.Println("***********Exiting query_by_composite_key***********")

	return shim.Success([]byte(list_of_records))
}

// =========================================================================================
// query_for_record_history - querying for record history (base function)
// =========================================================================================
func (t *IPDCChaincode) query_for_record_history(stub shim.ChaincodeStubInterface, primary_key string, pageno int) pb.Response {

	fmt.Println("***********Entering query_for_record_history***********")

	var i int

	if pageno < 1 {

		fmt.Println("Invalid page number")

		fmt.Println("***********Exiting query_for_record_history***********")

		return shim.Error("Invalid page number")
	}

	startingindex := (pageno - 1) * PROCESSING_LIMIT

	endingindex := pageno * PROCESSING_LIMIT

	// First fetch the history iterator

	recordhistoryIterator, err0 := stub.GetHistoryForKey(primary_key)

	if err0 != nil {

		fmt.Println(err0.Error())

		fmt.Println("***********Exiting query_for_record_history***********")

		return shim.Error(err0.Error())
	}

	defer recordhistoryIterator.Close()

	var buffer bytes.Buffer

	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for i = 0; recordhistoryIterator.HasNext(); i++ {

		if i >= endingindex {

			break

		} else if i < startingindex {

			_, err := recordhistoryIterator.Next()

			if err != nil {

				fmt.Println(err.Error())

				fmt.Println("***********Exiting query_for_record_history***********")

				return shim.Error(err.Error())
			}

			continue

		} else {

			response, err := recordhistoryIterator.Next()

			if err != nil {

				fmt.Println(err.Error())

				fmt.Println("***********Exiting query_for_record_history***********")

				return shim.Error(err.Error())

			}

			// Add a comma before array members, suppress it for the first array member
			if bArrayMemberAlreadyWritten == true {

				buffer.WriteString(",")

			}

			buffer.WriteString("{\"TxId\":")

			buffer.WriteString("\"")

			buffer.WriteString(response.TxId)

			buffer.WriteString("\"")

			buffer.WriteString(", \"Value\":")

			// if it was a delete operation on given key, then we need to set the
			//corresponding value null. Else, we will write the response.Value
			//as-is (as the Value itself a JSON marble)

			if response.IsDelete {

				buffer.WriteString("null")

			} else {

				buffer.WriteString(string(response.Value))

			}

			buffer.WriteString(", \"Timestamp\":")

			buffer.WriteString("\"")

			buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())

			buffer.WriteString("\"")

			buffer.WriteString(", \"IsDelete\":")

			buffer.WriteString("\"")

			buffer.WriteString(strconv.FormatBool(response.IsDelete))

			buffer.WriteString("\"")

			buffer.WriteString("}")
			bArrayMemberAlreadyWritten = true

		}

	}

	buffer.WriteString("]")

	fmt.Println("***********Exiting query_for_record_history***********")

	return shim.Success(buffer.Bytes())

}

// =========================================================================
// HELPER FUNCTIONS
// =========================================================================

/*
Function to fetch all primary keys and corresponding values from a given composite key
*/
func (t *IPDCChaincode) fetchRecordsFromCompositeKeys(stub shim.ChaincodeStubInterface, tableRowsIterator shim.StateQueryIteratorInterface, pageno int) ([]byte, error) {

	fmt.Println("***********Entering fetchRecordsFromCompositeKeys***********")

	var i int

	if pageno < 1 {

		fmt.Println("Invalid Page Number")

		fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

		return nil, errors.New("Invalid Page Number")
	}

	startingindex := (pageno - 1) * PROCESSING_LIMIT

	endingindex := pageno * PROCESSING_LIMIT

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false

	defer tableRowsIterator.Close()

	for i = 0; tableRowsIterator.HasNext(); i++ {

		if i >= endingindex {

			break

		} else if i < startingindex {

			_, err := tableRowsIterator.Next()

			if err != nil {

				fmt.Println(err.Error())

				fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

				return nil, err
			}

			continue

		} else {

			indexKey, err := tableRowsIterator.Next()

			if err != nil {

				fmt.Println(err.Error())

				fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

				return nil, err
			}

			objectType, compositeKeyParts, err := stub.SplitCompositeKey(indexKey.Key)

			if err != nil {

				fmt.Println(err.Error())

				fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

				return nil, err
			}

			length_of_compositekey := len(compositeKeyParts)

			fmt.Println("ObjectType is ", objectType)

			returnedKey := compositeKeyParts[length_of_compositekey-1]

			// Add a comma before array members, suppress it for the first array member

			if bArrayMemberAlreadyWritten == true {

				buffer.WriteString(",")

			}

			valAsBytes, err2 := stub.GetState(returnedKey)

			if err2 != nil {

				fmt.Println(fmt.Sprintf("Failed to get state: %s", returnedKey))

				fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

				return nil, errors.New(fmt.Sprintf("Failed to get state: %s", returnedKey))

			} else if valAsBytes == nil {

				fmt.Println("No value for : " + returnedKey)

				fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

				return nil, errors.New(fmt.Sprintf("No value for : %s", returnedKey))
			}

			_, err_buffer := buffer.WriteString(string(valAsBytes))

			if err_buffer != nil {

				fmt.Println("Error in writing to buffer: " + err_buffer.Error())

				fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

				return nil, errors.New("Error in writing to buffer: " + err_buffer.Error())

			}

			bArrayMemberAlreadyWritten = true
		}
	}

	fmt.Println("***********Exiting fetchRecordsFromCompositeKeys***********")

	return []byte(buffer.Bytes()), nil
}

/*
Create Primary key where the 'key' consists of multiple fields i.e. key is an interface
*/
func (t *IPDCChaincode) createInterfacePrimaryKey(map_record map[string]interface{}, field_type []interface{}) (string, error) {

	fmt.Println("***********Entering createInterfacePrimaryKey***********")

	var key string

	// fmt.Println("map record ")
	// fmt.Println(map_record)

	// fmt.Println("field type ")
	// fmt.Println(field_type)

	for i, k := range field_type {

		var kstring, ok = k.(string)

		if !ok {
			fmt.Println("There is a problem in specification for primary key")

			fmt.Println("***********Exiting createInterfacePrimaryKey***********")

			return "", errors.New("Primary Key cannot be created! There is a problem in specification for primary key.")
		}

		var value, ok1 = map_record[kstring]

		if !ok1 {
			fmt.Println("Nil value in map_record")

			fmt.Println("***********Exiting createInterfacePrimaryKey***********")

			return "", errors.New("Primary Key cannot be created! Incomplete record details.")
		}

		var valuestring string

		valuestring, ok = value.(string)

		if !ok {
			fmt.Println("Cannot handle interface as a value!!! Not a string!")

			fmt.Println("***********Exiting createInterfacePrimaryKey***********")

			return "", errors.New("Primary Key cannot be created! Incomplete record details.")
		}

		if i == 0 {
			// fmt.Println("i ")
			// fmt.Println(i)
			// fmt.Println("k")
			// fmt.Println(k.(string))
			//value := map_record[kstring]
			// fmt.Println("VAL IS")
			// fmt.Println(value)
			// fmt.Println(reflect.TypeOf(value))
			//if(reflect.TypeOf(value)!=reflect.TypeOf("str")) {
			//	fmt.Println("Cannot handle interface as a value!!!")
			//	return "","Primary Key cannot be created! Incomplete record details."
			//}
			key = valuestring + "`" + kstring
		} else {
			//value := map_record[k.(string)]
			// fmt.Println("VAL IS11")
			// fmt.Println(value)
			// fmt.Println(reflect.TypeOf(value))
			//if(reflect.TypeOf(value)!=reflect.TypeOf("str")) {
			//	fmt.Println("Cannot handle interface as a value???")
			//	return "","Primary key cannot be created! Incomplete record details."
			//}
			key = key + "^" + valuestring + "`" + kstring
		}
	}

	fmt.Println("***********Exiting createInterfacePrimaryKey***********")

	return key, nil
}

/*
Creates structure of composite key where the 'key' consists of multiple fields i.e. key is an interface
*/
func (t *IPDCChaincode) createInterfaceCompositeKeyStruct(map_record map[string]interface{}, field_type []interface{}, key string, fullOrPartialCompositeKey int) (string, []string, int, error) {

	fmt.Println("***********Entering createInterfaceCompositeKeyStruct***********")

	arr_value := make([]string, 0)

	fmt.Println("----[createInterfaceCompositeKeyStruct] Key created is ---- " + key)

	if key == "" && fullOrPartialCompositeKey == 1 {

		fmt.Println("There is no primary key given!")

		fmt.Println("***********Exiting createInterfaceCompositeKeyStruct***********")

		return "", arr_value, -1, errors.New("There is no primary key given!")
	}

	var compositekey_struct string

	var gotemptyvalue_nomoreappending bool

	gotemptyvalue_nomoreappending = false

	for i, k := range field_type {

		var kstring, ok = k.(string)

		if !ok {
			fmt.Println("There is a problem in specification for composite key")

			fmt.Println("***********Exiting createInterfaceCompositeKeyStruct***********")

			return "", arr_value, -1, errors.New("There is a problem in specification for composite key")
		}

		var valuestring string

		if i == 0 {

			value := map_record[kstring]
			// fmt.Println("Value is ----- ")
			// fmt.Println(value)

			if value == nil {

				fmt.Println("Value missing for : " + kstring)

				compositekey_struct = kstring

				gotemptyvalue_nomoreappending = true

				continue
			}

			if !gotemptyvalue_nomoreappending {

				valuestring, ok = value.(string)

				if !ok {

					fmt.Println("There is a problem in fetching value for composite key. Make sure the value is passed in the args")

					fmt.Println("***********Exiting createInterfaceCompositeKeyStruct***********")

					return "", arr_value, -1, errors.New("There is a problem in fetching value for composite key. Make sure the value is passed in the args")
				}

				arr_value = append(arr_value, valuestring)
			}

			compositekey_struct = kstring

		} else {
			value := map_record[kstring]

			// fmt.Println("Value is ----- ")
			// fmt.Println(value)

			if value == nil {

				fmt.Println("Value missing for : " + kstring)

				compositekey_struct += "~" + kstring

				gotemptyvalue_nomoreappending = true

				continue
			}

			if !gotemptyvalue_nomoreappending {

				valuestring, ok = value.(string)

				if !ok {

					fmt.Println("There is a problem in fetching value for composite key. Make sure the value is passed in the args")

					fmt.Println("***********Exiting createInterfaceCompositeKeyStruct***********")

					return "", arr_value, -1, errors.New("There is a problem in fetching value for composite key. Make sure the value is passed in the args")
				}

				arr_value = append(arr_value, valuestring)
			}

			compositekey_struct += "~" + kstring

		}
	}

	if !gotemptyvalue_nomoreappending && fullOrPartialCompositeKey == 1 {

		arr_value = append(arr_value, key)

	}

	if gotemptyvalue_nomoreappending && fullOrPartialCompositeKey == 1 {

		fmt.Println("Some value not found to construct the full composite key")

		fmt.Println("***********Exiting createInterfaceCompositeKeyStruct***********")

		return "", arr_value, -2, errors.New("Some value not found to construct the full composite key")
	}

	compositekey_struct = compositekey_struct + "~key"

	length_of_array := len(arr_value)

	fmt.Println("***********Exiting createInterfaceCompositeKeyStruct***********")

	return compositekey_struct, arr_value, length_of_array, nil

}

/*
Function to delete all primary keys and corresponding values from a given iterator
*/
func (t *IPDCChaincode) deleteRecordsFromCompositeKeys(stub shim.ChaincodeStubInterface, tableRowsIterator shim.StateQueryIteratorInterface, specs map[string]interface{}, limit int) (error, int, bool) {

	fmt.Println("***********Entering deleteRecordsFromCompositeKeys***********")

	var record_specification map[string]interface{}

	/*var specs map[string]interface{}

	config, ok := map_specification["config"]

	if !ok {

		fmt.Println("Invalid function config.")

		fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

		return errors.New("Invalid function config."), 0, true
	}


	specs, ok = config.(map[string]interface{})

	if !ok {

		fmt.Println("Invalid function specification.")

		fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

		return errors.New("Invalid function specification."), 0, true
	}*/

	var i int

	defer tableRowsIterator.Close()

	for i = 0; tableRowsIterator.HasNext(); i++ {

		if i >= limit {

			fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

			return nil, i, true

		} else {

			indexKey, err := tableRowsIterator.Next()

			if err != nil {

				fmt.Println(err.Error())

				fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

				return err, i, true
			}

			objectType, compositeKeyParts, err := stub.SplitCompositeKey(indexKey.Key)

			if err != nil {

				fmt.Println(err.Error())

				fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

				return err, i, true
			}

			length_of_compositekey := len(compositeKeyParts)

			fmt.Println("ObjectType is ", objectType)

			returnedKey := compositeKeyParts[length_of_compositekey-1]

			var valAsBytes []byte

			valAsBytes, err = stub.GetState(returnedKey)

			if err != nil {

				fmt.Println(fmt.Sprintf("Failed to get state: %s", returnedKey))

				fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

				return errors.New(fmt.Sprintf("Failed to get state: %s", returnedKey)), i, true

			} else if valAsBytes == nil {

				fmt.Println("No value for : " + returnedKey)

			} else {

				err2 := json.Unmarshal([]byte(valAsBytes), &record_specification)

				if err2 != nil {

					fmt.Println("Error in decoding Record JSON " + string(valAsBytes))

					fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

					return errors.New("JSON format for Record is not correct" + string(valAsBytes)), i, true

				}

				err_del := t.delete_composite_keys(stub, specs, record_specification, returnedKey)

				if err_del != nil {

					fmt.Println("Received error while deleting composite keys: " + err_del.Error())

					fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

					return errors.New("Received error while deleting composite keys: " + err_del.Error()), i, true

				}

				err_del = stub.DelState(returnedKey)

				if err_del != nil {

					fmt.Println("Received error while deleting primary key: " + err_del.Error())

					fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

					return errors.New("Received error while deleting primary key: " + err_del.Error()), i, true

				}

			}

		}
	}

	fmt.Println("***********Exiting deleteRecordsFromCompositeKeys***********")

	return nil, i, false
}

func (t *IPDCChaincode) delete_by_composite_key(stub shim.ChaincodeStubInterface, args []string, map_specification map[string]interface{}, limit int) (error, int, bool) {

	fmt.Println("***********Entering delete_by_composite_key***********")

	if len(args) != 2 {

		fmt.Println("Operation requires two arguments - record and specification list")

		fmt.Println("***********Exiting delete_by_composite_key***********")

		return errors.New("Operation requires two arguments - record and specification list"), 0, true
	}

	err, tableRowsIterator := t.query_by_composite_key_primitive_string_args(stub, args)

	if err != nil {

		fmt.Println(err.Error())

		fmt.Println("***********Exiting delete_by_composite_key***********")

		return err, 0, true

	}

	fmt.Println("***********Exiting delete_by_composite_key***********")

	return t.deleteRecordsFromCompositeKeys(stub, tableRowsIterator, map_specification, limit)

}

func (t *IPDCChaincode) get_keys_map(stub shim.ChaincodeStubInterface, field_values_map map[string]interface{}) (interface{}, error) {

	fmt.Println("***********Entering get_keys_map***********")

	record_type_interface, ok := field_values_map["record-type"]

	if !ok {

		fmt.Println("Error: record-type is not defined in config.")

		fmt.Println("***********Exiting get_keys_map***********")

		return nil, errors.New("Error: record_type is not defined in config.")
	}

	var record_type_string string

	record_type_string, ok = record_type_interface.(string)

	if !ok {

		fmt.Println("Error: Invalid record_type in config.")

		fmt.Println("***********Exiting get_keys_map***********")

		return nil, errors.New("Error: Invalid record_type in config.")
	}

	key_for_map := "RecordType*" + record_type_string

	valAsBytes, err := stub.GetState(key_for_map)

	if err != nil {

		fmt.Println(fmt.Sprintf("Error: unable to retrieve keys map for " + record_type_string + " :" + err.Error()))

		fmt.Println("***********Exiting get_keys_map***********")

		return nil, errors.New(fmt.Sprintf("Error: unable to retrieve keys map for " + record_type_string + " :" + err.Error()))

	} else if valAsBytes == nil {

		fmt.Println("Error: no keys map for " + record_type_string)

		fmt.Println("***********Exiting get_keys_map***********")

		return nil, errors.New("Error: no keys map for " + record_type_string)
	}

	var interface_to_return interface{}

	err = json.Unmarshal(valAsBytes, &interface_to_return)

	if err != nil {

		fmt.Println("Error in unmarshaling keys map for " + record_type_string)

		fmt.Println("***********Exiting get_keys_map***********")

		return nil, errors.New("Error in unmarshaling keys map for " + record_type_string)
	}

	return interface_to_return, nil
}
