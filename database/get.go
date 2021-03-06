package database

import (
	"github.com/keftcha/markovchaingo/database/inmemorydatabase"
	"github.com/keftcha/markovchaingo/database/jsondatabase"
	"strings"
)

// Get the database that implement the base interface
func Get(connectionString string) Base {
	conStrSplit := strings.Split(connectionString, "://")
	dbType, resource := conStrSplit[0], conStrSplit[1]

	switch dbType {
	case "file": // For a JSON (ex: file:///path/to/file.json)
		jsnDb := jsondatabase.New(resource)
		return &jsnDb
	case "in-memory": // For in memory database (ex: in-memory://_)
		inMemDb := inmemorydatabase.New()
		return &inMemDb
	default:
		panic("Database not supported")
	}
}
