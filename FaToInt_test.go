package FaToInt
import (
	"testing"
	"fmt"


)
func Test_test(t *testing.T){
	string1:="۹۹۹۹,۰۰۰۰"
	number:=FaToint(string1)
	fmt.Printf("%d \n",number)
}