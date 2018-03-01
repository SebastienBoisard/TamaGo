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
	"reflect"
	"testing"
)

func TestAddNote(t *testing.T) {

	t.Skip("@todo: test to implement")

	var testCases = []struct {
		note        Note   // note to add
		expectedID  string // expected id of the note
		expectedErr error  // expected error
	}{
		{note: Note{Tags: []*Tag{{Label: "test_tag"}}, Content: "test"}, expectedID: "id", expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("AddNote(%s)", tc.note.String()), func(t *testing.T) {
			nm := NoteManager{}
			var actualID string

			actualErr := nm.AddNote(&tc.note, &actualID)

			if actualErr != tc.expectedErr && tc.expectedErr != nil {
				t.Errorf("expected err=%v, got err=%v", tc.expectedErr, actualErr)
			}

			if actualID != tc.expectedID {
				t.Errorf("expected id=%s, got id=%s", tc.expectedID, actualID)
			}
		})
	}
}

func TestGetNote(t *testing.T) {

	t.Skip("@todo: test to implement")

	var testCases = []struct {
		id           string // id of the note to find
		expectedNote Note   // expected note to retrieve
		expectedErr  error  // expected error
	}{
		{id: "id", expectedNote: Note{Tags: []*Tag{{Label: "test_tag"}}, Content: "test"}, expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("GetNote(%s)", tc.id), func(t *testing.T) {
			nm := NoteManager{}
			var actualNote Note

			actualErr := nm.GetNote(&tc.id, &actualNote)

			if actualErr != tc.expectedErr && tc.expectedErr != nil {
				t.Errorf("expected err=%v, got err=%v", tc.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(actualNote, tc.expectedNote) {
				t.Errorf("expected note=%s, got note=%s", tc.expectedNote.String(), actualNote.String())
			}
		})
	}
}

func TestFindNotes(t *testing.T) {

	t.Skip("@todo: test to implement")

	var testCases = []struct {
		query         string // id of the note to find
		expectedNotes []Note // expected notes to retrieve
		expectedErr   error  // expected error
	}{
		{query: "", expectedNotes: []Note{{Tags: []*Tag{{Label: "test_tag"}}, Content: "test"}}, expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("FindNotes(%s)", tc.query), func(t *testing.T) {
			nm := NoteManager{}
			var actualNotes []Note

			actualErr := nm.FindNotes(&tc.query, &actualNotes)

			if actualErr != tc.expectedErr && tc.expectedErr != nil {
				t.Errorf("expected err=%v, got err=%v", tc.expectedErr, actualErr)
			}

			if len(actualNotes) != len(tc.expectedNotes) {
				t.Errorf("expected to find %d notes, got %d notes", len(tc.expectedNotes), len(actualNotes))
			}

			for i := 0; i < len(actualNotes); i++ {
				if !reflect.DeepEqual(actualNotes[i], tc.expectedNotes[i]) {
					t.Errorf("expected note[%d]=%s, got note[%d]=%s",
						i, tc.expectedNotes[i].String(), i, actualNotes[i].String())
				}
			}
		})
	}
}
