package main

type Product struct {
    ID          int     `json:"id"`
    ProductName string  `json:"ProductName"`
    CategoryID  int     `json:"categoryId"`
    UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
    ID           int    `json:"id"`
    CategoryName string `json:"categoryName"`
}

var products = []Product{
    {ID: 1, ProductName: "Laptop", CategoryID: 1, UnitPrice: 5000.99},
    {ID: 2, ProductName: "Mouse", CategoryID: 1, UnitPrice: 50},
    {ID: 3, ProductName: "Water", CategoryID: 2, UnitPrice: 0.99},
}

var categories = []Category{
    {ID: 1, CategoryName: "Computers"},
    {ID: 2, CategoryName: "Beverages"},
}
