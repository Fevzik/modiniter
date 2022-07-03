package modiniter

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

func GetAccessWare() fiber.Handler {
	return jwtware.New(jwtware.Config{
		//Filter: DefaultMiddlewareFilter,
		ErrorHandler:  jwtError,
		SigningKey:    []byte(JWTSigningKey),
		SigningMethod: jwt.SigningMethodHS512.Name,
		TokenLookup:   "header:" + fiber.HeaderAuthorization,
		AuthScheme:    "Bearer",
		ContextKey:    CurrentUserKey,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Некорректный токен", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Токен не существует", "data": nil})
}

func GetUserIdFromToken(c *fiber.Ctx) string {
	userID := ""
	u := c.Locals(CurrentUserKey)
	if u != nil {
		claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
		userID = claims[UserIdKey].(string)
	}
	return userID
}

func GetUserRoleFromToken(c *fiber.Ctx) []int64 {
	userRoles := []int64{}
	u := c.Locals(CurrentUserKey)
	if u != nil {
		claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
		roles, ok := claims[UserRolesKey].([]interface{})
		if ok {
			for _, v := range roles {
				r, ok := v.(float64)
				if ok {
					userRoles = append(userRoles, int64(r))
				}
			}
		}
	}
	return userRoles
}

func GetTokenClaims(c *fiber.Ctx) jwt.MapClaims {
	u := c.Locals(CurrentUserKey)
	if u != nil {
		claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
		return claims
	}
	return jwt.MapClaims{}
}
