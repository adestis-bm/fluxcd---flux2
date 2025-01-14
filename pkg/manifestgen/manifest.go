/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package manifestgen

import (
	"fmt"
	"os"
	"path/filepath"

	securejoin "github.com/cyphar/filepath-securejoin"
)

// Manifest holds the data of a multi-doc YAML
type Manifest struct {
	// Relative path to the YAML file
	Path string
	// Content in YAML format
	Content string
}

// WriteFile writes the YAML content to a file inside the the root path.
// If the file does not exist, WriteFile creates it with permissions perm,
// otherwise WriteFile overwrites the file, without changing permissions.
func (m *Manifest) WriteFile(rootDir string) (string, error) {
	output, err := securejoin.SecureJoin(rootDir, m.Path)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Dir(output), os.ModePerm); err != nil {
		return "", fmt.Errorf("unable to create dir, error: %w", err)
	}

	if err := os.WriteFile(output, []byte(m.Content), os.ModePerm); err != nil {
		return "", fmt.Errorf("unable to write file, error: %w", err)
	}
	return output, nil
}
