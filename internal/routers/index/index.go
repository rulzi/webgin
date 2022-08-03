package index

import (
	controller "webgin/internal/pkg/controller/index"
)

func Controller() controller.Index {
	controller := controller.IndexController()

	return controller
}
