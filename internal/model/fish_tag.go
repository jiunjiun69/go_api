package model

import (
	"github.com/jinzhu/gorm"
)

type FishTag struct {
	*Model
	TagID  uint32 `json:"tag_id"`
	FishID uint32 `json:"fish_id"`
}

func (a FishTag) TableName() string {
	return "blog_fish_tag"
}

func (a FishTag) GetByAID(db *gorm.DB) (FishTag, error) {
	var fishTag FishTag
	err := db.Where("fish_id = ? AND is_del = ?", a.FishID, 0).First(&fishTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fishTag, err
	}

	return fishTag, nil
}

func (a FishTag) ListByTID(db *gorm.DB) ([]*FishTag, error) {
	var fishTags []*FishTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).Find(&fishTags).Error; err != nil {
		return nil, err
	}

	return fishTags, nil
}

func (a FishTag) ListByAIDs(db *gorm.DB, fishIDs []uint32) ([]*FishTag, error) {
	var fishTags []*FishTag
	err := db.Where("fish_id IN (?) AND is_del = ?", fishIDs, 0).Find(&fishTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return fishTags, nil
}

func (a FishTag) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a FishTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("fish_id = ? AND is_del = ?", a.FishID, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a FishTag) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a FishTag) DeleteOne(db *gorm.DB) error {
	if err := db.Where("fish_id = ? AND is_del = ?", a.FishID, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}

	return nil
}
