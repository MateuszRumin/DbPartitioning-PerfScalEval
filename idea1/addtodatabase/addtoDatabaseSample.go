package addtoDatabase

import (
	"perfscaleval/config"
)

func Add(query string, data any, pointer int) {

	config.Connections[pointer].Query(query)

}
