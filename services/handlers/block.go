package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/tendermint/tmlibs/common"
)

// queryBlock is to query a block by height
func queryBlock(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	height := args["height"]

	c := commands.GetNode()
	h := cast.ToInt64(height)
	block, err := c.Block(&h)
	if err != nil {
		common.WriteError(w, err)
		return
	}
	if err := printResult(w, block); err != nil {
		common.WriteError(w, err)
	}
}

// queryValidators is to query validators by height
func queryValidators(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	height := args["height"]

	c := commands.GetNode()
	h := cast.ToInt64(height)
	block, err := c.Validators(&h)
	if err != nil {
		common.WriteError(w, err)
		return
	}
	if err := printResult(w, block); err != nil {
		common.WriteError(w, err)
	}
}

// queryRecentBlocks is to query recent 20 blocks
func queryRecentBlocks(w http.ResponseWriter, r *http.Request) {
	c := commands.GetNode()
	blocks, err := c.BlockchainInfo(0, 0)
	if err != nil {
		common.WriteError(w, err)
		return
	}
	if err := printResult(w, blocks); err != nil {
		common.WriteError(w, err)
	}
}

// mux.Router registrars

func RegisterQueryBlock(r *mux.Router) error {
	r.HandleFunc("/block/{height}", queryBlock).Methods("GET")
	return nil
}

func RegisterQueryValidators(r *mux.Router) error {
	r.HandleFunc("/validators/{height}", queryValidators).Methods("GET")
	return nil
}

func RegisterQueryRecentBlocks(r *mux.Router) error {
	r.HandleFunc("/blocks/recent", queryRecentBlocks).Methods("GET")
	return nil
}

func RegisterBlock(r *mux.Router) error {
	funcs := []func(*mux.Router) error{
		RegisterQueryBlock,
		RegisterQueryValidators,
		RegisterQueryRecentBlocks,
	}

	for _, fn := range funcs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

// End of mux.Router registrars
