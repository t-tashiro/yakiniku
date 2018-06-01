// @APIVersion 1.0.0
// @APITitle yakiniku blog
// @APIDescription yakiniku blog sample
// @BasePath http://localhost:8081/swagger-ui

package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	setup(e)

	// swaggerui.NewSwaggerController(e)

	e.Logger.Fatal(e.Start(":8081"))
}
