package query

import (
	"time"

	"github.com/doggystylez/interstellar-proto/query/pool"
	"github.com/doggystylez/interstellar/client/grpc"
)

func SpotPrice(poolId int, base string, quote string, g grpc.Client) (SpotPriceRes, error) {
	// var res *pool.SpotPriceResponse
	var res *pool.QuerySpotPriceResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		//		res, err = client.SpotPrice(g.Ctx, &pool.SpotPriceRequest{
		res, err = client.SpotPrice(g.Ctx, &pool.QuerySpotPriceRequest{
			PoolId: uint64(poolId),
			//	BaseAssetDenom:  base,
			//	QuoteAssetDenom: quote,
			BaseAssetDenom:  quote,
			QuoteAssetDenom: base, // fixed in v16
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

// func EstimateSwapSinglePool(poolId int, amountIn string, tokenIn string, denomOut string, g grpc.Client) (SwapEstimate, error) {
// 	var res *pool.EstimateSwapExactAmountInResponse
// 	err := g.Open()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer g.Close()
// 	client := pool.NewQueryClient(g.Conn)
// 	for tries := -1; tries < g.Retries; tries++ {
// 		res, err = client.EstimateSinglePoolSwapExactAmountIn(g.Ctx, &pool.EstimateSinglePoolSwapExactAmountInRequest{
// 			PoolId:        uint64(poolId),
// 			TokenIn:       amountIn + tokenIn,
// 			TokenOutDenom: denomOut,
// 		})
// 		if err != nil {
// 			time.Sleep(time.Duration(g.Interval) * time.Second)
// 		} else {
// 			return SwapEstimate{
// 				TokenIn:  Token{Denom: tokenIn, Amount: amountIn},
// 				TokenOut: Token{Denom: denomOut, Amount: res.TokenOutAmount},
// 			}, nil
// 		}
// 	}
// 	return SwapEstimate{}, RetryErr{retries: g.Retries, err: err}
// }

// wip
// func EstimateSwap(poolId int, amountIn string, tokenIn string, route []SwapRoute, g grpc.Client) (SwapEstimate, error) {
func EstimateSwap(sender string, poolId int, amountIn string, tokenIn string, route []SwapRoute, g grpc.Client) (SwapEstimate, error) {
	// var res *pool.EstimateSwapExactAmountInResponse
	var (
		res         *pool.QuerySwapExactAmountInResponse
		queryRoutes []*pool.SwapAmountInRoute
	)
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	for _, hop := range route {
		queryRoutes = append(queryRoutes, &pool.SwapAmountInRoute{PoolId: hop.PoolId, TokenOutDenom: hop.DenomOut})
	}
	for tries := -1; tries < g.Retries; tries++ {
		//	res, err = client.EstimateSwapExactAmountIn(g.Ctx, &pool.EstimateSwapExactAmountInRequest{
		res, err = client.EstimateSwapExactAmountIn(g.Ctx, &pool.QuerySwapExactAmountInRequest{
			Sender:  sender, //not in v16
			PoolId:  uint64(poolId),
			TokenIn: amountIn + tokenIn,
			Routes:  queryRoutes,
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
