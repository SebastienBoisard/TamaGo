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
	"fmt"
	"gopkg.in/mgo.v2"
)

// PrepareDB ensures presence of persistent and immutable data in the DB.
func PrepareDB(session *mgo.Session) {

	// Optional. Switch the session to a monotonic behavior.
	// Monotonic mode means that the client opens a single connection to some secondary node.
	// All reads happen through this connection. When a write happens, the client drops the connection and connects
	// to the primary node, and then performs the write. Reads following a write are performed from the primary node.
	// Cf. https://stackoverflow.com/questions/38572332/compare-consistency-models-used-in-mgo
	// Cf. http://docs.mongodb.org/manual/reference/read-preference/
	session.SetMode(mgo.Monotonic, true)

	// Define the indexes to set up in TamaGo.
	indexes := make(map[string]mgo.Index)
	indexes["notes"] = mgo.Index{
		Key: []string{"$text:content"},
		//Unique:     true,
		//DropDups:   true,
		//Background: false,
	}

	for collectionName, index := range indexes {
		err := session.DB("tamago").C(collectionName).EnsureIndex(index)
		if err != nil {
			panic(fmt.Sprintf("Error while initializing the database (err=%s)", err))
		}
	}
}
