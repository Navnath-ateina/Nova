package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
)

//type IPDCChaincode struct {
//}

func (t *IPDCChaincode) getUserNameusingMsp(stub shim.ChaincodeStubInterface) (string, error) {

	creator, err := stub.GetCreator()

	if err != nil {

		return "", errors.New(fmt.Sprintf("getcreator operation failed. Error: %s", err.Error()))

	} else {

		fmt.Println("GET CREATOR returned")

		fmt.Println(string(creator))
	}

	sId := &msp.SerializedIdentity{}

	err = proto.Unmarshal(creator, sId)

	if err != nil {

		return "", errors.New(fmt.Sprintf("Could not deserialize a SerializedIdentity, err %s", err))
	}

	block, _ := pem.Decode(sId.IdBytes)

	if block == nil {

		fmt.Printf("Error: PEM not parsed")

		return "", errors.New("Error: PEM not parsed!")

	} else {

		c, c_err := x509.ParseCertificate(block.Bytes)

		if c_err != nil {

			fmt.Printf("Errorin ParseCertificate : %s", c_err.Error())

			return "", errors.New(fmt.Sprintf("Error in ParseCertificate : %s", c_err.Error()))
		}

		userName := string((c.Subject).CommonName)

		fmt.Println("---------------------- User is ----------------------")

		fmt.Println(userName)
		/*
			fmt.Println("--------------------Organization is -----------------")

			fmt.Println((c.Subject).Organization)

			fmt.Println("--------------------Certificate  is -----------------")

			fmt.Printf("%+v\n", (c.Subject).ToRDNSequence())
		*/

		return userName, nil
	}

	return "", errors.New("Error in getting username from get creator!")

}

func (t *IPDCChaincode) updatestatusvaliditycheck(updatestatus_tochange string, updatestatusvalue string, map_specification map[string]interface{}) (int, error) {

	fmt.Println("***********Entering updatestatusvaliditycheck***********")

	var updatestatusvalidity_interface, ok = map_specification["update_status_validity"]

	if !ok {
		fmt.Println("Error in update status value validity spec")

		fmt.Println("***********Exiting updatestatusvaliditycheck***********")

		return -1, errors.New("Error in update status value validity spec")
	}

	var updatestatusvalidity_spec map[string]interface{}

	updatestatusvalidity_spec, ok = updatestatusvalidity_interface.(map[string]interface{})

	if !ok {
		fmt.Println("Error in update status value validity spec")

		fmt.Println("***********Exiting updatestatusvaliditycheck***********")

		return -1, errors.New("Error in update status value validity map")
	}

	var accessmap_interface interface{}

	accessmap_interface, ok = updatestatusvalidity_spec[updatestatus_tochange]

	if !ok {
		fmt.Println(fmt.Sprintf("Update status %s is invalid. ", updatestatus_tochange))

		fmt.Println("***********Exiting updatestatusvaliditycheck***********")

		return -1, errors.New(fmt.Sprintf("Update status %s is invalid. ", updatestatus_tochange))
	}

	var accessmap map[string]interface{}

	accessmap, ok = accessmap_interface.(map[string]interface{})

	if !ok {
		fmt.Println(fmt.Sprintf("Error in update status value validity spec for %s", updatestatus_tochange))

		fmt.Println("***********Exiting updatestatusvaliditycheck***********")

		return -1, errors.New(fmt.Sprintf("Error in update status value validity spec for %s", updatestatus_tochange))
	}

	if accessmap[updatestatusvalue] == nil {

		fmt.Println(fmt.Sprintf("Invalid update status value %s for %s", updatestatusvalue, updatestatus_tochange))

		fmt.Println("***********Exiting updatestatusvaliditycheck***********")

		return 1, errors.New(fmt.Sprintf("Invalid update status value %s for %s", updatestatusvalue, updatestatus_tochange))
	}

	fmt.Println("***********Exiting updatestatusvaliditycheck***********")

	return 0, nil
}
