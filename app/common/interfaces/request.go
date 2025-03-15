package interfaces

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PrevPage    *int  `json:"prev_page"`
	NextPage    *int  `json:"next_page"`
	TotalPages  int   `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
	Limit       int   `json:"limit"`
}
