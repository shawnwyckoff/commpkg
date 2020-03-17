package gweb

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Ctx struct {
		ctx *gin.Context
	}

	ErrFormatter func(err error) string
)

func (c Ctx) GetParam(name string) string {
	return c.ctx.Params.ByName(name)
}

func (c Ctx) WriteString(code int, format string, values... interface{}) {
	c.ctx.String(code, format, values...)
}

func (c Ctx) WriteMapJSON(code int, values map[string]interface{}) {
	c.ctx.JSON(code, values)
}

func (c Ctx) WriteStructJSON(code int, output interface{}, errFmt ErrFormatter) {
	buf, err := json.Marshal(output)
	if err != nil {
		c.WriteString(http.StatusInternalServerError, errFmt(err))
		return
	}
	c.ctx.String(code, string(buf))
}