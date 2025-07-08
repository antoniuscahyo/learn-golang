package todo_list_category_usecase

import (
	"context"
	"fmt"
	"time"

	generalEntity "github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/entity"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/helper"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/repository/mysql"
	mentity "github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/repository/mysql/entity"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/usecase"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/usecase/todo_list_category/entity"
	errwrap "github.com/pkg/errors"
)

type CrudTodoListCategoryUsecase struct {
	todoCategoryRepo mysql.ITodoCategoryRepository
}

func NewCrudTodoListCategoryUsecase(repo mysql.ITodoCategoryRepository) *CrudTodoListCategoryUsecase {
	return &CrudTodoListCategoryUsecase{repo}
}

type ICrudTodoListCategoryUsecase interface {
	GetByID(ctx context.Context, categoryID int64) (*entity.TodoListCategoryResponse, error)
	GetAll(ctx context.Context) ([]*entity.TodoListCategoryResponse, error)
	Create(ctx context.Context, req entity.TodoListCategoryReq) (*entity.TodoListCategoryResponse, error)
	UpdateByID(ctx context.Context, req entity.TodoListCategoryReq) error
	DeleteByID(ctx context.Context, categoryID int64) error
}

func (u *CrudTodoListCategoryUsecase) GetAll(ctx context.Context) ([]*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.GetAll"
	result, err := u.todoCategoryRepo.GetAll(ctx)
	if err != nil {
		helper.LogError("todoCategoryRepo.GetAll", funcName, err, nil, "")
		return nil, err
	}

	var res []*entity.TodoListCategoryResponse
	for _, v := range result {
		res = append(res, &entity.TodoListCategoryResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   helper.ConvertToJakartaTime(v.CreatedAt),
		})
	}
	return res, nil
}

func (u *CrudTodoListCategoryUsecase) GetByID(ctx context.Context, categoryID int64) (*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.GetByID"
	capture := generalEntity.CaptureFields{"category_id": helper.ToString(categoryID)}

	data, err := u.todoCategoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		helper.LogError("todoCategoryRepo.GetByID", funcName, err, capture, "")
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.TodoListCategoryResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
	}, nil
}

func (u *CrudTodoListCategoryUsecase) Create(ctx context.Context, req entity.TodoListCategoryReq) (*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.Create"
	capture := generalEntity.CaptureFields{"payload": helper.ToString(req)}

	if errMsg := usecase.ValidateStruct(req); errMsg != "" {
		return nil, errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	payload := &mentity.TodoListCategory{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	err := u.todoCategoryRepo.Create(ctx, nil, payload, false)
	if err != nil {
		helper.LogError("todoCategoryRepo.Create", funcName, err, capture, "")
		return nil, err
	}

	return &entity.TodoListCategoryResponse{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		CreatedAt:   helper.ConvertToJakartaTime(payload.CreatedAt),
	}, nil
}

func (u *CrudTodoListCategoryUsecase) UpdateByID(ctx context.Context, req entity.TodoListCategoryReq) error {
	funcName := "CrudTodoListCategoryUsecase.UpdateByID"
	capture := generalEntity.CaptureFields{"payload": helper.ToString(req)}

	return mysql.DBTransaction(u.todoCategoryRepo, func(trx mysql.TrxObj) error {
		locked, err := u.todoCategoryRepo.LockByID(ctx, trx, req.ID)
		if err != nil {
			helper.LogError("todoCategoryRepo.LockByID", funcName, err, capture, "")
			return err
		}
		if locked == nil {
			return fmt.Errorf("DATA NOT FOUND")
		}

		changes := &mentity.TodoListCategory{
			Name:        req.Name,
			Description: req.Description,
		}

		if err := u.todoCategoryRepo.Update(ctx, trx, locked, changes); err != nil {
			helper.LogError("todoCategoryRepo.Update", funcName, err, capture, "")
			return err
		}

		return nil
	})
}

func (u *CrudTodoListCategoryUsecase) DeleteByID(ctx context.Context, categoryID int64) error {
	funcName := "CrudTodoListCategoryUsecase.DeleteByID"
	capture := generalEntity.CaptureFields{"category_id": helper.ToString(categoryID)}

	err := u.todoCategoryRepo.DeleteByID(ctx, nil, categoryID)
	if err != nil {
		helper.LogError("todoCategoryRepo.DeleteByID", funcName, err, capture, "")
		return err
	}

	return nil
}
