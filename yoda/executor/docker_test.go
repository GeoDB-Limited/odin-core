package executor

//import (
//	"fmt"
//	"github.com/stretchr/testify/require"
//	"testing"
//	"time"
//)
//
//var env = map[string]interface{}{
//	EnvVarChainID:     "test-chain-id",
//	EnvVarValidator:   "test-validator",
//	EnvVarRequestID:   "test-request-id",
//	EnvVarExternalID:  "test-external-id",
//	EnvVarReporter:    "test-reporter",
//	EnvVarSignature:   "test-signature",
//	EnvVarMaxDataSize: 1024,
//}
//
//func TestDockerSuccess(t *testing.T) {
//	e := NewDockerExec("odinprotocol/core:master", 10*time.Second)
//	res, err := e.Exec([]byte(`#!/usr/bin/env python3
//	import json
//	import urllib.request
//	import sys
//
//	BINANCE_URL = "https://api.binance.com/api/v1/depth?symbol={}USDT&limit=5"
//
//	def make_json_request(url):
//		req = urllib.request.Request(url)
//		req.add_header(
//			"User-Agent",
//			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36",
//		)
//		return json.loads(urllib.request.urlopen(req).read())
//
//	def main(symbol):
//		res = make_json_request(BINANCE_URL.format(symbol))
//		bid = float(res["bids"][0][0])
//		ask = float(res["asks"][0][0])
//		return (bid + ask) / 2
//
//	if __name__ == "__main__":
//		try:
//			print(main(*sys.argv[1:]))
//		except Exception as e:
//			print(str(e), file=sys.stderr)
//			sys.exit(1)
//	`), "BTC", env)
//
//	require.NoError(t, err)
//	fmt.Println(string(res.Output), res.Code, err)
//	require.True(t, false)
//}
//
//func TestDockerLongStdout(t *testing.T) {
//	e := NewDockerExec("odinprotocol/core:master", 10*time.Second)
//	result, err := e.Exec([]byte(`#!/usr/bin/env python3 print("A"*1000)`), "BTC", env)
//	require.NoError(t, err)
//	fmt.Println(string(result.Output), result.Code, err)
//	require.True(t, false)
//}
