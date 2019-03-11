package Entity

import "github.com/jinzhu/gorm"

type AbstractEntity struct {
	Entity interface{}
}

func ClassAbstractEntity(entity InterfaceEntity) AbstractEntity {
	var abstractEntity AbstractEntity
	abstractEntity.Entity = entity
	return abstractEntity
}

func (thisIs AbstractEntity) Save() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	if db.NewRecord(thisIs.Entity) {
		db.Create(&thisIs.Entity)
	} else {
		db.Save(&thisIs.Entity)
	}

	err = db.Close()
}

func (thisIs AbstractEntity) Remove() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	if !db.NewRecord(thisIs.Entity) {
		db.Delete(&thisIs.Entity)
	}

	err = db.Close()
}
