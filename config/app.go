package config

type App struct {
	usdrub string
}

func NewApp() App {
	return App{
		usdrub: "80.0",
	}
}
