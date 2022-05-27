// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	stdCtx "context"
	"github.com/linkall-labs/vanus/internal/primitive/vanus"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
	"time"
)

func TestBlock_Creation(t *testing.T) {
	Convey("test block creation", t, func() {
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
		capacity := (rd.Int63n(128) + 16) * 1024 * 1024
		id := vanus.NewID()
		blk, err := Create(stdCtx.Background(), "/tmp", id, capacity)
		So(err, ShouldBeNil)

		So(blk.ID(), ShouldEqual, id)
		So(blk.Path(), ShouldEqual, resolvePath("/tmp", id))
		So(blk.Appendable(), ShouldBeTrue)
		So(blk.size(), ShouldEqual, 0)
		So(blk.remaining(0, 0), ShouldEqual, capacity-headerSize)
		So(blk.persistHeader(stdCtx.Background()), ShouldBeNil)
		So(blk.loadHeader(stdCtx.Background()), ShouldBeNil)
	})
}

func TestBlock_Index(t *testing.T) {
	Convey("test block index", t, func() {

	})
}

func randomGenerateData(rd rand.Rand) []byte {
	size := rd.Int31n(1024*1024) + 1
	data := make([]byte, size)
	for idx := 0; idx < int(size); idx++ {
		data[idx] = 'a'
	}
	return data
}
