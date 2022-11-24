package model

type Product struct {
    Name string `json: name`
    Updated string `json: updated`
}

var Products []Product

func init() {
    Products = append(Products,
        Product{Name: "Green Willy Stranger", Updated: "2022-11-22 00:11:30"},
        Product{Name: "Catatonic Tegridy Bud", Updated: "2022-11-22 00:11:43"},
        Product{Name: "Organic House Blend", Updated: "2022-11-22 00:11:52"},
        Product{Name: "Tegridy Jungle Bud", Updated: "2022-11-22 00:12:00"},
        Product{Name: "Tegridy Weed", Updated: "2022-11-22 00:12:11"},
    )
}
