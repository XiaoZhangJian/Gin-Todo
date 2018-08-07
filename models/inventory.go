package models

// 清单表
type Inventory struct {
	Model
	UserID int  `json:"user_id" gorm:"index"`
	User   User `json:"user"`

	Title string `json:"title"`

	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}
