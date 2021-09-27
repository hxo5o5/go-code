/*练习：rot13Reader
有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。
例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。
编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。
rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。
*/
package main

import (
	"io"
	"os"
	"strings"
)

//rot13是一个简单的字母替换密码，用字母后的第13个字母替换当前字母
type rot13Reader struct {
	r io.Reader
}

//解密方法
//分析，字母一共26个，用字母后的第13个字母替换，遇到Z后循环从A在开始
//所以前13个字母往后加13 后13个字母减13
//基本类型byte型代表了ASCII码的一个字符
func rot13(out byte) byte {
	switch {
	case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
		out += 13
	case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
		out -= 13
	}
	return out
}

//重写Read方法
func (fz rot13Reader) Read(b []byte) (int, error) {
	n, e := fz.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
