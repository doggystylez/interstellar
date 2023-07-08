package query

import (
	"time"

	"github.com/cosmos/cosmos-sdk/types/query"
	legacy_pool "github.com/doggystylez/interstellar-proto/query/legacy-pool"
	"github.com/doggystylez/interstellar-proto/query/pool"

	"github.com/doggystylez/interstellar/client/grpc"
	"google.golang.org/protobuf/proto"
)

func SpotPrice(poolId uint64, base string, quote string, g grpc.Client) (SpotPriceRes, error) {
	var res *pool.SpotPriceResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.SpotPrice(g.Ctx, &pool.SpotPriceRequest{
			PoolId:          poolId,
			BaseAssetDenom:  base,
			QuoteAssetDenom: quote,
		})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
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

func EstimateSwapSinglePool(poolId uint64, amountIn string, tokenIn string, denomOut string, g grpc.Client) (SwapEstimate, error) {
	var res *pool.EstimateSwapExactAmountInResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.EstimateSinglePoolSwapExactAmountIn(g.Ctx, &pool.EstimateSinglePoolSwapExactAmountInRequest{
			PoolId:        poolId,
			TokenIn:       amountIn + tokenIn,
			TokenOutDenom: denomOut,
		})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return SwapEstimate{
				TokenIn:  Token{Denom: tokenIn, Amount: amountIn},
				TokenOut: Token{Denom: denomOut, Amount: res.TokenOutAmount},
			}, nil
		}
	}
	return SwapEstimate{}, RetryErr{retries: g.Retries, err: err}
}

func EstimateSwap(poolId uint64, amountIn string, tokenIn string, route []SwapRoute, g grpc.Client) (SwapEstimate, error) {
	var res *pool.EstimateSwapExactAmountInResponse
	queryRoutes := make([]*pool.SwapAmountInRoute, len(route))
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	for i := range route {
		queryRoutes[i] = &pool.SwapAmountInRoute{
			PoolId:        route[i].PoolId,
			TokenOutDenom: route[i].DenomOut,
		}
	}
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.EstimateSwapExactAmountIn(g.Ctx, &pool.EstimateSwapExactAmountInRequest{
			PoolId:  uint64(poolId),
			TokenIn: amountIn + tokenIn,
			Routes:  queryRoutes,
		})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return SwapEstimate{
				TokenIn:  Token{Denom: tokenIn, Amount: amountIn},
				TokenOut: Token{Denom: route[len(route)-1].DenomOut, Amount: res.TokenOutAmount},
			}, nil
		}
	}
	return SwapEstimate{}, RetryErr{retries: g.Retries, err: err}
}

func NumPools(g grpc.Client) (*pool.NumPoolsResponse, error) {
	var res *pool.NumPoolsResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := pool.NewQueryClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.NumPools(g.Ctx, &pool.NumPoolsRequest{})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return res, nil
		}
	}
	return &pool.NumPoolsResponse{}, RetryErr{retries: g.Retries, err: err}
}

func AllPools(g grpc.Client) (PoolsRes, error) {
	// var res *pool.AllPoolsResponse
	var res *legacy_pool.QueryPoolsResponse
	var poolList PoolsRes
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	// client := pool.NewQueryClient(g.Conn)
	client := legacy_pool.NewQueryClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		//	res, err = client.AllPools(g.Ctx, &pool.AllPoolsRequest{})
		res, err = client.Pools(g.Ctx, &legacy_pool.QueryPoolsRequest{Pagination: &query.PageRequest{Limit: 2000}})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			for _, p := range res.Pools {
				o := Pool{
					Type: p.TypeUrl,
				}
				if o.Type == "/osmosis.gamm.v1beta1.Pool" {
					var b pool.BalancerPool
					err = proto.Unmarshal(p.Value, &b)
					if err != nil {
						return PoolsRes{}, err
					}
					for _, asset := range b.PoolAssets {
						o.PoolAssets = append(o.PoolAssets, PoolAsset{
							Token: Token{
								Denom:  asset.Token.Denom,
								Amount: asset.Token.Amount,
							},
							Weight: asset.Weight,
						})
					}
					o.Id = b.Id
				} else if o.Type == "/osmosis.gamm.poolmodels.stableswap.v1beta1.Pool" {
					var s pool.StableswapPool
					err = proto.Unmarshal(p.Value, &s)
					if err != nil {
						return PoolsRes{}, err
					}
					for _, asset := range s.PoolLiquidity {
						o.PoolLiquidity = append(o.PoolLiquidity, Token{
							Denom:  asset.Denom,
							Amount: asset.Amount,
						})
					}
					o.Id, o.ScalingFactors = s.Id, s.ScalingFactors
				}
				poolList = append(poolList, o)
			}
			return poolList, nil
		}
	}
	return PoolsRes{}, RetryErr{retries: g.Retries, err: err}
}
