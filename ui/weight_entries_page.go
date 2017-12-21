package ui

type WeightEntriesPage struct {
	Page
}

func NewWeightEntriesPage(userID string) *WeightEntriesPage {
	page := WeightEntriesPage{
		Page: *NewPage(userID),
	}

	return &page
}
