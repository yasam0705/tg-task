package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var URL = "https://api.telegram.org/bot1822246375:AAFBs9rUJ1wHJpweTlFHSOPuVXUfJQoKpTc/"

func InitRoutes() {
	router := gin.Default()

	router.GET("/sendgroup", SendGroupChat)
	router.GET("/sendchannel", SendChannel)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()

}

// @Summary Send message
// @Description Send message to group chat
// @ID send-message-group-chat
// @Accept  json
// @Produce  json
// @Param text query string true "message text"
// @Router /sendgroup [get]
func SendGroupChat(c *gin.Context) {
	client := http.Client{}

	var tt time.Duration = 5 * time.Second
	time.AfterFunc(tt, func() {
		_, err := sendRequest(&client, URL, "-1001317290790", c.Query("text"))
		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Fprintf(c.Writer, "<h1>Send message to telegram group</h1>")
}

// @Summary Send message
// @Description Send message to channel
// @ID send-message-channel
// @Accept  json
// @Produce  json
// @Param text query string true "message text"
// @Router /sendchannel [get]
func SendChannel(c *gin.Context) {
	client := http.Client{}

	res, err := sendRequest(&client, URL, "-1001652337843", c.Query("text"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(c.Writer, "<h1>Send message to telegram channel</h1>")
	fmt.Fprint(c.Writer, res)
}

func sendRequest(cl *http.Client, url, chatId, text string) (*http.Response, error) {
	post_url := fmt.Sprintf("%s%schat_id=%s&text=%s", url, "sendMessage?", chatId, text)

	req, err := http.NewRequest("POST", post_url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	res, err := cl.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return res, err
}
