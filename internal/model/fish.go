package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Fish struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type FishSwagger struct {
	List  []*Fish
	Pager *app.Pager
}

func (a Fish) TableName() string {
	return "blog_fish"
}

func (a Fish) Create(db *gorm.DB) (*Fish, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Fish) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a Fish) Get(db *gorm.DB) (Fish, error) {
	var fish Fish
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	err := db.First(&fish).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fish, err
	}

	return fish, nil
}

func (a Fish) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

type FishRow struct {
	FishID        uint32
	TagID         uint32
	TagName       string
	FishTitle     string
	FishDesc      string
	CoverImageUrl string
	Content       string
}

func (a Fish) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*FishRow, error) {
	fields := []string{"ar.id AS fish_id", "ar.title AS fish_title", "ar.desc AS fish_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(FishTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Fish{}.TableName()+"` AS ar ON at.fish_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fishs []*FishRow
	for rows.Next() {
		r := &FishRow{}
		if err := rows.Scan(&r.FishID, &r.FishTitle, &r.FishDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}

		fishs = append(fishs, r)
	}

	return fishs, nil
}

func (a Fish) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(FishTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Fish{}.TableName()+"` AS ar ON at.fish_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
