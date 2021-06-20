package db

import "github.com/ssst0n3/awesome_libs"

func MapObjectById(objects []map[string]interface{}) map[int64]awesome_libs.Dict {
	result := map[int64]awesome_libs.Dict{}
	for _, object := range objects {
		id := object["id"].(int64)
		result[id] = object
	}
	return result
}
