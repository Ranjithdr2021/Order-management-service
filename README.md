# Order-management-service
Procedure to run Application.
************************************************************************************************************************************************
add-orders – URL to post the details of orders in the request body : [POST] -- http://localhost:8080/add-order
************************************************************************************************************************************************
i)

{
"id" : "abcdef-123458",
"status" : "completed",
"items" : [{
	"id" : "123458",
	"description" : "product-descritn",
	"price" : 12.40,
	"quantity" : 1
},
{
	"id" : "123459",
	"description" : "product-description-1",
	"price" : 20.90,
	"quantity" : 2
}],
"total" : 12.40,
"currencyunit" : "usd"
}

**********************************************************************************************************************************
EXPECTED RESPONSE: SUCCESS
**********************************************************************************************************************************
To check the details of orders, please redirect to URL [GET] –- http://localhost:8080/orders
**********************************************************************************************************************************
EXPECTED RESPONSE: All the above posted orders will be retrieved as response
**********************************************************************************************************************************
Change status Service TO change the order status please redirect to URL [PUT] -- http://localhost:8080/update-status/:order-id
you can use this example url --> http://localhost:8080/update-status/abcdef-123458
**********************************************************************************************************************************
Enter the following details in the request body

{
"status" : "Done"
}
**********************************************************************************************************************************
EXPECTED RESPONSE: Status updated Successfully...
**********************************************************************************************************************************
Get order based on order status please redirect to URL [GET] -- http://localhost:8080/GetOrdersonstatus/:orders-status
you can use this example url --> http://localhost:8080/GetOrdersonstatus/pending_invoice
**********************************************************************************************************************************
EXPECTED RESPONSE: ALl orders will be retrieved but orders with the status in param will appear on top.
**********************************************************************************************************************************
Get order based on order id please redirect to URL [GET] -- http://localhost:8080/GetOrdersById/:order-id
you can use this example url --> http://localhost:8080/GetOrdersById/abcdef-123456
**********************************************************************************************************************************
EXPECTED RESPONSE: order with the requested order id will be retrieved in response.
**********************************************************************************************************************************
SUMMARY:
**********************************************************************************************************************************
URL used in the application :

To post the details of order in request body please select POST request and enter

POST-- http://localhost:8080/add-order

NOTE: Initially to post the order details first check with the mysql username and password.

To retrieve details of orders please select GET request and enter

GET -- http://localhost:8080/orders

To update the order status please select PUT request and enter

PUT -- http://localhost:8080/update-status/Order-id

To get the orders with specific status please select GET request and enter

GET -- http://localhost:8080/GetOrdersonstatus/order-status

To get the orders with specific order id please select GET request and enter

GET -- http://localhost:8080/GetOrdersById/order-id

**********************************************************************************************************************************
Note := Testing is only done for the core logic part in this project.
