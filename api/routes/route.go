package routes

import (
	"github.com/devfurkankizmaz/iosclass-backend/api/handlers"
	"github.com/devfurkankizmaz/iosclass-backend/api/middleware"
	"github.com/devfurkankizmaz/iosclass-backend/repository"
	"github.com/devfurkankizmaz/iosclass-backend/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, g *echo.Echo) {
	// All Public APIs
	public := g.Group("/v1/auth")
	NewRegisterRouter(db, public)
	NewLoginRouter(db, public)
	NewRefreshRouter(db, public)

	// All Protected APIs
	protected := g.Group("/v1", middleware.MiddlewareJWT)
	NewProfileRouter(db, protected)
	NewTravelRouter(db, protected)
	NewAddressRouter(db, protected)
	NewGalleryRouter(db, protected)
}

func NewRegisterRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := handlers.RegisterHandler{
		RegisterService: service.NewRegisterService(r),
	}
	group.POST("/register", h.Register)
}

func NewLoginRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := &handlers.LoginHandler{
		LoginService: service.NewLoginService(r),
	}
	group.POST("/login", h.Login)
}

func NewRefreshRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := &handlers.RefreshHandler{
		RefreshService: service.NewRefreshService(r),
	}
	group.POST("/refresh", h.Refresh)
}

func NewProfileRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := &handlers.ProfileHandler{
		ProfileService: service.NewProfileService(r),
	}
	group.GET("/me", h.Fetch)
}

func NewTravelRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewTravelRepository(db)
	h := &handlers.TravelHandler{
		TravelService: service.NewTravelService(r),
	}
	group.POST("/travels", h.Create)
	group.GET("/travels", h.FetchAllByUserID)
	group.GET("/travels/:travelId", h.FetchByID)
	group.PUT("/travels/:travelId", h.UpdateByID)
	group.DELETE("/travels/:travelId", h.DeleteByID)
}

func NewAddressRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewAddressRepository(db)
	h := &handlers.AddressHandler{
		AddressService: service.NewAddressService(r),
	}
	group.POST("/addresses", h.Create)
	group.GET("/addresses", h.FetchAllByUserID)
	group.GET("/addresses/:addressId", h.FetchByID)
	group.PUT("/addresses/:addressId", h.UpdateByID)
	group.DELETE("/addresses/:addressId", h.DeleteByID)
}

func NewGalleryRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewGalleryRepository(db)
	h := &handlers.GalleryHandler{
		GalleryService: service.NewGalleryService(r),
	}
	group.POST("/galleries", h.AddImageToTravel)
	group.GET("/galleries/:travelId", h.GetImagesByTravelID)
	group.DELETE("/galleries/:travelId/:imageId", h.DeleteImageByTravelID)
}
