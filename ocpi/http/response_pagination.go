package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func ResponsePagination[T any](w http.ResponseWriter, r *http.Request, response *ocpi.PaginatedResponse[T]) {
	b, err := json.Marshal(ocpi.NewResponse(response.Data))
	if err != nil {
		b, _ = json.Marshal(ocpi.Response[any]{
			StatusCode:    ocpi.StatusCodeServerError,
			StatusMessage: err.Error(),
			Timestamp:     time.Now().UTC(),
		})
		w.Write(b)
		return
	}

	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Link", GetHostname(r)+"; rel=\"next\"")
	// w.Header().Set("X-Total-Count", strconv.FormatInt(response.Header.TotalCount, 10))
	// w.Header().Set("X-Limit", strconv.FormatInt(response.Header.Limit, 10))
	w.Write(b)
}
