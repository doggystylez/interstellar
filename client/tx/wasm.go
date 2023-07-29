package tx

import "encoding/json"

func MakeContractSwapMsg(amount, denomIn, denomOut, slippage string) (bytes []byte) {
	bytes, err := json.Marshal(ContractSwap{
		Swap: Swap{
			InputCoin:   &Coin{Denom: denomIn, Amount: amount},
			OutputDenom: denomOut,
			Slippage:    Slippage{Twap: Twap{SlippagePercentage: slippage, WindowSeconds: 5}},
		},
	})
	if err != nil {
		panic(err)
	}
	return
}

func MakeContractIbcSwapMsg(from, to, denomOut, slippage string) (bytes []byte) {
	bytes, err := json.Marshal(ContractIbcSwap{
		OsmosisSwap: Swap{
			OnFailedDelivery: &OnFailedDelivery{Addr: from},
			Receiver:         to,
			OutputDenom:      denomOut,
			Slippage:         Slippage{Twap: Twap{SlippagePercentage: slippage, WindowSeconds: 5}},
		},
	})
	if err != nil {
		panic(err)
	}
	return
}
