package reader

import (
	"log"
	"time"
	"encoding/hex"
	"errors"

	"periph.io/x/periph/experimental/devices/mfrc522"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)


const authMode 			byte 				= 0x60

//var key 		mfrc522.Key 		= [6]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF}
var Rfid 		*mfrc522.Dev
var P 			spi.PortCloser
var err 		error


func InitReader(){
	if _, err := host.Init(); err != nil {
	    log.Fatal(err)
	}

	// Using SPI as an example. See package "periph.io/x/periph/conn/spi/spireg" for more details.
	P, err = spireg.Open("")
	if err != nil {
	    log.Fatal(err)
	}

	Rfid, err = mfrc522.NewSPI(P, rpi.P1_22, rpi.P1_18)
	if err != nil {
	    log.Fatal(err)
	}

	// Setting the antenna signal strength.
	Rfid.SetAntennaGain(5)	
}

func GetId( timeOut time.Duration)( uid *string, err error){
	callBack := make(chan []byte)
	defer close(callBack)
	go func(){
		for {
			uidHex, err := Rfid.ReadUID(timeOut)
			if err != nil {
				continue
			}
			callBack <- uidHex
			break
		}
	}()

	select {
	case feedBack := <- callBack:
		rv := hex.EncodeToString(feedBack)
		return &rv, nil
	case <- time.After(timeOut):
		return nil, errors.New("Reading time out.")
	}
}
