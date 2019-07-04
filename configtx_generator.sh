#!/bin/bash
 . ./config.sh
# cd /home/$user/.
# num=7051;
	echo "---
Organizations:

    - &OrdererOrg
        Name: OrdererOrg
        ID: OrdererMSP
        MSPDir: crypto-config/ordererOrganizations/$domain/msp

        Policies:
            Readers:
                Type: Signature
                Rule: \"OR('OrdererMSP.member')\"
            Writers:
                Type: Signature
                Rule: \"OR('OrdererMSP.member')\"
            Admins:
                Type: Signature
                Rule: \"OR('OrdererMSP.admin')\" "> configtx.yaml

port=7051
while read org
do
org1=$(echo $org | tr '[:upper:]' '[:lower:]')
echo "    - &"$org"
        Name: "$org"MSP

        # ID to load the MSP definition as
        ID: "$org"MSP

        MSPDir: crypto-config/peerOrganizations/"$org1".$domain/msp
        Policies:
            Readers:
                Type: Signature
                Rule: \"OR('"$org"MSP.admin', '"$org"MSP.peer', '"$org"MSP.client')\"
            Writers:
                Type: Signature
                Rule: \"OR('"$org"MSP.admin', '"$org"MSP.client')\"
            Admins:
                Type: Signature
                Rule: \"OR('"$org"MSP.admin')\"

        # leave this flag set to true.
        AnchorPeers:
            - Host: peer0.$org1."$domain"
              Port: $port
" >> configtx.yaml
port=$((port+1000));
done < $org_file
# sed -i 's/OR(/"OR(/g' configtx.yaml
# sed -i 's/client)/client)"/g' configtx.yaml
# sed -i 's/admin)/admin)"/g' configtx.yaml
# sed -i 's/member)/member)"/g' configtx.yaml

echo "################################################################################
Capabilities:
    Channel: &ChannelCapabilities
        V1_3: true
    Orderer: &OrdererCapabilities
        V1_1: true

    Application: &ApplicationCapabilities
        V1_3: true
        V1_2: false
        V1_1: false

################################################################################
################################################################################
Application: &ApplicationDefaults
    Organizations:
    Policies:
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"

    Capabilities:
        <<: *ApplicationCapabilities
################################################################################

" >> configtx.yaml
port=7050
echo "################################################################################
Orderer: &OrdererDefaults

    # Orderer Type: The orderer implementation to start
    # Available types are "solo" and "kafka"
    OrdererType: solo

    Addresses:
        - orderer."$domain":"$port" " >> configtx.yaml
i=2
port=8050;
#echo "Orderer-no- line143" $orderer_no
while [ $i -le $orderer_no ]
do
    echo "        - orderer"$i"."$domain":"$port" " >> configtx.yaml
    i=$(( i + 1 ))
    port=$(( port + 1000 ))
done 

echo '    # Batch Timeout: The amount of time to wait before creating a batch
    BatchTimeout: 2s
    BatchSize:
        MaxMessageCount: 10
        AbsoluteMaxBytes: 99 MB
        PreferredMaxBytes: 512 KB
    Kafka:
        Brokers:
            - 127.0.0.1:9092
    Organizations:
    Policies:
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
        BlockValidation:
            Type: ImplicitMeta
            Rule: "ANY Writers"


################################################################################
################################################################################
Channel: &ChannelDefaults
    Policies:
        # Who may invoke the 'Deliver' API
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        # Who may invoke the 'Broadcast' API
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        # By default, who may modify elements at this config level
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
    Capabilities:
        <<: *ChannelCapabilities

################################################################################
' >> configtx.yaml

port=8050
echo "################################################################################
Profiles:

    SampleMultiNodeEtcdRaft:
        <<: *ChannelDefaults
        Capabilities:
            <<: *ChannelCapabilities
        Orderer:
            <<: *OrdererDefaults
            OrdererType: etcdraft
            EtcdRaft:
                Consenters:
                - Host: orderer."$domain"
                  Port: "$(( port -1000))"
                  ClientTLSCert: crypto-config/ordererOrganizations/"$domain"/orderers/orderer."$domain"/tls/server.crt
                  ServerTLSCert: crypto-config/ordererOrganizations/"$domain"/orderers/orderer."$domain"/tls/server.crt " >> configtx.yaml
i=2
while [ $i -le $orderer_no ]
do
echo "                - Host: orderer"$i"."$domain"
                  Port: "$port"
                  ClientTLSCert: crypto-config/ordererOrganizations/$domain/orderers/orderer"$i"."$domain"/tls/server.crt
                  ServerTLSCert: crypto-config/ordererOrganizations/$domain/orderers/orderer"$i"."$domain"/tls/server.crt " >> configtx.yaml 
    i=$(( i + 1 ))
    port=$(( port + 1000 ))
done
port=8050
echo "            Addresses:
                - orderer."$domain":"$(( port - 1000 ))" " >> configtx.yaml
i=2
while [ $i -le $orderer_no ]
do
    echo "                - orderer"$i"."$domain":"$port" " >> configtx.yaml
    i=$(( i + 1 ))
    port=$(( port + 1000 ))
done                 


echo "            Organizations:
            - *OrdererOrg
            Capabilities:
                <<: *OrdererCapabilities
        Application:
            <<: *ApplicationDefaults
            Organizations:
            - <<: *OrdererOrg " >> configtx.yaml 
echo "        Consortiums: " >> configtx.yaml

i=0
while read channel  ##### for the Consortium
do
channel1=$(echo $channel | awk -F'channel' '{print $2}')
a_num=$(echo $channel1 | awk -F'&' '{print NF}')
echo "            SampleConsortium"$i":
                Organizations:" >> configtx.yaml

if [[ $(echo $channel1 | grep -i '&' ) == ' ' ]]
then
    echo "                - *"$channel1" " >> configtx.yaml
    echo 'One Org Line 290';
else
    j=1;
    while [ $j -le $a_num ]
    do
        org1=$(echo $channel1 | cut -d'&' -f"$j")
        echo "                - *"$org1" " >> configtx.yaml
    j=$(( j+1 ));
    done
    #echo "$a_num" "Line 298" ; 
fi
    i=$(( i+1 ));
done < $channel_file 
i=0
while read channel   ##### For the Channel Profile 
do 
echo "    TwoOrgsChannel"$i": " >> configtx.yaml
channel1=$(echo $channel | awk -F'channel' '{print $2}')
a_num=$(echo $channel1 | awk -F'&' '{print NF}')
echo "        Consortium: SampleConsortium"$i"
        <<: *ChannelDefaults
        Application:
            <<: *ApplicationDefaults
            Organizations:" >> configtx.yaml 

    if [[ $(echo $channel1 | grep -i '&' ) == ' ' ]]
    then
        echo "                - *"$channel1" " >> configtx.yaml
    else
        j=1;
        while [ $j -le $a_num ]
        do
            org1=$(echo $channel1 | cut -d'&' -f"$j")
            echo "                - *"$org1" " >> configtx.yaml
            j=$(( j+1 ));
        done
    fi  

i=$(( i + 1 ));
echo "            Capabilities:
                <<: *ApplicationCapabilities " >> configtx.yaml
done < $channel_file
