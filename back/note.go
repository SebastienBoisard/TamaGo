// Package back provides the back-end server: an access to several RPC for storing and retrieving notes from MongoDB.
///////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 - Sébastien Boisard.
// All rights reserved.
//
// Licensed under the GNU Affero General Public License (AGPL v3.0).
// See LICENSE file in the project root for full license information.
//
// Author: Sébastien Boisard (sebastien.boisard@gmail.com)
///////////////////////////////////////////////////////////////////////////////
package back

import (
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Note describes the different data of a note
type Note struct {
	ID        bson.ObjectId `json:"id"        bson:"_id"`
	Timestamp int64         `json:"timestamp" bson:"timestamp"`
	Tags      []*Tag        `json:"tags"      bson:"tags"`
	Content   string        `json:"content"   bson:"content"`
}

func (n *Note) String() string {
	var buf bytes.Buffer

	tm := time.Unix(n.Timestamp, 0)

	buf.WriteString(fmt.Sprintf("{id: %v, ", n.ID.Hex()))

	buf.WriteString(fmt.Sprintf("timestamp: %s, ", tm.Format("2006-01-02 15:04:05")))
	buf.WriteString("tags: [")
	for _, t := range n.Tags {
		buf.WriteString(fmt.Sprintf(`"%s"`, t.Label))
		buf.WriteString(", ")
	}
	buf.WriteString("], ")
	buf.WriteString(fmt.Sprintf(`content: "%s"}`, n.Content))

	return buf.String()
}
