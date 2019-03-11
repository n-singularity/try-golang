package Service

type AbstractService struct {
	Service InterfaceService
	x int
}

func ClassAbstractService(service InterfaceService) AbstractService {
	var abstractEntity AbstractService
	abstractEntity.Service = service
	return abstractEntity
}

func (it *AbstractService) SetX(input int) {
	it.x = input
}

func (it AbstractService) GetX() int{
	return it.x
}

func (it *AbstractService) Sum() int {
	return it.GetX()+it.Service.GetY()+it.Service.GetZ()
}

func (it *AbstractService) UpdateService(service InterfaceService) {
	it.Service = service
}
