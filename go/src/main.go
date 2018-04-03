package main 

import (
	"log"
	"os"
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"math/big"
	"./eth_router"
)

var addresses []common.Address

func main() {

	file, err := os.Open("./addresses.txt")
	if err != nil {
		log.Fatal("error reading file ", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		addresses = append(addresses, common.HexToAddress(string(scanner.Text())))
	}
	file.Close()
	

	client, err := ethclient.Dial("/home/solidity/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal("error connecting to network ", err)
	} else {
		fmt.Println(client)
	}

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("please enter raw json file")
	scanner.Scan()
	key := scanner.Text()
	fmt.Println(key)

	fmt.Println("enter password to decrypt key")
	scanner.Scan()
	pass := scanner.Text()

	auth, err := bind.NewTransactor(strings.NewReader(key), string(pass))
	if err != nil {
		log.Fatal("error authenticating with network")
	} else {
		fmt.Println(auth)
	}

	fmt.Println("please enter address of eth router contract")
	scanner.Scan()
	contractAddress := common.HexToAddress(string(scanner.Text()))

	ethRouter, err := EthRouter.NewEthRouter(contractAddress, client)
	auth.Value = big.NewInt(3750000000000000000)
	tx, err := ethRouter.SendEth(auth, addresses, big.NewInt(150000000000000000))
	if err != nil {
		log.Fatal("error submitting transaction ", err)
	} else {
		fmt.Println("Transaction hash 0x%x\n", tx.Hash())
	}
}


