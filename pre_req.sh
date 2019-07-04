#!/bin/bash
. ./config.sh

function requirement()
{
	os_type=$( cat /etc/issue | grep -i 'ubuntu\|centos\|'| tr '[:upper:]' '[:lower:]' | awk -F' ' '{print $1}')

	if [[ $os_type == 'ubuntu' ]]
	then
		echo "proceeding for the $os_type";
		ubuntu_requirement;
	elif [[ $os_type == 'centos' ]]
	then
		echo "proceeding for the $os_type";
		centos_requirement;
	elif [[ $os_type == 'redhat' ]]
	then
		echo "proceeding for the $os_type";
		centos_requirement;
	else
		echo "OS-type not compatible for this type of setup......." ;
		exit 1; 
	fi
}


function ubuntu_requirement()
{
	touch setup.sh
	echo '#!/bin/bash
	which git
	if [ "$?" == 0 ]
	then
		echo "git is alredy installed......"
	else
		sudo apt-get update
		sudo apt-get install -y git
	fi
	sleep 10 
	which docker
	if [ "$?" == 0 ]
	then
		echo "The docker is already installed......!"
	else
		sudo apt-get update
		sudo apt-get install -y curl
		sudo apt install apt-transport-https ca-certificates curl software-properties-common
		curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
		sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
		sudo apt-get update
		sudo apt-get install docker -y
		sudo systemctl restart docker;
		sudo usermod -aG docker $user;
		systemctl restart docker;
		which docker-compose
		if [ "$?" == 0 ]
		then
			echo "Docker-compose already installed..........!!"
		else
			sudo apt-get update
			a=`sudo apt-cache policy docker-compose | head -1 | awk -F':' '{print $1}'`;
			if [ $a == "docker-compose" ]
			then
				sudo apt-get install -y docker-compose
			fi
		fi
	fi' >setup.sh
}

function centos_requirement()
{
	echo '#!/bin/bash
	which git
	if [ "$?" == 0 ]
	then
		echo "git is alredy installed......"
	else
		sudo yum update
		sudo yum install -y git
	fi
	sleep 10 
	which docker
	if [ "$?" == 0 ]
	then
		echo "The docker is already installed......!"
	else
		sudo yum update
		sudo yum install -y curl
		sudo yum install ca-certificates curl
		sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
		sudo yum install docker-ce
		sudo yum update
		sudo systemctl restart docker;
		sudo usermod -aG docker $user;
		systemctl restart docker;
		which docker-compose
		if [ "$?" == 0 ]
		then
			echo "Docker-compose already installed..........!!"
		else
			curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
			sudo chmod +x /usr/local/bin/docker-compose
		fi
	fi' >setup.sh

}


function exe_req()
{
for i in $(cat $hosts_file | awk -F' ' '{print $1}' | sed '/^$/d' | uniq )
do
	host=$(echo $i)
	requirement;
	ssh $user@$host 'bash -s' < setup.sh;
done
}

exe_req;
