rm -R crypto-config/*

/home/paperspace/fabric/fabric-samples/bin/cryptogen generate --config=crypto-config.yaml

rm config/
/home/paperspace/fabric/fabric-samples/bin/configtxgen -profile VLMOrgOrdererGenesis -outputBlock ./config/genesis.block

/home/paperspace/fabric/fabric-samples/bin/configtxgen -profile VLMOrgChannel -outputCreateChannelTx ./config/vlmchannel.tx -channelID vlmchannel
