package helpers 

type Protocol int

const (
	HTTP Protocol = iota
	WS
)

func (p Protocol) String() string {
	switch p {
	case HTTP:
		return "http"
	case WS:
		return "ws"
	}
	return "unknown"
}