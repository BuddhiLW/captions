// Maintainer 2025 captions Pedro G. Branquinho
package transcribe

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
	"gopkg.in/yaml.v3"
)

// var Data abs.YamlMap
// var PathData abs.YamlMapDevcontainer
// var NsMap *abs.Namespaces
// var DCPath string

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()

	// // YAML string
	// var YamlData, err = Z.Conf.Data()
	// if err != nil {
	// 	panic(err)
	// }

	// // Parse the YAML
	// err = yaml.Unmarshal([]byte(YamlData), &Data)
	// if err != nil {
	// 	panic(err)
	// }

	// NsMap = abs.NewNamespaces(Data)
}

// Captions Map of Dependencies
var Cmd = &Z.Cmd{
	Name:      `dc`,
	Summary:   `A Bonzai composite command tree, to facilitate transcribing media (audio and videos) to multiple languages`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2025 Pedro G. Branquinho`,
	License:   `Apache-2.0`,
	Commands: []*Z.Cmd{
		// standard external branch imports (see rwxrob/{help,conf,vars})
		help.Cmd, conf.Cmd, vars.Cmd,

		// local commands (in this module)
		WhisperCmd,
		// DockerComposeCmd, RegisterCmd, ListCmd, RemoveCmd, ContainerPathSetCmd,
	},

	Description: `
	**captions** (_captions_): Makes your life as easy and modular as possible, when dealing with generating video timestampped-subscriptions in multiple natural languages.

	Other useful commands:

	* captions **command** help	(every **command** has *help*)

	Read the **README.md** for more information and examples, or use **help** to see another manual page about the specific command-tree branch.
`,
}
