package tests

import (
	"awesomeProject/internal/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/stretchr/testify/assert"
)

// TestMain permet d'initialiser et de nettoyer la base de données pour les tests.
func TestMain(m *testing.M) {
	// Initialisation de la base de données (utilisez ici une configuration adaptée aux tests)
	database.InitDB()
	// Migration de la table Resource (les autres tables ne sont pas utilisées ici)
	database.DB.AutoMigrate(&models.Resource{})

	// Exécuter les tests
	code := m.Run()

	// Fermer la base de données
	database.CloseDB()
	os.Exit(code)
}

// TestResourceAPI teste l'ensemble des endpoints pour les ressources : création, récupération de la liste et récupération par ID.
func TestResourceAPI(t *testing.T) {
	// Récupérer le routeur configuré
	router := routes.SetupRouter()

	// --- Test de création d'une ressource (POST /api/resources) ---
	newResource := models.Resource{
		Title:  "Test Book",
		Type:   "Livre",
		Status: "disponible",
	}
	jsonValue, err := json.Marshal(newResource)
	assert.NoError(t, err)

	reqCreate, err := http.NewRequest("POST", "/api/resources", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err)
	reqCreate.Header.Set("Content-Type", "application/json")

	wCreate := httptest.NewRecorder()
	router.ServeHTTP(wCreate, reqCreate)

	// Vérifier que le code HTTP est 201 Created
	assert.Equal(t, http.StatusCreated, wCreate.Code)

	var createdResource models.Resource
	err = json.Unmarshal(wCreate.Body.Bytes(), &createdResource)
	assert.NoError(t, err)
	// Vérifier que les champs retournés correspondent à ceux envoyés et qu'un ID a été attribué
	assert.Equal(t, newResource.Title, createdResource.Title)
	assert.Equal(t, newResource.Type, createdResource.Type)
	assert.Equal(t, newResource.Status, createdResource.Status)
	assert.NotZero(t, createdResource.ID)

	// --- Test de récupération de toutes les ressources (GET /api/resources) ---
	reqList, err := http.NewRequest("GET", "/api/resources", nil)
	assert.NoError(t, err)

	wList := httptest.NewRecorder()
	router.ServeHTTP(wList, reqList)

	// Vérifier le code HTTP 200 OK
	assert.Equal(t, http.StatusOK, wList.Code)

	var resources []models.Resource
	err = json.Unmarshal(wList.Body.Bytes(), &resources)
	assert.NoError(t, err)
	// La liste doit contenir au moins la ressource que nous venons de créer
	assert.True(t, len(resources) >= 1)

	// --- Test de récupération d'une ressource par son ID (GET /api/resources/:id) ---
	resourceID := strconv.Itoa(int(createdResource.ID))
	reqGet, err := http.NewRequest("GET", "/api/resources/"+resourceID, nil)
	assert.NoError(t, err)

	wGet := httptest.NewRecorder()
	router.ServeHTTP(wGet, reqGet)

	// Vérifier le code HTTP 200 OK
	assert.Equal(t, http.StatusOK, wGet.Code)

	var fetchedResource models.Resource
	err = json.Unmarshal(wGet.Body.Bytes(), &fetchedResource)
	assert.NoError(t, err)
	// Vérifier que la ressource récupérée correspond à celle créée
	assert.Equal(t, createdResource.ID, fetchedResource.ID)
	assert.Equal(t, createdResource.Title, fetchedResource.Title)
	assert.Equal(t, createdResource.Type, fetchedResource.Type)
	assert.Equal(t, createdResource.Status, fetchedResource.Status)
}
