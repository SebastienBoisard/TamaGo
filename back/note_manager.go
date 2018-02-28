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
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

// NoteManager is the object for accessing all the remote procedures.
type NoteManager struct {
	db *mgo.Session
}

// AddNote adds a note.
func (nm *NoteManager) AddNote(args *Note, reply *int) error {

	log.Printf("AddNote BEGIN")
	log.Printf("AddNote - args=%+v", args)

	// Copy() vs Clone()
	// https://godoc.org/gopkg.in/mgo.v2#Session.Copy
	// https://godoc.org/gopkg.in/mgo.v2#Session.Clone
	dbsession := nm.db.Copy()
	defer dbsession.Close()

	// We work on the "notes" collection
	c := dbsession.DB("tamago").C("notes")

	now := time.Now()
	err := c.Insert(&Note{bson.NewObjectId(), now.Unix(), []*Tag{{"go"}}, args.Content})
	if err != nil {
		log.Printf("AddNote - Error while adding a note [err=%s]", err)
		log.Printf("AddNote END")
		return err
	}

	log.Printf("AddNote END")
	return nil
}

// GetNote retrieves a not from its id.
func (nm *NoteManager) GetNote(noteID *IDArgs, reply *Note) error {

	log.Printf("GetNote BEGIN")
	log.Printf("GetNote - noteID.ID=%s\n", noteID.ID.Hex())

	dbsession := nm.db.Copy()
	defer dbsession.Close()

	// We work on the "notes" collection
	c := dbsession.DB("tamago").C("notes")

	// Find the wanted note with the proper ID
	err := c.FindId(noteID.ID).One(reply)
	if err != nil {
		log.Printf("GetNote - Error while retreiving a note (id=%s) [err=%s]", noteID.ID.Hex(), err)
		log.Printf("GetNote END")
		return err
	}

	log.Printf("GetNote - note retrieved: %s", reply.String())
	log.Printf("GetNote END")

	return nil
}

// FindNotes retrieves all the notes matching the query.
// If the query is empty, it will return all the notes.
func (nm *NoteManager) FindNotes(query *QueryArgs, reply *[]Note) error {

	log.Printf("FindNotes BEGIN")
	log.Printf("FindNotes - query=%+s", query.Query)

	dbsession := nm.db.Copy()
	defer dbsession.Close()

	// We work on the "notes" collection
	c := dbsession.DB("tamago").C("notes")

	var q *mgo.Query
	if query.Query == "" {

		// Query all notes
		q = c.Find(nil)
	} else {

		// Query all the notes matching the query
		q = c.Find(bson.M{"$text": bson.M{"$search": query.Query}})
	}

	// Retrieve notes and sort them on the timestamp
	err := q.Sort("-timestamp").All(reply)
	if err != nil {
		log.Printf("FindNotes - Error while finding notes [err=%s]", err)
		log.Printf("FindNotes END")
		return err
	}

	// Just for the logs
	for i, n := range *reply {
		fmt.Printf("FindNotes - found note %d: %s\n", i, n.String())
	}

	log.Printf("FindNotes END")

	return nil
}
