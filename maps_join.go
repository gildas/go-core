package core

// MapJoin joins two or more maps of the same kind
//
// None of the maps are modified, a new map is created.
//
// If two maps have the same key, the latter map overwrites the value from the former map.
func MapJoin[K comparable, T any](maps ...map[K]T) map[K]T {
	results := map[K]T{}

	for _, m := range maps {
		for key, value := range m {
			results[key] = value
		}
	}
	return results
}
