package domain

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Picture   string `json:"picture"`
	Title     string `json:"title"`
}
