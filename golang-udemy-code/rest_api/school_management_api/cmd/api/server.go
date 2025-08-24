package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "school_management_api/internal/api/middlewares"
	"school_management_api/internal/api/router"
	"school_management_api/pkg/utils"
	"time"

	"github.com/joho/godotenv"
)

//go:embed .env
var envFile embed.FS

func loadEnvFromEmbeddedFile() {
	// Read the embedded .env file
	content, err := envFile.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Create a temp file to load the env vars
	tempfile, err := os.CreateTemp("", ".env")
	if err != nil {
		log.Fatalf("Error creating temp .env file: %v", err)
	}
	defer os.Remove(tempfile.Name())

	// Write the content of the embedded .env file to the time file
	_, err = tempfile.Write(content)
	if err != nil {
		log.Fatalf("Error writing to temp .env file: %v", err)
	}

	err = tempfile.Close()
	if err != nil {
		log.Fatalf("Error closing temp .env file: %v", err)
	}

	// Load env vars from the temp file
	err = godotenv.Load(tempfile.Name())
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	// Only in development for running source code
	// err := godotenv.Load()
	// if err != nil {
	// 	return
	// }

	// Load environment variables from the embedded .env file
	loadEnvFromEmbeddedFile()

	fmt.Println("Environment Variable CERT_FILE:", os.Getenv("CERT_FILE"))

	port := os.Getenv("API_PORT")
	// cert := "cert.pem"
	// key := "key.pem"
	cert := os.Getenv("CERT_FILE")
	key := os.Getenv("KEY_FILE")

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	// mux := router.MainRouter()
	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	router := router.MainRouter()

	jwtMiddleware := mw.MiddlewaresExcludePaths(mw.JWTMiddleware, "/execs/login", "/execs/forgotpassword", "/execs/resetpassword/reset")

	secureMux := utils.ApplyMiddlewares(router, mw.SecurityHeaders, mw.Compression, mw.Hpp(hppOptions), mw.XSSMiddleware, jwtMiddleware, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	// secureMux := jwtMiddleware(mw.SecurityHeaders(router))
	// secureMux := mw.SecurityHeaders(router)
	// secureMux := mw.XSSMiddleware(router)
	// create custom server
	server := &http.Server{
		Addr:    port,
		Handler: secureMux,
		// Handler:   mw.SecurityHeaders(mux),
		// Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
