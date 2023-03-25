package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
	"fmt"
	// "io/ioutil"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

type RBody struct {
    Email string `json:"email"`
}


// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postDate(c *gin.Context) {
	var rb RBody
	if err := c.BindJSON(&rb); err != nil {
		c.IndentedJSON(422, "email is missing")
		return
    }
	fmt.Println(rb)
	c.IndentedJSON(http.StatusOK, rb)
}

func getHome(c *gin.Context){
	c.IndentedJSON(http.StatusOK, "wlcome")
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
	router.GET("/", getHome)
	router.POST("/data", postDate)
    router.Run("localhost:8080")
}
