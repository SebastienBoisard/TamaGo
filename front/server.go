///////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 - Sébastien Boisard.
// All rights reserved.
//
// Licensed under the GNU Affero General Public License (AGPL v3.0).
// See LICENSE file in the project root for full license information.
//
// Author: Sébastien Boisard (sebastien.boisard@gmail.com)
///////////////////////////////////////////////////////////////////////////////
package front

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func manageHomePage(c *gin.Context) {

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title": "Home Page",
		},
	)
}

func manageAddNote(c *gin.Context) {

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"add_note.html",
		// Pass the data that the page uses
		gin.H{
			"title": "Add a note",
		},
	)
}

func manageShowNote(c *gin.Context) {

	noteID := c.Param("note_id")

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"show_note.html",
		// Pass the data that the page uses
		gin.H{
			"note_id": noteID,
		},
	)
}

func Run() {

	// Instantiate a new router
	router := gin.Default()

	router.Static("/js", "./front/js")
	router.Static("/css", "./front/css")

	// Load all the templates at the start so they don't have to be loaded from the disk again
	router.LoadHTMLGlob("front/templates/*")

	// Add a handler on /
	router.GET("/", manageHomePage)

	// Add a handler on /add_note
	router.GET("/add_note", manageAddNote)

	// Add a handler on /show_note
	router.GET("/show_note/:note_id", manageShowNote)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
