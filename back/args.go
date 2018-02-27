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
	"gopkg.in/mgo.v2/bson"
)


type QueryArgs struct {
	Query string `json:"query"`
}

type IDArgs struct {
	ID bson.ObjectId `json:"id"`
}
