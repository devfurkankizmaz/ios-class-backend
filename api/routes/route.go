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

	open := g.Group("/v1")
	NewPlaceOpenRouter(db, open)
	NewGalleryOpenRouter(db, open)

	// All Protected APIs
	protected := g.Group("/v1", middleware.MiddlewareJWT)
	NewProfileRouter(db, protected)
	NewAddressRouter(db, protected)
	NewGalleryRouter(db, protected)
	NewPlaceProtectedRouter(db, protected)
	NewVisitRouter(db, protected)

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
	group.PUT("/change-password", h.ChangePassword)
	group.PUT("/edit-profile", h.UpdateProfile)
}

func NewVisitRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewVisitRepository(db)
	pr := repository.NewPlaceRepository(db)
	h := &handlers.VisitHandler{
		VisitService: service.NewVisitService(r),
		PlaceService: service.NewPlaceService(pr),
	}
	group.POST("/visits", h.Create)
	group.GET("/visits", h.FetchAllByUserID)
	group.GET("/visits/:visitId", h.FetchByID)
	group.GET("/visits/user", h.FetchAllByUserID)
	group.GET("/visits/user/:placeId", h.FetchByPlaceID)
	group.DELETE("/visits/:visitId", h.DeleteByID)
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
	group.DELETE("/galleries/:placeId/:imageId", h.DeleteImageByPlaceID)
}

func NewGalleryOpenRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewGalleryRepository(db)
	h := &handlers.GalleryHandler{
		GalleryService: service.NewGalleryService(r),
	}
	group.GET("/galleries/:placeId", h.GetImagesByPlaceID)
}

func NewPlaceOpenRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewPlaceRepository(db)
	h := &handlers.PlaceHandler{
		PlaceService: service.NewPlaceService(r),
	}
	group.GET("/places", h.FetchAll)
	group.GET("/places/last", h.FetchLastN)
	group.GET("/places/popular", h.FetchRandomN)
	group.GET("/places/:placeId", h.FetchByID)
}

func NewPlaceProtectedRouter(db *gorm.DB, group *echo.Group) {
	r := repository.NewPlaceRepository(db)
	vr := repository.NewVisitRepository(db)
	h := &handlers.PlaceHandler{
		PlaceService: service.NewPlaceService(r),
		VisitService: service.NewVisitService(vr),
	}
	group.GET("/places/user", h.FetchAllByUserID)
	group.POST("/places", h.Create)
	group.PUT("/places/:placeId", h.UpdateByID)
	group.DELETE("/places/:placeId", h.DeleteByID)
}
