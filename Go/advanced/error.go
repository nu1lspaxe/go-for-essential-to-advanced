package advanced

import "sync"

// Implement Error

type InternalError struct {
	msg string
}

func (ie InternalError) Error() string {
	return ie.msg
}

func New(msg string) error {
	return InternalError{msg: msg}
}

// sync.Mutex to manager error

var (
	service map[string]string
	servMux sync.Mutex
)

func RegisterServ(name, addr string) {
	servMux.Lock()
	defer servMux.Unlock()
	service[name] = addr
}

func LookupServ(name string) string {
	servMux.Lock()
	defer servMux.Unlock()
	return service[name]
}
