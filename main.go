package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}

func getBlockHeaderByNum(client *rpc.Client, args ...interface{}) (*ethtypes.Header, error) {
	var raw json.RawMessage
	err := client.CallContext(context.Background(), &raw, "eth_getBlockByNumber", args...)
	if err != nil {
		return nil, err
	} else if len(raw) == 0 {
		return nil, ethereum.NotFound
	}
	// Decode header and transactions.
	var body ethtypes.Header
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	return &body, nil
}

func main() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	header1, err := client.HeaderByNumber(ctx, big.NewInt(int64(1)))
	if err != nil {
		panic(err)
	}
	fmt.Println("unmodified header1 hash: ", header1.Hash())
	header1.Coinbase = common.HexToAddress("0xd512181b3dfa6f819cb6c6ae64732542e2fb6002")
	fmt.Println("coinbase header1 hash: ", header1.Hash())
	header1.TxHash = types.EmptyRootHash
	fmt.Println("TxHash header1 hash: ", header1.Hash())
	header1.BaseFee = big.NewInt(0)
	header1JSON, err := json.MarshalIndent(header1, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("header1 %s\n", string(header1JSON))
	fmt.Println("modified header1 hash: ", header1.Hash())
	fmt.Println("header1 parent hash: ", header1.ParentHash)

	header2, err := client.HeaderByNumber(ctx, big.NewInt(int64(8)))
	if err != nil {
		panic(err)
	}
	fmt.Println("unmodified header2 hash: ", header2.Hash())
	header2.Coinbase = common.HexToAddress("0xd512181b3dfa6f819cb6c6ae64732542e2fb6002")
	fmt.Println("coinbase header2 hash: ", header2.Hash())
	// header2.TxHash = types.EmptyRootHash
	// fmt.Println("TxHash header2 hash: ", header2.Hash())
	header2.GasUsed = 0
	fmt.Println("GasUsed header2 hash: ", header2.Hash())
	header2.Bloom = ethtypes.BytesToBloom([]byte("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	fmt.Println("Bloom header2 hash: ", header2.Hash())
	// header2.BaseFee = big.NewInt(0)
	header2JSON, err := json.MarshalIndent(header2, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("header2 %s\n", string(header2JSON))
	fmt.Println("modified header2 hash: ", header2.Hash())
	fmt.Println("header2 parent hash: ", header2.ParentHash)

}
