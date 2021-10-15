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

	router.POST("/sendgroup", SendGroupChat)
	router.POST("/sendchannel", SendChannel)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()

}

var priorityTable = map[string][]*http.Request{
	"low":    make([]*http.Request, 0),
	"medium": make([]*http.Request, 0),
	"high":   make([]*http.Request, 0),
}

// @Summary Send message
// @Description Send message to group chat
// @ID send-message-group-chat
// @Accept  json
// @Produce  json
// @Param text query string true "message text"
// @Param priority query string true "priority message"
// @Router /sendgroup [post]
func SendGroupChat(c *gin.Context) {
	client := http.Client{}

	err := createRequest(&client, URL, "-1001317290790", c.Query("text"), c.Query("priority"))
	if err != nil {
		log.Fatal(err)
	}

	doRequest(&client)
}

// @Summary Send message
// @Description Send message to channel
// @ID send-message-channel
// @Accept  json
// @Produce  json
// @Param text query string true "message text"
// @Param priority query string true "priority message"
// @Router /sendchannel [post]
func SendChannel(c *gin.Context) {
	client := http.Client{}

	err := createRequest(&client, URL, "-1001652337843", c.Query("text"), c.Query("priority"))
	if err != nil {
		log.Fatal(err)
	}
	doRequest(&client)
}

func createRequest(cl *http.Client, url, chatId, text, priority string) error {
	post_url := fmt.Sprintf("%s%schat_id=%s&text=%s", url, "sendMessage?", chatId, text)

	req, err := http.NewRequest("POST", post_url, nil)
	if err != nil {
		return err
	}

	priorityTable[priority] = append(priorityTable[priority], req)

	return err
}

func doRequest(cl *http.Client) {
	time.AfterFunc(20*time.Second, func() {
		for i, v := range priorityTable["high"] {
			cl.Do(v)
			priorityTable["high"] = append(priorityTable["high"][:i], priorityTable["high"][i+1:]...)
			return
		}
		for i, v := range priorityTable["medium"] {
			cl.Do(v)
			priorityTable["medium"] = append(priorityTable["medium"][:i], priorityTable["medium"][i+1:]...)
			return
		}
		for i, v := range priorityTable["low"] {
			cl.Do(v)
			priorityTable["low"] = append(priorityTable["low"][:i], priorityTable["low"][i+1:]...)
			return
		}
	})
}
