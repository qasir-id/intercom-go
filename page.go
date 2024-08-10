package intercom

// PageParams determine paging information to and from the API
type PageParams struct {
	TypeParam  string `json:"type" url:"type,omitempty"`
	Page       int64  `json:"page" url:"page,omitempty"`
	PerPage    int64  `json:"per_page" url:"per_page,omitempty"`
	TotalPages int64  `json:"total_pages" url:"-"`
}
