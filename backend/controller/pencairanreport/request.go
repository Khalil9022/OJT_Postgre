package pencairanreport

type DataRequest struct {
	Branch  string `json:"branch"`
	Company string `json:"company"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

type ReqPpk struct {
	Ppk []string `json:"ppk"`
}
