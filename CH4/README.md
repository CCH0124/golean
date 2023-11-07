## 區塊
- 有宣告式的每個地方都稱為一個**區塊**
- 在所有函式外宣告的變數、常數、型態和函式都在 **package** 區塊裡
- 當宣告式使用的名稱與外面區塊的識別碼名稱一樣時，會遮蔽（shadow）外面區塊建立的識別碼

### 遮蔽變數

```go
func main() {
	x := 10
	if x > 5 {
		log.Println(x)
		x := 5
		log.Println(x) // 遮蔽變數
	}
	log.Println(x)
}
// 10
// 5
// 10
```

遮蔽變數（shadowing variable）就是名稱與外面區塊變數一樣的變數，只要有遮蔽變數，就無法存取被遮蔽的變數

使用 `:=` 非常容易遮蔽變數

>宇宙區塊（universe block），有些**預先宣告的識別碼**會被定義在其裡面像是 nil, make, close 等

### if
與其它相比是雷同的，不一樣的是他可以宣告條件式在 if 和 else 區塊之內的變數

```go
func main() {
	if n := rand.Intn(10); n == 0 {
		log.Println("too low")
	} else if n > 5 {
		log.Printf("too big %d", n)
	} else {
		log.Printf("good number %d", n)
	}
}
```

這可以建立在需要使用時才可以使用的變數，因此在 if/else 條件式完成時，`n` 的定義就會被移除．如果在條件區塊外進行 `n` 的存取將會編譯錯誤

### for
其風格有以下：
- C 風格
- 只有條件式的 for
- 無窮 for
- for-range

C 風格

```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```
格式與大部分一樣，第一部分是初始值，需使用 `:=` 設定值。第二部分是比較式，需是個布林值。最後是遞增式。

只有條件式的 for

相似於其它語言的 `while`

```go
i := 1
for i < 100 {
	fmt.Println(i)
	i = i * 2
}
```

無窮 for

```go
for {
	fmt.Println(i)
}
```

要跳出無窮迴圈，可以使用 `break` 關鍵字或是使用 `continue` 跳過這次回圈的迭代。

for-range 陳述式

類似其它語言的 `iterator`。

```go
evens := []int{2, 4, 6, 8, 10}
for i, v := range evens {
	fmt.Println(i, v)
}
```

特別的地方在於，他有兩個變數。第一個是 `index`;第二個是該 `index` 的 `value`。如果不希望使用索引鍵可以使用 `_` 關鍵字。

如果要索引鍵不要值其允許省略第二個變數

```go
	uniqName := map[string]bool{"a": true, "b": true, "c": true}
	for k := range uniqName {
		log.Println(k)
	}
```

### for-range 值是副本
使用 `for-range` 迭代複合型態時，會將複合型態的值*複製*到值變數。**修改值變數時不會修改複合型態裡面的值**。

```go
	evenVals := []int{2, 4, 6, 8}
	for _, v := range evenVals {
		v *= 2
	}
	log.Println(evenVals)
	// [2 4 6 8]
```


>for-range 是遍歷字串最佳做法，其可提供 `rune`，而非 `bytes`
