// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

import (
	"github.com/jaypipes/pcidb"
)

func pciFillInfo(info *PCIInfo) error {
	db, err := pcidb.New()
	if err != nil {
		return err
	}
	info.Classes = db.Classes
	info.Vendors = db.Vendors
	info.Products = db.Products
	return nil
}
