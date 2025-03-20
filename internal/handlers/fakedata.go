package handlers

import (
	"net/http"

	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FillWithFakeData ajoute un jeu de données dans la table des ressources
// en évitant d'insérer des doublons si une ressource avec le même titre existe déjà.
func FillWithFakeData(c *gin.Context) {
	// Jeu de données cohérent avec le domaine fonctionnel
	fakeResources := []models.Resource{
		// Livres
		{Title: "1984", Type: "Livre", Status: "disponible"},
		{Title: "Le Petit Prince", Type: "Livre", Status: "disponible"},
		{Title: "Harry Potter à l'école des sorciers", Type: "Livre", Status: "disponible"},
		{Title: "Les Misérables", Type: "Livre", Status: "disponible"},
		{Title: "L'Étranger", Type: "Livre", Status: "disponible"},
		{Title: "Don Quichotte", Type: "Livre", Status: "disponible"},
		{Title: "Moby Dick", Type: "Livre", Status: "disponible"},
		{Title: "Crime et Châtiment", Type: "Livre", Status: "disponible"},
		{Title: "Gatsby le Magnifique", Type: "Livre", Status: "disponible"},
		{Title: "Orgueil et Préjugés", Type: "Livre", Status: "disponible"},
		{Title: "Le Comte de Monte-Cristo", Type: "Livre", Status: "disponible"},
		{Title: "La Peste", Type: "Livre", Status: "disponible"},
		{Title: "Les Fleurs du mal", Type: "Livre", Status: "disponible"},
		{Title: "Le Rouge et le Noir", Type: "Livre", Status: "disponible"},
		{Title: "Voyage au centre de la Terre", Type: "Livre", Status: "disponible"},
		{Title: "Vingt mille lieues sous les mers", Type: "Livre", Status: "disponible"},
		{Title: "La Métamorphose", Type: "Livre", Status: "disponible"},
		{Title: "Les Trois Mousquetaires", Type: "Livre", Status: "disponible"},
		{Title: "Le Seigneur des Anneaux", Type: "Livre", Status: "disponible"},
		{Title: "Hunger Games", Type: "Livre", Status: "disponible"},
		{Title: "Dune", Type: "Livre", Status: "disponible"},
		{Title: "Sherlock Holmes : Une étude en rouge", Type: "Livre", Status: "disponible"},
		{Title: "L'Île mystérieuse", Type: "Livre", Status: "indisponible"},
		{Title: "Frankenstein", Type: "Livre", Status: "disponible"},
		{Title: "Dracula", Type: "Livre", Status: "indisponible"},
		{Title: "Le Parfum", Type: "Livre", Status: "disponible"},
		{Title: "Le Nom de la Rose", Type: "Livre", Status: "disponible"},
		{Title: "La Nuit des temps", Type: "Livre", Status: "disponible"},
		{Title: "L'Alchimiste", Type: "Livre", Status: "indisponible"},
		{Title: "Les Hauts de Hurlevent", Type: "Livre", Status: "disponible"},

		// Jeux de plateau
		{Title: "Catan", Type: "Jeu", Status: "indisponible"},
		{Title: "Risk", Type: "Jeu", Status: "disponible"},
		{Title: "Carcassonne", Type: "Jeu", Status: "disponible"},
		{Title: "Les Aventuriers du Rail", Type: "Jeu", Status: "disponible"},
		{Title: "Splendor", Type: "Jeu", Status: "disponible"},
		{Title: "Dixit", Type: "Jeu", Status: "disponible"},
		{Title: "7 Wonders", Type: "Jeu", Status: "disponible"},
		{Title: "Terraforming Mars", Type: "Jeu", Status: "indisponible"},
		{Title: "Azul", Type: "Jeu", Status: "disponible"},
		{Title: "Pandemic", Type: "Jeu", Status: "indisponible"},
		{Title: "Kingdomino", Type: "Jeu", Status: "disponible"},
		{Title: "Codenames", Type: "Jeu", Status: "disponible"},
		{Title: "Small World", Type: "Jeu", Status: "disponible"},
		{Title: "Scythe", Type: "Jeu", Status: "disponible"},
		{Title: "Agricola", Type: "Jeu", Status: "disponible"},
		{Title: "Everdell", Type: "Jeu", Status: "indisponible"},
		{Title: "Root", Type: "Jeu", Status: "indisponible"},
		{Title: "Wingspan", Type: "Jeu", Status: "disponible"},
		{Title: "Architectes du Royaume de l'Ouest", Type: "Jeu", Status: "indisponible"},
		{Title: "Brass: Birmingham", Type: "Jeu", Status: "disponible"},
		{Title: "Spirit Island", Type: "Jeu", Status: "disponible"},
		{Title: "Gloomhaven", Type: "Jeu", Status: "indisponible"},
		{Title: "Clank!", Type: "Jeu", Status: "indisponible"},
		{Title: "Paladins du Royaume de l'Ouest", Type: "Jeu", Status: "disponible"},
		{Title: "The Crew", Type: "Jeu", Status: "disponible"},
		{Title: "The Mind", Type: "Jeu", Status: "indisponible"},
		{Title: "Tapestry", Type: "Jeu", Status: "disponible"},
		{Title: "Anachrony", Type: "Jeu", Status: "indisponible"},
		{Title: "Project Gaia", Type: "Jeu", Status: "indisponible"},
		{Title: "Barrage", Type: "Jeu", Status: "disponible"},
	}

	var addedCount int = 0

	// Parcours de chaque ressource et insertion uniquement si elle n'existe pas déjà.
	for _, resource := range fakeResources {
		var existing models.Resource
		err := database.DB.Where("title = ?", resource.Title).First(&existing).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// La ressource n'existe pas, on peut l'insérer.
				if err := database.DB.Create(&resource).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de la ressource " + resource.Title})
					return
				}
				addedCount++
			} else {
				// Une erreur est survenue lors de la vérification de l'existence.
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la vérification de la ressource " + resource.Title})
				return
			}
		}
		// Si aucune erreur, c'est que la ressource existe déjà : on passe à la suivante.
	}

	c.JSON(http.StatusOK, gin.H{"message": "Données de test insérées", "added": addedCount})
}
