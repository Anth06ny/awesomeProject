package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/routes"
)

func main() {
	database.InitDB()
	defer database.CloseDB()
	
	router := routes.SetupRouter()
	// Lancement du serveur sur le port 8080
	router.Run(":8080")
}

/*
func main() {

	// Initialisation de la base de données
	database.InitDB()
	defer database.CloseDB()

	// Migration automatique des modèles (User, Resource et Loan)
	err := database.DB.AutoMigrate(&models.User{}, &models.Resource{}, &models.Loan{})
	if err != nil {
		log.Fatalf("Erreur lors de la migration de la base de données: %v", err)
	}

	// --- Création (Create) ---
	user := models.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "motdepasse", // N'oubliez pas de hacher ce mot de passe en production !
	}
	if err := database.DB.Create(&user).Error; err != nil {
		log.Fatalf("Erreur lors de la création de l'utilisateur: %v", err)
	}
	fmt.Printf("Création: Utilisateur créé avec succès: %+v\n", user)

	// --- Lecture (Read) ---
	var readUser models.User
	if err := database.DB.First(&readUser, user.ID).Error; err != nil {
		log.Fatalf("Erreur lors de la lecture de l'utilisateur: %v", err)
	}
	fmt.Printf("Lecture: Utilisateur récupéré: %+v\n", readUser)

	// --- Mise à jour (Update) ---
	// Ici, on met à jour le nom de l'utilisateur
	updatedName := "Jane Doe"
	if err := database.DB.Model(&readUser).Update("Name", updatedName).Error; err != nil {
		log.Fatalf("Erreur lors de la mise à jour de l'utilisateur: %v", err)
	}
	fmt.Printf("Mise à jour: Le nom de l'utilisateur a été modifié en: %s\n", updatedName)

	// --- Suppression (Delete) ---
	if err := database.DB.Delete(&readUser).Error; err != nil {
		log.Fatalf("Erreur lors de la suppression de l'utilisateur: %v", err)
	}
	fmt.Println("Suppression: Utilisateur supprimé avec succès")

	/*
		r := gin.Default()

		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Bienvenue dans la bibliothèque !"})
		})

		r.Run(":8080") // Démarre le serveur sur le port 8080


}
*/
