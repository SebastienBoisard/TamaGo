///////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 - Sébastien Boisard.
// All rights reserved.
//
// Licensed under the GNU Affero General Public License (AGPL v3.0).
// See LICENSE file in the project root for full license information.
//
// Author: Sébastien Boisard (sebastien.boisard@gmail.com)
///////////////////////////////////////////////////////////////////////////////
package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tamago",
	Short: "tamago is a simple note taking system",
	Long:  `tamago is a simple note taking system built by Sébastien Boisard`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// @todo: to fill
	},
}
