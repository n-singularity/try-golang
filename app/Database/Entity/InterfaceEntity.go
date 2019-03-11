package Entity

import "go/types"

type InterfaceEntity interface {
	ToArray(arrayType string, getParent bool, getChild bool, addFunction types.Slice)
	Rule()
	Save()
	Remove()
	GetId() uint
}
