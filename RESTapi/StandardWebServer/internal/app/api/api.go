package api

type API struct {
}

func New() *API {
	return &API{}
}

func (a *API) Start() error {
	return nil
}
