package enum

type NotificationSource string

var (
	Candidate NotificationSource = "candidate"
	Customer  NotificationSource = "customer"
	Interview NotificationSource = "interview"
)
