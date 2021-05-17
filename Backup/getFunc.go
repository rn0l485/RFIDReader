package router

import (
	//"fmt"
	"time"
	//"log"
	"net/http"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-RC522/Worker"
	"go-RC522/Reader"
)


const UserDB_url 			string 				= "http://18.180.56.128:9000/user"
const UserDB_url_Search 	string 				= "/search"
const authMode 				byte 				= 0x60


//var key 		mfrc522.Key 		= [6]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF}
var w 			*worker.Worker 		= worker.InitWorker()
//var Rfid 		*mfrc522.Dev
//var P 			spi.PortCloser
var err 		error

func init(){
	reader.InitReader()
}

func getID( c *gin.Context) {
	uid, err:= reader.GetId(10*time.Second)
	if err != nil {
		c.JSON( http.StatusInternalServerError, gin.H{
			"errorFunc":"getID_p1",
			"info": err.Error(),
		})
		return
	}

	payload := &payloadStruct{
		Payload: map[string]string{
			"Account": "abd",
			"uid": *uid,
		},
		Ticket : "Rfid_reader_getID",
	}

	targetURL := UserDB_url+UserDB_url_Search
	feedBack, err := w.Post( targetURL, payload)
	if err != nil {
		c.JSON( http.StatusInternalServerError, gin.H{
			"errorFunc":"getID_p2",
			"info": err.Error(),
			"uid" : *uid,
		})
		return
	}

	var feedBackBody map[string]interface{}
	json.Unmarshal( feedBack.Body, &feedBackBody)
	c.JSON( http.StatusOK, feedBackBody)
	//c.JSON( http.StatusOK, uid)		
}

type payloadStruct struct {
	Payload  		map[string]string 		`json:"Payload"`
	Ticket 			string 					`json:"Ticket"`
}