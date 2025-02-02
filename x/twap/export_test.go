package twap

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v13/x/twap/types"
)

type (
	TimeTooOldError = timeTooOldError
	TwapType        = twapType
)

const (
	ArithmeticTwapType = arithmeticTwapType
	GeometricTwapType  = geometricTwapType
)

var GeometricTwapMathBase = geometricTwapMathBase

func (k Keeper) StoreNewRecord(ctx sdk.Context, record types.TwapRecord) {
	k.storeNewRecord(ctx, record)
}

func (k Keeper) GetMostRecentRecordStoreRepresentation(ctx sdk.Context, poolId uint64, asset0Denom string, asset1Denom string) (types.TwapRecord, error) {
	return k.getMostRecentRecordStoreRepresentation(ctx, poolId, asset0Denom, asset1Denom)
}

func (k Keeper) GetAllMostRecentRecordsForPool(ctx sdk.Context, poolId uint64) ([]types.TwapRecord, error) {
	return k.getAllMostRecentRecordsForPool(ctx, poolId)
}

func (k Keeper) GetRecordAtOrBeforeTime(ctx sdk.Context, poolId uint64, time time.Time, asset0Denom string, asset1Denom string) (types.TwapRecord, error) {
	return k.getRecordAtOrBeforeTime(ctx, poolId, time, asset0Denom, asset1Denom)
}

func (k Keeper) GetAllHistoricalTimeIndexedTWAPs(ctx sdk.Context) ([]types.TwapRecord, error) {
	return k.getAllHistoricalTimeIndexedTWAPs(ctx)
}

func (k Keeper) GetAllHistoricalPoolIndexedTWAPs(ctx sdk.Context) ([]types.TwapRecord, error) {
	return k.getAllHistoricalPoolIndexedTWAPs(ctx)
}

func (k Keeper) TrackChangedPool(ctx sdk.Context, poolId uint64) {
	k.trackChangedPool(ctx, poolId)
}

func (k Keeper) GetChangedPools(ctx sdk.Context) []uint64 {
	return k.getChangedPools(ctx)
}

func (k Keeper) UpdateRecord(ctx sdk.Context, record types.TwapRecord) types.TwapRecord {
	return k.updateRecord(ctx, record)
}

func (k Keeper) UpdateRecords(ctx sdk.Context, poolId uint64) error {
	return k.updateRecords(ctx, poolId)
}

func (k Keeper) PruneRecordsBeforeTimeButNewest(ctx sdk.Context, lastKeptTime time.Time) error {
	return k.pruneRecordsBeforeTimeButNewest(ctx, lastKeptTime)
}

func (k Keeper) PruneRecords(ctx sdk.Context) error {
	return k.pruneRecords(ctx)
}

func (k Keeper) GetInterpolatedRecord(ctx sdk.Context, poolId uint64, asset0Denom string, asset1Denom string, t time.Time) (types.TwapRecord, error) {
	return k.getInterpolatedRecord(ctx, poolId, t, asset0Denom, asset1Denom)
}

func ComputeTwap(startRecord types.TwapRecord, endRecord types.TwapRecord, quoteAsset string, twapType twapType) (sdk.Dec, error) {
	return computeTwap(startRecord, endRecord, quoteAsset, twapType)
}

func ComputeArithmeticTwap(startRecord types.TwapRecord, endRecord types.TwapRecord, quoteAsset string) sdk.Dec {
	return computeArithmeticTwap(startRecord, endRecord, quoteAsset)
}

func ComputeGeometricTwap(startRecord types.TwapRecord, endRecord types.TwapRecord, quoteAsset string) sdk.Dec {
	return computeGeometricTwap(startRecord, endRecord, quoteAsset)
}

func RecordWithUpdatedAccumulators(record types.TwapRecord, t time.Time) types.TwapRecord {
	return recordWithUpdatedAccumulators(record, t)
}

func NewTwapRecord(k types.AmmInterface, ctx sdk.Context, poolId uint64, denom0, denom1 string) (types.TwapRecord, error) {
	return newTwapRecord(k, ctx, poolId, denom0, denom1)
}

func TwapLog(x sdk.Dec) sdk.Dec {
	return twapLog(x)
}

func TwapPow(x sdk.Dec) sdk.Dec {
	return twapPow(x)
}

func GetSpotPrices(
	ctx sdk.Context,
	k types.AmmInterface,
	poolId uint64,
	denom0, denom1 string,
	previousErrorTime time.Time,
) (sp0 sdk.Dec, sp1 sdk.Dec, latestErrTime time.Time) {
	return getSpotPrices(ctx, k, poolId, denom0, denom1, previousErrorTime)
}

func (k *Keeper) GetAmmInterface() types.AmmInterface {
	return k.ammkeeper
}

func (k *Keeper) SetAmmInterface(ammInterface types.AmmInterface) {
	k.ammkeeper = ammInterface
}

func (k *Keeper) AfterCreatePool(ctx sdk.Context, poolId uint64) error {
	return k.afterCreatePool(ctx, poolId)
}
