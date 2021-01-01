package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/dao"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type FishRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type FishListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateFishRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateFishRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=2,max=100"`
	Desc          string `form:"desc" binding:"min=2,max=255"`
	Content       string `form:"content" binding:"min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteFishRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Fish struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

func (svc *Service) GetFish(param *FishRequest) (*Fish, error) {
	fish, err := svc.dao.GetFish(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	fishTag, err := svc.dao.GetFishTagByAID(fish.ID)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(fishTag.TagID, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}

	return &Fish{
		ID:            fish.ID,
		Title:         fish.Title,
		Desc:          fish.Desc,
		Content:       fish.Content,
		CoverImageUrl: fish.CoverImageUrl,
		State:         fish.State,
		Tag:           &tag,
	}, nil
}

func (svc *Service) GetFishList(param *FishListRequest, pager *app.Pager) ([]*Fish, int, error) {
	fishCount, err := svc.dao.CountFishListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	fishs, err := svc.dao.GetFishListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var fishList []*Fish
	for _, fish := range fishs {
		fishList = append(fishList, &Fish{
			ID:            fish.FishID,
			Title:         fish.FishTitle,
			Desc:          fish.FishDesc,
			Content:       fish.Content,
			CoverImageUrl: fish.CoverImageUrl,
			Tag:           &model.Tag{Model: &model.Model{ID: fish.TagID}, Name: fish.TagName},
		})
	}

	return fishList, fishCount, nil
}

func (svc *Service) CreateFish(param *CreateFishRequest) error {
	fish, err := svc.dao.CreateFish(&dao.Fish{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.CreateFishTag(fish.ID, param.TagID, param.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateFish(param *UpdateFishRequest) error {
	err := svc.dao.UpdateFish(&dao.Fish{
		ID:            param.ID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		ModifiedBy:    param.ModifiedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.UpdateFishTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteFish(param *DeleteFishRequest) error {
	err := svc.dao.DeleteFish(param.ID)
	if err != nil {
		return err
	}

	err = svc.dao.DeleteFishTag(param.ID)
	if err != nil {
		return err
	}

	return nil
}
