package storageutil

import (
	"cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
)

func ShouldRetry(err error) (b bool) {
	b = storage.ShouldRetry(err)
	if b {
		return
	}

	// HTTP 401 errors - Invalid Credentials
	// This is a work-around to fix the corner case where GCSFuse checks the token
	// as valid but GCS says invalid. This might be due to client-server timer
	// issues. Actual fix will be refresh the token earlier than 1 hr.
	// Changes will be done post resolution of the below issue:
	// https://github.com/golang/oauth2/issues/623
	// TODO: Please incorporate the correct fix post resolution of the above issue.
	if typed, ok := err.(*googleapi.Error); ok {
		if typed.Code == 401 {
			b = true
			return
		}
	}
	return
}
