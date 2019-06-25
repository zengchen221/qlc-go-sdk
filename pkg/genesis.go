// +build  !testnet

/*
 * Copyright (c) 2019 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package pkg

import (
	"encoding/json"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

var (
	jsonMintage = `{
        	"type": "ContractSend",
        	"token": "45dd217cd9ff89f7b64ceda4886cc68dde9dfa47a8a422d165e2ce6f9a834fad",
        	"address": "qlc_3qjky1ptg9qkzm8iertdzrnx9btjbaea33snh1w4g395xqqczye4kgcfyfs1",
        	"balance": "0",
        	"vote": "0",
        	"network": "0",
        	"storage": "0",
        	"oracle": "0",
        	"previous": "0000000000000000000000000000000000000000000000000000000000000000",
        	"link": "681bf5253c64672fc54c93d3b5b9a20d28965cb8f80ba70460ed3f99cb547bd5",
        	"message": "0000000000000000000000000000000000000000000000000000000000000000",
        	"data": "6TrdxEXdIXzZ/4n3tkztpIhsxo3enfpHqKQi0WXizm+ag0+tAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADVKa6ehgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAhoG/UlPGRnL8VMk9O1uaINKJZcuPgLpwRg7T+Zy1R71QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADUUxDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA1FMQwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
        	"povHeight": 0,
        	"timestamp": 1553990401,
        	"extra": "0000000000000000000000000000000000000000000000000000000000000000",
        	"representative": "qlc_1t1uynkmrs597z4ns6ymppwt65baksgdjy1dnw483ubzm97oayyo38ertg44",
        	"work": "000000000048f5b9",
        	"signature": "0d09a049ae11a005648b9ea9f4d7f2ce2fed985a9ec0b7009c87970b99c4de7b3c53c44768e6c9565babd627c83971bae8a32001f20f7acf888a80a0897f8409"
        }`
	jsonGenesis = `{
        	"type": "ContractReward",
        	"token": "45dd217cd9ff89f7b64ceda4886cc68dde9dfa47a8a422d165e2ce6f9a834fad",
        	"address": "qlc_1t1uynkmrs597z4ns6ymppwt65baksgdjy1dnw483ubzm97oayyo38ertg44",
        	"balance": "60000000000000000",
        	"vote": "0",
        	"network": "0",
        	"storage": "0",
        	"oracle": "0",
        	"previous": "0000000000000000000000000000000000000000000000000000000000000000",
        	"link": "c0d330096ec4ab6ccf5481e06cc54e74b14f534e99e38df486f47d1123cbd1ae",
        	"message": "0000000000000000000000000000000000000000000000000000000000000000",
        	"data": "Rd0hfNn/ife2TO2kiGzGjd6d+keopCLRZeLOb5qDT60AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANUprp6GAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACGgb9SU8ZGcvxUyT07W5og0olly4+AunBGDtP5nLVHvVAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGgb9SU8ZGcvxUyT07W5og0olly4+AunBGDtP5nLVHvVAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANRTEMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADUUxDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        	"povHeight": 0,
        	"timestamp": 1553990410,
        	"extra": "0000000000000000000000000000000000000000000000000000000000000000",
        	"representative": "qlc_1t1uynkmrs597z4ns6ymppwt65baksgdjy1dnw483ubzm97oayyo38ertg44",
        	"work": "000000000048f5b9",
        	"signature": "76fb8d1e5decf145616ee1b5713ad81b5792cadcc054a9d999d5acf026924e8ccb23f96994cf8a056b5e0a74706e4ccb7314dc989b1fccbcd88c024bb31c6b0d"
        }`

	jsonMintageQGAS = `{
        	"type": "ContractSend",
        	"token": "ea842234e4dc5b17c33b35f99b5b86111a3af0bd8e4a8822602b866711de6d81",
        	"address": "qlc_3qjky1ptg9qkzm8iertdzrnx9btjbaea33snh1w4g395xqqczye4kgcfyfs1",
        	"balance": "0",
        	"vote": "0",
        	"network": "0",
        	"storage": "0",
        	"oracle": "0",
        	"previous": "0000000000000000000000000000000000000000000000000000000000000000",
        	"link": "dd803c6bdb9878a35a3e447e4d88b8df4f65eecd35d7506157750ea46e2c6133",
        	"message": "0000000000000000000000000000000000000000000000000000000000000000",
        	"data": "6TrdxOqEIjTk3FsXwzs1+ZtbhhEaOvC9jkqIImArhmcR3m2BAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAjhvJvwQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAjdgDxr25h4o1o+RH5NiLjfT2XuzTXXUGFXdQ6kbixhMwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEUUdBUwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABFFHQVMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
        	"povHeight": 0,
        	"timestamp": 1553990401,
        	"extra": "0000000000000000000000000000000000000000000000000000000000000000",
        	"representative": "qlc_1t1uynkmrs597z4ns6ymppwt65baksgdjy1dnw483ubzm97oayyo38ertg44",
        	"work": "000000000048f5b9",
        	"signature": "175df7302eb8d830da4d86047819b4680f5a3c425ed8bc6593a4dd724c847aceac68b63bebfbde94d368b4e24182dd503f05ac474d231c6a17455fa131f0d10f"
        }`
	jsonGenesisQGAS = `{
        	"type": "ContractReward",
        	"token": "ea842234e4dc5b17c33b35f99b5b86111a3af0bd8e4a8822602b866711de6d81",
        	"address": "qlc_3qe19joxq85rnff5wj5ybp6djqtheqqetfgqc3iogxagnjq4rrbmbp1ews7d",
        	"balance": "10000000000000000",
        	"vote": "0",
        	"network": "0",
        	"storage": "0",
        	"oracle": "0",
        	"previous": "0000000000000000000000000000000000000000000000000000000000000000",
        	"link": "bdac41b3ff7ac35aee3028d60eabeb9578ea6f7bd148d611133a3b26dfa6a9be",
        	"message": "0000000000000000000000000000000000000000000000000000000000000000",
        	"data": "6oQiNOTcWxfDOzX5m1uGERo68L2OSogiYCuGZxHebYEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACOG8m/BAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACN2APGvbmHijWj5Efk2IuN9PZe7NNddQYVd1DqRuLGEzAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAN2APGvbmHijWj5Efk2IuN9PZe7NNddQYVd1DqRuLGEzAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAARRR0FTAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEUUdBUwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        	"povHeight": 0,
        	"timestamp": 1553990410,
        	"extra": "0000000000000000000000000000000000000000000000000000000000000000",
        	"representative": "qlc_1t1uynkmrs597z4ns6ymppwt65baksgdjy1dnw483ubzm97oayyo38ertg44",
        	"work": "000000000048f5b9",
        	"signature": "4e605d4ccfe3f061c1139eecec25b1129dbe491b12c42eb34e3febadc40971ff42e7240f9f42432591172cc5a91292406fe7b3115c757945e9c8848b83836001"
        }`

	//main net
	chainToken, _       = types.NewHash("45dd217cd9ff89f7b64ceda4886cc68dde9dfa47a8a422d165e2ce6f9a834fad")
	genesisAddress, _   = types.HexToAddress("qlc_1t1uynkmrs597z4ns6ymppwt65baksgdjy1dnw483ubzm97oayyo38ertg44")
	genesisMintageBlock types.StateBlock
	genesisMintageHash  types.Hash
	genesisBlock        types.StateBlock
	genesisBlockHash    types.Hash

	//main net gas
	gasToken, _     = types.NewHash("ea842234e4dc5b17c33b35f99b5b86111a3af0bd8e4a8822602b866711de6d81")
	gasAddress, _   = types.HexToAddress("qlc_3qe19joxq85rnff5wj5ybp6djqtheqqetfgqc3iogxagnjq4rrbmbp1ews7d")
	gasMintageBlock types.StateBlock
	gasMintageHash  types.Hash
	gasBlock        types.StateBlock
	gasBlockHash    types.Hash
)

func init() {
	_ = json.Unmarshal([]byte(jsonMintage), &genesisMintageBlock)
	_ = json.Unmarshal([]byte(jsonGenesis), &genesisBlock)
	genesisMintageHash = genesisMintageBlock.GetHash()
	genesisBlockHash = genesisBlock.GetHash()
	//main net gas
	_ = json.Unmarshal([]byte(jsonMintageQGAS), &gasMintageBlock)
	_ = json.Unmarshal([]byte(jsonGenesisQGAS), &gasBlock)
	gasMintageHash = gasMintageBlock.GetHash()
	gasBlockHash = gasBlock.GetHash()
}

func GenesisAddress() types.Address {
	return genesisAddress
}

func GasAddress() types.Address {
	return gasAddress

}

func ChainToken() types.Hash {
	return chainToken
}

func GasToken() types.Hash {
	return gasToken
}

func GenesisMintageBlock() types.StateBlock {
	return genesisMintageBlock
}

func GasMintageBlock() types.StateBlock {
	return gasMintageBlock
}

func GenesisMintageHash() types.Hash {
	return genesisMintageHash
}

func GasMintageHash() types.Hash {
	return gasMintageHash
}

func GenesisBlock() types.StateBlock {
	return genesisBlock
}

func GasBlock() types.StateBlock {
	return gasBlock

}

func GenesisBlockHash() types.Hash {
	return genesisBlockHash
}

func GasBlockHash() types.Hash {
	return gasBlockHash
}

// IsGenesis check block is chain token genesis
func IsGenesisBlock(block *types.StateBlock) bool {
	h := block.GetHash()
	return h == genesisMintageHash || h == genesisBlockHash || h == gasMintageHash || h == gasBlockHash
}

// IsGenesis check token is chain token genesis
func IsGenesisToken(hash types.Hash) bool {
	return hash == chainToken || hash == gasToken
}

func AllGenesisBlocks() []types.StateBlock {
	return []types.StateBlock{genesisMintageBlock, genesisBlock, gasMintageBlock, gasBlock}
}