echo "Installing vlm chaincode to peer0.manufacturer.vlm.com..."

# install chaincode
# Install code on manufacturer peer
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlm.com/users/Admin@manufacturer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlm.com:7051" cli peer chaincode install -n vlm -v 1.4 -p github.com/vlm -l golang

echo "Installed vlm chaincode to peer0.manufacturer.vlm.com"

echo "Installing vlm chaincode to peer0.dealer.vlm.com...."

# Install code on dealer peer
docker exec -e "CORE_PEER_LOCALMSPID=DealerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dealer.vlm.com/users/Admin@dealer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.dealer.vlm.com:7051" cli peer chaincode install -n vlm -v 1.4 -p github.com/vlm -l golang

echo "Installed vlm chaincode to peer0.dealer.vlm.com"

echo "Installing vlm chaincode to peer0.rta.vlm.com..."
# Install code on rta peer
docker exec -e "CORE_PEER_LOCALMSPID=RtaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rta.vlm.com/users/Admin@rta.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.rta.vlm.com:7051" cli peer chaincode install -n vlm -v 1.4 -p github.com/vlm -l golang

sleep 5

echo "Installed vlm chaincode to peer0.dealer.vlm.com"

echo "Instantiating vlm chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlm.com/users/Admin@manufacturer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlm.com:7051" cli peer chaincode instantiate -o orderer.vlm.com:7050 -C vlmchannel -n vlm -l golang -v 1.4 -c '{"Args":[""]}' -P "OR ('ManufacturerMSP.member','DealerMSP.member','RtaMSP.member')"

echo "Instantiated vlm chaincode."

echo "Following is the docker network....."
