package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmamino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"gopkg.in/yaml.v2"

	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

var _ exported.Account = (*EthAccount)(nil)
var _ exported.GenesisAccount = (*EthAccount)(nil)

// ----------------------------------------------------------------------------
// Main Ethermint account
// ----------------------------------------------------------------------------

// ProtoAccount defines the prototype function for BaseAccount used for an
// AccountKeeper.
func ProtoAccount() exported.Account {
	return &EthAccount{
		BaseAccount:  &auth.BaseAccount{},
		CodeHash:     ethcrypto.Keccak256(nil),
		SubAddresses: map[string]sdk.AccAddress{},
	}
}

type ethermintAccountPretty struct {
	Address       sdk.AccAddress `json:"address" yaml:"address"`
	Coins         sdk.Coins      `json:"coins" yaml:"coins"`
	PubKey        []byte         `json:"public_key" yaml:"public_key"`
	AccountNumber uint64         `json:"account_number" yaml:"account_number"`
	Sequence      uint64         `json:"sequence" yaml:"sequence"`
	CodeHash      string         `json:"code_hash" yaml:"code_hash"`

	Name         string                    `json:"name" yaml:"name"`
	SubAddresses map[string]sdk.AccAddress `json:"sub_addresses" yaml:"sub_addresses"`
}

// MarshalYAML returns the YAML representation of an account.
func (acc EthAccount) MarshalYAML() (interface{}, error) {
	alias := ethermintAccountPretty{
		Address:       acc.Address,
		PubKey:        acc.PubKey,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
		CodeHash:      ethcmn.Bytes2Hex(acc.CodeHash),
		Name:          acc.Name,
		SubAddresses:  acc.SubAddresses,
	}

	bz, err := yaml.Marshal(alias)
	if err != nil {
		return nil, err
	}

	return string(bz), err
}

// MarshalJSON returns the JSON representation of an EthAccount.
func (acc EthAccount) MarshalJSON() ([]byte, error) {
	alias := ethermintAccountPretty{
		Address:       acc.Address,
		PubKey:        acc.PubKey,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
		CodeHash:      ethcmn.Bytes2Hex(acc.CodeHash),
		Name:          acc.Name,
		SubAddresses:  acc.SubAddresses,
	}

	return json.Marshal(alias)
}

// UnmarshalJSON unmarshals raw JSON bytes into an EthAccount.
func (acc *EthAccount) UnmarshalJSON(bz []byte) error {
	acc.BaseAccount = &authtypes.BaseAccount{}
	var alias ethermintAccountPretty
	if err := json.Unmarshal(bz, &alias); err != nil {
		return err
	}

	if alias.PubKey != nil {
		pubk, err := tmamino.PubKeyFromBytes(alias.PubKey)
		if err != nil {
			return err
		}

		acc.BaseAccount.PubKey = pubk.Bytes()
	}

	acc.BaseAccount.Address = alias.Address
	acc.BaseAccount.AccountNumber = alias.AccountNumber
	acc.BaseAccount.Sequence = alias.Sequence
	acc.CodeHash = ethcmn.Hex2Bytes(alias.CodeHash)
	acc.Name = alias.Name
	acc.SubAddresses = alias.SubAddresses

	return nil
}
