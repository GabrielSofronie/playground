package main

import (
	"server"
)

func main() {
	server.Start()
}

/*
Scenario:
* request to a certain route
* will trigger an action which will use a certain handler
* the handler can adapt to either a DB, FS or mock ?
* handler uses a repository which will return a response
* for a Page it will return content fields
* response is passed to views (which are using templates)
* views will render the template replacing needed fields
*/