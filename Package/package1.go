//package name
package Package

import "fmt"

/* 
variables and function names starting with capital letter
are automatically exported and can be used from other packages
*/

var PackageVariable string = "this is variable from another package"

func PackageFunction() {
	fmt.Println("This is function from another package")	
}