package main

/**
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
在 S 上反复执行重复项删除操作，直到无法继续删除。
在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

示例：
输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。
之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。

提示：
1 <= S.length <= 20000
S 仅由小写英文字母组成。
*/
/**
思路：栈
    数字栈 和 字符栈
*/
func decodeString(s string) string {
	length := len(s)
	var numStack []int
	var strStack []byte
	for i := 0; i < length; i++ {
		if isNum(s[i]) {
			// 维护数字栈
			num := int(s[i] - '0')
			for isNum(s[i+1]) {
				num = num*10 + int(s[i+1]-'0')
				i++
			}
			numStack = append(numStack, num)
		} else {
			// 维护字符串栈
			if s[i] == ']' {
				var tempStr []byte
				for strStack[len(strStack)-1] != '[' {
					last := len(strStack) - 1
					tempStr = append([]byte{strStack[last]}, tempStr...)
					strStack = strStack[:last]
				}
				// 将 [ 出栈
				strStack = strStack[:len(strStack)-1]
				// 根据数字进行字符串复制和拼接
				numLast := len(numStack) - 1
				for i := 0; numStack != nil && i < numStack[numLast]; i++ {
					strStack = append(strStack, tempStr...)
				}
				numStack = numStack[:numLast]
			} else {
				strStack = append(strStack, s[i])
			}
		}
	}

	return string(strStack)
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {

}
