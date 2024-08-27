package routes

import "github.com/pascalallen/carline/internal/carline/application/http/action"

func (r Router) Default() {
	r.engine.NoRoute(action.HandleDefault())
}
