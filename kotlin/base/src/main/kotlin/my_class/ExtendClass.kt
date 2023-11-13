package my_class

open class Parent(private val parentName: String) {
    fun getName() {
        println(parentName)
    }

    open fun getNameOverride() {
        println(parentName)
    }

    fun superDemo() {
        println("super")
    }
}

class SubClass: Parent("pName") {
    override fun getNameOverride() {
        println("sName")
        // 调用超类实现
        super.superDemo()
    }
}

fun main() {
    val subClass = SubClass()
    subClass.getName()
    subClass.getNameOverride()
}