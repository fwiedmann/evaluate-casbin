package user

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepositoryMariaDB struct {
}

func (r RepositoryMariaDB) Create(m Model) error {
	db, err := gorm.Open(sqlite.Open("eval.db"))
	if err != nil {
		return err
	}

	db.AutoMigrate(&Model{})
	return db.Create(&m).Error
}

func (r RepositoryMariaDB) FindById(id string) (Model, error) {
	db, err := gorm.Open(sqlite.Open("eval.db"))
	if err != nil {
		return Model{}, err
	}

	db.AutoMigrate(&Model{})
	var m Model
	err = db.First(&m, "ID = ?", id).Error
	return m, err
}
