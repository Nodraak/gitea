// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package private

import (
	"fmt"
	"strings"

	"code.gitea.io/git"
	"code.gitea.io/gitea/models"
)

// PushUpdate update publick key updates
func PushUpdate(opt models.PushUpdateOptions) error {
	errMessageFormat := "PushUpdate: failed to update public key: %s"

	branch := strings.TrimPrefix(opt.RefFullName, git.BranchPrefix)
	if len(branch) == 0 || opt.PusherID <= 0 {
		return fmt.Errorf(errMessageFormat, "branch or secret is empty, or pusher ID is not valid")
	}

	err := models.PushUpdate(branch, opt)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			return fmt.Errorf(errMessageFormat, "user does not exist")
		} else {
			return fmt.Errorf(errMessageFormat, err.Error())
		}
	}

	return nil
}
