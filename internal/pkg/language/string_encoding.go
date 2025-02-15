package language

import (
	"github.com/allanpk716/ChineseSubFinder/internal/logic/charset"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/log_helper"
	"github.com/axgle/mahonia"
	nzlov "github.com/nzlov/chardet"
)

// ConvertToString 将字符串从原始编码转换到目标编码，需要配合字符串检测编码库使用 chardet.NewTextDetector()
func ConvertToString(src string, srcCode string, tagCode string) string {
	defer func() {
		if err := recover(); err != nil {
			log_helper.GetLogger().Errorln("ConvertToString panic:", err)
		}
	}()
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 感谢: https://blog.csdn.net/gaoluhua/article/details/109128154，解决了编码问题

// ChangeFileCoding2UTF8 自动检测文件的编码，然后转换到 UTF-8
func ChangeFileCoding2UTF8(inBytes []byte) ([]byte, error) {
	best, err := detector.DetectBest(inBytes)
	utf8String := ""
	if err != nil {
		return nil, err
	}
	if best.Confidence < 90 {
		detectBest := nzlov.Mostlike(inBytes)
		utf8String, err = charset.ToUTF8(charset.Charset(detectBest), string(inBytes))
	} else {
		utf8String, err = charset.ToUTF8(charset.Charset(best.Charset), string(inBytes))
	}
	if err != nil {
		return nil, err
	}
	if utf8String == "" {
		return inBytes, nil
	}
	return []byte(utf8String), nil
}
