# BlackRabbit

Creating a basic JSON API server in Go for a side-project

# Routes available

| Route | HTTP Method | Description |
|:-----:|:-----------:|:----------:|
| /accounts | GET | gets all user accounts |
| /accounts/{id} | GET | gets specific user account |
| /accounts | POST | create new account. Expects json object with string name, int age, and string email properties | 

# Todo

## High 
- Add authentication for accounts route
- Create Templates Model
- Create Groups Model
- Create Create Search Interface

## Medium
- Push to Heroku
- use real database, remove mocks
- add validation to http routes
- refactor routes into different folder structures

## Low
- add tests to existing code
- use actual logging lib 