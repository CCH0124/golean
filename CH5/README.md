# 函式
## 宣告與呼叫函式
Go 語言程式起點是 `main` 函式，它不接收任何參數和回傳值。

對於有回傳和參數的函示如同下面

[sample code](CH5/example1/main.go)
```go
func div(numerator int, denominator int) int {
    if denominator == 0 {
        return 0
    }
    return numerator/denominator
}
```

Go 是有型態的語言，所以要指定參數的回傳型態。如果有定義回傳值則必須使用 `return` 關鍵字。

## 模擬具名（named）與選用（optional）參數
對於 Go 來說沒有具名（named）與選用（optional）參數，如果要模擬可以使用 `struct` 方式。
