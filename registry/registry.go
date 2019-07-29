package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/utrow/golang-web-base/application/usecase"
	"github.com/utrow/golang-web-base/domain/repository"
	"github.com/utrow/golang-web-base/interface/controller"
	"log"
)

type Registry interface {
	NewController() controller.Controller
}

type registry struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRegistry(db *gorm.DB, logger *log.Logger) Registry {
	return &registry{
		db:     db,
		logger: logger,
	}
}

func (r *registry) NewController() controller.Controller {
	it := usecase.NewInteractor(repository.NewRepository(r.db))
	return controller.NewController(it, *r.logger)
}
