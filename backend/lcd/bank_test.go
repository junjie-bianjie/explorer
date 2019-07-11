package lcd

import (
	"testing"
	"encoding/json"
)

func  TestGetBankTokenStats(t *testing.T) {
	res,err := GetBankTokenStats()
	if err != nil {
		t.Fatal(err)
	}

	bytesData,_ := json.Marshal(res)
	t.Log(string(bytesData))
}

func TestGetTokenStatsCirculation(t *testing.T) {
	res,err := GetTokenStatsCirculation()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetTokenStatsSupply(t *testing.T) {
	res,err := GetTokenStatsSupply()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetBuredTokens(t *testing.T) {
	res := []*Coin{{Denom:"iris-atto",Amount:"925050600000000000000000"},{Denom:"kai-min",Amount:"800"}}
	coin := GetBuredTokens(res)
	t.Log(coin)
}