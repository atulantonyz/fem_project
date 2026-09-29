package routes

import (
	"github.com/atulantonyz/femProject/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)
		r.Use(app.Middleware.RequireUser)

		r.Get("/workouts/{id}", app.WorkoutHandler.HandlerGetWorkoutByID)
		r.Post("/workouts", app.WorkoutHandler.HandlerCreateWorkout)
		r.Put("/workouts/{id}", app.WorkoutHandler.HandlerUpdateWorkoutByID)
		r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkoutByID)
	})

	r.Get("/health", app.HealthCheck)
	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r
}
