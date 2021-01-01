package dao

import "github.com/go-programming-tour-book/blog-service/internal/model"

func (d *Dao) GetFishpowerTagByAID(fishpowerID uint32) (model.FishpowerTag, error) {
	fishpowerTag := model.FishpowerTag{FishpowerID: fishpowerID}
	return fishpowerTag.GetByAID(d.engine)
}

func (d *Dao) GetFishpowerTagListByTID(tagID uint32) ([]*model.FishpowerTag, error) {
	fishpowerTag := model.FishpowerTag{TagID: tagID}
	return fishpowerTag.ListByTID(d.engine)
}

func (d *Dao) GetFishpowerTagListByAIDs(fishpowerIDs []uint32) ([]*model.FishpowerTag, error) {
	fishpowerTag := model.FishpowerTag{}
	return fishpowerTag.ListByAIDs(d.engine, fishpowerIDs)
}

func (d *Dao) CreateFishpowerTag(fishpowerID, tagID uint32, createdBy string) error {
	fishpowerTag := model.FishpowerTag{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		FishpowerID: fishpowerID,
		TagID:       tagID,
	}
	return fishpowerTag.Create(d.engine)
}

func (d *Dao) UpdateFishpowerTag(fishpowerID, tagID uint32, modifiedBy string) error {
	fishpowerTag := model.FishpowerTag{FishpowerID: fishpowerID}
	values := map[string]interface{}{
		"fishpower_id": fishpowerID,
		"tag_id":       tagID,
		"modified_by":  modifiedBy,
	}
	return fishpowerTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteFishpowerTag(fishpowerID uint32) error {
	fishpowerTag := model.FishpowerTag{FishpowerID: fishpowerID}
	return fishpowerTag.DeleteOne(d.engine)
}
