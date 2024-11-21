package dto

type Cep struct {
	Cep          string
	State        string
	City         string
	Neighborhood string
	Street       string
	Resource     string
}

func NewCEP(cep string, state string, city string, neighborhood string, street string, resource string) *Cep {
	return &Cep{
		Cep:          cep,
		State:        state,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
		Resource:     resource,
	}
}
