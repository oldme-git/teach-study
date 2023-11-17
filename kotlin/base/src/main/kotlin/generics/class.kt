package generics

open class MyParam {}
class MyParam2: MyParam() {}

// 泛型约束
open class Abc1<V: MyParam> {}

// 泛型参数继承
class Abc2(): Abc1<MyParam2>() {}