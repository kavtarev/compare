package sender

import (
	"http_nats_rabbit_grpc/types"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) ShowTotalTimeHandler(w http.ResponseWriter, r *http.Request) {
	totalTimeStr := strconv.FormatInt(s.totalTime.Microseconds(), 10)
	w.Write([]byte(totalTimeStr))
}

func (s *Server) ShowFullCircleTimeHandler(w http.ResponseWriter, r *http.Request) {
	totalTimeStr := strconv.FormatInt(s.endTime.Sub(s.startTime).Microseconds(), 10)
	w.Write([]byte(totalTimeStr))
}

func (s *Server) ResetTimerHandler(w http.ResponseWriter, r *http.Request) {
	s.totalTime = 0
	s.startTime = time.Now()
	s.endTime = time.Now()
	s.ReceivedObjects = 0
	w.Write([]byte("reset to 0"))
}

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
