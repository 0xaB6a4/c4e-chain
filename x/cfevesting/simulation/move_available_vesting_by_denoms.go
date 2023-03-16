package simulation

import (
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"math/rand"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMoveAvailableVestingByDenoms(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allVestingAccounts := k.GetAllAccountVestingPools(ctx)
		if len(allVestingAccounts) == 0 {
			return simtypes.NewOperationMsg(&types.MsgMoveAvailableVestingByDenoms{}, false, "", nil), nil, nil
		}
		randInt := helpers.RandomInt(r, len(allVestingAccounts))
		accAddress := allVestingAccounts[randInt].Owner
		simAccount2Address := testcosmos.CreateRandomAccAddressNoBalance(randInt)
		msgSplitVesting := &types.MsgMoveAvailableVestingByDenoms{
			FromAddress: accAddress,
			ToAddress:   simAccount2Address,
			Denoms:      []string{testenv.DefaultTestDenom},
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.MoveAvailableVestingByDenoms(msgServerCtx, msgSplitVesting)

		if err != nil {
			if err != nil {
				k.Logger(ctx).Error("SIMULATION: Move available vesting be denoms error", err.Error())
			}

			return simtypes.NewOperationMsg(msgSplitVesting, false, "", nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Move available vesting be denoms - FINISHED")
		return simtypes.NewOperationMsg(msgSplitVesting, true, "", nil), nil, nil
	}
}
