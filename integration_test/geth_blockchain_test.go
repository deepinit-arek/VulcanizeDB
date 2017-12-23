package integration_test

import (
	"io/ioutil"
	"log"

	"github.com/vulcanize/vulcanizedb/pkg/blockchain_listener"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

var _ = Describe("Reading from the Geth blockchain", func() {

	var listener blockchain_listener.BlockchainListener
	var observer *fakes.BlockchainObserver
	var blockchain *geth.GethBlockchain

	BeforeEach(func() {
		observer = fakes.NewFakeBlockchainObserver()
		cfg, _ := config.NewConfig("private")
		blockchain = geth.NewGethBlockchain(cfg.Client.IPCPath)
		observers := []core.BlockchainObserver{observer}
		listener = blockchain_listener.NewBlockchainListener(blockchain, observers)
	})

	AfterEach(func() {
		listener.Stop()
	})

	It("reads two blocks", func(done Done) {
		go listener.Start()

		<-observer.WasNotified
		firstBlock := observer.LastBlock()
		Expect(firstBlock).NotTo(BeNil())

		<-observer.WasNotified
		secondBlock := observer.LastBlock()
		Expect(secondBlock).NotTo(BeNil())

		Expect(firstBlock.Number + 1).Should(Equal(secondBlock.Number))

		close(done)
	}, 15)

	It("retrieves the genesis block and first block", func(done Done) {
		genesisBlock := blockchain.GetBlockByNumber(int64(0))
		firstBlock := blockchain.GetBlockByNumber(int64(1))
		lastBlockNumber := blockchain.LastBlock()

		Expect(genesisBlock.Number).To(Equal(int64(0)))
		Expect(firstBlock.Number).To(Equal(int64(1)))
		Expect(lastBlockNumber.Int64()).To(BeNumerically(">", 0))
		close(done)
	}, 15)

	It("retrieves the node info", func(done Done) {
		node := blockchain.Node()
		devNetworkGenesisBlock := "0xe5be92145a301820111f91866566e3e99ee344d155569e4556a39bc71238f3bc"
		devNetworkNodeId := float64(1)

		Expect(node.GenesisBlock).To(Equal(devNetworkGenesisBlock))
		Expect(node.NetworkId).To(Equal(devNetworkNodeId))

		close(done)
	}, 15)

})
