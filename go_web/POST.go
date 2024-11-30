package go_web

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	Name    string `json:"name"`
// 	Address string `json:"address"`
// 	Age     int    `json:"age"`
// }

// func addUser(c *gin.Context) {
// 	user := User{}
// 	if error := c.ShouldBindJSON(&user); error != nil {
// 		fmt.Println("出错了")
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg":    error.Error(),
// 			"status": "400",
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, user)

// 	}
// }
// func main() {
// 	r := gin.Default()
// 	r.POST("/addUser", addUser)
// 	r.Run()
// }
