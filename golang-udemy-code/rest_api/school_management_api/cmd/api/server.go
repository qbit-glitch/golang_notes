package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "school_management_api/internal/api/middlewares"
	"school_management_api/internal/api/router"
	"school_management_api/internal/repository/sqlconnect"
	"school_management_api/pkg/utils"
	"time"
)

func main() {

	_, err := sqlconnect.ConnectDb("dbeaver_testdb")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	port := ":3000"
	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"allowedParam"},
	}

	mux := router.Router()
	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)

	// create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux
		// Handler:   mw.SecurityHeaders(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
