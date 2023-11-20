# go-keno
Keno impl in Go, for displaying on a TUI




Todo:
- Finish rest routes
- do the DB queries 
- ???
- profit












Placing a Pick
Request:

POST /api/v1/picks

{
    "picks_per_game": 10,
    "picks": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    "price_per_game": 1,
    "number_games": 5
}


Response:

{
    "card_id": 12,
    "selection": "picks": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    "start_game_num": 2,
    "last_game_num": 6
}


200 Ok
400 Invalid Picks

Checking your Card
GET /api/v1/check/:card_id

Response:

{
    "amount": 10 // How much you won
}


200 Games have finished
400 Games haven't finished
404 Invalid Card ID

Other Notes
Simulating the game can be just via stdout logs, unless you want to play with websockets and make a webapp graphic.

I'm going to be using gin (https://gin-gonic.com/) for the API handling. And using sqlite as a small DB, might use gorm https://gorm.io/ as an ORM or just do sql calls directly.
