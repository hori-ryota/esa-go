package esa

// PageResp is resp for pager
type PageResp struct {
	PrevPage   *uint `json:"prev_page"`
	NextPage   *uint `json:"next_page"`
	TotalCount uint  `json:"total_count"`
	Page       uint  `json:"page"`
	PerPage    uint  `json:"per_page"`
	MaxPerPage uint  `json:"max_per_page"`
}
