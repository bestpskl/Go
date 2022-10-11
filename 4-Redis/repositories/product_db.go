package repositories

import "gorm.io/gorm"

// 	adapter

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate([]product{})
	mockData(db)
	return productRepositoryDB{db: db}
}

//	method

func (r productRepositoryDB) GetProducts() (products []product, err error) {
	err = r.db.Order("quantity desc").Limit(20).Find(&products).Error
	return products, err
}
