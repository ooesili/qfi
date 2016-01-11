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
	if err := ensureShellSupport(shell); err != nil {
		return nil, err
	}
	if err := ensureValidScriptType(scriptType); err != nil {
		return nil, err
	}
	path := getAssetPath(shell, scriptType)
	return Asset(path)
}

func ensureShellSupport(shell string) error {
	for _, supportedShell := range SupportedShells {
		if supportedShell == shell {
			return nil
		}
	}
	return fmt.Errorf("unsupported shell: %s", shell)
}

func ensureValidScriptType(scriptType string) error {
	if !(scriptType == "comp" || scriptType == "wrapper") {
		return fmt.Errorf("invalid script type: %s", scriptType)
	}
	return nil
}

func getAssetPath(shell, scriptType string) string {
	fileName := shell + "_" + scriptType
	if alias, ok := aliases[fileName]; ok {
		fileName = alias
	}
	return filepath.Join("_assets", fileName)
}
