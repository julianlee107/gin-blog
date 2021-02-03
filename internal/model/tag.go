package model

import (
	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)

	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where(" id = ? AND state = ? AND is_del = ?", t.ID, t.State, 0).First(&tag).Error
	if err != nil {
		global.Logger.Error("model.Tag.Get err:", t.ID, t.State)
		return Tag{}, err
	}
	return tag, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}
