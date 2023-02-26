package main

import (
	"github.com/SuyunovJasurbek/url_shorting/src/handler"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.basic  BasicAuth
func main() {
	r := handler.SetupRouter()
	r.Run(":8080")
}
