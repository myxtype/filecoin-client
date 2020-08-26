# filecoin-client

需要自行部署Lotus Node节点：https://lotu.sh/en+getting-started

此库仅添加部分方法，但已经满足钱包充值提现逻辑了，如果需要其他方法，请Fork后自行添加。

充值流程：获取头部高度，从本地高度遍历到头部高度，再根据高度获取区块CID，根据区块CID获取区块的所有消息，判断消息的类型是否0(0为发送Fil)，和接收地址是否是本地的地址。

说明请查询client_test文件。

# 使用

```go
package main

func main() {
    
    client := NewClient("http://127.0.0.1:1234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.cF__3r_0IR9KwZ2nLkqcBW8vuPePruZieJAVvTAoUA4")

    b, err := client.WalletBalance(context.Background(), "t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
    if err != nil {
        //
    }

    fmt.Println(b.String())
}
```

# 离线签名版

我将在后续提供离线签名版（不需要自行搭建Node），因为现在https://infura.io 还未开放Lotus Node的API，所有离线签名版将在后续更新。

具体实现逻辑在官方库中：https://github.com/filecoin-project/go-filecoin

# Lotus文档

https://lotu.sh/en+api-methods
