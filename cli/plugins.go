package cli

import (
	"jochum.dev/orb/orb/util/container"
)

var Plugins = container.NewPlugins(
	func(opts ...Option) Cli { return nil }, // Plugin factory
	func() Config { return nil },            // Config factory
)