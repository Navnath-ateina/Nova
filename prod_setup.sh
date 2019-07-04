#!/bin/bash
. ./config.sh
if [ ! -d channel-artifacts ] 
then
    mkdir channel-artifacts
elif [ -d crypto-config ]
then 
    rm -fr crypto-config ;
    rm -rf cannel-artifacts/* ;
else  
    echo NULL
fi 
export PATH=${PWD}/bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}
export VERBOSE=false

# Print the usage message
function printHelp() {
  echo "	byfn.sh generate"
}

function generateCerts() {
  which cryptogen
  if [ "$?" -ne 0 ]; then
    echo "cryptogen tool not found. exiting"
    exit 1
  fi
  echo
  echo "##### Generate certificates using cryptogen tool #########"
  if [ -d "crypto-config" ]; then
    rm -Rf crypto-config
  fi
  cryptogen generate --config=./crypto-config.yaml
  if [ "$?" -ne 0 ]; then
    echo Failed to generate certificates...
    exit 1
  fi
  echo
}
function generateChannelArtifacts() {
  which configtxgen
  if [ "$?" -ne 0 ]; then
    echo configtxgen tool not found. exiting
    exit 1
  fi
  echo '#########  Generating Orderer Genesis block ##############'
  if [ $orderer_type == raft ]; then
    configtxgen -profile SampleMultiNodeEtcdRaft -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
  else
    echo unrecognized CONSESUS_TYPE=sCONSENSUS_TYPE. exiting
    exit 1
  fi
  if [ "$?" -ne 0 ]; then
    echo Failed to generate orderer genesis block...
    exit 1
  fi 
configtxgen -profile TwoOrgsChannel0 -outputCreateChannelTx ./channel-artifacts/channel0.tx -channelID channelyblandvendorandanchor
if [ "$?" -ne 0 ]; then
echo Failed to generate channel configuration transaction...
exit 1
fi 
configtxgen -profile TwoOrgsChannel1 -outputCreateChannelTx ./channel-artifacts/channel1.tx -channelID channelybl
if [ "$?" -ne 0 ]; then
echo Failed to generate channel configuration transaction...
exit 1
fi 
}
if [ "generate" == "generate" ]; then ## Generate Artifacts
  generateCerts
  generateChannelArtifacts
elif [ "${MODE}" == "upgrade" ]; then ## Upgrade the network from version 1.2.x to 1.3.x
  upgradeNetwork
else
  printHelp
  exit 1
fi 
