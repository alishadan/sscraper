package FaToInt
import(
	"fmt"
	"strings"
	"strconv"

)

//how to use:
//example
// string1:="۹۹۹۹,۰۰۰۰"
// number:=FaToint(string1)


func FaToint(string1 string)int {


	var rawNumber []string
	type pNumber struct{
		sefr	rune
		yek rune
		do	rune
		se  rune
		chahar rune
		panj	rune
		shesh	rune
		haft	rune
		hasht	rune
		noh		rune
		

	}
	number:=pNumber{'۰','۱','۲','۳','۴','۵','۶','۷','۸','۹'}
	var temNumber string

	for _,value:=range(string1){
		if value==number.sefr{
			temNumber="0"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.yek{
			temNumber="1"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.do{
			temNumber="2"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.se{
			temNumber="3"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.chahar{
			temNumber="4"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.panj{
			temNumber="5"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.shesh{
			temNumber="6"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.haft{
			temNumber="7"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.hasht{
			temNumber="8"
			rawNumber=append(rawNumber,temNumber)
		}
		if value==number.noh{
			temNumber="9"
			rawNumber=append(rawNumber,temNumber)
		}

	}
	finalNumber := strings.Join(rawNumber, "")
	n,err:=strconv.Atoi(finalNumber)
	if err!=nil{
		fmt.Printf("error in strconv funciton \n")
	}
	return n
	
}