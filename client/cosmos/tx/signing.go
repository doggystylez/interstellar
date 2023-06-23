package tx

import (
	txclient "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/doggystylez/interstellar/keys/hexkey"
	"github.com/doggystylez/interstellar/types"
)

func SignFromPrivkey(msg []sdk.Msg, config types.InterstellarConfig) (txBytes []byte, err error) {
	txConfig, err := buildTx(msg, config.TxInfo)
	if err != nil {
		return
	}
	sigData := signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	priv, err := hexkey.PrivKeyfromString(config.TxInfo.KeyInfo.KeyRing.HexPriv)
	if err != nil {
		return
	}
	sig := signing.SignatureV2{
		PubKey:   priv.PubKey(),
		Data:     &sigData,
		Sequence: config.TxInfo.KeyInfo.SeqNum,
	}
	err = txConfig.TxBuilder.SetSignatures(sig)
	if err != nil {
		return
	}
	signingData := authsigning.SignerData{
		Address:       config.TxInfo.Address,
		ChainID:       config.TxInfo.KeyInfo.ChainId,
		AccountNumber: config.TxInfo.KeyInfo.AccNum,
	}
	sig, err = txclient.SignWithPrivKey(
		txConfig.TxConfig.SignModeHandler().DefaultMode(), signingData,
		txConfig.TxBuilder, priv, txConfig.TxConfig, config.TxInfo.KeyInfo.SeqNum)
	if err != nil {
		return
	}
	err = txConfig.TxBuilder.SetSignatures(sig)
	if err != nil {
		return
	}
	txBytes, err = txConfig.TxConfig.TxEncoder()(txConfig.TxBuilder.GetTx())
	if err != nil {
		return
	}
	return
}
