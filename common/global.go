package common

// Struct API
type Message struct {
	Code   int         `json:"code"`
	Remark string      `json:"remark"`
	Result interface{} `json:"result"`
}

// Start order struct
type Order struct {
	OrderID    string        `json:"orderID"`
	CustomerID string        `json:"customerID,omitempty"`
	EmployeeID string        `json:"employeeID,omitempty"`
	OrderDate  string        `json:"orderDate,omitempty"`
	OrderDet   []OrderDetail `json:"ordersDetail,omitempty"`
}

type OrderDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

// End order struct

// Start customer struct
type Customer struct {
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

// End customer struct

// Start product struct
type Product struct {
	ProductID       string         `json:"productID"`
	ProductName     string         `json:"productName,omitempty"`
	QuantityPerUnit string         `json:"quantityPerUnit,omitempty"`
	UnitPrice       string         `json:"unitPrice,omitempty"`
	UnitsInStock    string         `json:"unitsInStock,omitempty"`
	UnitsOnOrder    string         `json:"unitsOnOrder,omitempty"`
	SupplierDet     SupplierDetail `json:"supplierDetail,omitempty"`
	CategoryDet     CategoryDetail `json:"categoryDetail,omitempty"`
}

type SupplierDetail struct {
	SupplierID   string `json:"supplierID"`
	CompanyName  string `json:"companyName"`
	ContactName  string `json:"contactName"`
	ContactTitle string `json:"contactTitle"`
	Address      struct {
		Street     string `json:"street"`
		City       string `json:"city"`
		PostalCode string `json:"postalCode"`
		Country    string `json:"country"`
	} `json:"address"`
}

type CategoryDetail struct {
	CategoryID   string `json:"categoryID"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}

// End product struct

// Start faspay struct

type RequestFaspay struct {
	Request       string `json:"request"`
	TransactionID string `json:"trx_id"`
	MerchantID    string `json:"merchant_id"`
	BillNumber    string `json:"bill_no"`
	Signature     string `json:"signature"`
}

type ResponseFaspay struct {
	Response          string `json:"response"`
	TransactionID     string `json:"trx_id"`
	MerchantID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNumber        string `json:"bill_not"`
	PaymentReff       string `json:"payment_reff"`
	PaymentDate       string `json:"payment_date"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	ResponseCode      string `json:"response_code"`
	ResponseDesc      string `json:"response_desc"`
}

// End faspay struct

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

//End Struct API
