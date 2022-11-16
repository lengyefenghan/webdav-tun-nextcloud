package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func helloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello q1mi!",
// 	})
// }

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"installed":       true,
		"maintenance":     false,
		"needsDbUpgrade":  false,
		"version":         "22.1.1.2",
		"versionstring":   "22.1.1",
		"edition":         "",
		"productname":     "Nextcloud",
		"extendedSupport": false,
	})
}

// {
// "installed": true,
// "maintenance": false,
// "needsDbUpgrade": false,
// "version": "22.1.1.2",
// "versionstring": "22.1.1",
// "edition": "",
// "productname": "Nextcloud",
// "extendedSupport": false
// }

func remote_dav(c *gin.Context) {
	log.Println(c)
	c.JSON(http.StatusOK, gin.H{
		"installed":       true,
		"maintenance":     false,
		"needsDbUpgrade":  false,
		"version":         "22.1.1.2",
		"versionstring":   "22.1.1",
		"edition":         "",
		"productname":     "Nextcloud",
		"extendedSupport": false,
	})
}
