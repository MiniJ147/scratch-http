package server

type HeaderData struct {
	data map[string]string
}

// returns value and bool to see if it exist.
// up to user to handle error.
func (h *HeaderData) Find(k string) (string, bool) {
	val, ok := h.data[k]
	return val, ok
}

// inserts into map at key with value
func (h *HeaderData) Insert(k string, v string) {
	h.data[k] = v
}

// creates struct
func createHeaderData() *HeaderData {
	return &HeaderData{
		data: make(map[string]string),
	}
}
