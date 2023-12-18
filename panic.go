package main

import "net/http"

func main() {
	// Panic is a built-in function
	// that stops the ordinary flow of control and begins panicking.
	// Panic will be executed after defer statements.

	/*
		The http.HandleFunc("/", ...)
		line sets up a handler function for the root URL ("/") of the server.
		The handler function takes in two parameters: w http.ResponseWriter and r *http.Request.
		The http.ResponseWriter is used to send the HTTP response back to the client,
		and the http.Request contains information about the incoming request.
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*
			Inside the handler function, w.Write([]byte("Hello World")) sends the string "Hello World" as the response body.
		*/
		w.Write([]byte("Hello World"))
	})
	/*
		Next, http.ListenAndServe(":8080", nil) starts the server on port 8080.
		It accepts incoming requests and dispatches them to the appropriate handler functions.
	*/
	err := http.ListenAndServe(":8080", nil)

	/*
		If there is an error starting the server,
		the err variable will be non-nil,and the program will panic by calling panic(err.Error()).
		This will cause the program to exit and print the error message.
	*/
	if err != nil {
		panic(err.Error())
	}
}
