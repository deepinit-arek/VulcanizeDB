package observers_test

import (
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/observers"
	"github.com/vulcanize/vulcanizedb/pkg/repositories"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Saving blocks to the database", func() {

	var repository *repositories.InMemory

	BeforeEach(func() {
		repository = repositories.NewInMemory()
	})

	It("implements the observer interface", func() {
		var observer core.BlockchainObserver = observers.NewBlockchainDbObserver(repository)
		Expect(observer).NotTo(BeNil())
	})

	It("saves a block with one transaction", func() {
		block := core.Block{
			Number:       123,
			Transactions: []core.Transaction{{}},
		}

		observer := observers.NewBlockchainDbObserver(repository)
		observer.NotifyBlockAdded(block)

		savedBlock, err := repository.FindBlockByNumber(123)
		Expect(err).ToNot(HaveOccurred())
		Expect(len(savedBlock.Transactions)).To(Equal(1))
	})

})
