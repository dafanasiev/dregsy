/*
	Copyright 2022 Alexander Vollschwitz <xelalex@gmx.net>

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	  http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package relays

import (
	"encoding/json"
	"github.com/xelalexv/dregsy/internal/pkg/tags"
)

type SyncResult struct {
	Errors []error
}

func (r *SyncResult) EncodeJSON(js *json.Encoder) error {
	errs2strings := func(errors []error) []string {
		if errors == nil {
			return nil
		}
		s := make([]string, len(errors))
		for i, err := range errors {
			s[i] = err.Error()
		}
		return s
	}

	m := make(map[string]interface{}, 1)
	if errors := errs2strings(r.Errors); errors != nil {
		m["errors"] = errors
	}

	return js.Encode(m)
}

func (r *SyncResult) Err() []error {
	return r.Errors
}

type SyncOptions struct {
	//
	SrcRef           string
	SrcAuth          string
	SrcSkipTLSVerify bool
	//
	TrgtRef           string
	TrgtAuth          string
	TrgtSkipTLSVerify bool
	//
	Tags     *tags.TagSet
	Platform string
	Verbose  bool
}

type Support interface {
	Platform(p string) error
}
