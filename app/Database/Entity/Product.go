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

func (p Product) GetId()  uint{
	return p.ID
}

func (p Product) ToArray(arrayType string, getParent bool, getChild bool, addFunction types.Slice)  {

}

func (p Product) Rule()  {

}