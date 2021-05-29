package main

import (
	"Website_1/database"
	"Website_1/router"
	"Website_1/service"
	"encoding/gob"
	"fmt"

	//"bufio"
	"Website_1/util"
	"text/template"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Connecting to database, starting the application ... ")
	util.Tpl = template.Must(template.ParseGlob("Templates/*"))
	util.DB = database.StartDatabase()
	defer util.DB.Close()
	err := util.DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database, now connecting rabbit mq")
	util.MessageBroker = database.StartMessageBroker()
	defer util.MessageBroker.Close()

	util.Store = sessions.NewCookieStore(
		[]byte("website-project"),
	)

	util.Store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(service.Users{})

	go func() {
		//	service.ConsumerService() //to run the function without letting main thread goes off
	}()

	router.HttpEndpoint()

}
