package handlers

import (
	"errors"
	"net/http"

	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetResources récupère la liste des ressources (livres et jeux).
func GetResources(c *gin.Context) {
	var resources []models.Resource
	// Récupère toutes les ressources de la base de données.
	if err := database.DB.Find(&resources).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les ressources"})
		return
	}
	c.JSON(http.StatusOK, resources)
}

// GetResource récupère les détails d'une ressource spécifique à partir de son ID.
func GetResource(c *gin.Context) {
	id := c.Param("id")
	var resource models.Resource
	// Recherche la ressource par son ID.
	if err := database.DB.First(&resource, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ressource non trouvée"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la ressource"})
		}
		return
	}
	c.JSON(http.StatusOK, resource)
}

// CreateResource permet d'ajouter une nouvelle ressource (livre ou jeu).
func CreateResource(c *gin.Context) {
	var resource models.Resource

	// On tente de lier le JSON de la requête à notre structure Resource.
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mauvaise requête, vérifiez le format des données"})
		return
	}

	// On insère la ressource dans la base de données.
	if err := database.DB.Create(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de la ressource"})
		return
	}

	// Retourner la ressource créée avec le code HTTP 201 Created.
	c.JSON(http.StatusCreated, resource)
}

// DisableResource met à jour une ressource en "indisponible".
func DisableResource(c *gin.Context) {
	var resource models.Resource
	id := c.Param("id")

	// Récupérer la ressource dans la base de données
	if err := database.DB.First(&resource, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ressource introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de la ressource"})
		return
	}

	// Mettre à jour le statut
	resource.Status = "indisponible"

	if err := database.DB.Save(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de la ressource"})
		return
	}

	// Retourner la ressource mise à jour
	c.JSON(http.StatusOK, resource)
}

// EnableResource met à jour une ressource en "disponible".
func EnableResource(c *gin.Context) {
	var resource models.Resource
	id := c.Param("id")

	// Récupérer la ressource dans la base de données
	if err := database.DB.First(&resource, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ressource introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de la ressource"})
		return
	}

	// Mettre à jour le statut
	resource.Status = "disponible"

	if err := database.DB.Save(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de la ressource"})
		return
	}

	// Retourner la ressource mise à jour
	c.JSON(http.StatusOK, resource)


}
