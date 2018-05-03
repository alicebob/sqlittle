// +build ci
// Tests against some database files found in the wild

package ci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alicebob/sqlittle"
)

func TestNorthwind(t *testing.T) {
	file := "../testdata/northwind.sqlite"

	test := func(table string, cols ...string) {
		sql := fmt.Sprintf("SELECT %s from '%s'", strings.Join(cols, ","), table)
		little := func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			if err := db.Select(table, cb, cols...); err != nil {
				t.Fatal(err)
			}
			return rows
		}
		CompareSelect(t, file, sql, little)
	}
	test("Category",
		"Id",
		"CategoryName",
		"Description",
	)
	test("Customer",
		"Id",
		"CompanyName",
		"ContactName",
		"ContactTitle",
		"Address",
		"City",
		"Region",
		"PostalCode",
		"Country",
		"Phone",
		"Fax",
	)
	test("CustomerCustomerDemo",
		"Id",
		"CustomerTypeId",
	)
	test("CustomerDemographic",
		"Id",
		"CustomerDesc",
	)
	test("Employee",
		"Id",
		"LastName",
		"FirstName",
		"Title",
		"TitleOfCourtesy",
		"BirthDate",
		"HireDate",
		"Address",
		"City",
		"Region",
		"PostalCode",
		"Country",
		"HomePhone",
		"Extension",
		"Photo",
		"Notes",
		"ReportsTo",
		"PhotoPath",
	)
	test("EmployeeTerritory",
		"Id",
		"EmployeeId",
		"TerritoryId",
	)
	test("Order",
		"Id",
		"CustomerId",
		"EmployeeId",
		"OrderDate",
		"RequiredDate",
		"ShippedDate",
		"ShipVia",
		// "Freight" // float
		"ShipName",
		"ShipAddress",
		"ShipCity",
		"ShipRegion",
		"ShipPostalCode",
		"ShipCountry",
	)
	test("OrderDetail",
		"Id",
		"OrderId",
		"ProductId",
		// "UnitPrice" // float
		"Quantity",
		// "Discount" // float
	)
	test("Product",
		"Id",
		"ProductName",
		"SupplierId",
		"CategoryId",
		"QuantityPerUnit",
		// "UnitPrice" // float
		"UnitsInStock",
		"UnitsOnOrder",
		"ReorderLevel",
		"Discontinued",
	)
	test("Region",
		"Id",
		"RegionDescription",
	)
	test("Shipper",
		"Id",
		"CompanyName",
		"Phone",
	)
	test("Supplier",
		"Id",
		"CompanyName",
		"ContactName",
		"ContactTitle",
		"Address",
		"City",
		"Region",
		"PostalCode",
		"Country",
		"Phone",
		"Fax",
		"HomePage",
	)
	test("Territory",
		"Id",
		"TerritoryDescription",
		"RegionId",
	)
}
