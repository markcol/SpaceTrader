package main

import "github.com/markcol/SpaceTrader/sdk"

const sdkToken = "foobar"

func main() {
	cfg := sdk.Configuration.NewConfiguration()
	cfg.AddHeader("Bearer", sdkToken)
	client := sdk.NewAPIClient(cfg)
	_ = client
}
