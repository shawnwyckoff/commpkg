package gperf

import (
	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/shawnwyckoff/gopkg/container/gnum"
	"net/http"
)

func Serve(port uint16) error {
	http.HandleFunc("/api/stats", stats_api.Handler)
	return http.ListenAndServe(":"+gnum.FormatUint16(port), nil)
}
