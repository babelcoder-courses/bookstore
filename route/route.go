package route

import "github.com/gin-gonic/gin"

func Serve(r *gin.Engine) {
	serveV1(r)
}
