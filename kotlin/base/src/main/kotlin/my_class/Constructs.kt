package my_class

open class Constructs(name: String) {
    init {
        println("主构造函数$name")
    }

    constructor(name: String, name2: String) : this(name) {
        println("次构造函数$name|$name2")
    }
}

class SubConstructs(name: String): Constructs(name) {
    init {
        println("子类构造函数$name")
    }
}

fun main() {
    // 会运行两个构造函数
    Constructs("name", "name2")
    println("--------------")
    // 只会运行主构造函数
    Constructs("name")
    println("--------------")
    // 子类
    SubConstructs("subName")
}
