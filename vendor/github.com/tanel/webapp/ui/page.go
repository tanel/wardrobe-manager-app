package ui

import (
	"github.com/tanel/webapp/model"
)

// Page represents a page
type Page struct {
	UserID      string
	CurrentUser *model.User
}

// NewPage returns a new page
func NewPage() *Page {
	return &Page{}
}

// NewPageWithUserID returns a new page with user ID
func NewPageWithUserID(userID string) *Page {
	return &Page{
		UserID: userID,
	}
}

// NewPageWithUser returns a new page with user
func NewPageWithUser(currentUser *model.User) *Page {
	return &Page{
		UserID:      currentUser.ID,
		CurrentUser: currentUser,
	}
}
