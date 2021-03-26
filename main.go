package main

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func add(a int, b int) int {
	return a + b
}
func creatNewVar() {
	var a = 1
	var b = 2
	a, b = b, a
	fmt.Println("a = ", a, " , b = ", b)
	fmt.Println("a + b = ", add(a, b))

	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(r.Int())
}

func getNewRandNum() int {
	rand.New(rand.NewSource(time.Now().Unix()))
	return rand.Int()
}

func tryPointer() {
	ptr := new(int)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
func _test(v int) {
	if v == 0 {
		return
	}
	creatNewVar()
	getNewRandNum()
	tryPointer()
}
func test(v int) {
	if v == 0 {
		return
	}
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("%T\n", arr)
	fmt.Printf("%d%T\n", arr[0:3], arr[0:3])
	anotherArr := make([]int, 3)
	fmt.Println(anotherArr)
	fmt.Printf("len = %d, cap = %d\n", len(anotherArr), cap(anotherArr))
	for i := 1; i < 10; i++ {
		anotherArr = append(anotherArr, i)
	}
	fmt.Println(anotherArr)
	fmt.Printf("len = %d, cap = %d\n", len(anotherArr), cap(anotherArr))
}

func testMap(v int) {
	if v == 0 {
		return
	}
	scores := make(map[string]int)
	scores["a"] = 1
	scores["aa"] = 11
	fmt.Println(scores)
	delete(scores, "aa")
	fmt.Println(scores)
	if math, ok := scores["a"]; ok {
		fmt.Printf("math 的值是: %d\n", math)
		fmt.Printf("%T\n", ok)
	} else {
		fmt.Println("math 不存在")
	}

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	sum := 0
	pre := 1
	return func() int {
		temp := sum
		sum += pre
		pre = temp
		return pre
	}
}

func testClosure(v int) {
	if v == 0 {
		return
	}
	p1, p2 := adder(), adder()
	for i := 1; i < 10; i++ {
		fmt.Println(p1(i),
			p2(i*2))
	}
	temp := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", temp())
	}

}

// interface experiment

type avActress struct {
	name string
	age  int
}

type work interface {
	Fuck()
	Cry()
}

func (p avActress) Fuck() {
	fmt.Println("come on baby!")
}

func (p avActress) Cry() {
	fmt.Println("no! i'm coming!")
}

type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

func (phone Phone) settleAccount() int {
	return phone.price * phone.quantity
}

func (phone Phone) orderInfo() string {
	return "you are going to buy " + strconv.Itoa(phone.quantity) +
		phone.name + " with " + strconv.Itoa(phone.settleAccount())
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

func (gift FreeGift) settleAccount() int {
	return 0
}

func (gift FreeGift) orderInfo() string {
	return "you are going to buy " + strconv.Itoa(gift.quantity) +
		gift.name + " with " + strconv.Itoa(gift.settleAccount())
}

func calculateAllPrice(goods []Good) int {
	allPrice := 0
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

func testInterface(v int) {
	if v == 0 {
		return
	}
	sensai := avActress{
		name: "sensai",
		age:  31,
	}
	i := work(sensai)
	i.Fuck()
	i.Cry()

	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    10000,
	}
	earPods := FreeGift{
		name:     "airPods",
		quantity: 1,
		price:    1000,
	}
	goods := []Good{iPhone, earPods}
	cost := calculateAllPrice(goods)
	fmt.Println(cost)

	fmt.Println("try type-assertion")
	item, ok := goods[0].(Phone)
	fmt.Println(item, ok)

	fmt.Println("try type-switch")
	switch goods[0].(type) {
	case Phone:
		fmt.Println("yep, you're right")
	case FreeGift:
		fmt.Println("nope, you're wrong")
	default:
		fmt.Println("actually, i don't know the answer")
	}
	boy := object{name: "abc"}
	if _, err := boy.testfun(); err != nil {
		fmt.Println(err)
	}
}

type object struct {
	name string
}

func (o *object) testfun() (string, error) {
	if o.name == "abc" {
		return "fuck", o
	}
	return o.name, nil
}
func (o *object) Error() string {
	return fmt.Sprintf("oh no!!! %s", o.name)
}

// interface experiment ends

func testIO(v int) {
	if v == 0 {
		return
	}
	r := strings.NewReader("fuck me! baby!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	_test(0)
	test(0)
	testMap(0)
	testClosure(0)
	testInterface(0)
	testIO(0)
	tryRoutine(0)
	tryChannel(0)
	tryMutex(0)
	tryReflect(0)
	tryWait(0)
}
