package main

import (
	"log"
	"time"
	"encoding/hex"
	"periph.io/x/periph/experimental/devices/mfrc522"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

var key mfrc522.Key = [6]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF}

type cardInfo struct {
	uId 		[]byte 
	uAuth		[]byte
	uData 		[][]byte
}

func main (){
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
	    log.Fatal(err)
	}

	// Using SPI as an example. See package "periph.io/x/periph/conn/spi/spireg" for more details.
	p, err := spireg.Open("")
	if err != nil {
	    log.Fatal(err)
	}
	defer p.Close()

	rfid, err := mfrc522.NewSPI(p, rpi.P1_13, rpi.P1_11)
	if err != nil {
	    log.Fatal(err)
	}

	// Idling device on exit.
	defer rfid.Halt()

	// Setting the antenna signal strength.
	rfid.SetAntennaGain(5)

	cb := make(chan []byte)

	// Stopping timer, flagging reader thread as timed out
	defer func() {
		close(cb)
	}()

	go func() {
		log.Printf("Started %s", rfid.String())

	    for {
	   	    // Trying to read card UID.

	   	    uid , err := rfid.ReadUID(5 * time.Second)
	        if err != nil {
	        	log.Printf("Error: %s", err.Error())
	            continue
	        }
	        //uauth, err := rfid.ReadAuth(10 * time.Second, 0x60, 8, key)
	        //if err != nil {
	        //	continue
	        //}

        // Some devices tend to send wrong data while RFID chip is already detected
        // but still "too far" from a receiver.
        // Especially some cheap CN clones which you can find on GearBest, AliExpress, etc.
        // This will suppress such errors.
        	cb <- uid
    	}
	}()

	for {
    	select {
    	case data := <-cb:
    	    log.Println("UID:", hex.EncodeToString(data))
    	    return

    	}
	}
}

/*func readBlock( sector, block int ) ([]byte, error){
	udata, err := rfid.ReadCard(10 * time.Second, 0x60, 8, 0, key)
	if err != nil {
		return nil, err
	} else {
		return udata, nil
	}
}*/