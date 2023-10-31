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