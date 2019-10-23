#!/bin/bash
org_file=/home/dell/git/1.4.1/org.txt
hosts_file=/home/dell/git/1.4.1/hosts.txt
channel_file=/home/dell/git/1.4.1/channel.txt
path1=/home/dell/git/1.4.1
user="$USER";               ## set the common user name on each host
domain="yesbank.in"         ## set the domain
peer_type='solo'            ## set solo or bi
couchdb_required='yes'      ## set yes or no
orderer_type='raft'         ## set solo or raft
orderer_no='3'              ## set the number of orderers WRT number of hosts
#type='multi'
image="1.4.1"
deploy="no"                 ## only conf files