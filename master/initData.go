package master

import (
	"database/sql"
	"linkAja/master/controllers"
	"linkAja/master/repositories"
	"linkAja/master/usecases"

	"github.com/gorilla/mux"
)

func InitData(r *mux.Router, db *sql.DB)  {
	accRepo := repositories.InitAccountRepoImpl(db)
	accUsecase := usecases.InitAccountUsecase(accRepo,db)
	controllers.AccountController(r,accUsecase)
}