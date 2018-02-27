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
	"github.com/sebastienboisard/TamaGo"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TamaGo",
	Long:  `All software has versions. This is TamaGo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Tamago version %s\n", TamaGo.TamaGoVersion)
	},
}
