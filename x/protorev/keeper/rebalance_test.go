package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
	"github.com/osmosis-labs/osmosis/v13/x/protorev/types"
)

func (suite *KeeperTestSuite) TestFindMaxProfitRoute() {

	type param struct {
		route          gammtypes.SwapAmountInRoutes
		expectedAmtIn  sdk.Int
		expectedProfit sdk.Int
	}

	tests := []struct {
		name       string
		param      param
		expectPass bool
	}{
		{name: "Mainnet Arb Route - 2 Asset, Same Weights (Block: 5905150)",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        22,
						TokenOutDenom: "ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        23,
						TokenOutDenom: "ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        24,
						TokenOutDenom: "uosmo",
					}},
				expectedAmtIn:  sdk.NewInt(10100000),
				expectedProfit: sdk.NewInt(24852)},
			expectPass: true},
		{name: "Mainnet Arb Route - Multi Asset, Same Weights (Block: 6906570)",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        26,
						TokenOutDenom: "ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        28,
						TokenOutDenom: "ibc/D189335C6E4A68B513C10AB227BF1C1D38C746766278BA3EEB4FB14124F1D858",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        27,
						TokenOutDenom: "uosmo",
					}},
				expectedAmtIn:  sdk.NewInt(4800000),
				expectedProfit: sdk.NewInt(4547)},
			expectPass: true},
		{name: "Arb Route - Multi Asset, Same Weights - Pool 22 instead of 26 (Block: 6906570)",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        22,
						TokenOutDenom: "ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        28,
						TokenOutDenom: "ibc/D189335C6E4A68B513C10AB227BF1C1D38C746766278BA3EEB4FB14124F1D858",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        27,
						TokenOutDenom: "uosmo",
					}},
				expectedAmtIn:  sdk.NewInt(519700000),
				expectedProfit: sdk.NewInt(67511701)},
			expectPass: true},
		{name: "Mainnet Arb Route - Multi Asset, Different Weights (Block: 6908256)",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        31,
						TokenOutDenom: "ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        32,
						TokenOutDenom: "ibc/A0CC0CF735BFB30E730C70019D4218A1244FF383503FF7579C9201AB93CA9293",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        33,
						TokenOutDenom: "ibc/D189335C6E4A68B513C10AB227BF1C1D38C746766278BA3EEB4FB14124F1D858",
					}},
				expectedAmtIn:  sdk.NewInt(4100000),
				expectedProfit: sdk.NewInt(5826)},
			expectPass: true},
		{name: "StableSwap Test Route",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        29,
						TokenOutDenom: "usdc",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        34,
						TokenOutDenom: "busd",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        30,
						TokenOutDenom: "uosmo",
					}},
				expectedAmtIn:  sdk.NewInt(137600000),
				expectedProfit: sdk.NewInt(56585438)},
			expectPass: true},
		{name: "No Arbitrage Opportunity",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        7,
						TokenOutDenom: "akash",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        12,
						TokenOutDenom: "juno",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        8,
						TokenOutDenom: "uosmo",
					}},
				expectedAmtIn:  sdk.NewInt(0),
				expectedProfit: sdk.NewInt(0)},
			expectPass: false},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {

			amtIn, profit, err := suite.App.ProtoRevKeeper.FindMaxProfitForRoute(
				suite.Ctx,
				test.param.route,
				test.param.route[2].TokenOutDenom,
			)

			if test.expectPass {
				suite.Require().NoError(err)
				suite.Require().Equal(test.param.expectedAmtIn, amtIn.Amount)
				suite.Require().Equal(test.param.expectedProfit, profit)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestExecuteTrade() {

	type param struct {
		route          gammtypes.SwapAmountInRoutes
		inputCoin      sdk.Coin
		expectedProfit sdk.Int
	}

	tests := []struct {
		name       string
		param      param
		poolId     uint64
		arbDenom   string
		expectPass bool
	}{
		{
			name: "Mainnet Arb Route",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        22,
						TokenOutDenom: "ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        23,
						TokenOutDenom: "ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        24,
						TokenOutDenom: "uosmo",
					},
				},
				inputCoin:      sdk.NewCoin("uosmo", sdk.NewInt(10100000)),
				expectedProfit: sdk.NewInt(24852),
			},
			poolId:     23,
			arbDenom:   types.OsmosisDenomination,
			expectPass: true,
		},
		{
			name: "No arbitrage opportunity - expect error at multihopswap due to profitability invariant",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        7,
						TokenOutDenom: "akash",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        12,
						TokenOutDenom: "juno",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        8,
						TokenOutDenom: "uosmo",
					}},
				inputCoin:      sdk.NewCoin("uosmo", sdk.NewInt(1000000)),
				expectedProfit: sdk.NewInt(0),
			},
			poolId:     12,
			arbDenom:   types.OsmosisDenomination,
			expectPass: false,
		},
		{
			name: "0 input amount - expect error at multihopswap due to amount needing to be positive",
			param: param{
				route: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        7,
						TokenOutDenom: "akash",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        12,
						TokenOutDenom: "juno",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        8,
						TokenOutDenom: "uosmo",
					}},
				inputCoin:      sdk.NewCoin("uosmo", sdk.NewInt(0)),
				expectedProfit: sdk.NewInt(0),
			},
			poolId:     12,
			arbDenom:   types.OsmosisDenomination,
			expectPass: false,
		},
	}

	for _, test := range tests {

		err := suite.App.ProtoRevKeeper.ExecuteTrade(
			suite.Ctx,
			test.param.route,
			test.param.inputCoin,
			test.poolId,
		)

		if test.expectPass {
			suite.Require().NoError(err)

			// Check the protorev statistics
			numberOfTrades, err := suite.App.ProtoRevKeeper.GetTradesByRoute(suite.Ctx, test.param.route.PoolIds())
			suite.Require().NoError(err)
			suite.Require().Equal(sdk.OneInt(), numberOfTrades)

			routeProfit, err := suite.App.ProtoRevKeeper.GetProfitsByRoute(suite.Ctx, test.param.route.PoolIds(), test.arbDenom)
			suite.Require().NoError(err)
			suite.Require().Equal(test.param.expectedProfit, routeProfit.Amount)

			profit, err := suite.App.ProtoRevKeeper.GetProfitsByDenom(suite.Ctx, test.arbDenom)
			suite.Require().NoError(err)
			suite.Require().Equal(test.param.expectedProfit, profit.Amount)

			totalNumberOfTrades, err := suite.App.ProtoRevKeeper.GetNumberOfTrades(suite.Ctx)
			suite.Require().NoError(err)
			suite.Require().Equal(sdk.OneInt(), totalNumberOfTrades)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestIterateRoutes() {
	type paramm struct {
		routes                     []gammtypes.SwapAmountInRoutes
		expectedMaxProfitAmount    sdk.Int
		expectedMaxProfitInputCoin sdk.Coin
		expectedOptimalRoute       gammtypes.SwapAmountInRoutes

		arbDenom string
	}

	tests := []struct {
		name       string
		params     paramm
		expectPass bool
	}{
		{name: "Single Route Test",
			params: paramm{
				routes: []gammtypes.SwapAmountInRoutes{
					{
						gammtypes.SwapAmountInRoute{
							PoolId:        22,
							TokenOutDenom: "ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC",
						},
						gammtypes.SwapAmountInRoute{
							PoolId:        23,
							TokenOutDenom: "ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0",
						},
						gammtypes.SwapAmountInRoute{
							PoolId:        24,
							TokenOutDenom: "uosmo",
						},
					},
				},
				expectedMaxProfitAmount:    sdk.NewInt(24852),
				expectedMaxProfitInputCoin: sdk.NewCoin("uosmo", sdk.NewInt(10100000)),
				expectedOptimalRoute: gammtypes.SwapAmountInRoutes{
					gammtypes.SwapAmountInRoute{
						PoolId:        22,
						TokenOutDenom: "ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        23,
						TokenOutDenom: "ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0",
					},
					gammtypes.SwapAmountInRoute{
						PoolId:        24,
						TokenOutDenom: "uosmo",
					}},
				arbDenom: types.OsmosisDenomination,
			},
			expectPass: true,
		},
	}

	for _, test := range tests {
		maxProfitInputCoin, maxProfitAmount, optimalRoute := suite.App.ProtoRevKeeper.IterateRoutes(suite.Ctx, test.params.routes)

		if test.expectPass {
			suite.Require().Equal(test.params.expectedMaxProfitAmount, maxProfitAmount)
			suite.Require().Equal(test.params.expectedMaxProfitInputCoin, maxProfitInputCoin)
			suite.Require().Equal(test.params.expectedOptimalRoute, optimalRoute)
		}
	}
}

// Test logic that compares proftability of routes with different assets
func (suite *KeeperTestSuite) TestConvertProfits() {
	type param struct {
		inputCoin           sdk.Coin
		profit              sdk.Int
		expectedUosmoProfit sdk.Int
	}

	tests := []struct {
		name       string
		param      param
		expectPass bool
	}{
		{name: "Convert atom to uosmo",
			param: param{
				inputCoin:           sdk.NewCoin(types.AtomDenomination, sdk.NewInt(100)),
				profit:              sdk.NewInt(10),
				expectedUosmoProfit: sdk.NewInt(8),
			},
			expectPass: true,
		},
		{name: "Convert juno to uosmo (random denom)",
			param: param{
				inputCoin:           sdk.NewCoin("juno", sdk.NewInt(100)),
				profit:              sdk.NewInt(10),
				expectedUosmoProfit: sdk.NewInt(9),
			},
			expectPass: true,
		},
		{name: "Convert denom without pool to uosmo",
			param: param{
				inputCoin:           sdk.NewCoin("random", sdk.NewInt(100)),
				profit:              sdk.NewInt(10),
				expectedUosmoProfit: sdk.NewInt(10),
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		profit, err := suite.App.ProtoRevKeeper.ConvertProfits(suite.Ctx, test.param.inputCoin, test.param.profit)

		if test.expectPass {
			suite.Require().NoError(err)
			suite.Require().Equal(test.param.expectedUosmoProfit, profit)
		} else {
			suite.Require().Error(err)
		}
	}
}
