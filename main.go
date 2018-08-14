package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/docgen/raml"
	"github.com/go-chi/render"
	"github.com/ofonimefrancis/brigg/features/photographer"
	"github.com/ofonimefrancis/brigg/internal/config"
	yaml "gopkg.in/yaml.v2"
)

// Routes returns a chi router instance which includes all routes needed for the application to run
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
		middleware.Timeout(60*time.Second),            // Timeout requests after 60 seconds
	)
	chiCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-Auth-Token", "*"},
		Debug:            false,
	})
	router.Use(chiCors.Handler)

	router.Mount("/api/users", photographer.Routes()) // Mount Golang Program debug/profiling route

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) { //TODO(tonyalaribe): Confirm if this is important, or can be replaced with notfound handler
		http.ServeFile(w, r, "./public/index.html")
	})
	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets")))
	router.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		fileServer.ServeHTTP(w, r)
	})

	return router
}

// GenerateRoutesDocumentation uses the built in chi router docgen to generate a routes documentatio. This can be triggered by running a `-gen` flag
func GenerateRoutesDocumentation(router *chi.Mux) error {
	ramlDocs := &raml.RAML{
		Title:     "Seemars API Doc",
		BaseUri:   "https://ceemars.com",
		Version:   "v1.0",
		MediaType: "application/json",
	}
	if err := chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		handlerInfo := docgen.GetFuncInfo(handler)
		resource := &raml.Resource{
			Description: handlerInfo.Comment,
		}
		return ramlDocs.Add(method, route, resource)
	}); err != nil {
		return err
	}
	ramlByt, err := yaml.Marshal(ramlDocs)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./doc/routes.yml", ramlByt, os.ModePerm)
	return err
}

func main() {
	var routes = flag.Bool("routes", false, "Generate router documentation")
	// var generateContent = flag.Bool("gen", false, "Run generator function")
	// var uploadContent = flag.Bool("upload", false, "Upload Users data")
	// var isProduction = flag.Bool("production", false, "Set Production Mode")
	// var isTest = flag.Bool("test", false, "Set test mode to use ephemeral data storage")
	// flag.Parse()

	config.Init()

	router := Routes()
	if *routes {
		err := GenerateRoutesDocumentation(router)
		if err != nil {
			log.Panicln(err)
		}
		return
	}

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("👉 %s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("⚠️  Logging err: %s\n", err.Error())
	}

	// This block of code will allow graceful shutdown of our server, using the server Shurdown method which is a part lf the standard library
	PORT := ":" + config.Get().Port
	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("😔 Shutting down. Goodbye..")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("⚠️  HTTP server Shutdown error: %v", err)
		}
		close(idleConnsClosed)
	}()
	log.Printf("Serving at 🔥 %s \n", PORT)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("⚠️  HTTP server ListenAndServe error: %v", err)
	}

	<-idleConnsClosed
}
