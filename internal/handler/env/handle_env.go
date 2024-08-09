package env

import (
	"net/http"
	"os"

	"backendgo/internal/response"
)

func HandleEnv(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, map[string]any{
		"authSecret":           os.Getenv("AUTH_SECRET"),
		"publicDomain":         os.Getenv("FE_PUBLIC_DOMAIN"),
		"apiUrl":               os.Getenv("API_URL"),
		"authEnabled":          os.Getenv("AUTH_ENABLED"),
		"billingType":          os.Getenv("BILLING_TYPE"),
		"stripePublishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
	}, nil)
}
