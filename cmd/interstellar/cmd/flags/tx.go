package flags

import (
	"github.com/doggystylez/interstellar/types"
	"github.com/spf13/cobra"
)

func ProcessTxFlags(cmd *cobra.Command) (txInfo types.TxInfo, err error) {
	txInfo.FeeAmount, err = cmd.Flags().GetUint64("fee-amount")
	if err != nil {
		return
	}
	txInfo.FeeDenom, err = cmd.Flags().GetString("fee-denom")
	if err != nil {
		return
	}
	txInfo.Gas, err = cmd.Flags().GetUint64("gas")
	if err != nil {
		return
	}
	txInfo.Address, err = cmd.Flags().GetString("from")
	if err != nil {
		return
	}
	txInfo.Memo, err = cmd.Flags().GetString("memo")
	if err != nil {
		return
	}
	txInfo.KeyInfo.ChainId, err = cmd.Flags().GetString("chain-id")
	if err != nil {
		return
	}
	txInfo.KeyInfo.AccNum, err = cmd.Flags().GetUint64("account")
	if err != nil {
		return
	}
	txInfo.KeyInfo.SeqNum, err = cmd.Flags().GetUint64("sequence")
	if err != nil {
		return
	}
	return
}

func TxFlags(rawCmds ...*cobra.Command) (cmds []*cobra.Command) {
	for _, cmd := range rawCmds {
		cmd.Flags().StringP("chain-id", "c", "", "chain-id")
		cmd.Flags().StringP("from", "f", "", "from address")
		cmd.Flags().StringP("channel-id", "i", "", "ibc channel")
		cmd.Flags().Uint64P("account", "u", 0, "account number")
		cmd.Flags().Uint64P("sequence", "s", 0, "sequence number")
		cmd.Flags().Uint64P("fee-amount", "a", 0, "fee amount")
		cmd.Flags().StringP("fee-denom", "d", "", "fee denom")
		cmd.Flags().Uint64P("gas", "g", 300000, "gas")
		cmd.Flags().StringP("memo", "m", "", "memo")
		cmds = append(cmds, cmd)
	}
	return
}
