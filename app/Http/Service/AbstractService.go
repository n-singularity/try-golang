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

func (thisIs *AbstractService) SetX(input int) {
	thisIs.x = input
}

func (thisIs AbstractService) GetX() int{
	return thisIs.x
}

func (thisIs *AbstractService) Sum() int {
	return thisIs.GetX()+thisIs.Service.GetY()+thisIs.Service.GetZ()
}

func (thisIs *AbstractService) UpdateService(service InterfaceService) {
	thisIs.Service = service
}
