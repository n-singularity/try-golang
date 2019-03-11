package Entity

import "github.com/jinzhu/gorm"

type AbstractEntity struct {
	Entity InterfaceEntity
}

func ClassAbstractEntity(entity InterfaceEntity) AbstractEntity {
	var abstractEntity AbstractEntity
	abstractEntity.Entity = entity
	return abstractEntity
}

func (it AbstractEntity) Save() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	if db.NewRecord(it.Entity) {
		db.Create(&it.Entity)
	} else {
		db.Save(&it.Entity)
	}

	err = db.Close()
}

func (it AbstractEntity) Remove() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	if !db.NewRecord(it.Entity) {
		db.Delete(&it.Entity)
	}

	err = db.Close()
}

func (it *AbstractEntity) UpdateEntityValue(entity InterfaceEntity) {
	it.Entity = entity
}
