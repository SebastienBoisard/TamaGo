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
	"github.com/sebastienboisard/tamago/back"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(backCmd)
}

var backCmd = &cobra.Command{
	Use:   "back",
	Short: "Run the back server of tamago",
	Long:  `Run the back server (i.e. the json-rpc server) of tamago`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		back.Run()
	},
}
