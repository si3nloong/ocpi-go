package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func ResponsePagination[T any](w http.ResponseWriter, r *http.Request, response *ocpi.PaginationResponse[T]) {
	b, err := json.Marshal(ocpi.NewResponse(response.Data))
	if err != nil {
		// ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Link", GetHostname(r)+"; rel=\"next\"")
	w.Header().Set("X-Total-Count", strconv.FormatInt(response.Header.TotalCount, 10))
	w.Header().Set("X-Limit", strconv.FormatInt(response.Header.Limit, 10))
	w.Write(b)
}
