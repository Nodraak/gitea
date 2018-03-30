// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package private

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code.gitea.io/gitea/models"
)

// Response internal request response
type Response struct {
	Err string `json:"err"`
}

func decodeJSONError(resp *http.Response) *Response {
	var res Response
	err := json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		res.Err = err.Error()
	}
	return &res
}

// UpdatePublicKeyUpdated update publick key updates
func UpdatePublicKeyUpdated(keyID int64) error {
	if err := models.UpdatePublicKeyUpdated(keyID); err != nil {
		return fmt.Errorf("Failed to update public key: %s", err.Error())
	}

	return nil
}
