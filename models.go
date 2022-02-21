package acloudguru

import "time"

type User struct {
	UserId            string    `json:"userId"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Admin             bool      `json:"admin"`
	LastSeenTimestamp time.Time `json:"lastSeenTimestamp"`
	Status            string    `json:"status"`
}

type Subscription struct {
	OrganisationId string    `json:"organisationId"`
	Name           string    `json:"name"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	TotalSeats     int       `json:"totalSeats"`
	SeatsInUse     int       `json:"seatsInUse"`
}
