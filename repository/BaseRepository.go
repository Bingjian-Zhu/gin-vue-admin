package repository

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
)

//BaseRepository 注入IDb
type BaseRepository struct {
	Source datasource.IDb `inject:""`
}

// Create 创建实体
func (b *BaseRepository) Create(value interface{}) error {
	return b.Source.DB().Create(value).Error
}

// Save 保存实体
func (b *BaseRepository) Save(value interface{}) error {
	return b.Source.DB().Save(value).Error
}

// // Updates 更新实体
// func (b *BaseRepository) Updates(model interface{}, value interface{}) error {
// 	return b.Source.DB().Model(model).Updates(value).Error
// }

// DeleteByWhere 根据条件删除实体
func (b *BaseRepository) DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := b.Source.DB().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByID 根据id删除实体
func (b *BaseRepository) DeleteByID(model interface{}, id uint64) (count int64, err error) {
	db := b.Source.DB().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByIDS 根据多个id删除多个实体
func (b *BaseRepository) DeleteByIDS(model interface{}, ids []uint64) (count int64, err error) {
	db := b.Source.DB().Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// First 根据条件获取一个实体
func (b *BaseRepository) First(where interface{}, out interface{}, selects ...string) error {
	db := b.Source.DB().Where(where)
	if len(selects) > 0 {
		for _, sel := range selects {
			db = db.Select(sel)
		}
	}
	return db.First(out).Error
}

// Find 根据条件返回数据
func (b *BaseRepository) Find(where interface{}, out interface{}, sel string, orders ...string) error {
	db := b.Source.DB().Where(where)
	if sel != "" {
		db = db.Select(sel)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// GetPage 分页返回数据
func (b *BaseRepository) GetPage(model interface{}, out interface{}, pageIndex, pageSize int, totalCount *uint64, maps map[string]interface{}, orders ...string) error {
	db := b.Source.DB().Model(model).Where(model)
	db = db.Where(maps)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// PluckList 查询 model 中的一个列作为切片
func (b *BaseRepository) PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return b.Source.DB().Model(model).Where(where).Pluck(fieldName, out).Error
}
