package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type Fish struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateFish(param *Fish) (*model.Fish, error) {
	fish := model.Fish{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}
	return fish.Create(d.engine)
}

func (d *Dao) UpdateFish(param *Fish) error {
	fish := model.Fish{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return fish.Update(d.engine, values)
}

func (d *Dao) GetFish(id uint32, state uint8) (model.Fish, error) {
	fish := model.Fish{Model: &model.Model{ID: id}, State: state}
	return fish.Get(d.engine)
}

func (d *Dao) DeleteFish(id uint32) error {
	fish := model.Fish{Model: &model.Model{ID: id}}
	return fish.Delete(d.engine)
}

func (d *Dao) CountFishListByTagID(id uint32, state uint8) (int, error) {
	fish := model.Fish{State: state}
	return fish.CountByTagID(d.engine, id)
}

func (d *Dao) GetFishListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.FishRow, error) {
	fish := model.Fish{State: state}
	return fish.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}
