package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, OpenFaaS Cloud. You said: %s", string(req))
}
