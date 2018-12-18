# fabric-multiorg-multipeer-network setup

Download the prerequisites : fabric, fabric-ca and dependent third party (couchdb, kafka etc.) docker images and the binaries. Add additional parameter [-d] if you want to skip docker image download or [-b] to skip binaries downlaod

     curl -sSL http://bit.ly/2ysbOFE | bash -s <fabric version> <fabric-ca version> <third-party version> -s
     for example: curl -sSL http://bit.ly/2ysbOFE | bash -s 1.4.0-rc1 1.4.0-rc1 0.4.13 -s

Inside your root directory a bin folder will be created along with the binaries. Use following command to check the downloaded docker images.

     docker images

