package request

type PageQuery struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
