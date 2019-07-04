#!/bin/bash
. ./config.sh

function novadir {

hosts=$1
echo "if [ ! -d "/opt/nova/hyperledger/fabric" ]
then
    sudo mkdir -p /opt/nova/hyperledger/fabric
    echo "creating directory: /opt/nova" 
    sudo chown -R $user:$user /opt/nova
else
    echo "Recreating the /opt/nova directory"
    sudo rm -fr /opt/nova/hyperledger/fabric
    sudo mkdir -p /opt/nova/hyperledger/fabric
    sudo chown -R $user:$user /opt/nova
fi
" >mkdir.sh
ssh $user@"$hosts" 'bash -s' < mkdir.sh

rm -rf mkdir.sh
}

function copy_files()
{
    for host in $( cat $hosts_file | awk -F' ' '{print $1}' | sed '/^$/d' | uniq)
    do
        if [[ $( echo $deploy | tr '[:upper:]' '[:lower:]' ) == 'yes' ]] 
        then
            novadir "$host";
            echo "$host from copy"
            scp -r channel-artifacts crypto-config peer-base.yaml $(ls order*.yaml) $(cat $org_file)  "$user"@"$host":/opt/nova/hyperledger/fabric/. 
            ssh "$user"@"$host" 'ls -lthr /opt/nova/hyperledger/fabric/'
        else
            tar -zcvf network_setup.tar crypto-config channel-artifacts $(ls order*.yaml) fabric-ca.yaml peer-base.yaml $(cat $org_file) blockchain-middleware
            mv network_setup.tar /home/$user/.
            echo -e "\n\nCopy the network.tar file from the /home/"$user" Directory........."
            echo "If you want to deploy the network please mention 'yes' it in config file."
        fi
    done < $hosts_file
}
copy_files

function deployment
{
    if [[ $deploy == 'yes' ]]
    then
        for hosts in $( cat $hosts_file | awk -F' ' '{print $1}')
        do
            echo "2"
            for host in $( cat $hosts_file | grep -i "$hosts" | awk -F' ' '{print $2}')
            do
                
                # echo "$hosts"
                if [[ $(echo $host | grep -i 'peer0' ) != ' ' ]]
                then
                    org=$(echo $host | grep -i 'peer0' | awk -F'.' '{print $2}');
                    ssh "$user"@"$hosts" "docker-compose -f /opt/nova/hyperledger/fabric/"$host".yaml up -d "
                    ssh "$user"@"$hosts" "docker-compose -f /opt/nova/hyperledger/fabric/"couchdb0$org".yaml up -d "
                else    
                    ssh "$user"@"$hosts" "docker-compose -f /opt/nova/hyperledger/fabric/"$host".yaml up -d "
                fi
            done
        done
    fi
}
deployment;

#su - fabricusr bash -c 'pm2 restart app'
