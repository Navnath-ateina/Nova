#!/bin/bash
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
			a=`sudo apt-cache policy docker-compose | head -1 | awk -F: {print }`;
			if [ $a == "docker-compose" ]
			then
				sudo apt-get install -y docker-compose
			fi
		fi
	fi
