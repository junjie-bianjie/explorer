package service

import (
	"github.com/irisnet/explorer/backend/utils"
	"testing"

	"encoding/json"
	"github.com/irisnet/explorer/backend/model"
	"gopkg.in/mgo.v2/bson"
)

func TestQueryTxList(t *testing.T) {
	txPage := new(TxService).QueryTxList(nil, 0, 100, true)

	t.Logf("total: %v \n", txPage.Count)

	bytestr, _ := json.Marshal(txPage.Data)
	t.Logf("items: %v \n", string(bytestr))
}

func TestTxQueryList(t *testing.T) {

	txPage := new(TxService).QueryList(nil, 0, 100, true)
	t.Logf("total: %v \n", txPage.Count)
	t.Logf("items: %v \n", txPage.Data)
}

func TestQueryRecentTx(t *testing.T) {

	txList := new(TxService).QueryRecentTx()

	for k, v := range txList {
		t.Logf("idx: %v v: %v \n", k, v)
	}
}

func TestQueryTxByHash(t *testing.T) {
	tx := new(TxService).Query("89D8527FC5CB56B79E02EBDFFCA47A1FAB246A8CEBBB6F455B1D44D5F8A39396")
	t.Logf("tx: %v\n", string(utils.MarshalJsonIgnoreErr(tx)))
}

func TestServiceTxfetchLogMessage(t *testing.T) {
	log := "Msg 0 failed: {\"codespace\":\"sdk\",\"code\":10,\"message\":\"12097471760000000000000iris-atto is less than 100000000000000000000000iris-atto\"}"
	ret := fetchLogMessage(log)
	t.Log(ret)
}

func TestQueryByAcc(t *testing.T) {

	txPage := new(TxService).QueryByAcc("faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn", 0, 10, true)

	t.Logf("total: %v \n", txPage.Count)
	if modelV, ok := txPage.Data.([]model.CommonTx); ok {
		for k, v := range modelV {
			t.Logf("idx: %v  v: %v \n", k, v)
		}
	}
}

func TestCountByType(t *testing.T) {

	statistic := new(TxService).CountByType(bson.M{})
	t.Logf("tx statistic by type: %v \n", statistic)
}

func TestQueryTxNumGroupByDay(t *testing.T) {

	txCountByDay := new(TxService).QueryTxNumGroupByDay()

	for k, v := range txCountByDay {
		t.Logf("idx: %v  txCountByDay: %v \n", k, v)
	}
}
