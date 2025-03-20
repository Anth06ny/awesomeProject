package models

import (
	"time"
)

// Utilisateur
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`          // À hacher avec bcrypt
	Loans    []Loan `gorm:"foreignKey:UserID"` // Relation avec les prêts
}

// Ressource (Livre ou Jeu)
type Resource struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Type   string `gorm:"not null"`              // "Livre" ou "Jeu"
	Status string `gorm:"default:disponible"`    // "disponible" ou "indisponible"
	Loans  []Loan `gorm:"foreignKey:ResourceID"` // Historique des prêts
}

// Prêt d'un livre ou jeu
type Loan struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	ResourceID uint      `gorm:"not null"`
	LoanDate   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReturnDate time.Time `gorm:"not null"`
	Status     string    `gorm:"default:en_cours"` // "en_cours" ou "retourné"
}
