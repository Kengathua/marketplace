> Specify item categories

-> Shop1

	-> Phones and Accessories
		-> ItemType
			-> TypeName 	-> Phone
			-> TypeCode	-> 001
			-> Brands 	-> (Samsung, LG, Nokia, Apple)
		-> Brand
			-> BrandName	-> Samsung
			-> Brandcode	-> 001
		-> Model
			-> ModelNumber	-> Galaxy S3
			-> ModelCode	-> 001
		-> Item
			-> Item Name	-> Samsung Galaxy S3 Phone
			-> Item code	-> 001
		-> ItemUnit
			-> Item 	-> Samsung Galaxy S3
			-> Units	-> Pieces
		-> ItemImage
			-> Item		-> Samsung Galaxy S3
			-> Image	-> img001.png
			-> IsHeroImage	-> True

		-> CatalogItem 
			-> Item 		-> Samsung Galaxy S3
			-> Quantity		-> 5
			-> Marked Price		-> 90,000
			-> Discount Amount	-> 3,000
			-> Selling Price	-> 87,000

		-> Cart
			-> Customer		-> Mr Mark Makau
			-> CartNumber		-> C-001
			-> TotalPrice		-> 174,000
			-> OrderGuid		-> UUID

		-> CartItem
			-> Cart			-> Mr Mark Makau C-001
			-> CatalogItem		-> Samsung Galaxy S3
			-> SellingPrice		-> 87,000
			-> Quantity		-> 2
			-> TotalPrice		-> 174,000

		-> Order
			-> Customer		-> Mr Mark Makau
			-> Cart			-> Mr Mark Makau C-001
			-> OrderNumber		-> O-001
			-> IsCleared
			-> Status (Pending, Confirmed, Awaiting Delivery, In Transit, Delivered)

		-> OrderItem
			-> Order		-> Mr Mark Makau O-001
			-> CartItem		-> Samsung Galaxy S3
			-> Quantity		-> 2
			-> SellingPrice		-> 87,000
			-> TotalPrice		-> 174,000
			-> IsCleared
			-> Status (Pending, Confirmed, Awaiting Collection, Collected, Awaiting Delivery, In Transit, Delivered)

		-> OrderStatusLog
			-> Order		-> Mr Mark Makau O-001
			-> StatusFrom		-> Pending
			-> StatusTo		-> Confirmed
			-> TransitionTime	-> 06/02/2023 17:45:55

		-> Payments
			-> Customer
			-> Amount
			-> Payment Method(CASH, MPESA TILL, MPESA PAYBILL)
			-> Account Number
			-> IsConfirmed
			-> IsProcessed

		-> Transaction
			-> Payment
			-> Amount

		-> OrderTransaction
			-> Order
			-> Transaction

		-> Sales
			-> Order
			-> TotalAmount

		-> SaleItem
			-> CatalogItem
			-> SellingPrice
			-> Quantity
			-> TotalPrice
			-> Profit

