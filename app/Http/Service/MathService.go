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

func (it *MathService) SetY(input int){
	it.y = input
	it.updateParent()
}

func (it *MathService) SetZ(input int) {
	it.z = input
	it.updateParent()
}

func (it MathService) GetY() int{
	return it.y
}

func (it MathService) GetZ() int {
	return it.z
}

func (it *MathService) updateParent() {
	it.AbstractService.UpdateService(it)
}

