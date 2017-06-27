// Copyright © 2017 Julien Garcia Gonzalez <garciagonzalez.julien@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"html/template"
	"net/http"

	"fmt"

	"github.com/jgsqware/a-beer-a-day/breweryDB"
	"github.com/spf13/cobra"
)

var beerTmpl = `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>A Beer A Day</title>

  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
</head>

<body>
  <div class="container-fluid">
    <div class="row">
      <div class="col-md-12">
        <h1 class="text-center">
          A Beer A Day
        </h1>
        <div class="row">
          <div class="col-md-4">
          </div>
          <div class="col-md-4">
            <dl>
              <dt>
                Name
              </dt>
              <dd>
                {{.Beer.NameDisplay}}
              </dd>
              <dt>
                Description
              </dt>
              <dd>
                {{.Beer.Description}}
              </dd>
              <dt>
                Build Time:
              </dt>
              <dd>
                {{.BuiltTime}}
              </dd>
            </dl>
          </div>
          <div class="col-md-4">
          </div>
        </div>
      </div>
    </div>
  </div>

  <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa"
    crossorigin="anonymous"></script>
</body>

</body>
</html>
`

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a Beer a day server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting a-beer-a-day server...")
		c := breweryDB.NewClient("98bf6f0fc30df4fa79495150c567d07c")
		http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
			b, err := c.RandomBeer()

			if err != nil {
				panic(err)
			}

			template.Must(template.New("beerTmpl").Parse(beerTmpl)).Execute(rw, struct {
				Beer      breweryDB.Beer
				BuiltTime string
			}{
				b,
				buildTime,
			})
		})

		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
