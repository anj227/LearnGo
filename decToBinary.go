package main

import (
    "fmt"
    "strconv"
    )

func concat_strings(str1 string, str2 string) string {
    if str1 == ""{
        return str2 
    } else if str2 == ""{
        return str1 
    } else {
        return str1 + str2
    }
}
func get_div_remainder(n int, base int) (int, int) {
    var remainder, div int 
    remainder = n % base 
    div = int(n / base)
    return div, remainder
}

func getBinaryFromDecimal(n int) string {
    var rr, nn int 
    var m string 
    m = ""
    nn = n
    for nn >= 2 {
        nn, rr = get_div_remainder(nn, 2)
        m = concat_strings(strconv.Itoa(rr), m) 
    }
    m = concat_strings(strconv.Itoa(nn), m)
    return m
    //return int(m) 
}
func main(){
    fmt.Println("Convert Decimal to Binary - give me a decimal")
    var input_dec int 
    fmt.Scan(&input_dec)
    fmt.Println("Provided decimal: ", input_dec)

    var out_bin string 
    out_bin = getBinaryFromDecimal(input_dec)
    fmt.Println("Binary # is: ", out_bin)
}