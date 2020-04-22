package main

import (
	contract "constellation-swap/contracts"
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ethAPI := os.Getenv("ETH_API_URL")
	tokenAddress := common.HexToAddress("0xa8258abc8f2811dd48eccd209db68f25e3e34667")
	swapAddress := common.HexToAddress("0x947Bc555CD04c2a9035c3E2EB084938BE3EC1673")

	client, err := ethclient.Dial(ethAPI)
	checkError("Cannot create an ETH client", err)

	tokenInstance, err := contract.NewToken(tokenAddress, client)
	checkError("Cannot get Constellation DAG token smart contract instance", err)

	swapInstance, err := contract.NewSwap(swapAddress, client)
	checkError("Cannot get Constellation Swap smart contract instance", err)

	ethAddresses, err := swapInstance.GetKeyList(nil)
	checkError("Cannot load swap ETH addresses", err)

	dagAddresses, err := swapInstance.GetMappedAddresses(nil)
	checkError("Cannot load swap mapped addresses", err)

	fmt.Println("Received", len(ethAddresses), "mapped addresses")

	file, err := os.Create("result.csv")
	checkError("Cannot create csv file", err)
	defer file.Close()

	wrongFile, err := os.Create("wrong.csv")
	checkError("Cannot create csv file", err)
	defer wrongFile.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	wrongWriter := csv.NewWriter(wrongFile)
	defer wrongWriter.Flush()

	wrongDAGAddress := 0

	for i, address := range ethAddresses {
		dagAddress := dagAddresses[i]
		if len(dagAddress) > 30 && strings.HasPrefix(dagAddress, "DAG") {
			balance, _ := tokenInstance.BalanceOf(nil, address)
			err := writer.Write([]string{dagAddress, balance.String()})
			checkError("Cannot write to csv file", err)
		} else {
			wrongDAGAddress++
			err := wrongWriter.Write([]string{address.Hex(), dagAddress})
			checkError("Cannot write to csv file", err)
		}
	}

	fmt.Println("Received", wrongDAGAddress, "wrong DAG addresses")
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}