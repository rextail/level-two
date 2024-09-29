package entities

type Event struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type CreateEventRequest struct {
	Title       string `json:"title" validate:"required"`
	Date        string `json:"date" validate:"required"`
	Description string `json:"description"`
}

type UpdateEventRequest struct {
	Title       string `json:"title" validate:"required"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type DeleteRequest struct {
	Title string `json:"title" validate:"required"`
}

type EventsByTimeRequest struct {
	TimeFilter string `json:"time_filter"`
}
