package my_annotation

// 定义一个自定义注解
annotation class MyAnnotation(val message: String)

// 在类中使用自定义注解
@MyAnnotation("This is a custom annotation")
class MyClass

fun main() {
    // 获取类的注解信息
    val annotation = MyClass::class.java.getAnnotation(MyAnnotation::class.java)

    // 判断注解是否存在
    if (annotation != null) {
        // 打印注解中的信息
        println("Annotation message: ${annotation.message}")
    } else {
        println("Annotation not found")
    }
}