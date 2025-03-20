package routes

import (
	"awesomeProject/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Autorise le frontend en dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		// Routes d'authentification et gestion d'utilisateur
		api.POST("/register", handlers.RegisterUser)
		api.POST("/login", handlers.LoginUser)
		api.GET("/profile", handlers.GetProfile)
		api.PUT("/profile", handlers.UpdateProfile)

		// Routes de gestion des ressources (livres et jeux)
		//fakedata http://localhost:8080/api/resources
		api.GET("/resources", handlers.GetResources)
		api.GET("/resources/:id", handlers.GetResource)
		api.POST("/resources", handlers.CreateResource)
		api.PUT("/resources/:id/disable", handlers.DisableResource)  // Passer en indisponible
		api.PUT("/resources/:id/enable", handlers.EnableResource)    // Passer en disponible
		//fakedata http://localhost:8080/api/resources/fill
		api.GET("/resources/fill", handlers.FillWithFakeData)

		// Routes de gestion des prêts
		api.POST("/loans", handlers.CreateLoan)
		api.GET("/loans", handlers.GetLoans)
		api.PUT("/loans/:id/return", handlers.ReturnLoan)
		// Optionnel : Suppression d'un prêt en attente
		// api.DELETE("/loans/:id", handlers.DeleteLoan)
	}

	// Déclaration du dossier des assets
	//http://localhost:8080/assts
	//router.Static("/assets", "./assets")
	router.Static("/static", "./awsome_front/dist")

	router.NoRoute(func(c *gin.Context) {
		c.File("./awsome_front/dist/index.html")
	})

	return router
}
