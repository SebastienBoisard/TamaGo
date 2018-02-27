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
	"github.com/sebastienboisard/tamago/front"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(frontCmd)
}

var frontCmd = &cobra.Command{
	Use:   "front",
	Short: "Run the front server of tamago",
	Long:  `Run the front server (i.e. the webserver) of tamago`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		front.Run()
	},
}
