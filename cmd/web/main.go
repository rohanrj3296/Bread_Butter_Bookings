package main

import (
	"database/sql/driver"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/config"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/handlers"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/helpers"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/models"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {
	
	err:=run()
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error{
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	//creating a  info logger:
	infoLog = log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)//log.New: This function creates a new logger. It takes three arguments,os.Stdout:
	// This is the destination where the log messages will be written."INFO\t": 
	//This is the prefix for each log message. The prefix "INFO\t" will be added at the beginning of every log entry.
	// The \t represents a tab character, which helps format the output nicely.log.Ldate | log.Ltime: These are flags that control what extra information gets included in each log entry.
	//log.Ldate: Adds the date to the log entry (e.g., 2024/08/21).
	//log.Ltime: Adds the time to the log entry (e.g., 15:04:05).
	//The | operator combines these flags so both the date and time are included in each log message.
	app.InfoLog = infoLog

	//creating an error logger and assigning it
	errorLog = log.New(os.Stdout,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile) //Lshortfile gives information about the error
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//cnnect to database
	log.Println("Connecting To DataBase")
	

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	helpers.NewHelpers(&app)//this line populates the app in helpers with a pointer to app config
	return nil
}