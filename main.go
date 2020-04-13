// Programmer name: Vincent G.
// Date completed:  4/13/2020
// Description: Section 1 3.6.1

package main

import "fmt"

// HoursMinutesSeconds computes the number of hours, minutes,
// and seconds given a number of seconds.
// EXAMPLE: 3661 seconds is 1 hours, 1 minutes, 1 seconds

func HoursMinutesSeconds(seconds int) int {
 
 var h int 
 var m int
 var s int

 h = (seconds/3600)
 m = (seconds - 3600*h)/60
 s = ( seconds -(3600*h)-(m*60))



 
 fmt.Println(seconds, "is", h,"hours", m,"minutes",s,"seconds")
 return 0
}

func main() {
//call your function
HoursMinutesSeconds(7890)
}