package constant

const (
	///? ====================== Master Data Employee ======================
	EMPLOYEE_INSERT = "INSERT INTO employees(id,name,phone_number,address)VALUES($1, $2, $3, $4)"
	EMPLOYEE_LIST   = "SELECT * FROM employees"
	EMPLOYEE_GET    = "SELECT * FROM employees where id = $1"
	EMPLOYEE_UPDATE = "UPDATE employees SET name = $2, phone_number = $3, address = $4 WHERE id = $1"
	EMPLOYEE_DELETE = "DELETE FROM employees WHERE id = $1"
	EMPLOYEE_COUNT  = "SELECT COUNT(*) FROM employees"

	///? ====================== Master Data Product ======================
	PRODUCT_INSERT = "INSERT INTO products (id,name,price,uom) VALUES($1, $2, $3, $4)"
	PRODUCT_LIST   = "SELECT * FROM products"
	PRODUCT_GET    = "SELECT * FROM products where id = $1"
	PRODUCT_UPDATE = "UPDATE products SET name = $2, price = $3, uom = $4 WHERE id = $1"
	PRODUCT_DELETE = "DELETE FROM products WHERE id = $1"
	PRODUCT_COUNT  = "SELECT COUNT(*) FROM products"

	///? ====================== Master Data Customer ======================
	CUSTOMER_INSERT = "INSERT INTO customers(id,name,phone_number,address)VALUES($1, $2, $3, $4)"
	CUSTOMER_LIST   = "SELECT * FROM customers"
	CUSTOMER_GET    = "SELECT * FROM customers where id = $1"
	CUSTOMER_UPDATE = "UPDATE customers SET name = $2, phone_number = $3, address = $4 WHERE id = $1"
	CUSTOMER_DELETE = "DELETE FROM customers WHERE id = $1"
	CUSTOMER_COUNT  = "SELECT COUNT(*) FROM customers"
	//...
	///? ====================== Data Bill ======================
	BILL_CREATE        = "INSERT INTO bills (id,bill_date,entry_date,employee_id,customer_id) VALUES ($1,$2,$3,$4,$5)"
	BILL_DETAIL_CREATE = "INSERT INTO bill_details (id,bill_id,product_id,product_price,qty,finish_date) VALUES ($1,$2,$3,$4,$5,$6)"
	BILL_LIST          = "SELECT b.id,b.bill_date,b.entry_date,c.id,c.name,c.phone_number,c.address,e.id,e.name,e.phone_number,e.address,d.id,d.bill_id,p.id,p.name,p.price,p.uom,d.qty,d.finish_date FROM bill_details AS d JOIN billS AS b ON d.bill_id = b.id JOIN customers AS c ON b.customer_id = c.id JOIN employees AS e on b.employee_id = e.id JOIN products AS p ON d.product_id = p.id"
	BILL_COUNT         = "SELECT COUNT(*) FROM bills"
	///? ====================== Data Bill Details ======================
	BILL_GET        = "SELECT * FROM bills WHERE id= $1"
	BILL_DETAIL_GET = "SELECT * FROM bill_details WHERE bill_id = $1"
	///? ====================== Master User ======================
	USER_LIST         = "SELECT * FROM users"
	USER_CREATE       = "INSERT INTO users (id,username,password,is_active) VALUES ($1,$2,$3,$4)"
	USER_USERNAME_GET = "SELECT * FROM users WHERE username = $1 AND is_active = $2"
)
