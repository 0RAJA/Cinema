package model

// Page 需要当前页的数据,当前页数,每页记录数,总页数,总记录数
type Page struct {
	Movies      []*Movie  `json:"movies,omitempty"` //每页查询出来的电影存放的切片
	Screens     []*Screen `json:"-"`
	Plans       []*Plan   `json:"-"`
	PageNo      int       `json:"page_no,omitempty"`       //当前页
	PageSize    int       `json:"page_size,omitempty"`     //每页显示的条数
	TotalPageNo int       `json:"total_page_no,omitempty"` //总页数
	TotalRecord int       `json:"total_record,omitempty"`  //总记录数
	Min         string    `json:"min,omitempty"`           //最小
	Max         string    `json:"max,omitempty"`           //最大
}

// IsHasPrev 是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// IsHasNext 是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// GetPrevPageNo 获取前一页
func (p *Page) GetPrevPageNo() int {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	return 1
}

// GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.PageSize
}
