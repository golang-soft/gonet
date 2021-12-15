package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//websocket实现
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

//func main() {
//	r := gin.Default()
//	gin.SetMode("debug")
//	r.GET("/", ping)
//	r.GET("/socket.io/", func(c *gin.Context) {
//		user := c.Query("user") //查询请求URL后面的参数
//		role := c.Query("role") //查询请求URL后面的参数
//		fmt.Sprintf("%s, %s", user, role)
//		c.JSON(200,"connect")
//	})
//	r.Run("192.168.1.233:3000")
//}
