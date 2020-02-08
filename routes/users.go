package users

import (
	"log"

	"github.com/gin-gonic/gin"
)

func createPost(c *gin.Context) {
	var form main.Tweet
	if err := c.Bind(&form); err != nil {
		log.Println("errは、、")
		log.Println(err)
		tweets := dbGetAll()
		c.HTML(200, "index.html", gin.H{"tweets": tweets, "err": err})
		c.Abort()
	} else {
		content := c.PostForm("content")
		dbInsert(content)
		c.Redirect(302, "/")
	}

}
