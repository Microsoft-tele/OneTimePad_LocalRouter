package CharsetUtil

import (
	"github.com/axgle/mahonia"
	"strings"
)

func CoverGBKToUTF8(src string) string {
	return mahonia.NewDecoder("gbk").ConvertString(src)
}

/**
 * 替换乱码
 */
func replaceNullHtml(src string) string {
	temp := strings.Replace(src, "聽", "", -1)
	return temp
}

/**
 * gbk转utf8
 */
func coverString(src string) string {
	return replaceNullHtml(CoverGBKToUTF8(src))
}
