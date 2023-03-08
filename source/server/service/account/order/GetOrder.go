package serviceOrder

import (
	"server/service/h"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	h.Ok(c, "ok")
}
