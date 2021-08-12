package sync

// Harbor is user image register
type Harbor struct {
	url      string
	userName string
	password string
}

// Image ...
type Image struct {
	Name    string
	Project string
	Repo    Harbor
}

// NewHarbor return object of Harbor
func NewHarbor(url, userName, password string) Harbor {
	return Harbor{
		url,
		userName,
		password,
	}
}

// syncImage pull
func (h *Harbor) syncImage(source Image) {

}
