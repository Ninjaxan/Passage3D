package v050

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/envadiv/Passage3D/app/upgrades"
	claim "github.com/envadiv/Passage3D/x/claim/keeper"
)

// Name is the on-chain upgrade name for the Cosmos SDK v0.47 -> v0.50 migration.
const Name = "v050"

// Upgrade migrates the chain from Cosmos SDK v0.47 to v0.50.
//
// No new module stores are introduced: every store (incl. x/consensus and
// x/crisis, which v047 added) already exists, so StoreUpgrades is empty. That
// matters — an empty Added set means this upgrade does NOT re-trigger the
// IAVL store-version-divergence bug (#13477/#21087) that the v047 store
// additions exposed under iavl 0.20.1. SDK 0.50 ships iavl v1.x, whose fixed
// versioned-snapshot access is the actual cure for the "version does not
// exist" query outage shipped in Passage v3.0.0. The handler itself is just
// the per-module state migrations run by RunMigrations.
var Upgrade = upgrades.Upgrade{
	UpgradeName:          Name,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        storetypes.StoreUpgrades{},
}

// CreateUpgradeHandler returns the v0.50 upgrade handler. The migration is a
// technical SDK bump; no keeper-specific state surgery is required, so the
// keepers are unused and RunMigrations drives every module from its current
// on-chain consensus version to its v0.50 version.
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	_ distribution.Keeper,
	_ bank.Keeper,
	_ auth.AccountKeeper,
	_ claim.Keeper,
	_ consensusparamkeeper.Keeper,
	_ paramskeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
