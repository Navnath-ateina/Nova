#!/bin/bash
. ./config.sh

port=7054

echo "version: '2'
services:" > fabric-ca.yaml
while read org
do
org1=$( echo $org | tr '[:upper:]' '[:lower:]' )
file1=$( ls -lthr ./crypto-config/peerOrganizations/$org1.$domain/ca/ | grep -i 'pem' | awk -F' ' '{print $NF}')
file2=$( ls -lthr ./crypto-config/peerOrganizations/$org1.$domain/ca/ | grep -i $user |grep -vi 'pem' | awk -F' ' '{print $NF}')

echo "  ca-"$org1":
    image: hyperledger/fabric-ca:$image
    environment:
      - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
      - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=host
      - FABRIC_CA_SERVER_CA_NAME=ca-"$org1"
      - FABRIC_CA_SERVER_CA_CERTFILE=/etc/hyperledger/fabric-ca-server-config/"$file1"
      - FABRIC_CA_SERVER_CA_KEYFILE=/etc/hyperledger/fabric-ca-server-config/"$file2"
      - FABRIC_CA_SERVER_TLS_ENABLED=true
      - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/"$file1"
      - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/"$file2"
    ports:
      - "$port:$port"
    command: sh -c 'fabric-ca-server start -b admin:adminpw -d'
    volumes:
      - ./crypto-config/peerOrganizations/"$org1"."$domain"/ca/:/etc/hyperledger/fabric-ca-server-config
    container_name: ca-$org1 
" >> fabric-ca.yaml

port=$(( port + 1000 ))
done < $org_file

