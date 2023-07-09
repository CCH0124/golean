# 複合型態
## Array

宣告

```go
// 指定陣列大小與元素型態
var x [3]int
```

沒有指定值時，會被設定該型態的預設值

陣列常值宣告，其可以省略元素數量，可以使用 `...` 即 `[...]{1,2,3}`
```go
var x = [3]int{1,2,3}
```

稀疏陣列

```go
var x = [12]int{1, 5:4, 6, 10:100, 15}
```

建立 12 個 int 型態元素，值會是 `[1, 0, 0, 0, 0, 4, 6 ,0, 0, 0, 100, 15]`

當要讀取長度，可以使用內建函數 `len`，其接收一個陣列並回傳長度。


**陣列大小無法用變數來決定，因為型態必須在編譯期決定，而非執行期。**且，不能使用型態轉換將不同大小的陣列轉換成一樣大小。


## Slice

`slice` 的長度非型態一部分，因此沒有陣列的限制。

宣告 `slice` 時無須定義大小

```go
var x = []int{10, 20, 30}
```

使用索引建立值

```go
var x = []int{1, 5:4, 6, 10:100, 15}
```

也可以宣告多維的 `slice`，`var x [][]int`

讀取和寫入值和陣列是相同的

```go
x[0] = 10
fmt.Println(x[2])
```

不使用常值進行宣告

```go
var x []int
```

其初始值會是 `nil`。對於 `slice` 而言他是無法進行 `==` 或 `!=` 比較的，但其可以和 `nil` 比較。

函式
- len
- append
    - 擴增 slice
- cap
    - 保留連續空間記憶體的數量
    - 當 slice 長度等於 cap 時，底層會配置更大的 slice 容量
- make


雖然 slice 可以自動增長長度，但是更有效率的方式是一次設定好大小。那至於如何設定就要透過 `make` 函示。

對於 slice 常值或是 nil 宣告來說，不能建立特定長度或是容量的 slice。


>Go 是**以值呼叫**，因此傳入參數時，會被底層製作一個副本。

[example](CH3/slice/main.go)
### make

宣告長度為 5，容量為 5 的 int 型態 slice，且初始值都為 0。
```go
x := make([]int, 5)
```

定義初始容量

```go
x := make([]int, 5, 10)
```

再有宣告 `len` 時，再用 `append` 方式進行賦值，會是該 `len` 之後。除非使用 `len` 為 0，但大於 0 的容量。

```go
var z := make([]int, 0, 3)
z = append(z, 10)
log.Printf("value: %v", z) //  [10]
```

[example](CH3/make/main.go)

### 切割 slice

slice 運算式可以用 slice 來建立 slice。

```go
    x := []int{1, 2, 3, 4}
	y := x[:2]
	log.Printf("y value: %v", y)
	z := x[1:]
	log.Printf("z value: %v", z)
	d := x[1:3]
	log.Printf("d value: %v", d)
	a := x[:]
	log.Printf("a value: %v", a)
// 2023/06/25 21:26:17 y value: [1 2]
// 2023/06/25 21:26:17 z value: [2 3 4]
// 2023/06/25 21:26:17 d value: [2 3]
// 2023/06/25 21:26:17 a value: [1 2 3 4]
```

### slice 共用儲存
從 slice 裡面取出的 slice 不會製作資料的複本，*而會得到兩個共享記憶體的變數*。因此變更會相互影響。

```go
    x := []int{1, 2, 3, 4}
    x[1] = 20
	y[0] = 10
	z[2] = 30
	log.Printf("x value: %v", x)
	log.Printf("y value: %v", y)
	log.Printf("z value: %v", z)
// 2023/06/25 21:30:09 x value: [10 20 3 30]
// 2023/06/25 21:30:09 y value: [10 20]
// 2023/06/25 21:30:09 z value: [20 3 30]
```

如果經過 `slice` 切割後再用 `append` 增加元素，會變得很困惑

```go
	x1 := []int{1, 2, 3, 4}
	y1 := x1[:2]
	log.Printf("x1 Cap: %d, y1 Cap: %d", cap(x1), cap(y1))

	y1 = append(y1, 30)
	log.Printf("x1 value: %v", x1)
	log.Printf("y1 value: %v", y1)
// 2023/06/25 21:40:07 x1 Cap: 4, y1 Cap: 4
// 2023/06/25 21:40:07 x1 value: [1 2 30 4]
// 2023/06/25 21:40:07 y1 value: [1 2 30]
```

在共用儲存下，原始 slice 未使用的*容量*也會與子 slice 共用。以範例 `y1` slice 來說，該 slice 的容量是原始 slice 容量減去子 slice 在原始 slice 內的位移。


```go
	log.Println("----------------------------------------")
	x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y = x[:2]
	z = x[2:]
	log.Printf("x : %d, y: %d, z: %d", cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	log.Printf("x value: %v", x)
	log.Printf("y value: %v", y)
	log.Printf("z value: %v", z)
// 2023/07/06 20:30:52 x : 5, y: 5, z: 3
// 2023/07/06 20:30:52 x value: [1 2 30 40 70]
// 2023/07/06 20:30:52 y value: [1 2 30 40 70]
// 2023/07/06 20:30:52 z value: [30 40 70]
```

上面範例說明*絕對不要使用 append 與子 slice*，否則使用*完整 slice 運算式* 確保 append 不會造成覆寫，如下範例。

```go
	x1 = make([]int, 0, 5)
	x1 = append(x1, 1, 2, 3, 4)
	y1 = x1[:2:2]
	z1 := x1[2:4:4]
	log.Printf("y1 cap : %d, z1 cap: %d", cap(y1), cap(z1))
	y1 = append(y1, 30, 40, 50)
	z1 = append(z1, 70)
	log.Printf("y1 value: %v", y1)
	log.Printf("z1 value: %v", z1)
// 2023/07/09 14:56:17 y1 cap : 2, z1 cap: 2
// 2023/07/09 14:56:17 y1 value: [1 2 30 40 50]
// 2023/07/09 14:56:17 z1 value: [3 4 70]
```

原因是因為將子 slice 容量限制為它們的長度，所以新增額外元素會產生新的 slice。

## 陣列轉成 slice
在陣列上可以使用 slice 函式取出 slice。同樣的取出的 slice 也會參造共用的記憶體。

```go
	x := [4]int{1, 2, 3, 4}
	y := x[:2]
	z := x[2:]
	log.Printf("x: %v", x)
	log.Printf("y: %v", y)
	log.Printf("z: %v", z)
```

### copy

如果要建立和來源相互獨立的 slice，可以使用 copy 函示

```go
	x1 := []int{1, 2, 3, 4}
	y1 := make([]int, 4)
	num := copy(y1, x1) // 將 x1 複製到 y1
	log.Printf("y: %v", y1)
	log.Printf("num length: %v", num)
// 2023/07/09 15:17:26 y: [1 2 3 4]
// 2023/07/09 15:17:26 num: 4
```
*copy 回傳為複製元素的數量*

複製可以有以下變化
```go
	x1 := []int{1, 2, 3, 4}
	y1 := make([]int, 2) // x1 最多兩個元素被複製到 y1
	num := copy(y1, x1)
```

```go
	x1 := []int{1, 2, 3, 4}
	y1 := make([]int, 2) // x1 最多兩個元素被複製到 y1
	copy(y1, x1[2:])
```

```go
	x1 := []int{1, 2, 3, 4}
	copy(x1[:3], x1[1:])
```

## 字串與 rune 與 byte

可以使用*索引運算式(index expression)* 從字串中取值
```go
	var s string = "Hello String"
	var b byte = s[6]
```


*slice 運算式*可用於處理字串。由於字串是不可變的，因此沒有 slice 共享記憶體問題。但有其它問題，字串由 byte 組成，但使用非英文或表情符號，會遇到更常 byte 的 UTF-8 編碼。這會在做切片運算式時可能出現非預期行為。

對於字串長度可以使用 `len` 函式獲取。*對於字串索引和 slice 運算式都用 bytes 來計算位置，所以長度是 bytes 長度*。

`rune` 或 `byte` 可以轉換成字串。

```go
		var a rune = 'x'
	s = string(a)
	var a1 rune = 'y'
	s2 = string(a1)
	log.Printf("s: %s", s)
	log.Printf("s2: %s", s2)

	s = "Hello, 🥱"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	log.Printf("bs: %v", bs)
	log.Printf("rs: %v", rs)

// 2023/07/09 20:39:43 s: x
// 2023/07/09 20:39:43 s2: y
// 2023/07/09 20:39:43 bs: [72 101 108 108 111 44 32 240 159 165 177]
// 2023/07/09 20:39:43 rs: [72 101 108 108 111 44 32 129393]	
```

## map
宣告 `string` 類型索引鍵，`int` 類型的值，預設值為 `nil`，長度為 0。
```go
var nilMap map[string]int
```

使用 `:=` 將一個*map 常值*指派給 map 變數來宣告

```go
totalWins := map[string]int{}
```
