package main

import (
	"os"

	"github.com/vikaskumar1187/publisher_saasv2/services/publisher/app/services/publisher-api/v1/cmd"
	"github.com/vikaskumar1187/publisher_saasv2/services/publisher/app/services/publisher-api/v1/cmd/all"
)

/*
	Need to figure out timeouts for http service.
*/

var build = "develop"
var routes = "all" // go build -ldflags "-X main.routes=crud"

func main() {

	// The idea here is that we can build different versions of the binary
	// with different sets of exposed web APIs. By default we build a single
	// instance with all the web APIs.
	//
	// Here is the scenario. It would be nice to build two binaries, one for the
	// transactional APIs (CRUD) and one for the reporting APIs. This would allow
	// the system to run two instances of the database. One instance tuned for the
	// transactional database calls and the other tuned for the reporting calls.
	// Tuning meaning indexing and memory requirements. The two databases can be
	// kept in sync with replication.

	switch routes {
	case "all":
		if err := cmd.Main(build, all.Routes()); err != nil {
			os.Exit(1)
		}
	}
}
