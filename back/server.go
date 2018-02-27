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
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Run prepares the MongoDB session, and runs the websocket handler through which the remote procedure calls will
// be received.
func Run() {

	nm := new(noteManager)

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("Error while initializing the database connection (err=%s)", err)
	}

	// Optional. Switch the session to a monotonic behavior.
	// Monotonic mode means that the client opens a single connection to some secondary node.
	// All reads happen through this connection. When a write happens, the client drops the connection and connects
	// to the primary node, and then performs the write. Reads following a write are performed from the primary node.
	// Cf. https://stackoverflow.com/questions/38572332/compare-consistency-models-used-in-mgo
	// Cf. http://docs.mongodb.org/manual/reference/read-preference/
	session.SetMode(mgo.Monotonic, true)

	// Create an index on the note content to allow query.
	c := session.DB("tamago").C("notes")
	index := mgo.Index{
		Key: []string{"$text:content"},
	}
	err = c.EnsureIndex(index)
	if err != nil {
		log.Fatalf("Error while initializing the database (err=%s)", err)
	}

	// Store the MongoDB session in the noteManager
	nm.db = session

	rpc.Register(nm)

	http.Handle("/ws", websocket.Handler(serve))
	http.ListenAndServe("localhost:7000", nil)
}

func serve(ws *websocket.Conn) {
	log.Printf("Handler starting")
	jsonrpc.ServeConn(ws)
	log.Printf("Handler exiting")
}
