package main

type User struct {
	Id int
	Valid bool
}

type CoinBaseParams struct {
	username string
}

type CoinBaseResponse struct {
	Code int
	Balance int64
}

type Error struct {
	Code int
	Message string
}

