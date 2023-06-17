package handlers

import (
	"encoding/json"
	"errors"
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
	"strings"
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
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		helpers.ServerError(w, errors.New("cannot get reservation from session"))
		return
	}

	room, err := m.DB.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res.Room.RoomName = room.RoomName

	layout := "2006-01-02"
	sd := res.StartDate.Format(layout)
	ed := res.EndDate.Format(layout)

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
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
	err := r.ParseForm()
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse start date!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse end date!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	rooms, err := m.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't get availability for rooms")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if len(rooms) == 0 {
		// no availability
		m.App.Session.Put(r.Context(), "error", "No availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	m.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

type jsonResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	RoomID    string `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	// need to parse request body
	err := r.ParseForm()
	if err != nil {
		// can't parse form, so return appropriate json
		resp := jsonResponse{
			OK:      false,
			Message: "Internal server error",
		}

		out, _ := json.MarshalIndent(resp, "", "     ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	sd := r.Form.Get("start")
	ed := r.Form.Get("end")

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	roomID, _ := strconv.Atoi(r.Form.Get("room_id"))

	available, err := m.DB.SearchAvailabilityByDatesByRoomID(startDate, endDate, roomID)
	if err != nil {
		// got a database error, so return appropriate json
		resp := jsonResponse{
			OK:      false,
			Message: "Error querying database",
		}

		out, _ := json.MarshalIndent(resp, "", "     ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}
	resp := jsonResponse{
		OK:        available,
		Message:   "",
		StartDate: sd,
		EndDate:   ed,
		RoomID:    strconv.Itoa(roomID),
	}

	// I removed the error check, since we handle all aspects of
	// the json right here
	out, _ := json.MarshalIndent(resp, "", "     ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// ReservationSummary displays the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}

// ChooseRoom displays list of available rooms
func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	// split the URL up by /, and grab the 3rd element
	exploded := strings.Split(r.RequestURI, "/")
	roomID, err := strconv.Atoi(exploded[2])
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "missing url parameter")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.RoomID = roomID

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)

}

// BookRoom takes URL parameters, builds a sessional variable, and takes user to make res screen
func (m *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	roomID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	var res models.Reservation

	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Can't get room from db!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.Room.RoomName = room.RoomName
	res.RoomID = roomID
	res.StartDate = startDate
	res.EndDate = endDate

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
