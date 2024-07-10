package memyandex

type SenderAll interface {
	// Sending metrics to server
	SendAll(addr string) error
}

type Sender interface {
	// Sending metrics to server
	Send(addr, name string) error
}
