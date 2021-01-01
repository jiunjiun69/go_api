package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/dao"
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

type FishpowerRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type FishpowerListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateFishpowerRequest struct {
	TagID     uint32 `form:"tag_id" binding:"required,gte=1"`
	Title     string `form:"title" binding:"required,min=2,max=100"`
	Power     string `form:"power" binding:"required,min=2,max=4294967295"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateFishpowerRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	TagID      uint32 `form:"tag_id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"min=2,max=100"`
	Power      string `form:"power" binding:"min=2,max=4294967295"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteFishpowerRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Fishpower struct {
	ID    uint32     `json:"id"`
	Title string     `json:"title"`
	Power string     `json:"power"`
	State uint8      `json:"state"`
	Tag   *model.Tag `json:"tag"`
}

func (svc *Service) GetFishpower(param *FishpowerRequest) (*Fishpower, error) {
	fishpower, err := svc.dao.GetFishpower(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	/*
		fishpowerTag, err := svc.dao.GetFishpowerTagByAID(fishpower.ID)
		if err != nil {
			return nil, err
		}

		tag, err := svc.dao.GetTag(fishpowerTag.TagID, model.STATE_OPEN)
		if err != nil {
			return nil, err
		}
	*/

	return &Fishpower{
		ID:    fishpower.ID,
		Title: fishpower.Title,
		Power: fishpower.Power,
		State: fishpower.State,
		//Tag:   &tag,
	}, nil
}

/*
func (svc *Service) GetFishpowerList(param *FishpowerListRequest, pager *app.Pager) ([]*Fishpower, int, error) {
	fishpowerCount, err := svc.dao.CountFishpowerListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	fishpowers, err := svc.dao.GetFishpowerListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var fishpowerList []*Fishpower
	for _, fishpower := range fishpowers {
		fishpowerList = append(fishpowerList, &Fishpower{
			ID:    fishpower.FishpowerID,
			Title: fishpower.FishpowerTitle,
			Power: fishpower.Power,
			Tag:   &model.Tag{Model: &model.Model{ID: fishpower.TagID}, Name: fishpower.TagName},
		})
	}

	return fishpowerList, fishpowerCount, nil
}
*/

func (svc *Service) CreateFishpower(param *CreateFishpowerRequest) error {
	//fishpower, err := svc.dao.CreateFishpower(&dao.Fishpower{
	_, err := svc.dao.CreateFishpower(&dao.Fishpower{
		Title:     param.Title,
		Power:     param.Power,
		State:     param.State,
		CreatedBy: param.CreatedBy,
	})
	if err != nil {
		return err
	}
	/*
		err = svc.dao.CreateFishpowerTag(fishpower.ID, param.TagID, param.CreatedBy)
		if err != nil {
			return err
		}
	*/

	return nil
}

func (svc *Service) UpdateFishpower(param *UpdateFishpowerRequest) error {
	err := svc.dao.UpdateFishpower(&dao.Fishpower{
		ID:         param.ID,
		Title:      param.Title,
		Power:      param.Power,
		State:      param.State,
		ModifiedBy: param.ModifiedBy,
	})
	if err != nil {
		return err
	}

	/*
		err = svc.dao.UpdateFishpowerTag(param.ID, param.TagID, param.ModifiedBy)
		if err != nil {
			return err
		}
	*/

	return nil
}

func (svc *Service) DeleteFishpower(param *DeleteFishpowerRequest) error {
	err := svc.dao.DeleteFishpower(param.ID)
	if err != nil {
		return err
	}

	/*
		err = svc.dao.DeleteFishpowerTag(param.ID)
		if err != nil {
			return err
		}
	*/

	return nil
}
