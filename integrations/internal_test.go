// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"

	"github.com/stretchr/testify/assert"
)

func TestInternal_GetProtectedBranch(t *testing.T) {
	prepareTestEnv(t)

	assertProtectedBranch(t, 1, "master", false, true)
	assertProtectedBranch(t, 1, "dev", false, true)
	assertProtectedBranch(t, 1, "lunny/dev", false, true)
}
