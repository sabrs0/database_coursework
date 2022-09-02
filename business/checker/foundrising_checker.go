package checker

type FoundrisingMainParams struct {
	Found_id   uint64
	Descr      string
	ReqSum     float64
	CreateDate string
}

func NewFoundrisingMainParams(found_id uint64, descr string, reqSum float64, createDate string) FoundrisingMainParams {
	return FoundrisingMainParams{Found_id: found_id, Descr: descr, ReqSum: reqSum, CreateDate: createDate}
}
