package common

//Struct API
// Order struct (Model) ...

type Message struct {
	Code    int         `json:"code"`
	Remark  string      `json:"remark"`
	OrderID string      `json:"orderID"`
	Orders  *Orders     `json:"orders,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type Orders struct {
	OrderID    string         `json:"orderID"`
	CustomerID string         `json:"customerID"`
	EmployeeID string         `json:"employeeID"`
	OrderDate  string         `json:"orderDate"`
	OrdersDet  []OrdersDetail `json:"ordersDetail"`
}

type OrdersDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

type Customers struct {
	CustomerID   string `json:"customerID"`
	CompanyName  string `json:"companyName,omitempty"`
	ContactName  string `json:"contactName,omitempty"`
	ContactTitle string `json:"contactTitle,omitempty"`
	Address      string `json:"address,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
}

//End Struct API
