#!/bin/bash

KEY="mykey"
CHAINID=8272
MONIKER="localtestnet"

# remove existing daemon and client
rm -rf ~/.qmint*

make install

qmintcli config keyring-backend test

# Set up config for CLI
qmintcli config chain-id $CHAINID
qmintcli config output json
qmintcli config indent true
qmintcli config trust-node true

# if $KEY exists it should be deleted
qmintcli keys add $KEY

# Set moniker and chain-id for Quantmint (Moniker can be anything, chain-id must be an integer)
qmintd init $MONIKER --chain-id $CHAINID

# Allocate genesis accounts (cosmos formatted addresses)
qmintd add-genesis-account $(qmintcli keys show $KEY -a) 1000000000000000000photon,1000000000000000000stake

# Sign genesis transaction
qmintd gentx --name $KEY --keyring-backend test

# Collect genesis tx
qmintd collect-gentxs

# Enable faucet
cat  $HOME/.qmintd/config/genesis.json | jq '.app_state["faucet"]["enable_faucet"]=true' >  $HOME/.qmintd/config/tmp_genesis.json && mv $HOME/.qmintd/config/tmp_genesis.json $HOME/.qmintd/config/genesis.json

echo -e '\n\ntestnet faucet enabled'
echo -e 'to transfer tokens to your account address use:'
echo -e "qmintcli tx faucet request 100photon --from $KEY\n"


# Run this to ensure everything worked and that the genesis file is setup correctly
qmintd validate-genesis

# Command to run the rest server in a different terminal/window
echo -e '\nrun the following command in a different terminal/window to run the REST server and JSON-RPC:'
echo -e "qmintcli rest-server --laddr \"tcp://localhost:8545\" --unlock-key $KEY --chain-id $CHAINID --trace\n"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
qmintd start --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info" --trace

