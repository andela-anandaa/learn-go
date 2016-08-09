package main


import (
    // "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // r := gin.Default()
    // fmt.Printf("Type %T\n", r)

    router := load_routes()

    router.Run(":9090") // listen and server on 0.0.0.0:9090
}

func load_routes() *gin.Engine {
     // set mode, default is Debug mode
    // gin.SetMode(gin.ReleaseMode)

    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    router.GET("/admin", func(c *gin.Context) {
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": "unauthorized access",
        })
    })

    // parameters in path
    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })

    router.GET("/user/:name/*action", func(c *gin.Context) {
        name := c.Param("name")
        action := c.Param("action")
        message := name + " is " + action
        c.String(http.StatusOK, message)
    })

    // query string
    router.GET("/welcome", func(c *gin.Context) {
        first_name := c.DefaultQuery("fname", "Guest")
        last_name := c.Query("lname") // shortcut for c.Request.URL.Query().Ger("lname")

        c.String(http.StatusOK, "Hello %s %s", first_name, last_name)
    })

    // mutli-part/urlencoded form
    router.POST("/form_post", func(c *gin.Context) {
        message := c.PostForm("message")
        nick := c.DefaultPostForm("nick", "anonymous")

        c.JSON(200, gin.H{
            "status": "posted",
            "message": message,
            "nick": nick,
        })
    })

    // more on JSON response
    router.GET("/json", func(c *gin.Context) {
        var msg struct {
            Name string `json:"user"`
            Message string
            Number int
        }

        msg.Name = "John"
        msg.Message = "Hello gin"
        msg.Number = 800

        c.JSON(200, msg)
    })

    return router
}
