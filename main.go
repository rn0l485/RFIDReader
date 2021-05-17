package main

import (
	"log"
	"net/http"
	"time"


	v1	"go-RC522/Router"
	"go-RC522/Reader"
	"go-RC522/Router/Models/Persion"
	"go-RC522/Router/Models/Assets"
	"go-RC522/Router/Models/Products"
	
	"golang.org/x/sync/errgroup"

)

var (
	g errgroup.Group
)

func main() {

	// route 1
	v1.InitRouter()
	v1.R = persion.InitRouter(v1.R)
	v1.R = assets.InitRouter(v1.R)
	v1.R = products.InitRouter(v1.R)


	defer reader.P.Close()
	defer reader.Rfid.Halt()



	srv1 := &http.Server{
		Addr:		":9000",
		Handler:	v1.R,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := srv1.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}


}
