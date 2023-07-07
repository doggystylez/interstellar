package query

import (
	"time"

	"github.com/doggystylez/interstellar-proto/query/pool"
	"github.com/doggystylez/interstellar/client/grpc"
)

func SpotPrice(poolId int, base string, quote string, g grpc.Client) (SpotPriceRes, error) {
	var res *pool.SpotPriceResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {

		res, err = client.SpotPrice(g.Ctx, &pool.SpotPriceRequest{
			PoolId:          uint64(poolId),
			BaseAssetDenom:  base,
			QuoteAssetDenom: quote,
		})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			return SpotPriceRes{
				Base:  base,
				Quote: quote,
				Price: res.SpotPrice,
			}, nil
		}
	}
	return SpotPriceRes{}, RetryErr{retries: g.Retries, err: err}
}

func EstimateSwapSinglePool(poolId int, amountIn string, tokenIn string, denomOut string, g grpc.Client) (SwapEstimate, error) {
	var res *pool.EstimateSwapExactAmountInResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.EstimateSinglePoolSwapExactAmountIn(g.Ctx, &pool.EstimateSinglePoolSwapExactAmountInRequest{
			PoolId:        uint64(poolId),
			TokenIn:       amountIn + tokenIn,
			TokenOutDenom: denomOut,
		})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			return SwapEstimate{
				TokenIn:  Token{Denom: tokenIn, Amount: amountIn},
				TokenOut: Token{Denom: denomOut, Amount: res.TokenOutAmount},
			}, nil
		}
	}
	return SwapEstimate{}, RetryErr{retries: g.Retries, err: err}
}

// wip
func EstimateSwap(poolId int, amountIn string, tokenIn string, route []SwapRoute, g grpc.Client) (SwapEstimate, error) {
	var res *pool.EstimateSwapExactAmountInResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.EstimateSwapExactAmountIn(g.Ctx, &pool.EstimateSwapExactAmountInRequest{
			PoolId:  uint64(poolId),
			TokenIn: amountIn + tokenIn,
			Routes:  []*pool.SwapAmountInRoute{},
		})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			return SwapEstimate{
				TokenIn:  Token{Denom: tokenIn, Amount: amountIn},
				TokenOut: Token{Denom: "", Amount: res.TokenOutAmount},
			}, nil
		}
	}
	return SwapEstimate{}, RetryErr{retries: g.Retries, err: err}
}
