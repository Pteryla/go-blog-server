package model
type ArticleCollection struct{
	*Model
	CollectionID uint32 `json:"collection_id"`
	ArticleID uint32 `json:"article_id"`
	State uint8 `json:"state"`
}

func(ac * ArticleCollection) TableName() string{
	return "blog_article_collection"
}