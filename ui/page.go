package ui

type Page struct {
	UserID string
}

func NewPage(userID string) *Page {
	page := Page{
		UserID: userID,
	}

	return &page
}
