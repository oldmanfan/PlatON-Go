package pposm

import (
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"math/big"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
)

type CandidatePoolContext struct {
	configs *params.PposConfig
}

var cContext *CandidatePoolContext

// Initialize the global candidate pool context object
func NewCandidatePoolContext(configs *params.PposConfig) *CandidatePoolContext {
	cContext = &CandidatePoolContext{
		configs: configs,
	}
	return cContext
}

func GetCandidateContextPtr() *CandidatePoolContext {
	return cContext
}

func (c *CandidatePoolContext)initCandidatePool () *CandidatePool {
	return NewCandidatePool(c.configs)
}



func (c *CandidatePoolContext) SetCandidate(state vm.StateDB, nodeId discover.NodeID, can *types.Candidate) error {
	return c.initCandidatePool().SetCandidate(state, nodeId, can)
}

func (c *CandidatePoolContext) GetCandidate (state vm.StateDB, nodeId discover.NodeID) (*types.Candidate, error) {
	return c.initCandidatePool().GetCandidate(state, nodeId)
}

func (c *CandidatePoolContext) GetCandidateArr (state vm.StateDB, nodeIds ...discover.NodeID) (types.CandidateQueue, error) {
	return c.initCandidatePool().GetCandidateArr(state, nodeIds...)
}

func (c *CandidatePoolContext) WithdrawCandidate (state vm.StateDB, nodeId discover.NodeID, price, blockNumber *big.Int) error {
	return c.initCandidatePool().WithdrawCandidate(state, nodeId, price, blockNumber)
}

func (c *CandidatePoolContext) GetChosens (state vm.StateDB) []*types.Candidate {
	return c.initCandidatePool().GetChosens(state)
}

func (c *CandidatePoolContext) GetChairpersons(state vm.StateDB) []*types.Candidate {
	return c.initCandidatePool().GetChairpersons(state)
}

func (c *CandidatePoolContext) GetDefeat(state vm.StateDB, nodeId discover.NodeID) ([]*types.Candidate, error) {
	return c.initCandidatePool().GetDefeat(state, nodeId)
}

func (c *CandidatePoolContext) IsDefeat(state vm.StateDB, nodeId discover.NodeID) (bool, error) {
	return c.initCandidatePool().IsDefeat(state, nodeId)
}

func (c *CandidatePoolContext) RefundBalance(state vm.StateDB, nodeId discover.NodeID, blockNumber *big.Int) error {
	return c.initCandidatePool().RefundBalance(state, nodeId, blockNumber)
}

func (c *CandidatePoolContext) GetOwner(state vm.StateDB, nodeId discover.NodeID) common.Address {
	return c.initCandidatePool().GetOwner(state, nodeId)
}

func (c *CandidatePoolContext) GetRefundInterval() uint64 {
	return c.initCandidatePool().GetRefundInterval()
}

func (c *CandidatePoolContext) MaxCount() uint64 {
	return c.initCandidatePool().MaxCount()
}

func (c *CandidatePoolContext) Election(state *state.StateDB, blocknumber *big.Int) ([]*discover.Node, error) {
	return c.initCandidatePool().Election(state)
}

func (c *CandidatePoolContext)  Switch(state *state.StateDB) bool {
	return c.initCandidatePool().Switch(state)
}

func (c *CandidatePoolContext) GetWitness(state *state.StateDB, flag int) ([]*discover.Node, error) {
	return c.initCandidatePool().GetWitness(state, flag)
}

func (c *CandidatePoolContext) GetAllWitness(state *state.StateDB) ([]*discover.Node, []*discover.Node, []*discover.Node, error) {
	return c.initCandidatePool().GetAllWitness(state)
}

func (c *CandidatePoolContext) SetCandidateExtra (state vm.StateDB, nodeId discover.NodeID, extra string) error {
	return c.initCandidatePool().SetCandidateExtra(state, nodeId, extra)
}
