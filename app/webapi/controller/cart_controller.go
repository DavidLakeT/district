package controller

import service "district/service/handler"

type CartController struct {
	servicePool *service.ServicePool
}

func NewCartController(servicePool *service.ServicePool) *CartController {
	return &CartController{servicePool: servicePool}
}
