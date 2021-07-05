package model

type Item struct {
	ShortLink string
	FullLink  string
	UserID    string
	TTL       uint
	Count     uint
}
