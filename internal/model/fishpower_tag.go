package model

import (
	"github.com/jinzhu/gorm"
)

type FishpowerTag struct {
	*Model
	TagID       uint32 `json:"tag_id"`
	FishpowerID uint32 `json:"fishpower_id"`
}

func (a FishpowerTag) TableName() string {
	return "blog_fishpower_tag"
}

func (a FishpowerTag) GetByAID(db *gorm.DB) (FishpowerTag, error) {
	var fishpowerTag FishpowerTag
	err := db.Where("fishpower_id = ? AND is_del = ?", a.FishpowerID, 0).First(&fishpowerTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fishpowerTag, err
	}

	return fishpowerTag, nil
}

func (a FishpowerTag) ListByTID(db *gorm.DB) ([]*FishpowerTag, error) {
	var fishpowerTags []*FishpowerTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).Find(&fishpowerTags).Error; err != nil {
		return nil, err
	}

	return fishpowerTags, nil
}

func (a FishpowerTag) ListByAIDs(db *gorm.DB, fishpowerIDs []uint32) ([]*FishpowerTag, error) {
	var fishpowerTags []*FishpowerTag
	err := db.Where("fishpower_id IN (?) AND is_del = ?", fishpowerIDs, 0).Find(&fishpowerTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return fishpowerTags, nil
}

func (a FishpowerTag) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a FishpowerTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("fishpower_id = ? AND is_del = ?", a.FishpowerID, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a FishpowerTag) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a FishpowerTag) DeleteOne(db *gorm.DB) error {
	if err := db.Where("fishpower_id = ? AND is_del = ?", a.FishpowerID, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}

	return nil
}
