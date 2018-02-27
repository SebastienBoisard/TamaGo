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
	"fmt"
	"github.com/SebastienBoisard/tamago"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tamago",
	Long:  `All software has versions. This is tamago's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Tamago version %s\n", tamago.TamaGoVersion)
	},
}
