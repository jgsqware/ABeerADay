package main

import (
	"html/template"
	"net/http"

	"github.com/jgsqware/ABeerADay/breweryDB"
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
                {{.NameDisplay}}
              </dd>
              <dt>
                Description
              </dt>
              <dd>
                {{.Description}}
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

func main() {

	c := breweryDB.NewClient("98bf6f0fc30df4fa79495150c567d07c")
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		b, err := c.RandomBeer()

		if err != nil {
			panic(err)
		}

		template.Must(template.New("beerTmpl").Parse(beerTmpl)).Execute(rw, b)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
