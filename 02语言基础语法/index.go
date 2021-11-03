/*  定义包名*/
package main

/* 导入外包*/
import "fmt"

/*程序开始执行*/
func main() {

	/* 行分隔符 */

	/* 在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ;
	结尾，因为这些工作都将由 Go 编译器自动完成。 */
	/* 如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。 */

	/*Go 语言的字符串可以通过 + 实现：*/
	fmt.Println("Google" + " GoLang")

	// %d 表示整型数字，%s 表示字符串
	// Go 语言中使用 fmt.Sprintf 格式化字符串并赋值给新串：
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	fmt.Println(url)
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}
