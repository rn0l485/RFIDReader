package main

import (
	"log"
	"time"
	"encoding/hex"

	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/mfrc522"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

func main() {
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

	rfid, err := mfrc522.NewSPI(p, rpi.P1_22, rpi.P1_18)
	if err != nil {
		log.Fatal(err)
	}

	// Idling device on exit.
	defer rfid.Halt()

	// Setting the antenna signal strength.
	rfid.SetAntennaGain(5)

	cb := make(chan []byte)
	defer close(cb)

	go func() {
		log.Printf("Started %s", rfid.String())

		for {
			// Trying to read data from sector 1 block 0
			data, err := rfid.ReadUID(5*time.Second)
			if err != nil {
				log.Printf(err.Error())
				continue
			}

			cb <- data
		}
	}()

	for {
		select {
		case data := <-cb:
			userID := hex.EncodeToString(data)
			log.Fatal(userID)
			return
		}
	}
}