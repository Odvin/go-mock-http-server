package app

type Application struct {
	store Store
}

func Init(store Store) *Application {
	return &Application{
		store: store,
	}
}
