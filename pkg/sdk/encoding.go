package sdk

type Item struct {
	Name  string
	Price int
}

type EncodeRequest struct {
	Items []Item
}

type EncodeResponse struct {
	Result []byte
}

type Encoder interface {
	Encode(req EncodeRequest, reply *EncodeResponse) error
}

const OpEncode = Operation("Encoder.Encode")
