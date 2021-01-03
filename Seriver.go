package main

import (
	"BitCoinConnect1214/entity"
	"encoding/json"
)

func GetBlockCount()interface{}{
	rpcBytes:=PrepareJson(GETBLOCKCOUNT)
	return SendPost(rpcBytes)
}

func GetDifficulty() interface{}  {
	rpcBytes:=PrepareJson(GETDIFFICULTY)
	return SendPost(rpcBytes)
}

func GetBestBlockHash() interface{} {
	rpcBytes:=PrepareJson(GETBESTBLOCKHASH)
	return SendPost(rpcBytes)
}

func GetBlockChainInfo() entity.BlockChainInfo {
	rpcBytes:=PrepareJson(GETBLOCKCHAININFO)

	result:=SendPost(rpcBytes)

	result1,err:=json.Marshal(result)
	result2:=entity.BlockChainInfo{}
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(result1,&result2)
	if err != nil {
		panic(err)
	}
	return result2
}

func GetBlockHash(height int) interface{} {
	rpcBytes:=PrepareJson(GETBLOCKHASH,height)
	return SendPost(rpcBytes)
}

func GetNewAddress(label string,address_type string)interface{}{
	rpcBytes:=PrepareJson(GETNEWADDRESS,label,address_type)
	return SendPost(rpcBytes)
}

func GetConnectionCount() interface{}{
	rpcBytes:=PrepareJson(GETCONNECTIONCOUNT)
	return SendPost(rpcBytes)
}

func GetBlockByHash(hashNum string) interface{} {
	rpcBytes:=PrepareJson(GETBLOCK,hashNum)
	return SendPost(rpcBytes)
}
