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

-   Faspay

    ```bash
    > localhost:8000/api/faspay
    ```

-   Trip

    ```bash
    > localhost:8000/api/trip
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

-   Faspay

    ```json
    {
        "request": "Pengecekan Status Pembayaran",
        "trx_id": "111111",
        "merchant_id": "1111115",
        "bill_no": "9988776655",
        "signature": "b03951a6051dfa90894eea48ccb964619e8b3474"
    }
    ```

-   Trip

    ```json
    {
        "depature_date_1": "2020-10-10",
        "depature_date_2": "2020-10-13",
        "provinsi": 31
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

-   Faspay

    ```json
    {
        "response": "Pengecekan Status Pembayaran",
        "trx_id": "111111",
        "merchant_id": "1111115",
        "merchant": "Zera Store",
        "bill_not": "123",
        "payment_reff": "123",
        "payment_date": "2020-11-16 07:17:28",
        "payment_status_code": "2",
        "payment_status_desc": "Payment Sukses",
        "response_code": "200",
        "response_desc": "Sukses ambil data"
    }
    ```

-   Trip
    ```json
    {
        "status": "200",
        "message": "Success",
        "data": [
            {
                "TripID": "288",
                "TravelID": "3",
                "TravelName": "PT. Baitullah Barakah Abadi",
                "Description": "Amazing Turkey",
                "Rating": "4.3",
                "Provinsi": "DKI Jakarta",
                "CityName": "Kota Jakarta Barat",
                "LicenseNumber": "748/2017",
                "DepartureDate": "11/02/2020",
                "ReturnDate": "19/04/2020",
                "Duration": "9",
                "OriginCity": "Kota Tangerang",
                "AirportName": "Bandar Udara Internasional Soekarno-Hatta",
                "Origin": "CGK",
                "Destination": "IST",
                "Transit": "No",
                "DetailTransit": "",
                "HotelName": "Tryp By Wyndham Istanbul Sancaktepe",
                "HotelRating": "",
                "Currency": "IDR",
                "Price": "19500000",
                "PromoCode": "",
                "PromoDescription": "",
                "AirlineName": "Turkish Airlines",
                "Goods": "",
                "TermCondition": "",
                "Lat": "",
                "Long": "",
                "DoubleType": "19500000",
                "TripleType": "0",
                "QuadType": "0",
                "Logo": "https://www.moslemtrip.com/traveladmin/apps/photo/3.png"
            }
        ]
    }
    ```
