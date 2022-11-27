package main

import (
	"net/http"
	"os"

	"github.com/DmitryLogunov/http-middlewares-proxy/helpers/files"
	"github.com/DmitryLogunov/http-middlewares-proxy/helpers/logger"
	"github.com/DmitryLogunov/http-middlewares-proxy/helpers/proxyhttp"
	"github.com/DmitryLogunov/http-middlewares-proxy/middlewares/auth"

	"github.com/justinas/alice"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80"
	}

	chainMiddlewares := alice.New(auth.ValidateJWT)

	logger.Info("\n  Http proxy server is listening on " + port + " port ... \n\n")
	routes, _ := files.ReadOneLevelYaml("./routes.yml")

	logger.Info("routes: \n\n")
	for route, proxyUrl := range routes {
		logger.Info("/" + route + " <=> " + proxyUrl)
		handler := proxyhttp.Handler(route, proxyUrl)
		http.Handle("/"+route, chainMiddlewares.ThenFunc(handler))
	}

	logger.Info("\n ----------------------------------- \n\n")

	logger.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
