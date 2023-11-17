package generics

// 泛型函数
fun <T> intToString(item: T): String {
    return item.toString()
}

// 泛型约束,上界约束
fun <T: Int> intToString2(item: T): String {
    return item.toString()
}

// 泛型约束,多重约束
fun <T> intToString3(item: T): String where T: String {
    return item.toString()
}

fun main() {
    val s: String = intToString(1)
    println(s::class.simpleName)
}
