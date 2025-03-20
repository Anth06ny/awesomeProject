package tests

import (
	"awesomeProject/internal/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIEndpoints(t *testing.T) {
	// Initialise le routeur avec nos routes
	router := routes.SetupRouter()

	// Définir une fonction helper pour tester un endpoint
	testEndpoint := func(method, path string, expectedStatus int) {
		req, err := http.NewRequest(method, path, nil)
		if err != nil {
			t.Fatalf("Erreur lors de la création de la requête: %v", err)
		}
		// Utilisation d'un ResponseRecorder pour capturer la réponse
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, expectedStatus, w.Code, "Erreur sur %s %s", method, path)
	}

	// Tester les endpoints d'authentification et de gestion d'utilisateur
	testEndpoint("POST", "/api/register", http.StatusNotImplemented)
	testEndpoint("POST", "/api/login", http.StatusNotImplemented)
	testEndpoint("GET", "/api/profile", http.StatusNotImplemented)
	testEndpoint("PUT", "/api/profile", http.StatusNotImplemented)

	// Tester les endpoints de ressources
	testEndpoint("GET", "/api/resources", http.StatusNotImplemented)
	testEndpoint("GET", "/api/resources/1", http.StatusNotImplemented)

	// Tester les endpoints de prêts
	testEndpoint("POST", "/api/loans", http.StatusNotImplemented)
	testEndpoint("GET", "/api/loans", http.StatusNotImplemented)
	testEndpoint("PUT", "/api/loans/1/return", http.StatusNotImplemented)
	// Si vous ajoutez DELETE plus tard
	// testEndpoint("DELETE", "/api/loans/1", http.StatusNotImplemented)
}
