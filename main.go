package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)
var (
	h bool
	defaultRobot string
)

func init()  {
	flag.BoolVar(&h,"h",false,"help")
	flag.StringVar(&defaultRobot,"defaultRobot","", "global dingtalk robot webhook")
}
func main()  {
	flag.Parse()
	if h{
		flag.Usage()
		return
	}
	router:=gin.New()
	router.POST("/webhook", func(context *gin.Context) {
		var notification  Notification
		err := context.BindJSON(&notification)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
			return
		}
		err = Send(notification, defaultRobot)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
		}
		context.JSON(http.StatusOK,gin.H{
			"message": "successful send alert notification message to DingTalk!",
		})
	})
	router.Run()
}


