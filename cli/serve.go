package cli

import (
	"github.com/jfdamy/swarmui/api"
	"github.com/codegangsta/cli"
)

func serve(c *cli.Context) {
    api.Serve();
}