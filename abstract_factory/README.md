## 概要

Golangには **継承**が存在せず、類似した機能として **Embed**があります。 Embedは Delegationを継承ライクに利用できる機能で一般的な用途では強力です。

しかし、通常の継承と異なり関数の親子関係は存在するため、Template Methodとして他の関数をコントロールしようとすると、関数の親子関係を意識しなければなりません。

具体的にはTemplateMethod（使う側）が親、それ以外の関数（使われる側）が子という形にしないと、TemplateMethodが対象を呼び出せないのです。

運の悪いことに、Templateにより共通化した汎用フローを、具象的な構造体に埋め込むという通常のデザインを採用すると、TemplateMethodが子側になるので、対象を呼び出せなくなる現象が発生します。

これを解決するためには2種類の実装があります。

### 1. Override先を 自身の関数の入れ替えで実現 (by delegation)
TemplateMethodの実装疎結合化を Overrideではなく関数の入れ替えで実現

```golang
package xxx
import "fmt"

type Template struct {
	a func()
	b func()
	c func()
}

func (t *Template) TemplateMethod() {
	t.a()
	t.b()
	t.c()
}

// 個別実装部分
func NewTemplateA() *Template {
	return &Template {
        a: func() {
            fmt.Println("a")
        },
        b: func() {
            fmt.Println("b")
        },
        c: func() {
            fmt.Println("c")
        },
    }
}
```

### 2. Override先を 内部のinterfaceに纏めて実現 (by_extension)
TemplateMethodの呼び出し先を 自身の関数ではなく移譲先Interfaceの関数で実現
 
```golang
package yyy
import "fmt"

type Template interface {
	TemplateMethod()
}

type templateImpl struct {
	target
}

func (t *templateImpl) TemplateMethod() {
    t.a()
    t.b()
    t.c()
}

type target interface {
    a()
    b()
    c()
}

// 個別実装部分
type myTemplateA struct {}
func (t *myTemplateA) a() {
    fmt.Println("a")
}
func (t *myTemplateA) b() {
    fmt.Println("b")
}
func (t *myTemplateA) c() {
    fmt.Println("c")
}

func NewTemplateA() Template {
    return &templateImpl{&myTemplateA{}}	
}
```