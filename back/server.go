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

	nm := new(NoteManager)

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("Error while initializing the database connection (err=%s)", err)
	}

	// Configure the database and its indexes
	PrepareDB(session)

	// Store the MongoDB session in the NoteManager
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
