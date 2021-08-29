package model

// Collection 文章栏目模型
type Collection struct{
	*Model
    Name string `json:"name"`
	State uint8 `json:"state"`
}

func (c Collection) TableName() string{
	return "blog_collection"
}