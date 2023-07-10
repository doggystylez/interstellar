package query

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/doggystylez/interstellar/client/grpc"
	"github.com/doggystylez/interstellar/proto/query/base"
	"github.com/doggystylez/interstellar/proto/query/wasm"
)

func (e RetryErr) Error() string {
	return "failed to query after " + strconv.Itoa(e.retries+1) + " attempts. last error: " + e.err.Error()
}

func (e TxNotFoundErr) Error() string {
	return e.message
}

func decodePrefix(address string) string {
	prefix, _, err := bech32.DecodeToBase256(address)
	if err != nil {
		panic(err)
	}
	return prefix
}

// LatesTendermintBlock
func LatestBlock(g grpc.Client) (BlockRes, error) {
	var res *base.GetLatestBlockResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := base.NewServiceClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.GetLatestBlock(g.Ctx, &base.GetLatestBlockRequest{})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return BlockRes{Block: *res.Block}, nil
		}
	}
	return BlockRes{}, RetryErr{retries: g.Retries, err: err}
}

// TendermintBlock
func Block(height int64, g grpc.Client) (BlockRes, error) {
	var res *base.GetBlockByHeightResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := base.NewServiceClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.GetBlockByHeight(g.Ctx, &base.GetBlockByHeightRequest{Height: height})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return BlockRes{Block: *res.Block}, nil
		}
	}
	return BlockRes{}, RetryErr{retries: g.Retries, err: err}
}

func ChainId(g grpc.Client) (ChainIdRes, error) {
	var res *base.GetNodeInfoResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := base.NewServiceClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.GetNodeInfo(g.Ctx, &base.GetNodeInfoRequest{})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return ChainIdRes{ChainId: res.DefaultNodeInfo.GetNetwork()}, nil
		}
	}
	return ChainIdRes{}, RetryErr{retries: g.Retries, err: err}
}

func AllContractData(address string, g grpc.Client) (ContractRes, error) {
	var res *wasm.QueryAllContractStateResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := wasm.NewQueryClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.AllContractState(g.Ctx, &wasm.QueryAllContractStateRequest{Address: address, Pagination: &query.PageRequest{Limit: 2000}})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			var data ContractRes
			for _, model := range res.Models {
				data.Models = append(data.Models, Model{Key: model.Key, Value: model.Value})
			}
			return data, nil
		}
	}
	return ContractRes{}, RetryErr{retries: g.Retries, err: err}
}

func ContractDataByQuery(address string, query interface{}, queryRes interface{}, g grpc.Client) error {
	var res *wasm.QuerySmartContractStateResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := wasm.NewQueryClient(g.Conn)
	data, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.SmartContractState(g.Ctx, &wasm.QuerySmartContractStateRequest{Address: address, QueryData: data})
		if err != nil {
			if strings.Contains(err.Error(), "query wasm contract failed") {
				return err
			} else {
				if tries < maxTries {
					time.Sleep(time.Duration(g.Interval) * time.Second)
				}
			}
		} else {
			err = json.Unmarshal(res.Data, &queryRes)
			if err != nil {
				panic(err)
			}
			return nil
		}
	}
	return RetryErr{retries: g.Retries, err: err}
}

func Jsonify(input interface{}) (output string) {
	jsonData, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		panic(err)
	}
	output = string(jsonData)
	return
}
