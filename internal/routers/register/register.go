package register

import (
	controller "webgin/internal/pkg/controller/register"
	repository "webgin/internal/pkg/repository/register"
	usecase "webgin/internal/pkg/usecase/register"
)

func Controller() controller.Register {
	repository := repository.Repository()
	usecase := usecase.Usecase(repository)
	controller := controller.RegisterController(usecase)

	return controller
}
