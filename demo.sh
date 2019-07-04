#!/bin/bash
. ./config.sh
. ./pre_req.sh
echo -e "Before proceed for the setup please check the config.sh file in details ?\n yes or no "
read check
if [[ $check == 'no' ]]
then
	echo "Thank you"
	exit 1;
fi

# if [[ $( cat $hosts_file | awk -F' ' '{print $1}' | sed '/^$/d' | uniq | wc -l) -lt 2 ]] 
# then 
# 	echo "For the the purpose of fault tolerance. The No. of hosts will be more than TWO." 
# 	exit 1; 
# fi

function generate_certs()
{
	bash config_generator.sh
	bash configtx_generator.sh;
	bash byfn_generator.sh;
	######configtx.yaml file change is done by another script#####
	which configtxgen
	if [ "$?" -ne 0 ]
	then
		echo -e "Before chaincode installation, Please execute test.sh ";
		exit 1;
	fi
	bash prod_setup.sh
	if [ "$?" -ne 0 ]
	then
		echo "Something we wrong........."
		exit 1;
	fi
}
######### generate certs 
generate_certs;

###### seperate for orderers
bash orderer_setup.sh

##### seperate for the peers
bash peer_setup.sh

#### Seperate for the CA
bash ca_setup.sh

### create network-config.yaml  ## pending in instantiate file 
bash network_config.sh

### build network
bash network_build.sh

### Install instantiate chaincode
#bash deploy.sh

