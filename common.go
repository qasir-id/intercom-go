package intercom

import "fmt"

// Tag represents an Tag in Intercom.
type CommonObject struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
}

// TagList, an object holding a list of Tags
type CommonObjectList struct {
	Type       string `json:"type"`
	Tags       []Tag  `json:"data,omitempty"`
	Url        string `json:"url"`
	TotalCount int64  `json:"total_count"`
	HasMore    bool   `json:"has_more"`
}

func (t CommonObject) String() string {
	return fmt.Sprintf("[intercom] CommonObject { id: %s type: %s url: %s }", t.ID, t.Type, t.Url)
}
