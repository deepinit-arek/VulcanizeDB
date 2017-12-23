package blockchain_listener_test

import (
	"github.com/vulcanize/vulcanizedb/pkg/blockchain_listener"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Blockchain listeners", func() {

	It("starts with no blocks", func(done Done) {
		observer := fakes.NewFakeBlockchainObserver()
		blockchain := fakes.NewBlockchain()

		blockchain_listener.NewBlockchainListener(blockchain, []core.BlockchainObserver{observer})

		Expect(len(observer.CurrentBlocks)).To(Equal(0))
		close(done)
	}, 1)

	It("sees when one block was added", func(done Done) {
		observer := fakes.NewFakeBlockchainObserver()
		blockchain := fakes.NewBlockchain()
		listener := blockchain_listener.NewBlockchainListener(blockchain, []core.BlockchainObserver{observer})
		go listener.Start()

		go blockchain.AddBlock(core.Block{Number: 123})

		wasObserverNotified := <-observer.WasNotified
		Expect(wasObserverNotified).To(BeTrue())
		Expect(len(observer.CurrentBlocks)).To(Equal(1))
		addedBlock := observer.CurrentBlocks[0]
		Expect(addedBlock.Number).To(Equal(int64(123)))
		close(done)
	}, 1)

	It("sees a second block", func(done Done) {
		observer := fakes.NewFakeBlockchainObserver()
		blockchain := fakes.NewBlockchain()
		listener := blockchain_listener.NewBlockchainListener(blockchain, []core.BlockchainObserver{observer})
		go listener.Start()

		go blockchain.AddBlock(core.Block{Number: 123})
		<-observer.WasNotified
		go blockchain.AddBlock(core.Block{Number: 456})
		wasObserverNotified := <-observer.WasNotified

		Expect(wasObserverNotified).To(BeTrue())
		Expect(len(observer.CurrentBlocks)).To(Equal(2))
		addedBlock := observer.CurrentBlocks[1]
		Expect(addedBlock.Number).To(Equal(int64(456)))
		close(done)
	}, 1)

	It("stops listening", func(done Done) {
		observer := fakes.NewFakeBlockchainObserver()
		blockchain := fakes.NewBlockchain()
		listener := blockchain_listener.NewBlockchainListener(blockchain, []core.BlockchainObserver{observer})
		go listener.Start()

		listener.Stop()

		Expect(blockchain.WasToldToStop).To(BeTrue())
		close(done)
	}, 1)

})
