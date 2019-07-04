package main

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//type IPDCChaincode struct {
//}

func (t *IPDCChaincode) Mapinputfieldstotarget(record_specification_initial map[string]interface{}, map_specification map[string]interface{}) (map[string]interface{}, error) {

	fmt.Println("***********Entering Mapinputfieldstotarget***********")

	var record_specification_return = make(map[string]interface{})

	fieldsmapstring, ok1 := map_specification["fields_map"]

	if !ok1 {
		fmt.Println("fields_map not present, no field mapping performed fields")

		//return record_specification_initial, nil

		fmt.Println("***********Exiting Mapinputfieldstotarget***********")

		return nil, errors.New("No Fields Mapping Present")

	} else {
		fieldsmap, ok2 := fieldsmapstring.(map[string]interface{})
		// fmt.Println("#######################################################")
		// fmt.Println(fieldsmap)
		// fmt.Println("#######################################################")

		if !ok2 {

			fmt.Println("Error in parsing fieldsmap")

			fmt.Println("***********Exiting Mapinputfieldstotarget***********")

			return nil, errors.New("Error in fields_map config")

		} else {

			for key, value := range fieldsmap {
				// fmt.Println("#######################################################")
				// fmt.Println(key)
				// fmt.Println(value)
				// fmt.Println("#######################################################")

				var domain, ok = value.(string)

				var _, okIn = value.([]interface{})
				// fmt.Println(okIn)
				// fmt.Println(domain_interface)

				if !ok {
					if !okIn {
						fmt.Println(fmt.Sprintf("Error in decoding domain of key %s", key))

						fmt.Println("***********Exiting Mapinputfieldstotarget***********")

						return nil, errors.New("Error in fields_map config")
					}

				}

				_, ok = record_specification_initial[domain].(string)

				// If string, apply trimspace

				if ok {
					record_specification_return[key] = strings.TrimSpace(record_specification_initial[domain].(string))

				} else if record_specification_initial[domain] != nil {

					//return nil, errors.New(fmt.Sprintf("Error in parsing value of %s. ", domain))
					record_specification_return[key] = record_specification_initial[domain]

				}

				if okIn {
					gob.Register(map[string]interface{}{})
					bytesData, err := json.Marshal(record_specification_initial[key])
					// var buf bytes.Buffer
					// enc := gob.NewEncoder(&buf)
					// err_json := enc.Encode(record_specification_initial[key])
					var list_args_interface []interface{}
					err_json := json.Unmarshal(bytesData, &list_args_interface)
					// bufOut := bytes.NewBuffer(buf.Bytes())
					// dec := gob.NewDecoder(bufOut)
					// err := dec.Decode(list_args_interface)

					if err != nil || err_json != nil {
						fmt.Println(list_args_interface)
						// fmt.Println([]byte(buf.Bytes()))
						fmt.Println(err)
						fmt.Println(err_json)

						fmt.Println(fmt.Sprintf("Error in decoding domain of key %s", key))

						fmt.Println("***********Exiting Mapinputfieldstotarget***********")

						return nil, err
					}

					record_specification_return[key] = list_args_interface
					fmt.Println("#######################################################")
					fmt.Println(record_specification_return[key])
					fmt.Println(key)
					fmt.Println(record_specification_initial[key])
					fmt.Println("#######################################################")
				}

			}

			/*
				anchor_code_string, ok3 := record_specification_return["anchor-code"]

				if !ok3 {
					 //record_specification_return["anchor-code"] = ANCHOR_CODE

				 } else if (anchor_code_string != ANCHOR_CODE) {

					 fmt.Println("Invalid Anchor Code")

					 fmt.Println("***********Exiting Mapinputfieldstotarget***********")

					 return nil, errors.New("Invalid Anchor Code")

				 }

				_, ok4 := record_specification_return["invoice-number"]

				if ok4 {

					 invoice_number_string, ok5 := record_specification_return["invoice-number"].(string)

					 if !ok5 {

						 fmt.Println("Invoice number is not valid")

						 fmt.Println("***********Exiting Mapinputfieldstotarget***********")

						 return nil, errors.New("Invoice number is not valid")

					} else if (len(invoice_number_string) > 16) {

						fmt.Println("Invoice Number length exceeds limit")

						fmt.Println("***********Exiting Mapinputfieldstotarget***********")

						return nil, errors.New("Invoice Number length exceeds limit")
					}
				 }

				 var invoice_date_interface interface{}

				 invoice_date_interface, ok4 = record_specification_return["invoice-date"]

				 if ok4 {

					 invoice_date_string, ok5 := invoice_date_interface.(string)

					 if !ok5 {
						 fmt.Println("Invalid format of invoice-date")

						 fmt.Println("***********Exiting Mapinputfieldstotarget***********")

						 return nil, errors.New("Invalid format of invoice-date")
					 }

					 _, err := time.Parse("2006-01-02", invoice_date_string)

					if err != nil {
							fmt.Println(fmt.Sprintf("Invalid date format for invoice-date", invoice_date_string))

							fmt.Println("***********Exiting Mapinputfieldstotarget***********")

							return nil, errors.New(fmt.Sprintf("Invalid date format for invoice-date", invoice_date_string))
					}

					record_specification_return["two-digit-invoice-financial-year"] = invoice_date_string[2:4]

				}
			*/
			fmt.Println("***********Exiting Mapinputfieldstotarget***********")

			return record_specification_return, nil

		}
	}
}

func (t *IPDCChaincode) Mandatoryfieldscheck(record_specification_tocheck map[string]interface{}, map_specification map[string]interface{}) (int, error) {

	fmt.Println("***********Entering Mandatoryfieldscheck***********")

	//var record_specification_tocheck map[string]interface{}

	mandatoryfieldsinterface, ok := map_specification["mandatory_fields"]

	if !ok {
		fmt.Println("Mandatory_fields not present, no check performed")

		fmt.Println("0***********Exiting Mandatoryfieldscheck***********")

		return -1, errors.New("Mandatory_fields not present, no check performed")

	}

	var mandatoryfields_interfacearray []interface{}

	mandatoryfields_interfacearray, ok = mandatoryfieldsinterface.([]interface{})

	if !ok {
		fmt.Println("Mandatory_fields bad spec, no check performed")

		fmt.Println("1***********Exiting Mandatoryfieldscheck***********")

		return -1, errors.New("Mandatory_fields bad spec, no check performed")

	} else {

		for i, value := range mandatoryfields_interfacearray {

			var value_string, ok1 = value.(string)

			if !ok1 {

				fmt.Println("Manatory field no. %d not a string", i)

				continue
			}
			_, okIn := record_specification_tocheck[value_string].([]interface{})
			_, okAm := record_specification_tocheck[value_string].(float64)
			_, okBl := record_specification_tocheck[value_string].(bool)

			if _, ok1 = record_specification_tocheck[value_string]; !ok1 {

				fmt.Println(fmt.Sprintf("Mandatory field %s missing", value_string))

				fmt.Println("3***********Exiting Mandatoryfieldscheck***********")

				return 1, errors.New(fmt.Sprintf("Mandatory field %s missing", value_string))
			}

			if _, ok1 = record_specification_tocheck[value_string].(string); !ok1 {
				// fmt.Println(record_specification_tocheck[value_string])
				// fmt.Println(value_interface)
				// fmt.Println(okIn)
				if !okIn {
					if !okAm {
						if !okBl {
							fmt.Println(fmt.Sprintf("Mandatory field %s has invalid value", value_string))

							fmt.Println("4***********Exiting Mandatoryfieldscheck***********")

							return 1, errors.New(fmt.Sprintf("Mandatory field %s has invalid value", value_string))
						}
					}
				}
			}

			if !okIn {
				if !okAm {
					if !okBl {
						if len(strings.TrimSpace(record_specification_tocheck[value_string].(string))) == 0 {

							fmt.Println(fmt.Sprintf("Mandatory field %s has empty value", value_string))

							fmt.Println("5***********Exiting Mandatoryfieldscheck***********")

							return 1, errors.New(fmt.Sprintf("Mandatory field %s has empty value", value_string))
						}
					}
				}
			}

		}

		fmt.Println("56***********Exiting Mandatoryfieldscheck***********")

		return 0, nil

	}

}

func (t *IPDCChaincode) Isfieldsmodified(record_prev map[string]interface{}, record_curr map[string]interface{}, map_specification map[string]interface{}) (int, error) {

	fmt.Println("***********Entering Isfieldsmodified***********")

	fieldstocheckformod_interface, ok := map_specification["fields_mod_check"]

	if !ok {
		fmt.Println("fields_mod_check not present, no check performed")

		fmt.Println("***********Exiting Isfieldsmodified***********")

		return -1, errors.New("Fields to check not present, no check performed")

	}

	var fieldstocheckformod []interface{}

	fieldstocheckformod, ok = fieldstocheckformod_interface.([]interface{})

	if !ok {
		fmt.Println("fields_mod_check spec bad, no check performed")

		fmt.Println("***********Exiting Isfieldsmodified***********")

		return -1, errors.New("Fields to check have bad spec, no check performed")

	} else {

		for _, value_interface := range fieldstocheckformod {

			//var prev, curr,

			var value string

			value, ok = value_interface.(string)

			if !ok {

				fmt.Println("Error in parsing fields_mod_check")

				fmt.Println("***********Exiting Isfieldsmodified***********")

				return -1, errors.New("Error in fields to check for modification list, no check performed")
			}

			if (record_prev[value] == nil) && (record_curr[value] == nil) {
				continue
			}

			if ((record_prev[value] == nil) && (record_curr[value] != nil)) || ((record_prev[value] != nil) && (record_curr[value] == nil)) {

				fmt.Println(fmt.Sprintf("Field %s is modified", value))

				fmt.Println("***********Exiting Isfieldsmodified***********")

				return 1, errors.New(fmt.Sprintf("Field %s is modified", value))
			}

			//prev, ok = record_prev[value].(string)

			//if !ok {
			//	fmt.Println("Bad format of record input, no check performed")

			//	return -1, errors.New("Bad format of record input, no check performed")
			//}

			//curr, ok  = record_curr[value].(string)

			//if !ok {

			//	fmt.Println("Bad format of record in ledger, no check performed")

			//	return -1, errors.New("Bad format of record in ledger, no check performed")
			//}

			//if (strings.Compare(prev, curr) != 0) {

			//	return 1, errors.New(fmt.Sprintf("Field %s is modified", value))
			//}

			if !(reflect.DeepEqual(record_prev[value], record_curr[value])) {

				fmt.Println(fmt.Sprintf("Field %s is modified", value))

				fmt.Println("***********Exiting Isfieldsmodified***********")

				return 1, errors.New(fmt.Sprintf("Field %s is modified", value))
			}

		}

		return 0, nil
	}

}
