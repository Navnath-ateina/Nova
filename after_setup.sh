#!/bin/bash
docker exec -ti orderer2.yesbank.in bash -c "sed -i s/8443/8445/g /etc/hyperledger/fabric/orderer.yaml"
docker restart orderer2.yesbank.in
docker exec -ti orderer1.yesbank.in bash -c "sed -i s/8443/8444/g /etc/hyperledger/fabric/orderer.yaml"
docker restart orderer1.yesbank.in
docker restart orderer0.yesbank.in
docker exec -ti peer0.vendor.yesbank.in bash -c "sed -i s/9443/9445/g /etc/hyperledger/fabric/core.yaml"
docker restart peer0.vendor.yesbank.in

docker exec -ti peer0.anchir.yesbank.in bash -c "sed -i s/9443/9444/g /etc/hyperledger/fabric/core.yaml"
docker restart peer0.anchor.yesbank.in
docker restart peer0.ybl.yesbank.in
docker exec -ti ca-ybl bash -c "sed -i s/org1/ybl/ $path1";
docker exec -ti ca-anchor bash -c "sed -i s/org1/anchor/g $path1";
docker exec -ti ca-vendor bash -c "sed -i s/org1/anchor/g $path1";
docker restart ca-ybl ca-anchor ca-vendor