package types

import (
	"github.com/flashbots/go-boost-utils/types"
)

type Slot = uint64
type Epoch = uint64

type PublicKey = types.PublicKey

type Hash = types.Hash

type Bid = types.SignedBuilderBid

type Root = types.Root

type ValidatorIndex = uint64

type Coordinate struct {
	Slot Slot
	Root Root
}
