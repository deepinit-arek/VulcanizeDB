package observers

import (
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/repositories"
)

type BlockchainDbObserver struct {
	repository repositories.Repository
}

func NewBlockchainDbObserver(repository repositories.Repository) BlockchainDbObserver {
	return BlockchainDbObserver{repository: repository}
}

func (observer BlockchainDbObserver) NotifyBlockAdded(block core.Block) {
	observer.repository.CreateOrUpdateBlock(block)
}
