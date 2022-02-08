# echo-mongo-api

Simple Golang REST application with Echo Framework & MongoDB

## Application

These are the list of endpoint:
Method       | URI                | Description
------------ | ------------------ | -------------
POST         | /player            | Create new player.
GET          | /player/<:playerID>  | Get player by ID.
PUT          | /player/<:playerID>  | Edit player.
GET          | /players           | Get list players.

## References

- [Medium] (<https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-echo-version-2gdg>)
- [GitHub] (<https://github.com/sangianpatrick/go-echo-mongo>)