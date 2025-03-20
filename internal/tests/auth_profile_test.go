package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupRouterForAuth configure un routeur avec les endpoints Register et Login.
func setupRouterForAuth() *gin.Engine {
	r := gin.Default()
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	return r
}

// setupRouterForProfile configure un routeur pour les endpoints de profil.
// On simule ici l'authentification en injectant l'ID utilisateur dans le contexte.
func setupRouterForProfile(userID uint) *gin.Engine {
	r := gin.Default()
	r.GET("/profile", func(c *gin.Context) {
		c.Set("userID", userID)
		handlers.GetProfile(c)
	})
	r.PUT("/profile", func(c *gin.Context) {
		c.Set("userID", userID)
		handlers.UpdateProfile(c)
	})
	return r
}

func TestAuthAndProfileEndpoints(t *testing.T) {
	// Initialisation de la base de données pour les tests
	database.InitDB()
	// Migration du modèle User (on peut également migrer les autres modèles si nécessaire)
	database.DB.AutoMigrate(&models.User{})
	// Nettoyage de la table des utilisateurs pour obtenir un environnement propre
	database.DB.Exec("DELETE FROM users")
	defer database.CloseDB()

	// --- Test de l'inscription (RegisterUser) ---
	routerAuth := setupRouterForAuth()
	registerPayload := map[string]interface{}{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonValue, _ := json.Marshal(registerPayload)
	reqRegister, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	reqRegister.Header.Set("Content-Type", "application/json")
	wRegister := httptest.NewRecorder()
	routerAuth.ServeHTTP(wRegister, reqRegister)

	// Vérifier que l'inscription renvoie un code HTTP 201 Created
	assert.Equal(t, http.StatusCreated, wRegister.Code)

	var registerResponse map[string]interface{}
	err := json.Unmarshal(wRegister.Body.Bytes(), &registerResponse)
	assert.NoError(t, err)
	// On s'attend à ce que la réponse contienne un champ "user"
	userData := registerResponse["user"].(map[string]interface{})
	// Récupérer l'ID de l'utilisateur créé
	testUserIDFloat, ok := userData["ID"].(float64)
	assert.True(t, ok, "L'ID utilisateur doit être un nombre")
	testUserID := uint(testUserIDFloat)

	// Tester une inscription en doublon
	reqDup, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	reqDup.Header.Set("Content-Type", "application/json")
	wDup := httptest.NewRecorder()
	routerAuth.ServeHTTP(wDup, reqDup)
	assert.Equal(t, http.StatusConflict, wDup.Code)

	// --- Test de la connexion (LoginUser) ---
	loginPayload := map[string]interface{}{
		"email":    "test@example.com",
		"password": "password123",
	}
	loginJson, _ := json.Marshal(loginPayload)
	reqLogin, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJson))
	reqLogin.Header.Set("Content-Type", "application/json")
	wLogin := httptest.NewRecorder()
	routerAuth.ServeHTTP(wLogin, reqLogin)
	assert.Equal(t, http.StatusOK, wLogin.Code)
	var loginResponse map[string]interface{}
	err = json.Unmarshal(wLogin.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)
	token, tokenExists := loginResponse["token"].(string)
	assert.True(t, tokenExists, "Le token JWT doit être présent")
	assert.NotEmpty(t, token)

	// --- Test de la récupération du profil (GetProfile) ---
	routerProfile := setupRouterForProfile(testUserID)
	reqGetProfile, _ := http.NewRequest("GET", "/profile", nil)
	wGetProfile := httptest.NewRecorder()
	routerProfile.ServeHTTP(wGetProfile, reqGetProfile)
	assert.Equal(t, http.StatusOK, wGetProfile.Code)

	var profileResponse models.User
	err = json.Unmarshal(wGetProfile.Body.Bytes(), &profileResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Test User", profileResponse.Name)
	assert.Equal(t, "test@example.com", profileResponse.Email)

	// --- Test de la mise à jour du profil (UpdateProfile) ---
	updatePayload := map[string]interface{}{
		"name":  "Updated User",
		"email": "updated@example.com",
	}
	updateJson, _ := json.Marshal(updatePayload)
	reqUpdate, _ := http.NewRequest("PUT", "/profile", bytes.NewBuffer(updateJson))
	reqUpdate.Header.Set("Content-Type", "application/json")
	wUpdate := httptest.NewRecorder()
	routerProfile.ServeHTTP(wUpdate, reqUpdate)
	assert.Equal(t, http.StatusOK, wUpdate.Code)

	var updatedUser models.User
	err = json.Unmarshal(wUpdate.Body.Bytes(), &updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, "Updated User", updatedUser.Name)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}
