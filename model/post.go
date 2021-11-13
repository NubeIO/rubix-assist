package model

import (
	"gorm.io/gorm"
)

// Data is mainle generated for filtering and pagination
type Data struct {
	TotalData    int64
	FilteredData int64
	Data         []Post
}

type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}

type Post struct {
	gorm.Model
	Name        string `json:"Name" gorm:"type:varchar(255)"`
	Description string `json:"Description"  gorm:"type:text"`
	Tags        []Tag  // One-To-Many relationship (has many - use Tag's UserID as foreign key)
}

type Tag struct {
	gorm.Model
	PostID      uint   `gorm:"index"` // Foreign key (belongs to)
	Name        string `json:"Name" gorm:"type:varchar(255)"`
	Description string `json:"Description" gorm:"type:text"`
}
