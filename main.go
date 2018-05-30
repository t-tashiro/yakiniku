// @APIVersion 1.0.0
// @APITitle yakiniku blog
// @APIDescription yakiniku blog sample
// @BasePath http://localhost:8081/swagger-ui

package main

import (
	"github.com/labstack/echo"
	swaggerui "github.com/t-tashiro/yakiniku/swagger-ui"
)

func main() {
	e := echo.New()

	swaggerui.NewSwaggerController(e)

	e.Logger.Fatal(e.Start(":8081"))
}
