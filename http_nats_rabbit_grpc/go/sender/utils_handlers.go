package sender

import (
	"http_nats_rabbit_grpc/types"
)

func (s *Server) GetStructByInput() any {
	switch s.opts.TypeOfObjects {
	case "s-number":
		return types.SmallNumber{}
	case "s-string":
		return types.SmallString{}
	case "s-mixed":
		return types.SmallMixed{}
	case "m-number":
		return types.MediumNumber{}
	case "m-string":
		return types.MediumString{}
	case "m-mixed":
		return types.MediumMixed{}
	case "l-number":
		return types.LargeNumber{}
	case "l-string":
		return types.LargeString{}
	case "l-mixed":
		return types.LargeMixed{}
	default:
		panic("invalid type")
	}
}
