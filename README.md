# Fruit_Slot_API using Golang

## Projcet Purpose

This is a backend slot game API written in Go. The "/play" endpoint returns 3 randomly selected fruits and a message based on the result:

"You win!"– if at least 2 of the 3 fruits are the same else "Try again!".

Basically here we have used:
* "crypto/rand" for secure random number generation.
* API handling with "net/http"
* response in JSON format.

## How to Run the Server

* install GO on your system.
* clone the repository(git clone url : https://github.com/SubhamKumar-07/Fruit_Slot_API ).
* go to the directory where you want to run (e.g, cd fruit_slot_api).
* go run main.go (to run the server) or you make the exe file to run the server(go build -o play.exe) and run directly using ("play.exe" if you are using windows cmd otherwise "./play.exe" for linux systems).
* Open your browser or use a tool like Postman to get the required output.
 " http://localhost:8080/play "

## Sample response format

Example 1: Win

{
  "fruits": ["Lemon", "Lemon", "Cherry"],
  "message": "You win!"
}

Example 2: Loss

{
  "fruits": ["Grape", "Orange", "Pineapple"],
  "message": "Try again!"
}

## Technologies used 

* Language: Go
* Libraries: net/http, crypto/rand, encoding/json



