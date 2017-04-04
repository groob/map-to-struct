Convert a Go `map[string]string` type to a struct.
This utility is not supposed to be a complete solution but rather a simple tool to save a lot of typing when creating map-to-struct conversions.

# Install

```
go get -u github.com/groob/map-to-struct
```

# Usage

`map-to-struct` expects a json object at stdin.

```
osqueryi --json 'select * from os_version'|jq '.[0]' | map-to-struct

	if v, ok := osMap["minor"]; ok {
		goStruct.Minor = v
	}

	if v, ok := osMap["name"]; ok {
		goStruct.Name = v
	}

	if v, ok := osMap["patch"]; ok {
		goStruct.Patch = v
	}

	if v, ok := osMap["platform"]; ok {
		goStruct.Platform = v
	}

	if v, ok := osMap["version"]; ok {
		goStruct.Version = v
	}

	if v, ok := osMap["build"]; ok {
		goStruct.Build = v
	}

	if v, ok := osMap["codename"]; ok {
		goStruct.Codename = v
	}

	if v, ok := osMap["major"]; ok {
		goStruct.Major = v
	}

	if v, ok := osMap["platform_like"]; ok {
		goStruct.PlatformLike = v
	}
```
