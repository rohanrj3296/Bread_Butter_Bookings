package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rohanrj3296/Bread_Butter_Bookings/pkg/config"
	"github.com/rohanrj3296/Bread_Butter_Bookings/pkg/models"
	"github.com/rohanrj3296/Bread_Butter_Bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w,r, "home.page.tmpl", &models.TemplateData{})
}
//General is the handler of generals room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	

	render.RenderTemplate(w,r, "generals.page.tmpl", &models.TemplateData{})
}
//Majors is the handler of majors room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w,r, "majors.page.tmpl", &models.TemplateData{})
}
//Availability is the handler of search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w,r, "search-availability.page.tmpl", &models.TemplateData{})
}
//Contact is the handler of contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r,"contact.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"make-reservation.page.tmpl", &models.TemplateData{})
}

//for posting data from search availability error
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start:=r.Form.Get("start") //here start means we are taking the input with name=start in form
	end:=r.Form.Get("end")//this will be a string by default

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s",start,end)))
}

type jsonResponse struct {
	Ok bool `json:"ok"`
	Message string `json:"message"`
}
//creating a handler for Availability check and it sends json response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok: true,
		Message: "Available!!",
	}
	out,err:=json.MarshalIndent(resp,"","         ")
	if err!=nil{
		fmt.Println("error in marshalling json file",err)
	}
	log.Println(string(out))
	//setting the header
	w.Header().Set("Content-Type","application/json")
	w.Write(out)
}


// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w,r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
