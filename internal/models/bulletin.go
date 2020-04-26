package models

import "time"

type Bulletin struct {
	Model

	Content string `json:"content"`
	Top     int    `json:"top"`
}

type QueryBulletinReq struct {
	Top int `json:"top"`
	Pagination
}

type BulletinResponse struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	Content   string     `json:"content"`
	Top       int        `json:"top"`
}

type BulletinsSerializer struct {
	Bulletins []Bulletin
}

func (b *Bulletin) Response() BulletinResponse {
	Bulletin := BulletinResponse{
		ID:        b.ID,
		CreatedAt: b.CreatedAt,
		Content:   b.Content,
		Top:       b.Top,
	}
	return Bulletin
}

func (s *BulletinsSerializer) Response() []BulletinResponse {
	var Bulletins []BulletinResponse
	for _, Bulletin := range s.Bulletins {
		Bulletins = append(Bulletins, Bulletin.Response())
	}
	return Bulletins
}
