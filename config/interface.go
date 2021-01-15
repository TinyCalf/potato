package config

import (
	"potato/piface"
)

// ICompnent ..
type ICompnent interface {
	piface.ICompnent
	Load(filePath string)
}