package cbft

import "math/big"

const (
	//BaseElection = 230
	//
	//BaseSwitchWitness = 250
	//
	//BaseRemoveFormerPeers = 270

	BaseElection = 50
	//
	BaseSwitchWitness = 60
	//
	BaseAddNextPeers = 50
	//
	//BaseRemoveFormerPeers = 70
	//
	BaseIrrCount = 20

	//BaseElection = 50000

	//BaseSwitchWitness = 60000

	//BaseAddNextPeers = 50000

	//BaseRemoveFormerPeers = 700000

	//BaseIrrCount = 2000

	FirstRound = 1

)
var (
	//ppos
	InitAmount = new(big.Int).SetBytes([]byte{3, 59, 46, 60, 159, 208, 128, 60, 232, 0, 0, 0}) //1,000,000,000 * 10^18 ADP -> 0x033b2e3c9fd0803ce8000000
	FirstYearReward = new(big.Int).SetBytes([]byte{1, 90, 241, 215, 139, 88, 196, 0, 0}) //25,000,000 * 10^18 ADP -> 0x015af1d78b58c40000
	YearBlocks = new(big.Int).SetInt64(31536000)//24*3600*365
	Rate = big.NewInt(1025)//Actual value Rate/Base
	Base = big.NewInt(1000)
	FeeBase = big.NewInt(10000)
)