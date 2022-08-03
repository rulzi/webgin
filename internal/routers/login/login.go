package login

import (
	controller "webgin/internal/pkg/controller/login"
	repository "webgin/internal/pkg/repository/login"
	usecase "webgin/internal/pkg/usecase/login"
)

func Controller() controller.Login {
	repository := repository.Repository()
	usecase := usecase.Usecase(repository)
	controller := controller.LoginController(usecase)

	return controller
}
