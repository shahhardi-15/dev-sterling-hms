package main

import (
	"log"
	"os"

	"sterling-hms/config"
	"sterling-hms/internal/controllers"
	"sterling-hms/internal/middlewares"
	"sterling-hms/internal/repositories"
	"sterling-hms/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := config.ConnectDB()

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	apptRepo := repositories.NewAppointmentRepository(db)
	apptService := services.NewAppointmentService(apptRepo)
	apptController := controllers.NewAppointmentController(apptService)

	billingRepo := repositories.NewBillingRepository(db)
	billingService := services.NewBillingService(billingRepo)
	billingController := controllers.NewBillingController(billingService)

	prescriptionRepo := repositories.NewPrescriptionRepository(db)
	prescriptionService := services.NewPrescriptionService(prescriptionRepo, billingRepo)
	prescriptionController := controllers.NewPrescriptionController(prescriptionService)

	medicineController := controllers.NewMedicineController(db)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174", "http://localhost:5173", "https://dev-sterling-hms.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Sterling HMS API is running"})
	})

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/change-password", middlewares.JWTAuth(), authController.ChangePassword)
		}

		users := v1.Group("/users")
		users.Use(middlewares.JWTAuth())
		{
			users.GET("", middlewares.RequireRoles("admin"), userController.GetAllUsers)
			users.POST("", middlewares.RequireRoles("admin"), userController.CreateUser)
			users.GET("/:id", middlewares.RequireRoles("admin"), userController.GetUserByID)
			users.PUT("/:id", middlewares.RequireRoles("admin"), userController.UpdateUser)
			users.DELETE("/:id", middlewares.RequireRoles("admin"), userController.DeleteUser)
		}

		patients := v1.Group("/patients")
		patients.Use(middlewares.JWTAuth())
		{
			patients.GET("", middlewares.RequireRoles("admin", "receptionist", "pharmacist", "patient", "doctor"), userController.GetAllPatients)
			patients.GET("/:id/appointments", middlewares.RequireRoles("admin", "doctor", "receptionist"), apptController.GetByPatientID)
			patients.GET("/:id/prescriptions", middlewares.RequireRoles("admin", "doctor", "receptionist", "pharmacist", "patient"), prescriptionController.GetByPatientID)
		}

		doctors := v1.Group("/doctors")
		doctors.Use(middlewares.JWTAuth())
		{
			doctors.GET("", middlewares.RequireRoles("admin", "receptionist", "patient", "doctor", "pharmacist"), userController.GetAllDoctors)
			doctors.GET("/me/:id", middlewares.RequireRoles("doctor"), userController.GetDoctorByUserID)
			doctors.GET("/:id/appointments", middlewares.RequireRoles("admin", "doctor", "receptionist"), apptController.GetByDoctorID)
		}

		appointments := v1.Group("/appointments")
		appointments.Use(middlewares.JWTAuth())
		{
			appointments.GET("", middlewares.RequireRoles("admin", "receptionist", "doctor", "patient"), apptController.GetAll)
			appointments.POST("", middlewares.RequireRoles("patient", "receptionist"), apptController.Create)
			appointments.GET("/:id", middlewares.RequireRoles("admin", "doctor", "receptionist"), apptController.GetByID)
			appointments.PATCH("/:id/approve", middlewares.RequireRoles("doctor"), apptController.Approve)
			appointments.PATCH("/:id/reject", middlewares.RequireRoles("doctor", "receptionist"), apptController.Reject)
			appointments.PATCH("/:id/cancel", middlewares.RequireRoles("patient", "receptionist"), apptController.Cancel)
			appointments.PATCH("/:id/complete", middlewares.RequireRoles("doctor"), apptController.Complete)
		}

		prescriptions := v1.Group("/prescriptions")
		prescriptions.Use(middlewares.JWTAuth())
		{
			prescriptions.POST("", middlewares.RequireRoles("doctor"), prescriptionController.Create)
			prescriptions.GET("/:id", middlewares.RequireRoles("admin", "doctor", "pharmacist"), prescriptionController.GetByID)
			prescriptions.PATCH("/:id/dispense", middlewares.RequireRoles("pharmacist"), prescriptionController.Dispense)
		}

		billing := v1.Group("/billing")
		billing.Use(middlewares.JWTAuth())
		{
			billing.GET("", middlewares.RequireRoles("admin", "receptionist"), billingController.GetAll)
			billing.POST("", middlewares.RequireRoles("admin", "receptionist"), billingController.Create)
			billing.GET("/:id", middlewares.RequireRoles("admin", "receptionist"), billingController.GetByID)
			billing.PATCH("/:id/pay", middlewares.RequireRoles("admin", "receptionist"), billingController.MarkAsPaid)
		}

		medicines := v1.Group("/medicines")
		medicines.Use(middlewares.JWTAuth())
		{
			medicines.GET("", middlewares.RequireRoles("admin", "doctor", "pharmacist", "patient"), medicineController.GetAll)
			medicines.POST("", middlewares.RequireRoles("admin", "pharmacist"), medicineController.Create)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
