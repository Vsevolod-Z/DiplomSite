package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func (h *Handler) main(c *gin.Context) {
	fmt.Println(c.ClientIP())
	http.ServeFile(c.Writer, c.Request, "./internal/markup/index.html")

}
func (h *Handler) game(c *gin.Context) {
	fmt.Println(c.ClientIP())
	http.ServeFile(c.Writer, c.Request, "./internal/markup/gamePage.html")
}
