//这是一个帮助理解go 的oop 的相关程序
//相关概念 基类(父类):baseBrid 派生类(子类):bigBrid
//go的封装 大写开头的包外可见 小写开头的包内可见
//go的继承 子类的struct 嵌入 父类的struct ,但实际上这并不是真正的继承 而是匿名组合,本质还是组合 go的组合(所有的struct)是静态绑定
//go的多态 由接口来实现 是一种duck typig的类型,如果一个struct定义了一个接口的所有方法,那么该类型就隐式的实现了该接口.   接口是一组方法的集合
package main

import "fmt"

type Animal interface {
	show()
	eat()
	//Add()
}

type baseBird struct {
	age int
}

//可以看到baseBird 实现了Animal接口的方法,那么我们可以说baseBird 是一种动物
func (this *baseBird) show() {
	fmt.Println("bird age is =", this.age)
}
func (this *baseBird) eat() {
	fmt.Println("bird eat bugs")
}
func (this *baseBird) Add() {
	fmt.Println("baseBird org age =", this.age)
	this.age = this.age + 1
	fmt.Println("baseBird after add age =", this.age)
}

func (this *baseBird) Cil() {
	//func Cil(this Animal) {
	this.Add()
}

type bigBird struct {
	baseBird
}

func (this *bigBird) Add() {
	fmt.Println("bigBird org age =", this.age)
	this.age = this.age + 2
	fmt.Println("bigBird after add age =", this.age)
}
func main() {
	var b1 baseBird
	var b2 bigBird

	b1 = baseBird{age: 1}
	b1.show()
	b1.eat()
	b1.Add()
	b1.Cil()
	//Cil(&b1)

	b2 = bigBird{baseBird{1}}
	b2.show()
	b2.eat()
	b2.Add()
	b2.Cil()
	//Cil(&b2)
}

/*总结:可以看到  1.子类可以调用父类的cil方法
			2.子类也重写了父类的add方法
  但是你可以发现 子类调用父类的cil方法时,调用的还是父类的add方法
  如果我们想要子类调用cil方法时,调用的还是子类的add方法 该怎么办呢?
  答案还得是 interface
  我们在Animal 接口中添加Add 方法 重写 Cli 方法 将以上代码注释的地方打开.之前的注释掉 运行结果就是我们想要的结果了
  所以GOLang的oop 和java 还有c++的并不相同 不能惯用之前的思想
*/
