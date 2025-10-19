package models

type Post struct {
	Id       int32  `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title" validate:"required,min=3"`
	Body     string `json:"body" validate:"required"`
	AuthorId string `json:"authorId" validate:"required"`
}
