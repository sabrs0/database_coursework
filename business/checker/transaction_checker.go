package checker

type TransactionMainParams struct {
	From_essence_type bool
	From_id           uint64
	To_essence_type   bool
	Sum_of_money      float64
	Comment           string
	To_id             uint64
}

func NewTransactionMainParams(from_e bool, fid uint64, tetype bool, reqSum float64, comm string, tid uint64) TransactionMainParams {
	return TransactionMainParams{From_essence_type: from_e, From_id: fid,
		To_essence_type: tetype, Sum_of_money: reqSum, Comment: comm, To_id: tid}
}
