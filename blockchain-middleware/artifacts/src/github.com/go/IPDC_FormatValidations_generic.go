package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"golang.org/x/exp/utf8string"
)

func (t *IPDCChaincode) Datefieldscheck(record_specification_tocheck map[string]interface{}, map_specification map[string]interface{}) (int, error) {

	//var record_specification_tocheck map[string]interface{}

	fmt.Println("***********Entering Datefieldscheck***********")

	datefieldsinterface, ok := map_specification["date_fields"]

	if !ok {
		fmt.Println("Date_fields not present, no check performed")

		fmt.Println("***********Exiting Datefieldscheck***********")

		return -1, errors.New("Date_fields not present, no check performed")

	}

	var datefields_interfacearray []interface{}

	datefields_interfacearray, ok = datefieldsinterface.([]interface{})

	if !ok {
		fmt.Println("Date_fields bad specification, no check performed")

		fmt.Println("***********Exiting Datefieldscheck***********")

		return -1, errors.New("Date_fields bad specification, no check performed")

	} else {

		for i, value := range datefields_interfacearray {

			var value_string string

			value_string, ok = value.(string)

			if !ok {

				fmt.Println(fmt.Sprintf("Date field no. %d not a string in specification", i))

				continue
			}

			if _, ok = record_specification_tocheck[value_string]; !ok {

				fmt.Println(fmt.Sprintf("Date field %s missing", value_string))

				//return 1, errors.New(fmt.Sprintf("Date field %s missing", value_string))
			} else {
				var date_value_string string

				date_value_string, ok = record_specification_tocheck[value_string].(string)

				if !ok {
					fmt.Println(fmt.Sprintf("Record value for Date field %s not a string", value_string))

					fmt.Println("***********Exiting Datefieldscheck***********")

					return 1, errors.New(fmt.Sprintf("Record value for Date field %s not a string", value_string))
				}

				if len(strings.TrimSpace(date_value_string)) == 0 {

					fmt.Println(fmt.Sprintf("Date field %s missing", value_string))

					continue
				}

				_, err := time.Parse("2006-01-02T15:04:05.999Z", date_value_string)

				if err != nil {
					fmt.Println(fmt.Sprintf("Invalid date format for field %s", value_string))

					fmt.Println("***********Exiting Datefieldscheck***********")

					return 1, errors.New(fmt.Sprintf("Invalid date format for field %s", value_string))
				}
			}
		}

		fmt.Println("***********Exiting Datefieldscheck***********")

		return 0, nil

	}

}

func (t *IPDCChaincode) Amountfieldscheck(record_specification_tocheck map[string]interface{}, map_specification map[string]interface{}) (int, error) {

	//var record_specification_tocheck map[string]interface{}

	fmt.Println("***********Entering Amountfieldscheck***********")

	amountfieldsinterface, ok := map_specification["amount_fields"]

	if !ok {
		fmt.Println("Amount_fields not present, no check performed")

		fmt.Println("***********Exiting Amountfieldscheck***********")

		return -1, errors.New("Amount_fields not present, no check performed")

	}

	var amountfields_interfacearray []interface{}

	amountfields_interfacearray, ok = amountfieldsinterface.([]interface{})

	if !ok {
		fmt.Println("Amount_fields bad spec, no check performed")

		fmt.Println("***********Exiting Amountfieldscheck***********")

		return -1, errors.New("Amount_fields bad spec, no check performed")

	} else {

		for i, value := range amountfields_interfacearray {

			var value_string string

			value_string, ok = value.(string)
			// fmt.Println("value . string ===>")
			// fmt.Println(value.(string))
			// fmt.Println("value_string ===>")
			// fmt.Println(value_string)
			// fmt.Println("ok ===>")
			// fmt.Println(ok)

			if !ok {

				fmt.Println(fmt.Sprintf("Amount field no. %d not a string in specification", i))

				continue
			}

			if _, ok = record_specification_tocheck[value_string]; !ok {

				fmt.Println(fmt.Sprintf("Amount field %s missing", value_string))

				//return 1, errors.New(fmt.Sprintf("Amount field %s missing", value_string))
			} else {
				var amount_value_string string

				amount_value_string, ok = record_specification_tocheck[value_string].(string)
				if !ok {
					_, ok = record_specification_tocheck[value_string].(float64)
				}

				if !ok {
					fmt.Println(fmt.Sprintf("Record value for Amount field %s not a string", value_string))

					fmt.Println("***********Exiting Amountfieldscheck***********")

					return 1, errors.New(fmt.Sprintf("Record value for Amount field %s not a string", value_string))
				}

				if len(strings.TrimSpace(amount_value_string)) == 0 {

					fmt.Println(fmt.Sprintf("Amount field %s missing", value_string))

					continue
				}

				ok1, err := regexp.MatchString(`^-?[0-9]\d*(\.[0-9][0-9]?)?$`, amount_value_string)

				if (err != nil) || (!ok1) {

					fmt.Println(fmt.Sprintf("Invalid amount format for field %s", value_string))

					fmt.Println("***********Exiting Amountfieldscheck***********")

					return 1, errors.New(fmt.Sprintf("Invalid amount format for field %s", value_string))
				}

			}

		}

		fmt.Println("***********Exiting Amountfieldscheck***********")

		return 0, nil

	}

}

func (t *IPDCChaincode) StringValidation(record_specification_tocheck map[string]interface{}, map_specification map[string]interface{}) error {

	fmt.Println("***********Entering StringValidation***********")
	fmt.Println(record_specification_tocheck)
	amountfieldsinterface, _ := map_specification["amount_fields"]
	// var amountfields_interfacearray []interface{}
	// amountfields_interfacearray, _ = amountfieldsinterface.([]interface{})
	fmt.Println(amountfieldsinterface)
	for key, value := range record_specification_tocheck {
		// fmt.Println(fmt.Sprintf(" String Validation  key %s : ", key))

		if !utf8string.NewString(key).IsASCII() {
			fmt.Println(fmt.Sprintf("Invalid  key %s : not ASCII", key))

			fmt.Println("***********Exiting StringValidation***********")

			return errors.New(fmt.Sprintf("Invalid key %s : not ASCII", key))
		}

		if strings.ContainsAny(key, "~*\\^_`?<>") {
			fmt.Println(fmt.Sprintf("Invalid  characters in key %s", key))

			fmt.Println("***********Exiting StringValidation***********")

			return errors.New(fmt.Sprintf("Invalid characters in key %s", key))
		}

		//var err error

		value_string, ok := value.(string)

		if !ok {
			_, ok = value.([]interface{})
		}

		if !ok {
			_, ok = value.(bool)
		}

		if !ok {
			_, ok = value.(float64)
		}

		if !ok {
			_, ok = value.(int)
		}

		if !ok {
			fmt.Println(fmt.Sprintf("Invalid value of key %s", key))

			fmt.Println("***********Exiting StringValidation***********")

			return errors.New(fmt.Sprintf("Invalid value of key %s", key))
		}

		if !utf8string.NewString(value_string).IsASCII() {
			fmt.Println(fmt.Sprintf("Invalid  value %s for key %s : not ASCII", value, key))

			fmt.Println("***********Exiting StringValidation***********")

			return errors.New(fmt.Sprintf("Invalid  value %s for key %s : not ASCII", value, key))
		}

		if strings.ContainsAny(value_string, "~*\\^`?<>") {
			fmt.Println(fmt.Sprintf("Invalid  characters in value %s for key %s", value_string, key))

			fmt.Println("***********Exiting StringValidation***********")

			return errors.New(fmt.Sprintf("Invalid  characters in value %s for key %s", value_string, key))
		}
	}
	return nil
}
