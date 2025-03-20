package database

import (
	"database/sql"
	_ "database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"

	// Import nécessaire pour enregistrer le driver modernc
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func InitDB() {
	// On définit ici le DSN (Data Source Name). Le paramètre "_fk=1" permet d'activer les clés étrangères.
	sqlDB, err := sql.Open("sqlite", "file:database.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la connexion SQL: %v", err)
	}

	// Initialisation de GORM en utilisant la connexion existante.
	DB, err = gorm.Open(sqlite.Dialector{
		Conn: sqlDB,
	}, &gorm.Config{})
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de GORM: %v", err)
	}


}

// CloseDB permet de fermer proprement la connexion à la base de données.
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du sql.DB: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Erreur lors de la fermeture de la base de données: %v", err)
	}
}