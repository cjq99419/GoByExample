## Go interface详解

根据oop的思想，接口定义了对象的行为，而具体行为的实现取决于对象。

Go中，接口可以理解为一组方法的集合，即一个接口定义一个或多个方法。当一个类型为接口中所有的方法提供定义的时候，称其实现了这个接口。

Go的核心：**根据类型可以执行的操作而不是其容纳的数据类型来设计抽象。**

```go
type Animal interface{
    Speak() string
}
```

定义一个名称为Animal的接口，并定义speak方法，不接受参数返回一个字符串

```go
type Dog struct {}
func (d Dog) Speak() string {
    return "汪！"
}

type Cat struct {}
func(c Cat) Speak() string {
    return "喵！"
}

func main() {
    animals := []Animal{Dog{}, Cat{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}

```

***

### interface{} 类型

interface{}类型为**空接口**，即没有实现方法的接口。在Go中没有implements关键字，因而**所有的类都可以理解为实现了interface{}接口**。这意味着interface{}作为函数参数，可以代表**任何值**。

```go
func say(words interface{}) {
    fmt.Println(words)
}
```

对于函数say，接受一个类型为interface{}的值，但words并不代表是任意类型的。words就是interface{}的类型，即当值传给say函数的时候，go运行时将执行强制类型转换，并转化为interface{}类型。

一个接口值由两个字来组成，其中一个字指向底层方法表，另一个指向实际数据







