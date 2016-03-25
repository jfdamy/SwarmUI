package cli

import (
	"github.com/jfdamy/swarmui/autoscaling"
	"github.com/codegangsta/cli"
)

func runautoscaling(c *cli.Context) {
    autoscaling.Run();
}