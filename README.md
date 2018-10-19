### hugo && GitHub webhook 的自动化推送工具
#### 依赖
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
    "github.com/mailru/easyjson"
#### 使用
make clean
make build

#### 文件描述

fileoperate.go 为在新一次构建docker容器后一键完成hugo的头部填充
