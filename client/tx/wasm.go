package tx

import "encoding/json"

func MakeSwapContractMsg(amount, denomIn, denomOut, slippage string) (bytes []byte) {
	bytes, err := json.Marshal(WasmSwap{
		Swap: Swap{
			InputCoin:   InputCoin{Denom: denomIn, Amount: amount},
			OutputDenom: denomOut,
			Slippage:    Slippage{Twap: Twap{SlippagePercentage: slippage, WindowSeconds: 5}},
		},
	})
	if err != nil {
		panic(err)
	}
	return
}
