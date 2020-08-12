package main

import (
    gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
    "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

    "github.com/gin-gonic/gin"
    "time"
    "math/rand"
 //   "log"
)

func main() {
    // tracer.Start(tracer.WithAnalytics(true))
    tracer.Start(
        tracer.WithAgentAddr("datadog-agent:8126"),
        tracer.WithAnalytics(true),
    )
    defer tracer.Stop()

    r := gin.Default()
    
    // Use the tracer middleware with your desired service name.
    r.Use( gintrace.Middleware("emu-back") )

	r.GET("/eat", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"message": "You ate new cookies.",
		})
	})

	r.GET("/sleep", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        length := rand.Intn(20)
        time.Sleep(time.Duration(length + 20) * time.Second)
		c.JSON(200, gin.H{
			"message": "you slept and feel refreshed",
		})
    })
    
    r.GET("/get-emu", func(c *gin.Context) {
        

        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON( 500, gin.H{ "result": "bad" } )
   //      c.JSON(200, gin.H{
			// "result": "/img/noemu.png", 
   //      })    

    })

    r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
