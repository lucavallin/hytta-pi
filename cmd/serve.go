// Copyright © 2018 Luca Cavallin <me@lucavall.in>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/lucavallin/hytta/pkg/api"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "SERVER: start API",

	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	http.HandleFunc("/", api.RootHandler)

	log.Println("Hytta API server is starting...")
	log.Fatal(http.ListenAndServe(":80", nil))
}