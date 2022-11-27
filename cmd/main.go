package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DmitryLogunov/go-proxy-server/internal/helpers/files"
	"github.com/DmitryLogunov/go-proxy-server/internal/helpers/logger"
	"github.com/DmitryLogunov/go-proxy-server/internal/http-server/http-proxy"
	"github.com/DmitryLogunov/go-proxy-server/internal/http-server/middlewares/auth"

	"github.com/joho/godotenv"
	"github.com/justinas/alice"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80"
	}

	jwtAuthMiddleware := alice.New(auth.ValidateJWT)
	tokenAuthMiddleware := alice.New(auth.ValidateToken)
	noneAuthMiddleware := alice.New(auth.NoneAuthenticate)

	logger.Info("\n  Http proxy server is listening on " + port + " port ... \n\n")
	routes, _ := files.ReadTwoLevelYaml("./routes.yml")

	logger.Info("Proxy routes: \n\n")

	for routeEndpoint, routeData := range routes {
		logger.Info("/" + routeEndpoint + ": \n")

		var authenticationStrategy = routeData["authentication"]
		var proxyUrl = routeData["url"]

		logger.Info(" --> url: " + proxyUrl)
		logger.Info(" --> authentication: " + authenticationStrategy)

		handler := http_proxy.Handler(routeEndpoint, proxyUrl)

		if authenticationStrategy == "jwt" {
			http.Handle("/"+routeEndpoint, jwtAuthMiddleware.ThenFunc(handler))
			continue
		}

		if authenticationStrategy == "token" {
			http.Handle("/"+routeEndpoint, tokenAuthMiddleware.ThenFunc(handler))
			continue
		}

		http.Handle("/"+routeEndpoint, noneAuthMiddleware.ThenFunc(handler))
	}

	logger.Info("\n ----------------------------------- \n\n")

	logger.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
