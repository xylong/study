package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"study/db/mysql/src/config"
	"study/db/mysql/src/models"
)

const Port = 8080

func GinBoot() {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.NewUserModel().LoadByID(1))
	})

	go func() {
		if err := r.Run(fmt.Sprintf(":%d", Port)); err != nil {
			config.ErrorChan <- err
		}
	}()
}
