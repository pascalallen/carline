package routes

func (r Router) Fileserver() {
	r.engine.Static("/static", "./web/static")
}
