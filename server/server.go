package server

import (
	"coord-converter/convert"
	"coord-converter/types"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func handleConvert(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	var convertedPayload types.ConversionBody
	if err := json.NewDecoder(r.Body).Decode(&convertedPayload); err != nil {
		w.WriteHeader(500)
		body, _ := io.ReadAll(r.Body)
		logger.Error("invalid body data", r.Header, string(body[:]))
		_, _ = w.Write([]byte("could not parse body data"))
		return
	}
	results := convert.Convert(convertedPayload.From, convertedPayload.To, convertedPayload.Coordinates, logger)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(types.ConversionResult{Coordinates: results})
}

func NewServer(logger *slog.Logger) *http.ServeMux {
	server := http.NewServeMux()
	server.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("[POST] /convert")
		handleConvert(w, r, logger)
	})
	return server
}
