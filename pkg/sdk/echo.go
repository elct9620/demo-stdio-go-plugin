package sdk

type EchoRequest struct {
	Msg string
}

type EchoResponse struct {
	Msg string
}

type EchoService interface {
	Ping(req EchoRequest, reply *EchoResponse) error
}

const OpPing = Operation("Echo.Ping")
