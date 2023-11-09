package controller

import (
	service "district/service/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UtilsController struct {
	servicePool *service.ServicePool
}

func NewUtilsController(servicePool *service.ServicePool) *UtilsController {
	return &UtilsController{servicePool: servicePool}
}

// Endpoint: DELETE /api/database
// Clears all the database (deleting tables).
func (uc *UtilsController) ClearDatabase(c echo.Context) error {
	if err := uc.servicePool.UtilsService.ClearDatabase(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "database successfully cleared.",
	})
}
