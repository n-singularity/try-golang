package Service

type MathService struct {
	AbstractService
	y int
	z int
}

func ClassMathService() MathService {
	var service MathService
	service.AbstractService = ClassAbstractService(service)
	return service
}

func (thisIs *MathService) SetY(input int){
	thisIs.y = input
	thisIs.updateParent()
}

func (thisIs *MathService) SetZ(input int) {
	thisIs.z = input
	thisIs.updateParent()
}

func (thisIs MathService) GetY() int{
	return thisIs.y
}

func (thisIs MathService) GetZ() int {
	return thisIs.z
}

func (thisIs *MathService) updateParent() {
	thisIs.AbstractService.UpdateService(thisIs)
}

