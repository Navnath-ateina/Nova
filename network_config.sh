#!/bin/bash
. ./config.sh

file="$(pwd)/blockchain-middleware/artifacts/network-config.yaml"
if [ ! -f $file ]
then
    touch $file;
fi
if [ -d $(pwd)/blockchain-middleware/artifacts/crypto-config ] || [ -d $(pwd)/blockchain-middleware/artifacts/channel-artifacts ]
then
    rm -rf $(pwd)/blockchain-middleware/artifacts/crypto-config
    rm -rf $(pwd)/blockchain-middleware/artifacts/channel-artifacts
fi
cp -ra crypto-config channel-artifacts $(pwd)/blockchain-middleware/artifacts/.

echo "---
name: "blockchain-setup"
x-type: "hlfv1"
description: "Balance Transfer Network"
version: "1.0"
channels:" > $file

while read channel
do
    channel1=$(echo $channel | awk -F'channel' '{print $2}')
    channel2=$(echo $channel| sed 's/&/and/g'| tr '[:upper:]' '[:lower:]')
    
    echo "  $channel2
    orderers:
      - orderer.$domain" >> $file
    i=2
    while [ $i -le $orderer_no ]
    do
        echo "      - orderer"$i".$domain" >> $file
        i=$(( i +1 ))
    done
    echo "    peers:" >> $file

    a_num=$(echo $channel1 | awk -F'&' '{print NF}')
    if [[ $(echo $channel1 | grep -i '&' ) == ' ' ]]
    then
        org1=$(echo $channel1 | awk -F'channel' '{print $2}' | tr '[:upper:]' '[:lower:]')
        echo "      peer0."$org1"."$domain":
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true " >> $file
    else
        j=1;
        while [ $j -le $a_num ]
        do
            org1=$(echo $channel1 | cut -d'&' -f"$j" | tr '[:upper:]' '[:lower:]')
            echo "      peer0."$org1".$domain:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true
" >> $file
            j=$(( j+1 ));
        done
        #echo "$a_num" "Line 298" ; 
    fi 

done < $channel_file

echo "organizations:" >> $file
while read org
do
org1=$( echo $org | tr '[:upper:]' '[:lower:]' )
file1=$( ls `pwd`/blockchain-middleware/artifacts/crypto-config/peerOrganizations/"$org1".$domain/users/Admin@"$org1.$domain"/msp/keystore/ | grep -vi "pem" )
echo "  $org:
    mspid: "$org"MSP
    peers:
      - peer0."$org1".$domain
    certificateAuthorities:
      - ca-"$org1"
    adminPrivateKey:
      path: artifacts/crypto-config/peerOrganizations/"$org1".$domain/users/Admin@"$org1.$domain"/msp/keystore/$file1
    signedCert:
      path: artifacts/crypto-config/peerOrganizations/"$org1".$domain/users/Admin@$org1.$domain/msp/signcerts/Admin@$org1.$domain-cert.pem
" >> $file

done < $org_file

echo "orderers:
  orderer."$domain":
    url: grpcs://orderer."$domain":7050
    grpcOptions:
      ssl-target-name-override: orderer.$domain
    tlsCACerts:
      path: artifacts/crypto-config/ordererOrganizations/$domain/orderers/orderer.$domain/tls/ca.crt " >>$file

i=2
port=8050
while [ $i -le $orderer_no ]
do
    echo "
  orderer"$i"."$domain":
    url: grpcs://orderer$i."$domain":$port
    grpcOptions:
      ssl-target-name-override: orderer$i.$domain
    tlsCACerts:
      path: artifacts/crypto-config/ordererOrganizations/$domain/orderers/orderer$i.$domain/tls/ca.crt " >>$file
    i=$(( i +1 ))
    port=$(( port + 1000 ))
done
echo "peers:" >> $file
port=7051
while read org
do
org1=$( echo $org | tr '[:upper:]' '[:lower:]' )
echo "  peer0."$org1"."$domain":
    url: grpcs://peer0."$org1".$domain:"$port"
    eventUrl: grpcs://peer0."$org1".$domain:"$((port + 2))"
    grpcOptions:
      ssl-target-name-override: peer0.$org1.$domain
    tlsCACerts:
      path: artifacts/crypto-config/peerOrganizations/"$org1".$domain/peers/peer0.$org1.$domain/tls/ca.crt
" >> $file
port=$(( port + 1000))
done < $org_file

echo "certificateAuthorities:" >> $file
port=7054
while read org 
do
org1=$( echo $org | tr '[:upper:]' '[:lower:]' )
echo "  ca-"$org1":
    url: https://ca-"$org1":$port
    httpOptions:
      verify: false
    tlsCACerts:
      path: artifacts/crypto-config/peerOrganizations/"$org1"."$domain"/ca/ca."$org1"."$domain"-cert.pem
    registrar:
      - enrollId: admin
        enrollSecret: adminpw
    caName: ca-"$org1"
" >> $file
port=$(( port + 1000))
done < $org_file

while read org
do
  org1=$(echo $org | tr '[:upper:]' '[:lower:]')
  cp -ra $(pwd)/blockchain-middleware/artifacts/org1.yaml $(pwd)/blockchain-middleware/artifacts/$org1.yaml
  sed -i "s/org1/$org1/g" $(pwd)/blockchain-middleware/artifacts/$org1.yaml
  sed -i "s/Org1/$org/g" $(pwd)/blockchain-middleware/artifacts/$org1.yaml
  if [[ "$(cat `pwd`/blockchain-middleware/config.js | grep -i "$org1")" == " " ]]
  then
    echo "hfc.setConfigSetting('"$org"-connection-profile-path',path.join(__dirname, 'artifacts', '"$org1".yaml')); " >> $(pwd)/blockchain-middleware/config.js 
  fi
done < $org_file


# while read channel ### pending in instantiate-chaincode.js file
# do
#   channel1=$(echo $channel | tr '[:upper:]' '[:lower:]')


# done



