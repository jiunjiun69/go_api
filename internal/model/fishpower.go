package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Fishpower struct {
	*Model
	Title string `json:"title"`
	Power string `json:"power"`
	State uint8  `json:"state"`
}

type FishpowerSwagger struct {
	List  []*Fishpower
	Pager *app.Pager
}

func (a Fishpower) TableName() string {
	return "blog_fishpower"
}

func (a Fishpower) Create(db *gorm.DB) (*Fishpower, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Fishpower) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a Fishpower) Get(db *gorm.DB) (Fishpower, error) {
	var fishpower Fishpower
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	err := db.First(&fishpower).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fishpower, err
	}

	return fishpower, nil
}

func (a Fishpower) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

type FishpowerRow struct {
	FishpowerID    uint32
	TagID          uint32
	TagName        string
	FishpowerTitle string
	Power          string
}

/*
func (a Fishpower) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*FishpowerRow, error) {
	fields := []string{"ar.id AS fishpower_id", "ar.title AS fishpower_title", "ar.cover_image_url", "ar.power"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

		rows, err := db.Select(fields).Table(FishpowerTag{}.TableName()+" AS at").
			Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
			Joins("LEFT JOIN `"+Fishpower{}.TableName()+"` AS ar ON at.fishpower_id = ar.id").
			Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
			Rows()
		if err != nil {
			return nil, err
		}
		defer rows.Close()


	var fishpowers []*FishpowerRow
	for rows.Next() {
		r := &FishpowerRow{}
		if err := rows.Scan(&r.FishpowerID, &r.FishpowerTitle, &r.Power, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}

		fishpowers = append(fishpowers, r)
	}

	return fishpowers, nil
}

func (a Fishpower) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(FishpowerTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Fishpower{}.TableName()+"` AS ar ON at.fishpower_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
*/
