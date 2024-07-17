package global

type Page struct {
	PageNum int `json:"page" form:"page"`
	Limit   int `json:"limit" form:"limit"`
	Total   int `json:"total" form:"total"`
}

func (p *Page) GetPage() (offset, page int) {
	var _limit = 0
	var _page = 0
	if p.PageNum == 0 {
		_limit = 10
	} else {
		_limit = p.Limit
	}

	if p.PageNum == 0 {
		_page = 0
	} else {
		_page = p.PageNum
	}

	return _limit, _page
}
