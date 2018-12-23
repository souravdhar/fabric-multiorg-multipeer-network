#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

docker-compose -f docker-compose.yaml down

docker-compose -f docker-compose.yaml up -d ca1.example.com ca2.example.com ca3.example.com orderer.example.com peer0.org1.example.com peer0.org2.example.com peer0.org3.example.com couchdb

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=10
#echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}

# Create the channel1
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/org1.example.com/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer.example.com:7050 -c mychannel1 -f /etc/hyperledger/configtx/channel1.tx
docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/org2.example.com/users/Admin@org2.example.com/msp" peer0.org2.example.com peer channel create -o orderer.example.com:7050 -c mychannel1 -f /etc/hyperledger/configtx/channel1.tx

# # Create the channel2
# docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/org1.example.com/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer.example.com:7050 -c mychannel2 -f /etc/hyperledger/configtx/channel2.tx


# Join peer0.org1.example.com to the channel1.
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/org1.example.com/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b mychannel1.block

# # Join peer0.org1.example.com to the channel2.
# docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/org1.example.com/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b mychannel2.block

# Join peer0.org2.example.com to the channel1.
docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/org2.example.com/users/Admin@org2.example.com/msp" peer0.org2.example.com peer channel join -b mychannel1.block

# # Join peer0.org3.example.com to the channel2.
# docker exec -e "CORE_PEER_LOCALMSPID=Org3MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/org1.example.com/users/Admin@org3.example.com/msp" peer0.org3.example.com peer channel join -b mychannel2.block