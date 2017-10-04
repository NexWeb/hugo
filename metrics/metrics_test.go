// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimilarPercentage(t *testing.T) {
	assert := require.New(t)

	assert.Equal(100, similarPercentage("Hugo Rules", "Hugo Rules"))
	assert.Equal(50, similarPercentage("Hugo Rules", "Hugo Rocks"))
	assert.Equal(67, similarPercentage("The Hugo Rules", "The Hugo Rocks"))
	assert.Equal(0, similarPercentage("Totally different", "Not Same"))
}
