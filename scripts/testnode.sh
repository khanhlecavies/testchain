KEY="mykey"
CHAINID="dev-chain"
MONIKER="testchain"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon
rm -rf ~/.testchain*

testchaind config keyring-backend $KEYRING
testchaind config chain-id $CHAINID

echo "taste shoot adapt slow truly grape gift need suggest midnight burger horn whisper hat vast aspect exit scorpion jewel axis great area awful blind" | testchaind keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

# if $KEY exists it should be deleted
testchaind init $MONIKER --chain-id $CHAINID 


# Allocate genesis accounts (testchaind formatted addresses)
testchaind genesis add-genesis-account $KEY 10000000000000000000stake --keyring-backend $KEYRING

# Sign genesis transaction cosmos1g5r2vmnp6lta9cpst4lzc4syy3kcj2lj0nuhmy
testchaind genesis gentx $KEY 100000000000000000stake --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
testchaind genesis collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
testchaind vgenesis alidate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# # Start the node (remove the --pruning=nothing flag if historical queries are not needed)
# testchaind start --pruning=nothing  --minimum-gas-prices=0stake 