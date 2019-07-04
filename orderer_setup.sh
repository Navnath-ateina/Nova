#!/bin/bash
. ./config.sh

cp -ra ./raft.yaml ./orderer.$domain.yaml
sed -i "s/cateina.in/$domain/g" ./orderer.$domain.yaml
port=8050
i=2
while [ $i -le $orderer_no ]
do
orderer=$(cat $hosts_file | grep -i "orderer.$domain"| uniq | awk -F' ' '{print $1}' )
echo "$i"
echo "version: '2'
services:
  orderer"$i"."$domain":
    extends:
      file: ./peer-base.yaml
      service: orderer-base
    environment:
      - ORDERER_GENERAL_LISTENPORT="$port"
      - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=host
    container_name: orderer"$i".$domain
    volumes:
        - ./channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
        - ./crypto-config/ordererOrganizations/$domain/orderers/orderer"$i".$domain/msp:/var/hyperledger/orderer/msp
        - ./crypto-config/ordererOrganizations/$domain/orderers/orderer"$i".$domain/tls/:/var/hyperledger/orderer/tls
    ports:
        - $port:$port 
    extra_hosts:
        - \"orderer.$domain: "$orderer"\" " > ./orderer"$i".$domain.yaml
	j=2
	while [ $j -le $orderer_no ]
	do
    #echo "in second while"
		orderer=$(cat $hosts_file | grep -i "orderer"$j".$domain"| uniq | awk -F' ' '{print $1}' )	
		echo "            - \"orderer"$j".yesbank.in: $orderer\" " >> ./orderer"$i".$domain.yaml
    j=$(( j + 1 ))

	done

port=$(( port + 1000 ))
i=$(( i + 1 ))
done

j=2
while [ $j -le $orderer_no ]
do
  #echo "in second while"
  orderer=$(cat $hosts_file | grep -i "orderer"$j".$domain"| uniq | awk -F' ' '{print $1}' );
  echo "orderer"
  echo "            - \"orderer"$j".yesbank.in: $orderer\" " >> ./orderer.$domain.yaml
  j=$(( j + 1 ))

done
orderer=$(cat $hosts_file | grep -i "orderer.$domain"| uniq | awk -F' ' '{print $1}' );
echo "            - \"orderer.yesbank.in: $orderer\" " >> ./orderer.$domain.yaml