package constant

import "github.com/gofiber/fiber/v2"

const (
	jeketiErrorCodeCommon     = "JKTEC_1"
	jeketiErrorCodeHandler    = "JKTEC_2"
	jeketiErrorCodeMiddleware = "JKTEC_3"
	jeketiErrorCodeService    = "JKTEC_4"
)

var (
	InternalServerError = fiber.Map{
		"code":       500,
		"jeketiCode": jeketiErrorCodeCommon + "_1",
		"message":    "Internal Server Error",
	}
)

var (
	BindingError = fiber.Map{
		"code":       401,
		"jeketiCode": jeketiErrorCodeHandler + "_1",
		"message":    "Binding validation error",
	}
)

var (
	InvalidPasswordError = fiber.Map{
		"code":       401,
		"jeketiCode": jeketiErrorCodeService + "_1",
		"message":    "Invalid Password",
	}
	DataNotFoundError = fiber.Map{
		"code":       404,
		"jeketiCode": jeketiErrorCodeService + "_2",
		"message":    "Data not found",
	}
)
