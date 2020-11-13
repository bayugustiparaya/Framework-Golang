# Framework - Golang

## Endpoint

-   Order

    ```bash
    > localhost:8000/api/order
    ```

-   Customer

    ```bash
    > localhost:8000/api/customer
    ```

-   Product

    ```bash
    > localhost:8000/api/product
    ```

## Request

-   Order

    ```json
    {
        "orderID": "10248"
    }
    ```

-   Customer

    ```json
    {
        "customerID": "ALFKI"
    }
    ```

-   Product

    ```json
    {
        "productID": "2"
    }
    ```

## Response

-   Order

    ```json
    {
        "code": 100,
        "remark": "Success",
        "result": {
            "orderID": "10248",
            "customerID": "SANTG",
            "employeeID": "5",
            "orderDate": "1971-01-15",
            "ordersDetail": [
                {
                    "orderID": "10248",
                    "ProductID": "11",
                    "ProductName": "Queso Cabrales",
                    "UnitPrice": 14,
                    "Quantity": 12
                },
                {
                    "orderID": "10248",
                    "ProductID": "42",
                    "ProductName": "Singaporean Hokkien Fried Mee",
                    "UnitPrice": 9.8,
                    "Quantity": 10
                },
                {
                    "orderID": "10248",
                    "ProductID": "72",
                    "ProductName": "Mozzarella di Giovanni",
                    "UnitPrice": 34.8,
                    "Quantity": 5
                }
            ]
        }
    }
    ```

-   Customer

    ```json
    {
        "code": 100,
        "remark": "Success",
        "orderID": "0",
        "result": {
            "customerID": "ALFKI",
            "companyName": "Alfreds Futterkistes",
            "contactName": "Maria Anders",
            "contactTitle": "Sales Representative",
            "address": "Obere Str. 57",
            "city": "Berlin",
            "country": "Germany",
            "phone": "030-0074321",
            "postalCode": "12209"
        }
    }
    ```

-   Product

    ```json
    {
        "code": 100,
        "remark": "Success",
        "result": {
            "productID": "5",
            "productName": "Chef Anton's Gumbo Mix",
            "quantityPerUnit": "36 boxes",
            "unitPrice": "21.3500",
            "unitsInStock": "0",
            "unitsOnOrder": "0",
            "supplierDetail": {
                "supplierID": "2",
                "companyName": "New Orleans Cajun Delights",
                "contactName": "Shelley Burke",
                "contactTitle": "Order Administrator",
                "address": {
                    "street": "P.O. Box 78934",
                    "city": "New Orleans",
                    "postalCode": "70117",
                    "country": "USA"
                }
            },
            "categoryDetail": {
                "categoryID": "2",
                "categoryName": "Condiments",
                "description": "Sweet and savory sauces, relishes, spreads, and seasonings"
            }
        }
    }
    ```
