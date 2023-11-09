package configuration

import (
	"os"
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/repos"
)

var (
	// Info for all secrets
	secrets SecretInfo
	// Info for all outlines
	outline RegularOutline
	// Info for all configs
	config RegularConfig
)

// StructureRootFolder returns the folder name for the root of the clone structure.
func (outline RegularOutline) StructureRootFolder() string {
	return "repos"
}

// StructureRootFolder returns the folder name for the root of the clone structure.
func (outline RegularOutline) StructureRootFolder() string {
	return "repos"
}

// StructureRoot returns the absolute path to the root of the clone structure.
func (outline RegularOutline) StructureRoot() string {
	return filepath.Join(outline.Path)
}

// GetSecrets returns the currently saved secrets.
func GetSecrets() SecretInfo {
	if secrets != (SecretInfo{}) {
		return secrets
	}
	path := filepath.Join(repos.Expand(fghPath), secretsName)
	if !exists(path) {
		return SecretInfo{}
	}
	err := read(path, &secrets)
	if err != nil {
		Error(fmt.Sprintf("Failed to read %v", secretsName), err, 1)
	}
	return secrets
}

// GetConfig returns the currently saved config.
func GetConfig() RegularConfig {
	if config != (RegularConfig{}) {
		return config
	}
	path := filepath.Join(repos.Expand(fghPath), configName)
	if !exists(path) {
		return RegularConfig{}
	}
	err := read(path, &config)
	if err != nil {
		Error(fmt.Sprintf("Failed to read %v", configName), err, 1)
	}
	return config
}

// SaveSecrets saves the secrets to the fs.
func SaveSecrets() {
	err := write(secretsPath(), secrets)
	if err != nil {
		Error(fmt.Sprintf("Failed to save %v", secretsName), err, 1)
	}
}

// SaveConfig saves the config to the fs.
func SaveConfig() {
	err := write(configPath(), config)
	if err != nil {
		Error(fmt.Sprintf("Failed to save %v", configName), err, 1)
	}
}

// Check that the path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
You should propagate the context wherever required.There seems to be an error in the code. The closing parentheses should be inside the `checkConfigPath` function. Here's the corrected version:

```go
// Check that the path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
