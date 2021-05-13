package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"study/db/mysql/src/config"
	"study/db/mysql/src/models"
)

const Port = 8080

func GinBoot() {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.NewUserModel().LoadByID(1))
	})
	r.GET("/users/:name", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.NewUserModel().LoadByName(ctx.Param("name")))
	})
	r.GET("/age/:age", func(ctx *gin.Context) {
		user := models.NewUserModel()
		age, _ := strconv.Atoi(ctx.Param("age"))
		user.Filter(user.AgeCompare(age, models.GraterThan))
		ctx.JSON(http.StatusOK, user)
	})

	go func() {
		if err := r.Run(fmt.Sprintf(":%d", Port)); err != nil {
			config.ErrorChan <- err
		}
	}()
}
