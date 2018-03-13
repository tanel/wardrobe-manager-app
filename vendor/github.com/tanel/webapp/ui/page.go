package ui

// Page represents a page
type Page struct {
	UserID string
}

// NewPage returns a new page
func NewPage(userID string) *Page {
	page := Page{
		UserID: userID,
	}

	return &page
}
