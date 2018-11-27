// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("apm-server", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
