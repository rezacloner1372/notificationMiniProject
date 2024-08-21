// this is the notifier interface
package externalServices

// A String as Input and A string message as Output
type Notifier interface {
	SendNotify(receiver string, message string)
}
