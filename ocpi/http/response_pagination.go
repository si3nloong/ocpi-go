package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func ResponsePagination[T ocpi.Timestamp, D any](w http.ResponseWriter, r *http.Request, ts T, response *ocpi.PaginatedResponse[T, D]) {
	b, err := json.Marshal(ocpi.NewResponse(ts, response.Data))
	if err != nil {
		b, _ = json.Marshal(ocpi.Response[T, any]{
			StatusCode:    ocpi.StatusCodeServerError,
			StatusMessage: err.Error(),
			Timestamp:     ts,
		})
		w.Write(b)
		return
	}

	if link := response.Link(); link != "" {
		w.Header().Set("Link", fmt.Sprintf(GetHostname(r)+"; rel=%q", link))
	}
	if totalCount, err := response.TotalCount(); err == nil {
		w.Header().Set("X-Total-Count", strconv.FormatUint(totalCount, 10))
	}
	if limit, _ := response.Limit(); limit > 0 {
		w.Header().Set("X-Limit", strconv.Itoa(limit))
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
