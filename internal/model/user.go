package model

type User struct {
	ID    uint32 `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Brief string `json:"brief"`
	AvatarUrl string `json:"avatar_url"`
	UID string `json:"uid"`
	Account string `json:"account"`
	WxAccount string `json:"wx_account"`
	GitHubAddress string `json:"github_address"`
	WeiBoAddress string `json:"weibo_address"`
	BilibiliAddress string `json:"bilibili_address"`
	PasswordSalt string `json:"password_salt"`
	Password string `json:"password"`
	WxID string `json:"wx_id"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel uint8 `json:"is_del"`
}

func (u User) TableName() string{
	return "blog_user"
}