package routes

import (
	"github.com/atulantonyz/femProject/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandlerGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandlerCreateWorkout)

	return r
}
