#!/bin/bash
. ./config.sh
function crypto_config_generator()
{

	echo -e "OrdererOrgs:
  - Name: OrdererOrg
    Domain: $domain
    Specs:
      - Hostname: orderer" >crypto-config.yaml
	i=2;
	if [[ $orderer_type == 'raft' ]] 
	then
		while [ $i -le $orderer_no ]
		do
			echo "      - Hostname: orderer"$i"" >> crypto-config.yaml
			i=$(( i + 1 ))
		done 
        #echo "raff"

	elif [[ $orderer_type == 'solo' ]]
	then
		echo "orderer_type is solo........"
	else
		echo "order_type not mentioned in the config.sh file..........."
	fi
	echo "PeerOrgs:" >> crypto-config.yaml ### PEER SECTION
	while read org;
	do
		echo "$org";
		org1="$(echo $org| tr '[:upper:]' '[:lower:]')" #<<< ${org:0:1})${org:1}"
		echo "  - Name: $org
    Domain: "$org1"."$domain"
    Template:
      Count: 1
    Users:
      Count: 1
" >>crypto-config.yaml
	done < $org_file
}

crypto_config_generator; 
