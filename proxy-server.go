package main

import (
	"net/http"
	"os"

	"app/helpers/files"
	"app/helpers/logger"
	"app/helpers/handlers"
	"app/middlewares/auth"
	"github.com/justinas/alice"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80"
	}

	logger.Header("\n  Http proxy server is listening on " + port + " port ... \n\n")

	http.HandleFunc("/healthCheck", handlers.HealthCheckHandler)

	routes, _ := files.ReadTwoLevelYaml("./routes.yml")

	logger.Info("\n\n  - transparent routes: \n\n")
	for route, proxyUrl := range routes["transparent"] {
		logger.Info("    " + route + " <=> " + proxyUrl)
		handler := handlers.HttpProxyHandler(route, proxyUrl)
		http.HandleFunc(route, handler)		
	}
	
	logger.Info("\n\n  - authentificate routes: \n\n")	
	chainMiddlewares := alice.New(auth.ValidateJWT)
	for route, proxyUrl := range routes["authentificate"] {
		logger.Info("    " + route + " <=> " + proxyUrl)
		handler := handlers.HttpProxyHandler(route, proxyUrl)
		http.Handle(route, chainMiddlewares.ThenFunc(handler))		
	}

	logger.Debug("\n ----------------------------------- \n\n")
	
	logger.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))	
}
