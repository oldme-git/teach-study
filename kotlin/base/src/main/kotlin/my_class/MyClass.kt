package my_class

class MyClass (id: Int) {
    // 属性
    private val name: String = "oldme"

    // 初始化块
    init {
        println(id)
    }

    // 函数
    fun getName(): String {
        return name
    }
}

fun main() {
    // 实例化
    val myClass = MyClass(1)
    // 调用函数
    val name = myClass.getName()
    println(name)
}