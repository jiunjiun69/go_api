package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

type Fishpower struct {
	ID         uint32 `json:"id"`
	TagID      uint32 `json:"tag_id"`
	Title      string `json:"title"`
	Power      string `json:"power"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      uint8  `json:"state"`
}

func (d *Dao) CreateFishpower(param *Fishpower) (*model.Fishpower, error) {
	fishpower := model.Fishpower{
		Title: param.Title,
		Power: param.Power,
		State: param.State,
		Model: &model.Model{CreatedBy: param.CreatedBy},
	}
	return fishpower.Create(d.engine)
}

func (d *Dao) UpdateFishpower(param *Fishpower) error {
	fishpower := model.Fishpower{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.Power != "" {
		values["power"] = param.Power
	}

	return fishpower.Update(d.engine, values)
}

func (d *Dao) GetFishpower(id uint32, state uint8) (model.Fishpower, error) {
	fishpower := model.Fishpower{Model: &model.Model{ID: id}, State: state}
	return fishpower.Get(d.engine)
}

func (d *Dao) DeleteFishpower(id uint32) error {
	fishpower := model.Fishpower{Model: &model.Model{ID: id}}
	return fishpower.Delete(d.engine)
}

/*
func (d *Dao) CountFishpowerListByTagID(id uint32, state uint8) (int, error) {
	fishpower := model.Fishpower{State: state}
	return fishpower.CountByTagID(d.engine, id)
}

func (d *Dao) GetFishpowerListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.FishpowerRow, error) {
	fishpower := model.Fishpower{State: state}
	return fishpower.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}
*/
