package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/config"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig
func TestMain(m *testing.M){
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false
	infoLog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)//log.New: This function creates a new logger. It takes three arguments,os.Stdout:
	// This is the destination where the log messages will be written."INFO\t": 
	//This is the prefix for each log message. The prefix "INFO\t" will be added at the beginning of every log entry.
	// The \t represents a tab character, which helps format the output nicely.log.Ldate | log.Ltime: These are flags that control what extra information gets included in each log entry.
	//log.Ldate: Adds the date to the log entry (e.g., 2024/08/21).
	//log.Ltime: Adds the time to the log entry (e.g., 15:04:05).
	//The | operator combines these flags so both the date and time are included in each log message.
	app.InfoLog = infoLog

	//creating an error logger and assigning it
	errorLog := log.New(os.Stdout,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile) //Lshortfile gives information about the error
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	testApp.Session = session
	app=&testApp
	os.Exit(m.Run())
}
type myWriter struct{}

func (tw *myWriter) Header() http.Header{
	var h http.Header
	return h
}
func (tw *myWriter) WriteHeader(i int){

} 
func (tw *myWriter) Write(b []byte) (int,error){
	length:=len(b)
	return length,nil

}