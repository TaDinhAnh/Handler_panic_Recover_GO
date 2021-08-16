package todotransport

import (
	"net/http"
	"strconv"
	"yearbook/common"
	"yearbook/component/appctx"
	todobus "yearbook/modules/todo/bus"
	todostorage "yearbook/modules/todo/storage"

	"github.com/gin-gonic/gin"
)

func GetTodo(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := todostorage.NewSQLStore(appCtx.GetMainDbConnection())
		bus := todobus.NewGetTodoBus(store)
		data, err := bus.GetTodo(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}
