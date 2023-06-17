package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dapetoo/go-bookings/internal/config"
	"github.com/dapetoo/go-bookings/internal/driver"
	"github.com/dapetoo/go-bookings/internal/forms"
	"github.com/dapetoo/go-bookings/internal/helpers"
	"github.com/dapetoo/go-bookings/internal/models"
	"github.com/dapetoo/go-bookings/internal/render"
	"github.com/dapetoo/go-bookings/internal/repository"
	"github.com/dapetoo/go-bookings/internal/repository/dbrepo"
	"log"
	"net/http"
	"strconv"
	"time"
)

//// Repo the repository used by the handlers
//var Repo *Repository
//
//// Repository is the repository type
//type Repository struct {
//	App *config.AppConfig
//	DB  repository.DatabaseRepo
//}
//
//// NewRepo creates a new repository
//func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
//	return &Repository{
//		App: a,
//		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
//	}
//}
//
//// NewHandlers sets the repository for the handlers
//func NewHandlers(r *Repository) {
//	Repo = r
//}
//
//// Home is the handler for the home page
//func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
//	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
//}
//
//func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
//	log.Println("Contact page")
//	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
//}
//
//func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
//	log.Println("Contact page")
//	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
//}
//

//
//// Reservation renders the make a reservation page and displays form
//func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
//	var emptyReservation models.Reservation
//	data := make(map[string]interface{})
//	data["reservation"] = emptyReservation
//
//	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
//		Form: forms.New(nil),
//		Data: data,
//	})
//}
//
//// PostReservation handles the posting of a reservation form
//func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
//	err := r.ParseForm()
//	if err != nil {
//		helpers.ServerError(w, err)
//		return
//	}
//
//	sd := r.Form.Get("start_date")
//	ed := r.Form.Get("end_date")
//
//	// layout := "2006-01-02" // the date format must be in this format
//	// 2020-01-01 01/02 03:04:05PM '06 -0700
//
//	layout := "2006-01-02"
//
//	startDate, err := time.Parse(layout, sd)
//	if err != nil {
//		helpers.ServerError(w, err)
//		return
//	}
//
//	endDate, err := time.Parse(layout, ed)
//	if err != nil {
//		helpers.ServerError(w, err)
//		return
//	}
//
//	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
//	if err != nil {
//		helpers.ServerError(w, err)
//		return
//	}
//
//	// grabbed data from whatever they entered in the field of the form
//	reservation := models.Reservation{
//		FirstName: r.Form.Get("first_name"),
//		LastName:  r.Form.Get("last_name"),
//		Email:     r.Form.Get("email"),
//		Phone:     r.Form.Get("phone"),
//		StartDate: startDate,
//		EndDate:   endDate,
//		RoomID:    roomID,
//	}
//
//	// Validate the form
//	form := forms.New(r.PostForm)
//
//	form.Required("first_name", "last_name", "email")
//	form.MinLength("first_name", 3)
//	form.IsEmail("email")
//
//	if !form.Valid() {
//		data := make(map[string]interface{})
//		data["reservation"] = reservation
//		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
//			Form: form,
//			Data: data,
//		})
//		return
//	}
//
//	//Add reservation to the database
//	newReservationID, err := m.DB.InsertReservation(reservation)
//	if err != nil {
//		helpers.ServerError(w, err)
//		return
//	}
//
//	restriction := models.RoomRestriction{
//		StartDate:     startDate,
//		EndDate:       endDate,
//		RoomID:        roomID,
//		ReservationID: newReservationID,
//		RestrictionID: 1,
//	}
//
//	err = m.DB.InsertRoomRestriction(restriction)
//
//	m.App.Session.Put(r.Context(), "reservation", reservation)
//	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
//}
//
//// Generals renders the room page
//func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
//	render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
//}
//
//// Majors renders the room page
//func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
//	render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
//}
//
//// Availability renders the search availability page
//func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
//	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
//}
//
//// PostAvailability handles post
//func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
//	start := r.Form.Get("start")
//	end := r.Form.Get("end")
//
//	w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
//}
//
//type jsonResponse struct {
//	OK      bool   `json:"ok"`
//	Message string `json:"message"`
//}
//
//// AvailabilityJSON handles request for availability and sends JSON response
//func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
//	resp := jsonResponse{
//		OK:      true,
//		Message: "Available!",
//	}
//
//	out, err := json.MarshalIndent(resp, "", "     ")
//	if err != nil {
//		helpers.ServerError(w, err)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(out)
//}
//

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// call renderTemplate
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})

}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	// 01/02 03:04:05PM '06 -0700

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
	}

	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// grabbed data from whatever they entered in the field of the form
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
	}

	// check my form
	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		// render the form
		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        roomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Service(w http.ResponseWriter, r *http.Request) {
	log.Println("Service page")
	render.Template(w, r, "service.page.tmpl", &models.TemplateData{})
}

// Majors  renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability  renders the availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability  renders the availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
	}

	endDate, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(w, err)
	}

	rooms, err := m.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(rooms) == 0 {
		// no availablity
		m.App.Session.Put(r.Context(), "error", "No availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send JSON responses
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// ReservationSummary
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Can't get error from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
