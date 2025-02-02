package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	epochtypes "github.com/osmosis-labs/osmosis/v13/x/epochs/types"
	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
)

// AccountKeeper defines the account contract that must be fulfilled when
// creating a x/protorev keeper.
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// BankKeeper defines the banking contract that must be fulfilled when
// creating a x/protorev keeper.
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
}

// GAMMKeeper defines the Gamm contract that must be fulfilled when
// creating a x/protorev keeper.
type GAMMKeeper interface {
	GetPoolAndPoke(ctx sdk.Context, poolId uint64) (gammtypes.CFMMPoolI, error)
	GetPoolsAndPoke(ctx sdk.Context) (res []gammtypes.CFMMPoolI, err error)
	GetPoolDenoms(ctx sdk.Context, poolId uint64) ([]string, error)
	SwapExactAmountIn(ctx sdk.Context, sender sdk.AccAddress, poolId uint64, tokenIn sdk.Coin, tokenOutDenom string, tokenOutMinAmount sdk.Int) (sdk.Int, error)
	MultihopSwapExactAmountIn(ctx sdk.Context, sender sdk.AccAddress, routes []gammtypes.SwapAmountInRoute, tokenIn sdk.Coin, tokenOutMinAmount sdk.Int) (tokenOutAmount sdk.Int, err error)
}

// EpochKeeper defines the Epoch contract that must be fulfilled when
// creating a x/protorev keeper.
type EpochKeeper interface {
	GetEpochInfo(ctx sdk.Context, identifier string) epochtypes.EpochInfo
}
