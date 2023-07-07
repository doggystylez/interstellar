package query

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/doggystylez/interstellar-proto/query/base"
	"github.com/doggystylez/interstellar-proto/query/wasm"
	"github.com/doggystylez/interstellar/client/grpc"
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

func ChainId(g grpc.Client) (ChainIdRes, error) {
	var res *base.GetNodeInfoResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := base.NewServiceClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.GetNodeInfo(g.Ctx, &base.GetNodeInfoRequest{})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
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
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.AllContractState(g.Ctx, &wasm.QueryAllContractStateRequest{Address: address, Pagination: &query.PageRequest{Limit: 2000}})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
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
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.SmartContractState(g.Ctx, &wasm.QuerySmartContractStateRequest{Address: address, QueryData: data})
		if err != nil {
			if strings.Contains(err.Error(), "query wasm contract failed") {
				return err
			} else {
				time.Sleep(time.Duration(g.Interval) * time.Second)
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
