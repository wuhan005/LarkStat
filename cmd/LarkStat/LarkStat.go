package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"

	"github.com/wuhan005/LarkStat/internal/handler"
	"github.com/wuhan005/LarkStat/internal/strutil"
)

func main() {
	kvNamespace := os.Getenv("KV_NAMESPACE")
	if kvNamespace == "" {
		kvNamespace = "larkstat"
	}

	workers.Serve(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		label := req.URL.Path
		referer := req.Header.Get("Referer")
		subDomain, err := strutil.ParseLarkSubDomain(referer)
		if err != nil {
			if errors.Is(err, strutil.ErrEmptySubDomain) {
				w.Header().Set("Content-Type", "text/plain")
				_, _ = fmt.Fprint(w, "设置一个独特的路径，再将链接以卡片链接的形式插入飞书文档~")
				return
			}
			logrus.WithError(err).Error("Failed to parse lark subdomain")
			return
		}

		if strings.HasSuffix(label, ".png") {
			label = label[:len(label)-4]
			counterKey := fmt.Sprintf("pv-%s-%s", subDomain, label)

			kv, err := cloudflare.NewKVNamespace(req.Context(), kvNamespace)
			if err != nil {
				handler.Error(w, err, "Failed to create KV namespace")
				return
			}

			// Read the count number from KV.
			countStr, err := kv.GetString(counterKey, nil)
			if err != nil {
				handler.Error(w, err, "Failed to get counter")
				return
			}
			count, _ := strconv.Atoi(countStr)
			if err := kv.PutString(counterKey, strconv.Itoa(count+1), nil); err != nil {
				handler.Error(w, err, "Failed to update counter")
				return
			}

			handler.DrawImage(req.Context(), handler.DrawImageOptions{
				Writer: w,
				Count:  count,
			})

		} else {
			handler.PageView(w, req)
		}
	}))
}
