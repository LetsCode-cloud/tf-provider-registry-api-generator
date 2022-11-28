package types

type Discovery struct {
	ProvidersV1 string `json:"providers.v1"`
}

func (c Discovery) Construct() *Discovery {
	c.ProvidersV1 = "/v1/providers/"
	return &c
}
