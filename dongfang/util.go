package dongfang

import (
	"compress/gzip"
	"compress/zlib"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/chentiangang/xlog"
	"github.com/klauspost/compress/zstd"
)

func ut() string {
	data := time.Now().Format("2006-01-02 15:04:05")

	// 生成 MD5 哈希
	hash := md5.Sum([]byte(data))

	// 将哈希值转换为 32 位的十六进制字符串
	return hex.EncodeToString(hash[:])
}

var (
	schema   = "https"
	quoteApi = "push2.eastmoney.com"
)

func Domain() string {
	i := rand.Int31n(100) + 1
	return fmt.Sprintf("%s://%d.%s", schema, i, quoteApi)
}

const (
	ListTotal = 5644
)

func Unzip(resp *http.Response) ([]byte, error) {
	var reader io.Reader
	var err error
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			xlog.Error("Error opening gzip reader: %s", err)
			return nil, nil
		}
	case "deflate":
		reader, err = zlib.NewReader(resp.Body)
		if err != nil {
			xlog.Error("Error opening deflate reader: %s", err)
			return nil, nil
		}
	case "br":
		reader = brotli.NewReader(resp.Body)
	case "zstd":
		decoder, err := zstd.NewReader(resp.Body)
		if err != nil {
			xlog.Error("Error opening zstd reader: %s", err)
			return nil, nil
		}
		defer decoder.Close()
		reader = decoder
	default:
		reader = resp.Body // 没有压缩时直接读取响应体
	}
	return io.ReadAll(reader)
}
