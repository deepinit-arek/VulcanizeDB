package main

import (
	"fmt"

	"flag"

	"github.com/vulcanize/vulcanizedb/cmd"
	"github.com/vulcanize/vulcanizedb/pkg/blockchain_listener"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
	"github.com/vulcanize/vulcanizedb/pkg/observers"
)

func main() {
	environment := flag.String("environment", "", "Environment name")
	flag.Parse()
	config := cmd.LoadConfig(*environment)
	fmt.Printf("Creating Geth Blockchain to: %s\n", config.Client.IPCPath)
	blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
	repository := cmd.LoadPostgres(config.Database, blockchain.Node())
	listener := blockchain_listener.NewBlockchainListener(
		blockchain,
		[]core.BlockchainObserver{
			observers.BlockchainLoggingObserver{},
			observers.NewBlockchainDbObserver(repository),
		},
	)
	listener.Start()
}
