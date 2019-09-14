package janken

import (
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

const appName = "janken"

var (
	// DefaultCLIHome default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.jankencli")

	// DefaultNodeHome the folder where the applcation data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.jankend")

	// ModuleBasics this is in charge of setting up basic module elemnets
	ModuleBasics = module.NewBasicManager()
)

type jankenApp struct {
	*baseapp.BaseApp
	cbc *codec.Codec
}

// NewJankenApp initialize an app instance
func NewJankenApp(logger log.Logger, db dbm.DB) *jankenApp {
	cbc := MakeCodec()

	bApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cbc))

	var app = &jankenApp{
		BaseApp: bApp,
		cbc:     cbc,
	}

	return app
}

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
