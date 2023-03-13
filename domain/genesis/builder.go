package genesis

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	difficulty     uint
	reward         uint
	halving        uint
	pMiningValue   *uint8
	maxTrxPerBlock uint
	blockDuration  time.Duration
	pCreatedOn     *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		difficulty:     0,
		reward:         0,
		halving:        0,
		pMiningValue:   nil,
		maxTrxPerBlock: 0,
		blockDuration:  0,
		pCreatedOn:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithDifficulty adds a difficulty to the builder
func (app *builder) WithDifficulty(difficulty uint) Builder {
	app.difficulty = difficulty
	return app
}

// WithReward adds a reward to the builder
func (app *builder) WithReward(reward uint) Builder {
	app.reward = reward
	return app
}

// WithHalving adds an halving to the builder
func (app *builder) WithHalving(halving uint) Builder {
	app.halving = halving
	return app
}

// WithMiningValue adds a miningValue to the builder
func (app *builder) WithMiningValue(miningValue uint8) Builder {
	app.pMiningValue = &miningValue
	return app
}

// WithMaxTrxPerBlock adds a maxTrxPerBlock to the builder
func (app *builder) WithMaxTrxPerBlock(maxTrxPerBlock uint) Builder {
	app.maxTrxPerBlock = maxTrxPerBlock
	return app
}

// WithBlockDuration adds a blockDuration to the builder
func (app *builder) WithBlockDuration(blockDuration time.Duration) Builder {
	app.blockDuration = blockDuration
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.difficulty <= 0 {
		return nil, errors.New("the difficulty must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.reward <= 0 {
		return nil, errors.New("the reward must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.halving <= 0 {
		return nil, errors.New("the halving must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.pMiningValue == nil {
		return nil, errors.New("the miningValue is mandatory in order to build a Genesis instance")
	}

	if *app.pMiningValue > maxMiningValue {
		str := fmt.Sprintf("the miningValue must be an integer between zero (0) and nine (%d), %d provided", maxMiningValue, *app.pMiningValue)
		return nil, errors.New(str)
	}

	if app.maxTrxPerBlock <= 0 {
		return nil, errors.New("the maxTrxPerBlock must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.blockDuration <= 0 {
		return nil, errors.New("the blockDuration must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the miningValue is mandatory in order to build a Genesis instance")
	}

	if app.pCreatedOn.Unix() <= 0 {
		return nil, errors.New("the creationTime's timestamp must be greater than zero (0) in order to build a Genesis instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%d", app.difficulty)),
		[]byte(fmt.Sprintf("%d", app.reward)),
		[]byte(fmt.Sprintf("%d", app.halving)),
		[]byte(fmt.Sprintf("%d", *app.pMiningValue)),
		[]byte(fmt.Sprintf("%d", app.maxTrxPerBlock)),
		[]byte(fmt.Sprintf("%d", app.blockDuration)),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
	})

	if err != nil {
		return nil, err
	}

	return createGenesis(
		*pHash,
		app.difficulty,
		app.reward,
		app.halving,
		*app.pMiningValue,
		app.maxTrxPerBlock,
		app.blockDuration,
		*app.pCreatedOn,
	), nil
}
