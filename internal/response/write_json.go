package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/stripe/stripe-go/v72"
)

type errResp struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func WriteJSON(w http.ResponseWriter, v interface{}, err error) {
	var respVal interface{}
	if err != nil {
		msg := err.Error()
		var stripeErr *stripe.Error
		if errors.As(err, &stripeErr) {
			msg = stripeErr.Msg
		}
		w.WriteHeader(http.StatusBadRequest)
		var e errResp
		e.Error.Message = msg
		respVal = e
	} else {
		respVal = v
	}

	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(respVal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Error("json encode failed", slog.Any("error", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = io.Copy(w, &buf); err != nil {
		slog.Error("copy body response failed", slog.Any("error", err))
		return
	}
}
