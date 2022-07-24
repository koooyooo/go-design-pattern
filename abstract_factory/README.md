# Abstract Factory の実装方針

## 概要

**継承**が無い Go言語において、AbstractFactoryをどう実現すべきかという話です。

Go言語には **継承**が存在せず、類似した機能として **Embed**があります。 Embedは委譲により機能を埋め込みつつも、階層化した構造を透過的に見せかけることで、外部からはあたかも継承しているかの様に利用できるという機能で、これは一般的な用途ではとても強力です。

しかし、当然ながら通常の継承と異なる訳で、継承した関数の親子関係は存在するため、問題は発生します。

### Embedの問題

問題は、Template Methodとなる関数が他の関数をコントロールしようとする際、関数の親子関係に注意しなければならない点です。 具体的にはTemplateMethod（使う側）が親、それ以外の関数（使われる側）が子という構造にしないと、TemplateMethodが利用先の関数を呼び出せないのです。

仮に継承が使えるのであれば共通化したTemplateMethodは親側に定義され対象の関数を呼び出せます。ところが、Embedで実現した場合は共通化したTemplateMethodが子側（委譲先）として埋め込まれる構造になるため、上位の関数は見えなく（呼び出せなく）なってしまいます。

これを解決するためには2種類の実装があります。

### 1. Override先を 自身の関数の入れ替えで実現 (by delegation)

Embedによる逆転した親子関係で呼び出しが不能になるのが問題なら、呼び出し先を階層化させずに自身の関数にすれば良いという解決策です。関数の更新はサブクラスによるオーバーライドを用いずに、単なる関数の入れ替えで済ませてしまいます。

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

### 2. Override先を 内部のinterfaceに纏めて実現 (by extension)

TemplateMethodの呼び出し先の継承関係に問題があるなら、継承ではなく委譲を使って実現しよう。委譲先は入れ替え可能な様に interfaceを定義しておこう。という解決策です。
 
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