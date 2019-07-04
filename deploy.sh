#!/bin/bash
. ./config.sh
i=0
while read channel
do
    echo "Generating Token"
    if [[ $(echo $channel | grep '&') == '' ]]
    then
        org=$(echo $org | awk -F'channel' '{print $2}')
    else
        org=$(echo $org | awk -F 'channel' '{print $2}' | awk -F'&' '{print $1}')
    fi
    token=$(curl -s -X POST http://localhost:4000/users -H "content-type: application/x-www-form-urlencoded" -d "username=Jim&orgName=$org"| awk -F'token":' '{print $2}' | cut -d'"' -f2)
    channel1=$( echo $channel | tr '[:upper:]' '[:lower:]' | sed 's/&/and/g')
    echo "curl -s -X POST http://localhost:4000/channels -H "authorization: Bearer $token" -H "content-type: application/json" -d %27{"channelName":"$channel1","channelConfigPath":"../artifacts/channel-artifacts/channel"$i".tx"}%27"
    
    
    #echo $channel
    i=$(( i + 1 ))
done < $channel_file

function join_channel
{
while read channel 
do
    if [[ $(echo $channel | grep '&') == '' ]]
    then
        org=$(echo $org | awk -F'channel' '{print $2}')
    else
        org=$(echo $org | awk -F 'channel' '{print $2}' | awk -F'&' '{print $1}')
    fi
    token=$(curl -s -X POST http://localhost:4000/users -H "content-type: application/x-www-form-urlencoded" -d "username=Jim&orgName=$org"| awk -F'token":' '{print $2}' | cut -d'"' -f2)
    channel1=$( echo $channel | tr '[:upper:]' '[:lower:]' | sed 's/&/and/g')
    org1=$( echo $org | tr '[:upper:]' '[:lower:]')
    curl -s -X POST http://localhost:4000/channels/$channel1/peers -H "authorization: Bearer $token" -H "content-type: application/json" -d %27{"peers": ["peer0.$org1.$domain"}%27

done < $channel_file

}
join_channel;

function install_chaincode
{
while read channel
do
    while read org
    do
        echo $channel1 | grep -i "$org"
        if [[ "$?" == 0 ]]
        then
            token=$(curl -s -X POST http://localhost:4000/users -H "content-type: application/x-www-form-urlencoded" -d "username=Jim&orgName=$org"| awk -F'token":' '{print $2}' | cut -d'"' -f2)
            channel1=$( echo $channel | tr '[:upper:]' '[:lower:]' | sed 's/&/and/g')
            org1=$( echo $org | tr '[:upper:]' '[:lower:]' )
            curl -s -X POST http://localhost:4000/chaincodes -H "authorization: Bearer $token" -H "content-type: application/json" -d %27{"peers": ["peer0."$org1".$domain"],"chaincodeName":"nova","chaincodePath":"github.com/example_cc/go","chaincodeType": "golang","chaincodeVersion":"v0"}%27
        fi
    done < $org_file
done < $channel_file
}
install_chaincode;

function instantiate
{
while read channel
do
    while read org
    do
        echo $channel1 | grep -i "$org"
        if [[ "$?" == 0 ]]
        then
            token=$(curl -s -X POST http://localhost:4000/users -H "content-type: application/x-www-form-urlencoded" -d "username=Jim&orgName=$org"| awk -F'token":' '{print $2}' | cut -d'"' -f2)
            channel1=$( echo $channel | tr '[:upper:]' '[:lower:]' | sed 's/&/and/g')
            org1=$( echo $org | tr '[:upper:]' '[:lower:]' )
            curl -s -X POST http://localhost:4000/channels/"$channel1"/chaincodes -H "authorization: Bearer $token" -H "content-type: application/json" -d '{"chaincodeName":"nova","chaincodeVersion":"v0","chaincodeType": "golang","args":["a","100","b","200"]}'            
        fi
    done < $org_file
done < $channel_file
}
instantiate;

echo "HERE we done the Network Setup............"
