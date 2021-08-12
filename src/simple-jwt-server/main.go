package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials a struct that models the structure of a user,
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims a struct used to represent claims to identify a user
// The embedded struct jwt.StandardClaims struct contains useful
// fields like expiry and issuer name
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key") // only suitable for dev
// In production make sure you use a real private key, preferably
// at least 256 bits in length

func Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting Signin function")
	fmt.Println(r.Method)
	var credentials Credentials

	// -- ------------------------------------------
	// Get user credentials and authenticate them
	// -- ------------------------------------------

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Look up stored password for user
	// We are using an inmemory map but could use database etc.
	expectedPassword, ok := users[credentials.Username]

	// Check if user is known and has correct password
	if !ok {
		fmt.Printf("Error Unknown user: %s", credentials.Username)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if expectedPassword != credentials.Password {
		fmt.Printf("Error incorrect password for user: %s", credentials.Username)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// -- ------------------------------------------------------
	// Now build the JWT token as we have authenticated the user
	// -- ------------------------------------------------------

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as milliseconds
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "SampleJWTService",
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	// That is where we create a token from the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT token - do this by signing the token using a
	// secure private key (jwtKey) it will create
	// {header}.{payload}.{signature}
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		fmt.Printf("Error creating the token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Println("Finished Signin")
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting Welcome function")

	// -- -------------------------------------------
	// Retrieve the JWT token form the cooke
	// -- -------------------------------------------

	// We can obtain the session token from the requests cookies, which come with every request
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Printf("Error no cookie set: %v\n", err)
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		fmt.Printf("Error accessing cookie: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tokenString := cookie.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT token string and store the result in `claims`.
	// This method will return an error if the token is invalid (if it
	// has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Printf("Error signature invalid %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Printf("Error processing JWT token %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		fmt.Printf("Error invalid token %v\n", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// -- --------------------------------------
	// If we made it this far all is ok
	// -- --------------------------------------

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

	fmt.Println("Finished Welcome function")
}

func main() {
	fmt.Println("Starting Server App")

	fmt.Println("Registering functions with urls")
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)

	// Start the server on port 8000
	fmt.Println("Starting Listen and Serve")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Printf("Error starting server %v", err)
	}
}
