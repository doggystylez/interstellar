package tx

import (
	txclient "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/doggystylez/interstellar/client/keys"
)

func SignFromPrivkey(msg []sdk.Msg, txInfo TxInfo) (txBytes []byte, err error) {
	txConfig, err := buildTx(msg, txInfo)
	if err != nil {
		return
	}
	sigData := signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	priv := keys.FromBytes(txInfo.KeyInfo.KeyRing.KeyBytes)
	sig := signing.SignatureV2{
		PubKey:   priv.PubKey(),
		Data:     &sigData,
		Sequence: txInfo.KeyInfo.SeqNum,
	}
	err = txConfig.TxBuilder.SetSignatures(sig)
	if err != nil {
		return
	}
	signingData := authsigning.SignerData{
		Address:       txInfo.Address,
		ChainID:       txInfo.KeyInfo.ChainId,
		AccountNumber: txInfo.KeyInfo.AccNum,
	}
	sig, err = txclient.SignWithPrivKey(
		txConfig.TxConfig.SignModeHandler().DefaultMode(), signingData,
		txConfig.TxBuilder, priv, txConfig.TxConfig, txInfo.KeyInfo.SeqNum)
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
