// Package scripts returns wrapper and completion scripts embeded into the
// binary with go-bindata.
package scripts

//go:generate go-bindata -nometadata -pkg scripts _assets
import (
	"fmt"
	"path/filepath"
)

// alias remove the need for duplicate files
var aliases = map[string]string{
	"bash_wrapper": "sh_wrapper",
	"zsh_wrapper":  "sh_wrapper",
}

// shells is a whitelist of all supported shells
var SupportedShells = []string{"zsh", "bash", "fish"}

// Scripts holds the GetScript method so that it can be passed into interfaces.
type Scripts struct{}

// GetScript returns the script flie for the given shell and type.
func (s Scripts) GetScript(shell, scriptType string) ([]byte, error) {
	// make sure the shell is supported
	for _, supportedShell := range SupportedShells {
		if supportedShell == shell {
			goto supported
		}
	}
	return nil, fmt.Errorf("unsupported shell: %s", shell)

supported:
	// validate script type
	if !(scriptType == "comp" || scriptType == "wrapper") {
		return nil, fmt.Errorf("invalid script type: %s", scriptType)
	}

	// assemble file name
	fileName := shell + "_" + scriptType

	// check for alias
	if alias, ok := aliases[fileName]; ok {
		fileName = alias
	}

	// get asset
	path := filepath.Join("_assets", fileName)
	return Asset(path)
}
