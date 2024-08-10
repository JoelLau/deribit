package main

import (
	"fmt"
	"log/slog"

	jsonrpchttp "github.com/JoelLau/deribit/pkgs/deribit-go-sdk/clients/jsonrpc-http"
)

func main() {
	client, err := jsonrpchttp.NewJsonRpcClient(jsonrpchttp.ClientConfig{
		Host:    "test.deribit.com",
		Verbose: true,
	})
	if err != nil {
		err = fmt.Errorf("error creating Deribit JSON RPC client: %w", err)
		panic(err)
	}

	res, err := client.Get(jsonrpchttp.MessageRequest{
		JsonRpc: "2.0",
		Id:      1111,
		Method:  "public/get_instruments",
		Params:  map[string]any{},
	})
	if err != nil {
		err = fmt.Errorf("error while making GET request: %w", err)
		panic(err)
	}

	slog.Info("received response", slog.Any("response", res))
}
