package up

func (up *Up) route() {
	up.router.POST("/upload/unlisted", up.uploadUnlisted)
}
