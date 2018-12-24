# fabric-multiorg-multipeer-network setup (without TLS) ... In progress

Clone this repo to a directory. cd to that repo directory [fabric-multiorg-multipeer-network] and do the necessary prerequisites as following.

Download the prerequisites : fabric, fabric-ca and dependent third party (couchdb, kafka etc.) docker images and the binaries. Add additional parameter [-d] if you want to skip docker image download or [-b] to skip binaries downlaod

     curl -sSL http://bit.ly/2ysbOFE | bash -s <fabric version> <fabric-ca version> <third-party version> -s
     for example: curl -sSL http://bit.ly/2ysbOFE | bash -s 1.4.0-rc1 1.4.0-rc1 0.4.13 -s

Inside your root directory a [bin] and [config] folder will be created. [bin] folder contains necessary binaries to create your blockchain network and client. [config] contains few sample yml file. 
It will take some time for the first time if docker images are not yet downloaded locally. 


## Network setup

     cd basic-network

Run following shell script to generate crypto and confic files. It will create two folders [config] and [crypto-config]. Refer [configtx.yaml] and [crypto-config.yaml] files for configuration setup. Make necessary changes to make your own organization if you wish.

     ./generate.sh

Open [docker-compose.yaml] file and [FABRIC_CA_SERVER_CA_KEYFILE] tag in ca1.example.com, ca2.example.com and ca3.example.com
and replace with the private key as following respectively:

     crypto-config/peerOrganizations/org1.example.com/ca/xxxxxxxx_sk

Automation script for replacing this in progress...
Similarly replace for other orgs

This network contains:

     Org1
          - peer
          - couchdb
          - ca
          - orderer
     Org2
          - peer
          - couchdb
          - ca
     Org3
          - peer
          - couchdb
          - ca

Once above configuration is done then run following shell script to the network and create channel and join peer to that channel

     ./start.sh

## Chaincode deployment

Chaincode is written in [golang] and in [chaincode] directory. Check [volumes] section of [cli] service in [docker-compose.yaml] file for chaincode reference. Run fillowing command to install, instantiate and invoke [purchaseorderchaincode] and [shipmentchaincode] in mychannel1 and mychannel2 respectively.

     cd ../fabric-example
     ./deploy-chaincode.sh



