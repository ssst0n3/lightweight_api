package db

import "github.com/ssst0n3/awesome_libs"

func MapObjectById(objects []map[string]interface{}) map[uint]awesome_libs.Dict {
	result := map[uint]awesome_libs.Dict{}
	for _, object := range objects {
		id := object["id"].(uint)
		result[id] = object
	}
	return result
}
