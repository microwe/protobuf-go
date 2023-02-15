// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoreflect build tag disables use of fast-path methods.
//go:build protoreflect
// +build protoreflect

package proto

import (
	"github.com/microwe/protobuf-go/reflect/protoreflect"
	"github.com/microwe/protobuf-go/runtime/protoiface"
)

const hasProtoMethods = false

func protoMethods(m protoreflect.Message) *protoiface.Methods {
	return nil
}
