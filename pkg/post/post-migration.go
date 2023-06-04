package post

import "gorm.io/gorm"

type ModelPost struct {
	gorm.Model
	Title         string                `gorm:"size:255;not null"`
	Slug          string                `gorm:"size:255;not null;uniqueIndex"`
	Author        string                `gorm:"size:255;not null"`
	Content       string                `gorm:"type:text;"`
	FeaturedImage string                `gorm:"type:text;"`
	Posts         []ModelPostToCategory `gorm:"foreignKey:Post"`
}

type ModelPostCategory struct {
	gorm.Model
	Name          string                `gorm:"size:255;not null"`
	Slug          string                `gorm:"size:255;not null;uniqueIndex"`
	FeaturedImage string                `gorm:"type:text;"`
	Categories    []ModelPostToCategory `gorm:"foreignKey:Category"`
}

type ModelPostToCategory struct {
	gorm.Model
	Post     int64
	Category int64
}
