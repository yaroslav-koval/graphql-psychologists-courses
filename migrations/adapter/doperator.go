package adapter

import "os"

func NewDirectoryOperator() *DirectoryOperator {
	return &DirectoryOperator{}
}

type DirectoryOperator struct {
}

func (do *DirectoryOperator) IsFileExist(name string) bool {
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		return true
	}

	return false
}

func (do *DirectoryOperator) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
