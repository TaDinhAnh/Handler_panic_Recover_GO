package todostorage

import (
	"context"
	"yearbook/common"
	todomodel "yearbook/modules/todo/model"
)

func (s *sqlStore) FindDataByCondition(ctx context.Context, conditions map[string]interface{}) (*todomodel.Todo, error) {
	var rs todomodel.Todo
	db := s.db
	if err := db.Where(conditions).First(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &rs, nil
}
