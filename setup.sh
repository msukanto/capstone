echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=RtaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rta.vlm.com/users/Admin@rta.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.rta.vlm.com:7051" cli peer channel create -o orderer.vlm.com:7050 -c vlmchannel -f /etc/hyperledger/configtx/vlmchannel.tx

sleep 5

echo "Channel genesis block created."

echo "peer0.manufacturer.vlm.com joining the channel..."
# Join peer0.manufacturer.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlm.com/users/Admin@manufacturer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlm.com:7051" cli peer channel join -b vlmchannel.block

echo "peer0.manufacturer.vlm.com joined the channel"

echo "peer0.dealer.vlm.com joining the channel..."

# Join peer0.dealer.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=DealerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dealer.vlm.com/users/Admin@dealer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.dealer.vlm.com:7051" cli peer channel join -b vlmchannel.block

echo "peer0.dealer.vlm.com joined the channel"

echo "peer0.rta.vlm.com joining the channel..."
# Join peer0.rta.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=RtaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rta.vlm.com/users/Admin@rta.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.rta.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5

echo "peer0.rta.vlm.com joined the channel"


echo "Installing vlm chaincode to peer0.manufacturer.vlm.com..."
# install chaincode
# Install code on manufacturer peer
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlm.com/users/Admin@manufacturer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlm.com:7051" cli peer chaincode install -n vlm -v 1.0 -p github.com/vlm/go/ -l golang
echo "Installed vlm chaincode to peer0.manufacturer.vlm.com"

echo "Installing vlm chaincode to peer0.dealer.vlm.com...."
# Install code on dealer peer
docker exec -e "CORE_PEER_LOCALMSPID=DealerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dealer.vlm.com/users/Admin@dealer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.dealer.vlm.com:7051" cli peer chaincode install -n vlm -v 1.0 -p github.com/vlm/go/ -l golang
echo "Installed vlm chaincode to peer0.dealer.vlm.com"

echo "Installing vlm chaincode to peer0.rta.vlm.com..."
# Install code on rta peer
docker exec -e "CORE_PEER_LOCALMSPID=RtaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rta.vlm.com/users/Admin@rta.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.rta.vlm.com:7051" cli peer chaincode install -n vlm -v 1.0 -p github.com/vlm/go -l golang

sleep 5

echo "Installed vlm chaincode to peer0.dealer.vlm.com"

echo "Instantiating vlm chaincode.."
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlm.com/users/Admin@manufacturer.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlm.com:7051" cli peer chaincode instantiate -o orderer.vlm.com:7050 -C vlmchannel -n vlm -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('ManufacturerMSP.member','DealerMSP.member','RtaMSP.member')"

echo "Instantiated vlm chaincode."
