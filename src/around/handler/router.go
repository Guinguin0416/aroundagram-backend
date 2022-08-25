package handler

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware" // 经过一系列的过滤，再调用handler
	// 验证token是否还有效
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	router := mux.NewRouter()

	router.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST")
	router.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET")
	router.Handle("/post/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE")

	router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
	router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")

	originsOk := handlers.AllowedOrigins([]string{"*"}) // 决定哪些前端网址可以访问后端
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	// 前端请求带Authorization header：
	// content type header:
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(router)

}
