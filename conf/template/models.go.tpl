package {{ .packagename }}

import(
    "github.com/jinzhu/gorm"
)

type {{.structname }} struct{
    {{ range .Fields}}
    {{ .Fieldname }} {{ .Fieldtype }}
    {{ end }}
}

func (c *{{.structname}}) Create(s {{.structname}}) error{

}

func (c *{{.structname}}) Update(s {{.structname}}) error{

}

func (c *{{.structname}}) Delete(s {{.structname}}) error{

}