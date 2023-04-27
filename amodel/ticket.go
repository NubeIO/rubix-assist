package amodel

import "time"

type TicketImportanceLevel string

const (
	TicketImportanceLevelLow      TicketImportanceLevel = "LOW"
	TicketImportanceLevelMedium   TicketImportanceLevel = "MEDIUM"
	TicketImportanceLevelHigh     TicketImportanceLevel = "HIGH"
	TicketImportanceLevelCritical TicketImportanceLevel = "CRITICAL"
)

// TicketStatus model.
type TicketStatus string

// Different ticket status instances.
const (
	TicketStatusNew      TicketStatus = "NEW"
	TicketStatusReplied  TicketStatus = "REPLIED"
	TicketStatusResolved TicketStatus = "RESOLVED"
	TicketStatusClosed   TicketStatus = "CLOSED"
	TicketStatusBlocked  TicketStatus = "BLOCKED"
)

type Comment struct {
	UUID       string    `json:"uuid" gorm:"primary_key"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	TicketID   string    `json:"ticket_id"`
	Owner      string    `json:"owner"`
	Content    string    `json:"content"`
	Metadata   string    `json:"metadata"`
}

type Ticket struct {
	UUID            string                `json:"uuid" gorm:"primary_key"`
	CreatedAt       time.Time             `json:"created_at"`
	ModifiedAt      time.Time             `json:"modified_at"`
	Issuer          string                `json:"issuer"`
	Owner           string                `json:"owner"`
	Subject         string                `json:"subject"`
	Content         string                `json:"content"`
	Metadata        string                `json:"metadata"`
	ImportanceLevel TicketImportanceLevel `json:"importance_level"`
	Status          TicketStatus          `json:"status"`
	Comments        []*Comment            `json:"comments"`
}
