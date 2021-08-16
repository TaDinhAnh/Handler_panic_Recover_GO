package todobus

import (
	"context"
	"yearbook/common"
	todomodel "yearbook/modules/todo/model"
)

type GetTodoStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}) (*todomodel.Todo, error)
}

type getTodoBus struct {
	store GetTodoStore
}

func NewGetTodoBus(store GetTodoStore) *getTodoBus {
	return &getTodoBus{store: store}
}

func (bus *getTodoBus) GetTodo(ctx context.Context, id int) (*todomodel.Todo, error) {
	result, err := bus.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {

		return nil, common.ErrCannotGetEntity(todomodel.EntityName, err)
	}
	return result, nil
}
