package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/noworldwar/bc-service-livestream-api-go/internal/pkg/utils"
)

func ApiService(c *gin.Context) {

	// data := c.PostForm("data")

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "action", "data"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check Data Format
	data := map[string]string{}
	err := json.Unmarshal([]byte(c.PostForm("data")), &data)
	if err != nil {
		utils.ErrorResponse(c, 400, "Data Format Error", nil)
		return
	}
	fmt.Println(data["data"])

	// Step 3: Take The Corresponding Action
	switch action := c.PostForm("action"); action {
	case "balance":
		getBalance(c, data)
	case "fund_transfer":
		fundTransfer(c, data)
	}
}

func getBalance(c *gin.Context, data map[string]string) {

	if missing := utils.CheckMapData(data, "username"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing Required Parameter: "+missing, nil)
		return
	}

	c.JSON(200, gin.H{"data": data})
	return
}

func fundTransfer(c *gin.Context, data map[string]string) {

	if missing := utils.CheckMapData(data, "username", "tran_id", "amount", "currency_code"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing Required Parameter: "+missing, nil)
		return
	}

	c.JSON(200, gin.H{"data": data})
	return
}
