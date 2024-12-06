package boot

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yeungon/corpora/internal/config"
	sqlite "github.com/yeungon/corpora/internal/database"
	"github.com/yeungon/corpora/pkg/logs"
)

func App() {
	logs.Log()

	// Might declare more value here in the future work. Use as dependency injection.
	app := config.NewApp(
		true, false, "hello world",
	)
	// fmt.Println(app)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)
	r.Use(securityHeaders)
	r.Use(middleware.Timeout(60 * time.Second))
	Register(r, app)

	// =============sqlite=============
	sqlite.BunConnect()
	db := sqlite.DB()
	defer db.Close()
	// Enable debugging for queries
	// =============sqlite=============

	http.ListenAndServe(":9999", r)
	defer logs.CloseLog()
}
