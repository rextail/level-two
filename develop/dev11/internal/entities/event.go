package entity

type Event struct {
	Title       string
	Date        string
	Description string
}

type CreateEventRequest struct {
	Title       string
	Date        string
	Description string
}

type UpdateEventRequest struct {
	Title       string
	Date        string
	Description string
}

type DeleteRequest struct {
	Title string
}

type EventsByTimeRequest struct {
	Time string
}
