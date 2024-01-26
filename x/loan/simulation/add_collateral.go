package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"loan/x/loan/keeper"
	"loan/x/loan/types"
)

func SimulateMsgAddCollateral(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddCollateral{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddCollateral simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddCollateral simulation not implemented"), nil, nil
	}
}
