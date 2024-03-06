// Copyright 2024 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handler

import (
	"fmt"
	"net/http"
)

func PageView(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	h := fmt.Sprintf(`<html>
<head>
	<meta charset="UTF-8" />
	<link rel="icon" sizes="256x256" href="//lf-scm-cn.feishucdn.com/ccm/pc/web/resource/bear/feishu.ico"/>
	<link rel="icon" sizes="32x32" href="//lf-scm-cn.feishucdn.com/ccm/pc/web/resource/bear/feishu.ico"/>
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="applicable-device" content="pc,mobile">
	<meta name="renderer" content="webkit">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=5.0">
	<meta name="twitter:image:src" content="%[1]s" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content="页面访问量统计" />
	<meta name="twitter:description" content="Powered by LarkStat.  https://github.com/wuhan005/LarkStat" />
	<meta property="og:image" content="%[1]s" />
	<meta property="og:image:alt" content="页面访问量统计" />
	<meta property="og:image:width" content="%[2]d" />
	<meta property="og:image:height" content="%[2]d" />
	<meta property="og:site_name" content="页面访问量统计" />
	<meta property="og:type" content="object" />
	<meta property="og:title" content="页面访问量统计" />
	<meta property="og:description" content="Powered by LarkStat.  https://github.com/wuhan005/LarkStat" />
</head>
<body>
</body>
</html>
`, req.URL.String()+".png", 700, 220)

	_, _ = fmt.Fprint(w, h)
}
