# Golang Oktawave SDK

This library provides SDK for Oktawave API. Written in Golang.

## Usage
```go
import "github.com/oktawave-code/odk"
```

## Examples
```go
package main

import (
	"context"
	"fmt"

	"github.com/oktawave-code/odk"
)

func main() {
	auth := context.WithValue(context.Background(), odk.ContextAccessToken, "USERACCESSTOKEN")
	conf := odk.NewConfiguration()
	client := odk.NewAPIClient(conf)
	coll, resp, err := client.OCIApi.InstancesGet(auth, nil)
	fmt.Println(coll, resp, err)
}
```

## Documentation
Sdk specification available [here](./docs.md).