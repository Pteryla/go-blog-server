package model

type Article struct{
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	ContentType uint8 `json:"content_type"`
	Content string `json:"content"`
	ContentUrl string `json:"content_url"`
	State uint8 `json:"state"`
}

func (a Article) TableName() string{
	return "blog_article"
}