package models

type Category struct {
	ID       int64
	Name     string
	ImageURL string
}

type CategoryCreate struct {
	Name     string
	ImageURL string
}
