package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

func customerName(first, last string) string {
	return fmt.Sprintf("%s %s", first, last)
}

func main() {
	text := `
Dear {{ fullName .FirstName .LastName }},

    Thank you for registering with {{ .Product }}!

    Here are you customer credentials:
      o UserName:     {{ .Login }}
      o Access Token: {{ .AccessToken }}
      o Access Roles: {{ join .Roles ", " }}

Best Regards,
Melisandre
VMWare Customer Satisfaction Specialist
`
	fMap := template.FuncMap{
		"fullName": customerName,
		"join":     strings.Join,
	}

	tpl := template.Must(
		template.New("Mailer").Funcs(fMap).Parse(text),
	)

	customer := struct {
		FirstName   string
		LastName    string
		Product     string
		Login       string
		AccessToken string
		Roles       []string
	}{
		"John",
		"Deer",
		"VMWare Cloud Foundation",
		"jdeer",
		"getAR0p3!",
		[]string{"Admin", "SuperUser", "AverageJohn"},
	}

	err := tpl.Execute(os.Stdout, customer)
	if err != nil {
		panic(err)
	}
}
