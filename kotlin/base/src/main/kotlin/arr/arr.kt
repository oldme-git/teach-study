package arr

fun main() {
    val m: Map<String, Map<String, Int>> = mapOf(
        "key1" to mapOf(
            "key11" to 1,
            "key12" to 2,
        ),
        "key2" to mapOf(
            "key21" to 2
        )
    )

    for ((k, v) in m) {
        println(k)
        for ((k1, v1) in v) {
            println(k1 + v1)
        }
    }
}