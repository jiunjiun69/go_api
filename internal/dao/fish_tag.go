package dao

import "github.com/go-programming-tour-book/blog-service/internal/model"

func (d *Dao) GetFishTagByAID(fishID uint32) (model.FishTag, error) {
	fishTag := model.FishTag{FishID: fishID}
	return fishTag.GetByAID(d.engine)
}

func (d *Dao) GetFishTagListByTID(tagID uint32) ([]*model.FishTag, error) {
	fishTag := model.FishTag{TagID: tagID}
	return fishTag.ListByTID(d.engine)
}

func (d *Dao) GetFishTagListByAIDs(fishIDs []uint32) ([]*model.FishTag, error) {
	fishTag := model.FishTag{}
	return fishTag.ListByAIDs(d.engine, fishIDs)
}

func (d *Dao) CreateFishTag(fishID, tagID uint32, createdBy string) error {
	fishTag := model.FishTag{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		FishID: fishID,
		TagID:  tagID,
	}
	return fishTag.Create(d.engine)
}

func (d *Dao) UpdateFishTag(fishID, tagID uint32, modifiedBy string) error {
	fishTag := model.FishTag{FishID: fishID}
	values := map[string]interface{}{
		"fish_id":     fishID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return fishTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteFishTag(fishID uint32) error {
	fishTag := model.FishTag{FishID: fishID}
	return fishTag.DeleteOne(d.engine)
}
