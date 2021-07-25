package notification

type Notifier interface {
	SendTextMessage(message string, recipients []string) (err error)
}
