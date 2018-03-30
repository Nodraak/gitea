// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package private

import (
	"fmt"

	"code.gitea.io/gitea/models"
)

// GetProtectedBranchBy get protected branch information
func GetProtectedBranchBy(repoID int64, branchName string) (*models.ProtectedBranch, error) {
	protectBranch, err := models.GetProtectedBranchBy(repoID, branchName)
	if err != nil {
		return nil, fmt.Errorf("GetProtectedBranchBy: failed to get protected branch: %s", err.Error())
	}

	if protectBranch == nil {
		protectBranch = &models.ProtectedBranch{ID: 0}
	}
	return protectBranch, nil
}

// CanUserPush returns if user can push
func CanUserPush(protectedBranchID, userID int64) (bool, error) {
	pbID := protectedBranchID

	protectBranch, err := models.GetProtectedBranchByID(pbID)
	if err != nil {
		return false, fmt.Errorf("CanUserPush: failed to retrieve push user: %s", err.Error())
	} else if protectBranch != nil {
		return protectBranch.CanUserPush(userID), nil
	} else {
		return false, nil
	}
}
