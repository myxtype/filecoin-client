# filecoin-client

需要自行部署Lotus Node节点：https://lotu.sh/en+getting-started

此库仅添加部分方法，但已经满足钱包充值提现逻辑了，如果需要其他方法，请Fork后自行添加。

充值流程：获取头部高度，从本地高度遍历到头部高度，再根据高度获取区块CID，根据区块CID获取区块的所有消息，判断消息的类型是否0(0为发送Fil)，和接收地址是否是本地的地址。

说明请查询client_test文件。


# 安装

`go get github.com/myxtype/filecoin-client`

# 使用

```go
package main

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client"
)

func main() {

	client := filecoin.NewClient("http://127.0.0.1:1234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.cF__3r_0IR9KwZ2nLkqcBW8vuPePruZieJAVvTAoUA4")

	addr, _ := address.NewFromString("t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
	b, err := client.WalletBalance(context.Background(), addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
}
```

# 离线签名

离线签名已支持，详情轻查看wallet_test.go中`TestClient_WalletNewLocal` 和 `TestClient_WalletSignMessageLocal`。

签名和地址生成目前仅支持`SigTypeSecp256k1`类型。

# 注意

使用本库前，需要设置地址的网络类型，否则地址是以`t`开头的测试地址。

```go
import (
	"github.com/filecoin-project/go-address"
)

func init() {
	address.CurrentNetwork = address.Mainnet
}
```

尽量放在init函数。

# Lotus文档

https://lotu.sh/en+api-methods
