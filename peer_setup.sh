#!/bin/bash
. ./config.sh

port=7051;
port1=7551;
c_port=5984;
while read org
do
    rm -rf $org
    org1=$( echo $org | tr '[:upper:]' '[:lower:]')
    #mkdir $org
    #echo "Creating directory";
    cp -ra ./peer/0_abc.yaml peer0."$org1.$domain".yaml
    cp -ra ./peer/1_abc.yaml peer0."$org1.$domain".yaml

    sed -i "s/abc/$org1/g" peer0."$org1.$domain".yaml
    sed -i "s/cateina.in/$domain/g" peer0."$org1.$domain".yaml
    sed -i "s/Abc/"$org"/g" peer0."$org1.$domain".yaml
    sed -i "s/7051/"$port"/g" peer0."$org1.$domain".yaml
    sed -i "s/7052/$((port + 1))/g" peer0."$org1.$domain".yaml
    sed -i "s/7053/$(( port + 2))/g" peer0."$org1.$domain".yaml
    if [ $couchdb_required == 'yes' ]
    then
        sed -i "s/      - CORE_PEER_LOCALMSPID="$org"MSP/      - CORE_PEER_LOCALMSPID="$org"MSP\n      - CORE_LEDGER_STATE_STATEDATABASE=CouchDB\n      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb0$org1:$c_port\n      - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=admin\n      - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=manchester/g" peer0."$org1.$domain".yaml
        cp -ra ./peer/couchdb.yaml couchdb0"$org1".yaml
        sed -i "s/couchdb0/couchdb0$org1/g" couchdb0"$org1".yaml
        sed -i "s/5984/$c_port/g" couchdb0"$org1".yaml


        if [ $peer_type == 'bi' ] 
        then
            cp -ra ./peer/couchdb.yaml couchdb1"$org1".yaml
            #echo "couch"
            c_port=$(( c_port + 500 ))
            sed -i "s/      - CORE_PEER_LOCALMSPID=AbcMSP/      - CORE_PEER_LOCALMSPID=AbcMSP\n      - CORE_LEDGER_STATE_STATEDATABASE=CouchDB\n      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb1$org1:$c_port\n      - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=admin\n      - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=manchester/g" peer1."$org1.$domain".yaml
            sed -i "s/couchdb0/couchdb1$org1/g" couchdb1"$org1".yaml
            sed -i "s/5984/$c_port/g" couchdb1"$org1".yaml
            c_port=$(( c_port - 500 ))
        fi
    fi
    if [ $peer_type == 'bi' ]
    then
        
        sed -i "s/abc/$org1/g" peer0."$org1.$domain".yaml
        sed -i "s/cateina.in/$domain/g" peer0."$org1.$domain".yaml
        sed -i "s/Abc/"$org"/g" peer0."$org1.$domain".yaml
        sed -i "s/8051/"$port1"/g" peer0."$org1.$domain".yaml
        sed -i "s/7051/$port/g" peer0."$org1.$domain".yaml
        sed -i "s/8052/$((port1 + 1))/g" peer0."$org1.$domain".yaml
        sed -i "s/8053/$(( port1 + 2))/g" peer0."$org1.$domain".yaml

    fi
port=$(( port + 1000 ))
port1=$(( port1 + 1000 ))
c_port=$(( c_port + 1000 ))
done < $org_file



