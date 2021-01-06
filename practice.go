package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)


type Deadline time.Time

type Task struct {
	Title string
	Deadline *Deadline
}

func (d *Deadline) OverDue() bool {
	return d != nil && time.Time(*d).Before(time.Now())
}

func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}

func ExampleDeadline_OverDue() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))

	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())
}

func Example_taskTestAll() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", &d1}
	t2 := Task{"4h later", &d2}
	t3 := Task{"no due", nil}

	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
}

func fac(n int) int {
	if n<= 0 {
		return 1
	}
	return n* fac(n-1)
}


func facItr(n int) int{
	result := 1
	for n > 0{
		result *= n
		n--
	}
	return result
}

func main() {
	fmt.Println("Hello")

	var i = 10
	b := "little"
	b += "Gophers"
	fmt.Println("Hello", i, b)
	fmt.Println(fac(5))
	fmt.Println(facItr(5))

	var num int
	fmt.Sscanf("57", "%d", &num)


	fruits := []string{"사과", "바나나", "토마토"}

	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다 \n", fruit)
	}

	fruits = append(fruits, "포도")

	fmt.Println(fruits)

	ExampleDeadline_OverDue()
	Example_taskTestAll()
}


func Example_sliceCap() {
	nums := make([]int, 5)
	nums = nums[:3]
}

func count(s string, codeCount map[rune] int){
	for _, r := range s{
		codeCount[r]++
	}
}

//func TestCount(t *testing.T) {
//	codeCount := map[rune] int{}
//
//	count("가나다나", codeCount)
//	if !reflect.DeepEqual(
//		map[rune] int{'가':1, '나':2, '다':1},
//		codeCount, )
//}

func WriteTo(w io.Writer, lines []string ) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil{
			return err
		}
	}

	return nil
}


func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func ExampleWriterTo() {
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
	}

	fmt.Println(lines)
}

