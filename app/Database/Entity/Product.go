package Entity

import (
	"github.com/jinzhu/gorm"
	"go/types"
)

type Product struct {
	AbstractEntity
	gorm.Model
	Code string
	Price uint
}

func ClassProduct() Product {
	var interfaceEntity InterfaceEntity
	var product Product
	product.AbstractEntity = ClassAbstractEntity(product)
	interfaceEntity = product
	_=interfaceEntity

	return product
}

func (it Product) GetId()  uint{
	return it.ID
}

func (it Product) ToArray(arrayType string, getParent bool, getChild bool, addFunction types.Slice)  {

}

func (it Product) Rule()  {

}

func (it *Product) updateParent() {
	it.AbstractEntity.Entity = it
}