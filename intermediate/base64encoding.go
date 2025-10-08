package intermediate

import (
	"encoding/base64"
	"fmt"
)

// 1.Why Base64?
// -Text Transmission
// -Storage
// -URLs and Data URLs
// Base64编码的主要作用是将二进制数据转换为文本格式，
// 这样可以方便地通过文本协议（如电子邮件和HTTP头）进行传输。这使得在不同的系统和应用程序之间处理数据变得更加简单和安全。
// 2.Why is Encoding lmportant?
// -Data Storage
// -Data Transmission
// -Data Interoperability
// 3.Common Examples of Encoding
// -Text Encoding
// --ASCII
// --UTF-8
// --UTF-16
// -Data Encoding
// --Base64
// --URL Encoding
// -File Encoding
// --Binary Encoding
// --Text Encoding

// 3. Base64 的核心理念是 “将 3 个字节（24 位）的二进制数据，重新编码为 4 个 6 位的单元”。

// 每个 6 位的单元（取值范围 0-63）对应一个特定的可打印字符。这就是“64”的由来（2的6次方=64）。
// 编码步骤：
// 分组： 将原始二进制数据按每 3 个字节（24 位）为一组进行划分。
// 重划分： 将 24 位的数据重新划分为 4 个 6 位的数据块。
// 映射： 将每个 6 位的数据块（值在 0 到 63 之间）根据 Base64 索引表，映射到一个对应的可打印字符。

// 4. 常见应用场景
// 电子邮件（MIME）： 这是 Base64 最早也是最经典的应用，用于在邮件中传输附件。
// 在网页中嵌入小文件： 通过 Data URLs 的方式，可以将图片、字体等小文件直接以 Base64 形式嵌入到 HTML 或 CSS 中，减少 HTTP 请求。
// 例如： data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA...
// 存储二进制数据在文本配置中： 如 XML、JSON 配置文件需要存储少量二进制信息时。
// URL 和文件名： 有时会对 Base64 进行变种（如 URL-safe Base64，将 + 和 / 替换为 - 和 _ ），用于在 URL 中安全地传递二进制参数。


func main() {

	data:=[]byte("He~lo,Base64 Encoding")

	fmt.Println(data)
	//Encode Base64
	encoded:=base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded) //SGV+bG8sQmFzZTY0IEVuY29kaW5n

	//decode from base64
	decoded,err:=base64.StdEncoding.DecodeString(encoded)
  if err!=nil {
		fmt.Println("Error in decoding",err)
		return
	}
	fmt.Println("Decoded:",decoded) //[72 101 126 108 111 44 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
  fmt.Println("Decoded:",string(decoded)) // He~lo,Base64 Encoding


	// URL safe, avoid '/' and '+'
	urlSafeEncoded:=base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe encoded:",urlSafeEncoded) //SGV-bG8sQmFzZTY0IEVuY29kaW5n
}
