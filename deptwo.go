package deptwo

import depone "github.com/jimmyjames85/depOne"

func Two() string { return "This is Two: One is: " + depone.One() }
