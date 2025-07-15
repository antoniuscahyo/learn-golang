package mysql

import (
	"context"

	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/config"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/helper"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/repository/mysql/entity"

	a "github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/error"

	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type ITodoCategoryRepository interface {
	TrxSupportRepo
	GetAll(ctx context.Context) ([]*entity.TodoListCategoryResponse, error)
	GetByID(ctx context.Context, ID int64) (*entity.TodoListCategoryResponse, error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error
	Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, changes *entity.TodoListCategory) error
	DeleteByID(ctx context.Context, dbTrx TrxObj, ID int64) error
	LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (*entity.TodoListCategory, error)
}

type TodoListCategoryRepository struct {
	GormTrxSupport
}

func NewTodoListCategoryRepository(mysql *config.Mysql) *TodoListCategoryRepository {
	return &TodoListCategoryRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TodoListCategoryRepository) GetAll(ctx context.Context) ([]*entity.TodoListCategoryResponse, error) {
	funcName := "TodoListCategoryRepository.GetAll"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var result []*entity.TodoListCategoryResponse
	err := r.db.Raw(`
		SELECT tlc.id, tlc.name, tlc.description, users.name as created_by, tlc.created_at 
		FROM todo_list_categories tlc
		JOIN users ON tlc.created_by = users.id`).Scan(&result).Error
	if err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	return result, nil
}

func (r *TodoListCategoryRepository) GetByID(ctx context.Context, ID int64) (*entity.TodoListCategoryResponse, error) {
	funcName := "TodoListCategoryRepository.GetByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var result *entity.TodoListCategoryResponse
	err := r.db.Raw(`SELECT tlc.id, tlc.name, tlc.description, users.name as created_by, tlc.created_at 
		FROM todo_list_categories tlc
		JOIN users ON tlc.created_by = users.id
		WHERE tlc.id = ? LIMIT 1`, ID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, a.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoListCategoryRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error {
	funcName := "TodoListCategoryRepository.Create"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *TodoListCategoryRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, changes *entity.TodoListCategory) error {
	funcName := "TodoListCategoryRepository.Update"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	db := r.Trx(dbTrx).Model(params)
	if changes != nil {
		return db.Updates(*changes).Error
	}
	return db.Updates(helper.StructToMap(params, false)).Error
}

func (r *TodoListCategoryRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, ID int64) error {
	funcName := "TodoListCategoryRepository.DeleteByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	return r.Trx(dbTrx).Where("id = ?", ID).Delete(&entity.TodoListCategory{}).Error
}

func (r *TodoListCategoryRepository) LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (*entity.TodoListCategory, error) {
	funcName := "TodoListCategoryRepository.LockByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var result *entity.TodoListCategory
	err := r.Trx(dbTrx).
		Raw("SELECT * FROM todo_list_categories WHERE id = ? FOR UPDATE", ID).
		Scan(&result).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, a.ErrRecordNotFound()
	}

	return result, err
}
