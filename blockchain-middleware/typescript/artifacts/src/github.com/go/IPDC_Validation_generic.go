package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// =================================================================================================
// validation_checks - perform respective validation checks as mentioned in the config file
// =================================================================================================
func (t *IPDCChaincode) validation_checks(stub shim.ChaincodeStubInterface, map_specification map[string]interface{}, map_record map[string]interface{}) (int, error) {

	fmt.Println("***********Entering validation_checks***********")

	//var return_val int

	var return_err string

	return_err = ""

	all_validations := map_specification["Validation_checks"]

	if all_validations == nil {

		fmt.Println("No validation checks specified")

		fmt.Println("***********Exiting validation_checks***********")

		return 0, nil
	}

	validations_map, ok := all_validations.(map[string]interface{})

	if !ok {
		fmt.Println("Error in validation_checks config")

		fmt.Println("***********Exiting validation_checks***********")

		return -2, errors.New("Error in validation_checks config")
	}

	for k, v := range validations_map {

		fmt.Println(k)

		// fmt.Println(v)

		var validation_ map[string]interface{}

		validation_, ok = v.(map[string]interface{})

		if !ok {
			fmt.Println(fmt.Sprintf("Error in parsing validation_check %s.", k))

			//return_err = return_err + fmt.Sprintf("Error in parsing validation check %s. ", k)

			fmt.Println("***********Exiting validation_checks***********")

			return -2, errors.New(fmt.Sprintf("Error in parsing validation check %s. ", k))

			//continue
		}

		fmt.Println(validation_)

		var target_pk_fields []interface{}

		if validation_["Target_primary_key"] == nil {

			fmt.Println("Missing Target_primary_key for validation check %s", k)

			return_err = return_err + fmt.Sprintf("Missing Target_primary_key for validation check %s. ", k)

			continue
		}

		target_pk_fields, ok = validation_["Target_primary_key"].([]interface{})

		if !ok {
			fmt.Println(fmt.Sprintf("Error in parsing Target_primary_key for validationcheck %s.", k))

			//return_err = return_err + fmt.Sprintf("Error in parsing Target_primary_key for validationcheck %s. ", k)

			fmt.Println("***********Exiting validation_checks***********")

			return -2, errors.New(fmt.Sprintf("Error in parsing Target_primary_key for validationcheck %s.", k))

			//continue
		}

		fmt.Println(target_pk_fields)

		var target_pk_values map[string]interface{}

		if validation_["Target_addtional_primary_key_fields_values"] == nil {

			target_pk_values = make(map[string]interface{})

		} else {

			// fmt.Println(target_pk_fields[0].(string))

			target_pk_values, ok = validation_["Target_addtional_primary_key_fields_values"].(map[string]interface{})

			if !ok {

				fmt.Println(fmt.Sprintf("Error in parsing Target_addtional_primary_key_fields_values for validationcheck %s", k))

				//return_err = return_err + fmt.Sprintf("Error in parsing Target_addtional_primary_key_fields_values for validationcheck %s. ", k)

				fmt.Println("***********Exiting validation_checks***********")

				return -2, errors.New(fmt.Sprintf("Error in parsing Target_addtional_primary_key_fields_values for validationcheck %s. ", k))

				//continue

			}

		}

		// fmt.Println(target_pk_values[target_pk_fields[0].(string)])

		var validation_fields_map map[string]interface{}

		if validation_["Validation_fields_map"] == nil {

			fmt.Println(fmt.Sprintf("Missing Validation_fields_map for validationcheck %s", k))

			return_err = return_err + fmt.Sprintf("Missing Validation_fields_map for validationcheck %s. ", k)

			continue

		} else {

			//var validation_fields_map map[string]interface{}

			validation_fields_map, ok = validation_["Validation_fields_map"].(map[string]interface{})

			if !ok {
				fmt.Println(fmt.Sprintf("Error in parsing Validation_fields_map for validation check %s", k))

				//return_err = return_err + fmt.Sprintf("Error in parsing Validation_fields_map for validationcheck %s. ", k)

				fmt.Println("***********Exiting validation_checks***********")

				return -2, errors.New(fmt.Sprintf("Error in parsing Validation_fields_map for validation check %s. ", k))

				//continue
			}
		}

		var target_fields_values_checks map[string]interface{}

		if validation_["Target_fields_values_checks"] == nil {

			fmt.Println(fmt.Sprintf("Missing Target_fields_values_checks for validationcheck %s", k))

			//return -3, errors.New(fmt.Sprintf("Missing Target_fields_values_checks for validationcheck %s", k))

			target_fields_values_checks = make(map[string]interface{})

		} else {

			//var target_fields_values_checks map[string]interface{}

			target_fields_values_checks, ok = validation_["Target_fields_values_checks"].(map[string]interface{})

			if !ok {
				fmt.Println(fmt.Sprintf("Error in parsing Target_fields_values_checks for validation check %s", k))

				//return_err = return_err + fmt.Sprintf("Error in parsing Target_fields_values_checks for validationcheck %s. ", k)

				fmt.Println("***********Exiting validation_checks***********")

				return -2, errors.New(fmt.Sprintf("Error in parsing Target_fields_values_checks for validation check %s. ", k))

				//continue
			}
		}

		fmt.Println("Validation Fields Map: ")
		fmt.Println(validation_fields_map)

		//return_val = 0

		found_err := false

		//for i := 0; i < len(target_pk_fields); i++ {
		for _, target_pk_fields_interface := range target_pk_fields {

			// fmt.Println(target_pk_fields[i].(string))

			// fmt.Println(validation_fields_map[target_pk_fields[i].(string)])

			target_pk_fields_string, ok := target_pk_fields_interface.(string)

			if !ok {

				fmt.Println(fmt.Sprintf("Error in decoding Target fields for validation check %s", k))

				//return_err = return_err + fmt.Sprintf("Error in decoding Target fields for validationcheck %s. ", k)

				fmt.Println("***********Exiting validation_checks***********")

				return -2, errors.New(fmt.Sprintf("Error in decoding Target fields for validation check %s. ", k))

				//found_err = true

				//break
			}

			// fmt.Println("------------" + target_pk_fields_string)
			// fmt.Println(validation_fields_map[target_pk_fields_string])
			// fmt.Println(reflect.TypeOf(validation_fields_map[target_pk_fields_string]))
			// fmt.Println(reflect.TypeOf(validation_fields_map[target_pk_fields_string].(string)))

			if target_pk_values[target_pk_fields_string] == nil {

				if validation_fields_map[target_pk_fields_string] == nil {

					fmt.Println(return_err + fmt.Sprintf("Missing Target field %s in validation fields map for validationcheck %s", target_pk_fields_string, k))

					return_err = return_err + fmt.Sprintf("Missing Target field %s in validation fields map for validationcheck %s. ", target_pk_fields_string, k)

					found_err = true

					break

				}

				validation_fields_map_string, ok := validation_fields_map[target_pk_fields_string].(string)

				if !ok {
					fmt.Println(fmt.Sprintf("Error in decoding validation fields map for Target field %s and validation check %s", target_pk_fields_string, k))

					//return_err = return_err + return_err + fmt.Sprintf("Error in decoding validation fields map for Target field %s and validation check %s. ", target_pk_fields_string, k)

					fmt.Println("***********Exiting validation_checks***********")

					return -2, errors.New(fmt.Sprintf("Error in decoding validation fields map for Target field %s and validation check %s. ", target_pk_fields_string, k))

					//found_err = true

					//break

				}

				if map_record[validation_fields_map_string] == nil {

					fmt.Println("Missing field %s in input record for validation check %s", validation_fields_map_string, k)

					return_err = return_err + return_err + fmt.Sprintf("Missing field %s in input record for validation check %s. ", validation_fields_map_string, k)

					found_err = true

					break

				}

				target_pk_values[target_pk_fields_string] = map_record[validation_fields_map_string]
			}
		}
		// fmt.Println("------------")
		// fmt.Println(target_pk_values)

		if found_err {

			continue
		}

		var key string

		var err_key error

		key, err_key = t.createInterfacePrimaryKey(target_pk_values, target_pk_fields)

		if err_key != nil {

			fmt.Println(fmt.Sprintf("Error in creating a primary key for validation check %s. ", k))

			return_err = return_err + fmt.Sprintf("Error in creating a primary key for validation check %s. ", k)

			//fmt.Println("***********Exiting validation_checks***********")

			//return -2, errors.New(fmt.Sprintf("Error in creating a primary key for validation check %s. ", k))

			continue

		}

		valAsBytes, err := stub.GetState(key)

		if err != nil {

			fmt.Println(fmt.Sprintf("Error in validation check: Failed to get state: for target primary key %s in validation check %s. ", key, k))

			//return_err = return_err + fmt.Sprintf("Error in validation check: Failed to get state: for target primary key %s in validation check %s. ", key, k)

			fmt.Println("***********Exiting validation_checks***********")

			return -2, errors.New(fmt.Sprintf("Error in validation check: Failed to get state: for target primary key %s in validation check %s. ", key, k))

			//continue

		} else if valAsBytes == nil {

			fmt.Println(fmt.Sprintf("Error in validation check: No value for target primary key %s. Validation check %s has failed. ", key, k))

			//return_val = -1

			//return_err = errors.New(fmt.Sprintf("Error in validation check: No value for target primary key %s", key))

			return_err = return_err + fmt.Sprintf("Validation check %s has failed: no value for target primary key. ", k)

			continue

			//return -1, errors.New(fmt.Sprintf("Error in validation check: No value for target primary key %s", key))
		}

		var record_to_be_validated_against map[string]interface{}

		err = json.Unmarshal([]byte(valAsBytes), &record_to_be_validated_against)

		if err != nil {

			fmt.Println(fmt.Sprintf("Error in decoding Input Record for validation check %s. ", k))

			//return_err = return_err +  fmt.Sprintf("Error in decoding Input Record for validation check %s. ", k)

			fmt.Println("***********Exiting validation_checks***********")

			return -2, errors.New(fmt.Sprintf("Error in decoding Input Record for validation check %s. ", k))

			//continue
		}

		for k_target, v_target := range target_fields_values_checks {

			v_target_string, ok3 := v_target.(string)

			if !ok3 {

				fmt.Println(fmt.Sprintf("Error in decoding value of %s in target_fields_values_checks in validation check %s. ", k_target, k))

				//return_err = return_err +  fmt.Sprintf("Error in decoding value of %s in target_fields_values_checks in validation check %s. ", k_target, k)

				fmt.Println("***********Exiting validation_checks***********")

				return -2, errors.New(fmt.Sprintf("Error in decoding value of %s in target_fields_values_checks in validation check %s. ", k_target, k))

				//continue
			}

			if record_to_be_validated_against[k_target] != v_target_string {

				fmt.Println(fmt.Sprintf("Validation check %s has failed  for field %s.", k, k_target))

				//return_val = -2

				//return_err = errors.New(fmt.Sprintf("Validation check %s has failed  for field %s", k, k_target))

				return_err = return_err + fmt.Sprintf("Validation check %s has failed  for field %s. ", k, k_target)

				continue
			}
		}
	}

	if return_err == "" {

		fmt.Println("***********Exiting validation_checks***********")

		return 0, nil

	} else {

		fmt.Println("***********Exiting validation_checks***********")

		return -1, errors.New(return_err)
	}
}
