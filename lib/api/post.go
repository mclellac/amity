package api

type Post struct {
	ID      int32  `json:"id"`
	Created int32  `sql:"size:25"                  json:"created,omitempty"`
	Title   string `sql:"size:250;not null"        json:"title" binding:"required"`
	Article string `sql:"size:8000;not null"       json:"article"  binding:"required"`
}
