package app

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	cms "github.com/PGITAb/bc-proto-entity-cms-go"
	"github.com/PGITAb/bc-service-cms-api-go/internal/srvclient"
	"github.com/gin-gonic/gin"
)

func CheckHasPrefix(path string, vals ...string) bool {
	for _, v := range vals {
		if strings.HasPrefix(path, v) {
			return true
		}
	}

	return false
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		header := ""
		body := ""

		// Do not print log
		if CheckHasPrefix(path, "/doc", "/static", "/favicon.ico") {
			c.Next()
			return
		}

		// body log
		if CheckHasPrefix(path, "/banner", "/dealer", "/avatar") {
			// Only form data without file
			if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
				log.Println("ParseMultipartForm err:", err)
			}

			for k, v := range c.Request.PostForm {
				body += k + ": " + fmt.Sprint(v) + "\r\n"
			}
		} else {
			// all info
			byteBody, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))
			body = string(byteBody)
		}

		c.Next()

		// header log
		for k, v := range c.Request.Header {
			header += k + ": " + fmt.Sprint(v) + "\r\n"
		}

		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}

		errorMsg := c.Keys["ErrorMsg"]
		staff := c.Keys["Staff"]

		msg := "------------------------------------------------------------\r\n"
		msg += "[%s] | %s\r\n\r\n"
		msg += "[Request] \r\n%s %s\r\n\r\n"
		msg += "[Header] \r\n%s\r\n"
		msg += "[Body] \r\n%s\r\n"
		msg += "[Staff] \r\n%v\r\n\r\n"
		msg += "[Status] \r\n%v\r\n\r\n"
		msg += "[ErrorMsg] \r\n%v\r\n\r\n"
		msg += "------------------------------------------------------------\r\n"
		fmt.Printf(msg,
			time.Now().Format("2006/01/02 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			path,
			header,
			body,
			staff,
			c.Writer.Status(),
			errorMsg)

		go func() {
			_, err := srvclient.CmsClient.AddLog(context.TODO(), &cms.AddLogRequest{
				StaffID:  fmt.Sprintf("%v", staff),
				Endpoint: fmt.Sprintf("%v %v", c.Request.Method, path),
				Details:  c.ClientIP(),
			})
			if err != nil {
				log.Println("AddLog Error:", err)
			}
		}()
	}
}
