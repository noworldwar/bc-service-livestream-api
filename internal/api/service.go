package api

import (
	"context"
	"encoding/json"
	"io/ioutil"

	ppf "github.com/PGITAb/bc-proto-entity-playerprofile-go"
	wal "github.com/PGITAb/bc-proto-wallet-go"
	"github.com/gin-gonic/gin"
	"github.com/noworldwar/bc-service-livestream-api-go/internal/pkg/utils"
	"github.com/noworldwar/bc-service-livestream-api-go/internal/srvclient"
)

func ApiService(c *gin.Context) {

	// Step 1: Read Json Data
	req, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.ErrorResponse(c, 103, "Data Content Error: "+err.Error())
		return
	}

	// Step 2: Check Data Format
	var m map[string]interface{}
	err = json.Unmarshal(req, &m)
	if err != nil {
		utils.ErrorResponse(c, 103, "JSON Format Error: "+err.Error())
		return
	}

	// Step 3: Take The Corresponding Action
	switch action := m["action"]; action {
	case "balance":
		getBalance(c, m["data"])
	case "fund_transfer":
		fundTransfer(c, m["data"])
	}

}

func getBalance(c *gin.Context, data interface{}) {

	// Step 1: Check Parameter
	if missing := utils.CheckJsonDataContent(data, "username"); missing != "" {
		utils.ErrorResponse(c, 103, "Missing parameter: "+missing)
		return
	}
	m := data.(map[string]interface{})
	playerID := m["username"].(string)

	// Step 2: Get Player Profile
	playerData, err := srvclient.PlayerClient.GetPlayerProfile(context.TODO(), &ppf.GetPlayerProfileRequest{
		PlayerID: playerID,
	})
	if err != nil {
		utils.ErrorResponse(c, 102, "Get Player Profile Error: "+err.Error())
		return
	}

	// Step 3: Get Balance
	walData, err := srvclient.WalletClient.GetBalance(context.TODO(), &wal.GetBalanceRequest{
		PlayerID:   playerID,
		OperatorID: playerData.PlayerProfile.OperatorID,
	})
	if err != nil {
		utils.ErrorResponse(c, 102, "Get Balance Error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{
		"success":    true,
		"error_code": 0,
		"data": gin.H{
			"username": playerID,
			"balance":  float64(walData.BalAmount) / 100,
			"currency": playerData.PlayerProfile.Currency.String(),
		},
	})

	return

}

func fundTransfer(c *gin.Context, data interface{}) {

	// Step 1: Check Parameter
	if missing := utils.CheckJsonDataContent(data, "username", "tran_id", "amount", "currency_code"); missing != "" {
		utils.ErrorResponse(c, 103, "Missing parameter: "+missing)
		return
	}

	// Initial Parameter
	m := data.(map[string]interface{})
	playerID := m["username"].(string)
	tranID := m["tran_id"].(string)
	intAmount := int64(0)
	if val, ok := m["amount"].(float64); ok {
		intAmount = int64(val * 100)
	} else {
		utils.ErrorResponse(c, 103, "Amount Format Error")
		return
	}

	// currency_code := m["currency_code"].(string)

	// Step 2: Get Player Profile
	playerData, err := srvclient.PlayerClient.GetPlayerProfile(context.TODO(), &ppf.GetPlayerProfileRequest{
		PlayerID: playerID,
	})
	if err != nil {
		utils.ErrorResponse(c, 102, "Get Player Profile Error: "+err.Error())
		return
	}

	// Step 2: Update Balance
	walData, err := srvclient.WalletClient.UpdateBalance(context.TODO(), &wal.UpdateBalanceRequest{
		PlayerID:   playerID,
		OperatorID: playerData.PlayerProfile.OperatorID,
		GainBal:    -intAmount,
	})

	if err != nil {
		utils.ErrorResponse(c, 101, "Update Balance Profile Error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{
		"success":    true,
		"error_code": 0,
		"data": gin.H{
			"username":      playerID,
			"tran_id":       tranID,
			"amount":        float64(walData.BalAmount) / 100,
			"currency_code": playerData.PlayerProfile.Currency.String(),
		},
	})
	return
}
