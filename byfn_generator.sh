#!/bin/bash
. ./config.sh
echo "#!/bin/bash
. ./config.sh
if [ ! -d channel-artifacts ] 
then
    mkdir channel-artifacts
elif [ -d crypto-config ]
then 
    rm -fr crypto-config ;
    rm -rf cannel-artifacts/* ;
else  
    echo "NULL"
fi " > prod_setup.sh
echo 'export PATH=${PWD}/bin:${PWD}:$PATH
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
    echo 'Failed to generate certificates...'
    exit 1
  fi
  echo
}' >> prod_setup.sh

echo "function generateChannelArtifacts() {
  which configtxgen
  if [ "$?" -ne 0 ]; then
    echo "configtxgen tool not found. exiting"
    exit 1
  fi
  echo '#########  Generating Orderer Genesis block ##############'
  if [ orderer_type == "raft" ]; then
    configtxgen -profile SampleMultiNodeEtcdRaft -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
  else
    echo "unrecognized CONSESUS_TYPE='sCONSENSUS_TYPE'. exiting"
    exit 1
  fi
  if [ "$?" -ne 0 ]; then
    echo "Failed to generate orderer genesis block..."
    exit 1
  fi " >> prod_setup.sh
i=0  
while read channel  ##### for channel 
do
channelid=$(echo $channel | sed 's/&/and/g' | tr '[:upper:]' '[:lower:]')
echo "configtxgen -profile TwoOrgsChannel"$i" -outputCreateChannelTx ./channel-artifacts/channel"$i".tx -channelID $channelid
if [ "$?" -ne 0 ]; then
echo "Failed to generate channel configuration transaction..."
exit 1
fi ">> prod_setup.sh
i=$(( i + 1 ));
done < $channel_file

echo "}" >> prod_setup.sh 

#Create the network using docker compose

echo 'if [ "generate" == "generate" ]; then ## Generate Artifacts
  generateCerts
  generateChannelArtifacts
elif [ "${MODE}" == "upgrade" ]; then ## Upgrade the network from version 1.2.x to 1.3.x
  upgradeNetwork
else
  printHelp
  exit 1
fi ' >> prod_setup.sh

sed -i 's/\[ 0/\[ "$?"/g' prod_setup.sh
sed -i 's/orderer_type/$orderer_type/g' prod_setup.sh
