package handlers

import (
	"net/http"
	"strconv"
	"time"

	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateLoanInput définit le format attendu pour la création d'un prêt.
type CreateLoanInput struct {
	ResourceID uint   `json:"resource_id" binding:"required"`
	BorrowType string `json:"borrow_type" binding:"required,oneof=sur_place a_emporter"`
}

// CreateLoan gère la création d'un nouveau prêt pour une ressource.
func CreateLoan(c *gin.Context) {
	// Récupérer l'ID de l'utilisateur depuis le contexte
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non authentifié"})
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne"})
		return
	}

	// Lier les données JSON à notre input
	var input CreateLoanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifier que la ressource existe et est disponible
	var resource models.Resource
	if err := database.DB.First(&resource, input.ResourceID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ressource non trouvée"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la recherche de la ressource"})
		}
		return
	}

	if resource.Status != "disponible" {
		c.JSON(http.StatusConflict, gin.H{"error": "La ressource n'est pas disponible"})
		return
	}

	// Définir la date de prêt et la date de retour en fonction du type d'emprunt
	loanDate := time.Now()
	var returnDate time.Time
	if input.BorrowType == "a_emporter" {
		returnDate = loanDate.Add(15 * 24 * time.Hour) // maximum de 15 jours
	} else {
		// Pour un prêt sur place, on peut définir la date de retour égale à la date d'emprunt (ou autre logique)
		returnDate = loanDate
	}

	// Créer le prêt
	loan := models.Loan{
		UserID:     userID,
		ResourceID: input.ResourceID,
		LoanDate:   loanDate,
		ReturnDate: returnDate,
		Status:     "en_cours",
	}
	if err := database.DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du prêt"})
		return
	}

	// Mettre à jour le statut de la ressource en "emprunté"
	resource.Status = "emprunté"
	if err := database.DB.Save(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de la ressource"})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

// GetLoans récupère la liste des prêts de l'utilisateur connecté.
func GetLoans(c *gin.Context) {
	// Récupérer l'ID de l'utilisateur depuis le contexte
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non authentifié"})
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne"})
		return
	}

	// Récupérer les prêts de l'utilisateur
	var loans []models.Loan
	if err := database.DB.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des prêts"})
		return
	}

	c.JSON(http.StatusOK, loans)
}

// ReturnLoan permet de marquer le retour d'une ressource empruntée.
func ReturnLoan(c *gin.Context) {
	// Récupérer l'ID du prêt depuis les paramètres d'URL
	loanIDStr := c.Param("id")
	loanID, err := strconv.ParseUint(loanIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de prêt invalide"})
		return
	}

	// Récupérer l'ID de l'utilisateur depuis le contexte (pour vérifier l'appartenance du prêt)
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non authentifié"})
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne"})
		return
	}

	// Récupérer le prêt
	var loan models.Loan
	if err := database.DB.First(&loan, uint(loanID)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Prêt non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la recherche du prêt"})
		}
		return
	}

	// Vérifier que le prêt appartient bien à l'utilisateur (optionnel selon votre logique métier)
	if loan.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Accès non autorisé"})
		return
	}

	// Vérifier que le prêt est en cours
	if loan.Status != "en_cours" {
		c.JSON(http.StatusConflict, gin.H{"error": "Le prêt est déjà retourné"})
		return
	}

	// Marquer le prêt comme retourné
	loan.Status = "retourné"
	if err := database.DB.Save(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du prêt"})
		return
	}

	// Mettre à jour le statut de la ressource associée en "disponible"
	var resource models.Resource
	if err := database.DB.First(&resource, loan.ResourceID).Error; err == nil {
		resource.Status = "disponible"
		_ = database.DB.Save(&resource)
	}

	c.JSON(http.StatusOK, loan)
}

// Optionnel : Suppression d'un prêt en attente.
// func DeleteLoan(c *gin.Context) {
// 	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
// }
