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
	GetAll(ctx context.Context) ([]*entity.TodoCategory, error)
	GetByID(ctx context.Context, ID int64) (*entity.TodoCategory, error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoCategory, nonZeroVal bool) error
	Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoCategory, changes *entity.TodoCategory) error
	DeleteByID(ctx context.Context, dbTrx TrxObj, ID int64) error
	LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (*entity.TodoCategory, error)
}

type TodoCategoryRepository struct {
	GormTrxSupport
}

func NewTodoCategoryRepository(mysql *config.Mysql) *TodoCategoryRepository {
	return &TodoCategoryRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TodoCategoryRepository) GetAll(ctx context.Context) ([]*entity.TodoCategory, error) {
	funcName := "TodoCategoryRepository.GetAll"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var result []*entity.TodoCategory
	err := r.db.Raw("SELECT * FROM todo_list_categories").Scan(&result).Error
	if err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	return result, nil
}

func (r *TodoCategoryRepository) GetByID(ctx context.Context, ID int64) (*entity.TodoCategory, error) {
	funcName := "TodoCategoryRepository.GetByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var result *entity.TodoCategory
	err := r.db.Raw("SELECT * FROM todo_list_categories WHERE id = ? LIMIT 1", ID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, a.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoCategoryRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoCategory, nonZeroVal bool) error {
	funcName := "TodoCategoryRepository.Create"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *TodoCategoryRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoCategory, changes *entity.TodoCategory) error {
	funcName := "TodoCategoryRepository.Update"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	db := r.Trx(dbTrx).Model(params)
	if changes != nil {
		return db.Updates(*changes).Error
	}
	return db.Updates(helper.StructToMap(params, false)).Error
}

func (r *TodoCategoryRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, ID int64) error {
	funcName := "TodoCategoryRepository.DeleteByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	return r.Trx(dbTrx).Where("id = ?", ID).Delete(&entity.TodoCategory{}).Error
}

func (r *TodoCategoryRepository) LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (*entity.TodoCategory, error) {
	funcName := "TodoCategoryRepository.LockByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	var result *entity.TodoCategory
	err := r.Trx(dbTrx).
		Raw("SELECT * FROM todo_list_categories WHERE id = ? FOR UPDATE", ID).
		Scan(&result).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, a.ErrRecordNotFound()
	}

	return result, err
}
