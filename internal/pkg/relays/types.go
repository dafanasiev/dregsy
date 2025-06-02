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
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/xelalexv/dregsy/internal/pkg/tags"
)

type SyncResultErrors []error

func (sre SyncResultErrors) MarshalJSON() ([]byte, error) {
	if len(sre) == 0 {
		return []byte("[]"), nil
	}

	rv := bytes.NewBuffer(nil)
	js := json.NewEncoder(rv)
	rv.WriteByte('[')
	for i, e := range sre {
		if i > 0 {
			rv.WriteByte(',')
		}
		switch e.(type) {
		case json.Marshaler:
			m, err := e.(json.Marshaler).MarshalJSON()
			if err != nil {
				return nil, errors.Wrap(err, "unable to marshal JSON error (using json.Marshaler)")
			}
			rv.Write(m)
			break
		default:
			err := js.Encode(e.Error())
			if err != nil {
				return nil, errors.Wrap(err, "unable to marshal JSON error")
			}
		}
	}
	rv.WriteByte(']')

	return rv.Bytes(), nil
}

type SyncResult struct {
	Errors SyncResultErrors `json:"errors,omitempty"`
}

func (sr *SyncResult) Err() []error {
	return sr.Errors
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
