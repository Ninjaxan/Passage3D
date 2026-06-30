package v047

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/envadiv/Passage3D/app/upgrades"
	claim "github.com/envadiv/Passage3D/x/claim/keeper"
)

// Name is the on-chain upgrade name for the Cosmos SDK v0.45 -> v0.47 migration.
const Name = "v047"

// Upgrade migrates the chain from Cosmos SDK v0.45 to v0.47.
//
// Two new module stores are introduced in v0.47 and must be created at the
// upgrade height:
//   - x/consensus: the dedicated home for Tendermint consensus params, which
//     previously lived in the x/params "baseapp" subspace.
//   - x/crisis: the constant fee moved from the x/params subspace into the
//     module's own store.
var Upgrade = upgrades.Upgrade{
	UpgradeName:          Name,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			consensusparamtypes.StoreKey,
			crisistypes.StoreKey,
		},
	},
}

// CreateUpgradeHandler returns the v0.47 upgrade handler. The distribution,
// bank, auth and claim keepers are unused for this purely-technical SDK bump;
// the consensus-params and params keepers are required to migrate the
// Tendermint consensus params out of the legacy x/params subspace.
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	_ distribution.Keeper,
	_ bank.Keeper,
	_ auth.AccountKeeper,
	_ claim.Keeper,
	consensusParamsKeeper consensusparamkeeper.Keeper,
	paramsKeeper paramskeeper.Keeper,
	_ *stakingkeeper.Keeper,
	_ govkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// v047 already ran on mainnet (Passage v3.0.0); historical no-op under v0.50.
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
