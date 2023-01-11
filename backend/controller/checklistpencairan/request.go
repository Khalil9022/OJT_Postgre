package checklistpencairan

type DataRequest struct {
	Branch  string `json:"branch"`
	Company string `json:"company"`
	Start   string `json:"start"`
	End     string `json:"end"`
}
