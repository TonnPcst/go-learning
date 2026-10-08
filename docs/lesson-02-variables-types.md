# ບົດທີ 2 — Variable, Constant, Type ພື້ນຖານ ແລະ ການແປງ Type

> Phase 1 · ພື້ນຖານ Go
> ໂຄດ: [`lesson02/`](../lesson02/)
> ເວີຊັນ Go: 1.27.1

---

## 1. Variable ແລະ zero value

```go
var count int            // declared, no value given → 0
var name = "Ton"         // type inferred: string
city := "Vientiane"      // short declaration: only inside functions
a, b := 1, 2             // multiple assignment
a, b = b, a              // swap, no temporary variable needed
```

**Zero value:** ທຸກ variable ມີຄ່າສະເໝີ. ບໍ່ມີ `undefined`, ແລະ `nil` ມີສະເພາະ pointer, map, slice, channel, function ແລະ interface ເທົ່ານັ້ນ.

| Type | Zero value |
|---|---|
| `int`, `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |
| pointer, slice, map, interface | `nil` |
| struct | ທຸກ field ເປັນ zero value ຂອງມັນ |

**ເປັນຫຍັງ Go ເຮັດແບບນີ້:** TypeScript ມີ "ບໍ່ມີຄ່າ" ສອງແບບ (`undefined` ແລະ `null`) ບວກກັບ optional chaining ເພື່ອຮັບມືກັບມັນ. Go ເລືອກ "ມີຄ່າເລີ່ມຕົ້ນສະເໝີ", ເພື່ອໃຫ້ zero value ສ່ວນຫຼາຍພ້ອມໃຊ້ເລີຍ. `var sb strings.Builder` ຫຼື `var mu sync.Mutex` ໃຊ້ໄດ້ທັນທີ ໂດຍບໍ່ຕ້ອງມີ constructor.

### `:=` ທຽບກັບ `=`

- `:=` **ປະກາດ** variable ໃໝ່; `=` **assign** ຄ່າໃຫ້ variable ທີ່ມີຢູ່ແລ້ວ.
- `:=` ຕ້ອງມີຊື່ໃໝ່ຢ່າງໜ້ອຍໜຶ່ງຊື່ຢູ່ເບື້ອງຊ້າຍ:

```go
x, err := f()   // declares x and err
y, err := g()   // OK: y is new, err is reused
```

- `:=` ໃຊ້ໄດ້ສະເພາະພາຍໃນ function. ລະດັບ package ຕ້ອງໃຊ້ `var`.

### ກັບດັກ shadowing (ຈະເຈີໃນໂຄດຈິງ)

```go
var cfg Config
if debug {
	cfg, err := loadDebug() // NEW cfg, scoped to this if block
	_ = err
	fmt.Println("inner:", cfg.Name) // inner cfg must be used, or it won't compile
}
fmt.Println(cfg) // {}: the outer cfg is still the zero value
```

ພາຍໃນ block, `:=` ສ້າງ `cfg` ຕົວທີສອງ ທີ່ຫາຍໄປເມື່ອ block ຈົບ. ຜົນທີ່ໄດ້: `inner: debug` ແລ້ວ `{}`.

- ຖ້າບໍ່ໃຊ້ `cfg` ຕົວໃນເລີຍ, ຈະເປັນ compile error `declared and not used: cfg`. ໃນໂຄດຈິງ ຕົວໃນມັກຖືກໃຊ້ພາຍໃນ block, ສະນັ້ນ compile ຜ່ານ ແລະ bug ນີ້ຈຶ່ງເບິ່ງບໍ່ເຫັນ.
- `go vet` **ບໍ່ກວດ** shadowing ໂດຍ default (ການກວດ shadow ເປັນເຄື່ອງມືແຍກຕ່າງຫາກ), ສະນັ້ນໃຫ້ລະວັງທຸກຄັ້ງທີ່ໃຊ້ `:=` ພາຍໃນ block (`if`, `for`, `switch`).

ວິທີແກ້: ປະກາດ `err` ກ່ອນ, ແລ້ວໃຊ້ `=` ເພື່ອ assign ໃສ່ `cfg` ຕົວນອກ:

```go
var cfg Config
if debug {
	var err error
	cfg, err = loadDebug() // = assigns to the OUTER cfg
	if err != nil {
		// handle the error
	}
}
```

## 2. Constant ແລະ `iota`

```go
const MaxChargers = 100           // untyped constant
const Timeout = 30 * time.Second  // typed (time.Duration)

type ConnectorStatus int

const (
	StatusAvailable ConnectorStatus = iota // 0
	StatusCharging                         // 1
	StatusFaulted                          // 2
)
```

- **Constant ຖືກຄຳນວນຕອນ compile.** ເປັນໄດ້ສະເພາະຕົວເລກ, string ຫຼື bool, ບໍ່ສາມາດເປັນ slice ຫຼື struct.
- **Untyped constant** ເຊັ່ນ `MaxChargers` ປັບຕົວຕາມບໍລິບົດ, ສະນັ້ນ `var f float64 = MaxChargers` compile ຜ່ານ. ນີ້ແມ່ນບ່ອນດຽວທີ່ Go ເຮັດຄ້າຍກັບການແປງ type ແບບອັດຕະໂນມັດ.
- **`iota`** ນັບຂຶ້ນຈາກ 0 ພາຍໃນ block `const (...)`. ນີ້ແມ່ນ `enum` ຂອງ TS ໃນແບບ Go. ບໍ່ມີ keyword `enum`: ໃຫ້ສ້າງ named type ບວກກັບ constant.
- **ຂໍ້ຄວນລະວັງ:** `ConnectorStatus(42)` ຍັງ compile ຜ່ານ. Enum ຂອງ Go ບໍ່ແມ່ນຊຸດປິດ, ສະນັ້ນຕ້ອງ validate ຂໍ້ມູນທີ່ມາຈາກພາຍນອກເອງ.

## 3. Type ພື້ນຖານ

| Go | TypeScript | ໝາຍເຫດ |
|---|---|---|
| `int`, `int8/16/32/64` | `number` | `int` ເປັນ 64-bit ເທິງ Mac ແລະ server. **Overflow ວົນກັບແບບງຽບໆ.** |
| `uint`, `uint8`…`uint64` | — | `byte` = `uint8` |
| `float32`, `float64` | `number` | ໃຊ້ `float64` ເປັນຄ່າເລີ່ມຕົ້ນ |
| `string` | `string` | **Byte ແບບ UTF-8 ທີ່ປ່ຽນແປງບໍ່ໄດ້ (immutable)** |
| `rune` | — | = `int32`; ໜຶ່ງ Unicode code point |
| `bool` | `boolean` | |

TypeScript ມີ type `number` ດຽວ (float 64-bit). Go ໃຫ້ເລືອກເອງ, ເພາະຂະໜາດ ແລະ ການມີເຄື່ອງໝາຍ (signed/unsigned) ສຳຄັນສຳລັບ memory, database (`BIGINT` ↔ `int64`) ແລະ binary protocol.

### String ແມ່ນ byte — ສຳຄັນຫຼາຍສຳລັບພາສາລາວ

```go
s := "ສະບາຍດີ"
fmt.Println(len(s))                    // 21: BYTES, not characters
fmt.Println(utf8.RuneCountInString(s)) // 7: code points
for i, r := range s {                  // range over a string yields runes
	fmt.Println(i, string(r))           // i is the byte offset: 0, 3, 6...
}
```

- ໃນ JavaScript, `"ສະບາຍດີ".length` ໄດ້ 7, ເພາະ JS ນັບເປັນ UTF-16 code unit.
- ໃນ Go, `len` ນັບ **byte**. ຕົວອັກສອນລາວແຕ່ລະຕົວໃຊ້ 3 byte ໃນ UTF-8, ສະນັ້ນ 7 × 3 = 21.
- `s[0:2]` ຕັດຕົວອັກສອນເຄິ່ງກາງ ແລະ ໄດ້ byte ທີ່ບໍ່ແມ່ນ UTF-8 ທີ່ຖືກຕ້ອງ.
- **Rune ບໍ່ແມ່ນ "ຕົວອັກສອນທີ່ເຫັນ" ສະເໝີ:** ສະຫຼະ ແລະ ວັນນະຍຸດ (ເຊັ່ນ `ະ`, `ີ`) ເປັນ rune ແຍກ. `ດີ` ເບິ່ງຄືຕົວດຽວ ແຕ່ເປັນ 2 rune. ການນັບສິ່ງທີ່ຕາເຫັນ (grapheme cluster) ຕ້ອງໃຊ້ library ເພີ່ມ.

## 4. ການແປງ type: ຕ້ອງຂຽນເອງສະເໝີ

```go
var a int = 7
var b float64 = 2
// a / b                    // compile error: mismatched types int and float64
fmt.Println(float64(a) / b) // 3.5
fmt.Println(a / 2)          // 3: integer division truncates

var big int64 = 300
small := int8(big)          // 44: compiles, silently wraps
```

- ຮູບແບບແມ່ນ `T(v)`: `float64(a)`, `int64(n)`.
- ການຫານ integer **ຕັດເສດຖິ້ມ** (`7 / 2 == 3`).
- ການແປງໄປ type ທີ່ນ້ອຍກວ່າ **ບໍ່ມີ error**: `int8(300)` ໄດ້ `44` (300 − 256).

### String ↔ ຕົວເລກ ໃຊ້ `strconv`, ບໍ່ແມ່ນການ cast

```go
n, err := strconv.Atoi("42")              // int, error
id, err := strconv.ParseInt("42", 10, 64) // base 10, 64-bit
s := strconv.Itoa(42)                     // "42"

string(65) // "A", NOT "65"! It converts a code point. go vet warns about this.
```

**ເປັນຫຍັງເຂັ້ມງວດຂະໜາດນີ້:** ໃນ JavaScript `"5" * 2 === 10` ແລະ `"5" + 2 === "52"` ເປັນແຫຼ່ງຂອງ bug. Go ບໍ່ຍອມປົນ type ແບບງຽບໆ. ການ parse ອາດລົ້ມເຫຼວໄດ້, ຈຶ່ງ return `error` ອອກມາ. `Number("abc")` ໄດ້ `NaN` ໂດຍບໍ່ມີການເຕືອນ; `strconv.Atoi("abc")` ໃຫ້ error ທີ່ຕ້ອງຈັດການ.

## 5. Verb ຂອງ `fmt` ທີ່ໃຊ້ເລື້ອຍ

```go
fmt.Printf("%v %T\n", x, x)              // value, type
fmt.Printf("%d %s %q\n", 42, "hi", "hi") // 42 hi "hi"
fmt.Printf("%.2f\n", 3.14159)            // 3.14
fmt.Printf("%+v\n", someStruct)          // struct with field names
fmt.Fprintln(os.Stderr, "error:", err)   // print to stderr
```

## ສະຫຼຸບ syntax

- `var x T` = zero value; `var x = v` = infer type; `x := v` = ປະກາດສັ້ນ (ພາຍໃນ function ເທົ່ານັ້ນ).
- `const` + `iota` = enum ແບບ Go; ຕ້ອງ validate ຄ່າທີ່ມາຈາກພາຍນອກເອງ.
- `len(string)` = ຈຳນວນ byte; `utf8.RuneCountInString` = ຈຳນວນ code point; `range` ເທິງ string ໃຫ້ rune.
- ແປງ type ດ້ວຍ `T(v)` ສະເໝີ; string ↔ ຕົວເລກ ໃຊ້ `strconv` ແລະ ກວດ `err`.
- ເງິນ = integer (ໜ່ວຍນ້ອຍສຸດ), ບໍ່ແມ່ນ `float64`.

---

## ແບບຝຶກຫັດ 2: ເຄື່ອງຄິດໄລ່ຄ່າສາກໄຟ ⚡

ສ້າງ module ໃໝ່ `lesson02` (`go mod init example.com/lesson02`).

**1. Package `internal/charging`:**

- `type ConnectorType int` ພ້ອມ constant ແບບ `iota` ສຳລັບ `Type2`, `CCS2` ແລະ `CHAdeMO`.
- **ລາຄາຕໍ່ connector type ເປັນ ກີບ/kWh**, ເຊັ່ນ Type2 = 1500, CCS2 = 2500, CHAdeMO = 2500. ໃຊ້ integer: **ຫ້າມໃຊ້ float ກັບເງິນ.** ເອົາລາຄາໄວ້ໃນ `switch` ຫຼື function.
- `func Cost(c ConnectorType, energyWh int64) int64` return ລາຄາເປັນກີບ. ພະລັງງານເຂົ້າມາເປັນ **Wh** (meter ຂອງ OCPP ລາຍງານເປັນ Wh), ລາຄາເປັນຕໍ່ **kWh**. **ປັດຂຶ້ນ** ເປັນກີບເຕັມ ໂດຍໃຊ້ integer ເທົ່ານັ້ນ: ບໍ່ໃຊ້ `float64` ແລະ ບໍ່ໃຊ້ `math.Ceil`.
- Table-driven test ທີ່ມີກໍລະນີພໍດີ (ເຊັ່ນ 1000 Wh), ກໍລະນີທີ່ຕ້ອງປັດຂຶ້ນ (ເຊັ່ນ 1 Wh) ແລະ 0 Wh.

**2. `cmd/cost/main.go`:**

- ວິທີໃຊ້: `go run ./cmd/cost 12345` ພິມລາຄາສຳລັບ CCS2.
- Parse argument ດ້ວຍ `strconv.ParseInt`. ຖ້າ parse ບໍ່ໄດ້ ຫຼື ບໍ່ມີ argument, ພິມຂໍ້ຄວາມໄປ **stderr** (`fmt.Fprintln(os.Stderr, ...)`) ແລ້ວອອກດ້ວຍ `os.Exit(1)`.

**3. ຕອບສັ້ນໆ:**

- **A.** `len("ສະບາຍດີ")` ແລະ `utf8.RuneCountInString("ສະບາຍດີ")` return ເທົ່າໃດ, ແລະ ເປັນຫຍັງຈຶ່ງຕ່າງກັນ? `.length` ຂອງ JavaScript ຈະ return ເທົ່າໃດ?
- **B.** ເປັນຫຍັງ integer (ກີບ, ຫຼື cent/satang ໃນສະກຸນເງິນອື່ນ) ດີກວ່າ `float64` ສຳລັບເງິນ? ລອງ `fmt.Println(0.1 + 0.2)` ໃນ Go ແລ້ວອະທິບາຍສິ່ງທີ່ເຫັນ.
- **C.** `string(65)` ໄດ້ຫຍັງ, ແລະ ວິທີທີ່ຖືກຕ້ອງໃນການແປງຕົວເລກ 65 ເປັນ string `"65"` ແມ່ນຫຍັງ?
- **D.** ໃນຕົວຢ່າງ shadowing, ເປັນຫຍັງ `cfg` ຕົວນອກຈຶ່ງຍັງຫວ່າງ? ຈະຂຽນ block ນັ້ນຄືນໃໝ່ແນວໃດ ເພື່ອໃຫ້ assign ໃສ່ `cfg` ຕົວນອກ?

**Hint ສຳລັບການປັດຂຶ້ນ:** ເຄັດລັບ integer ທີ່ໃຊ້ກັນທົ່ວໄປແມ່ນ `(a + b - 1) / b`. ຄິດໃຫ້ອອກກ່ອນວ່າ *ເປັນຫຍັງ* ມັນຈຶ່ງປັດຂຶ້ນ, ແລ້ວຈຶ່ງໃຊ້.

**ສິ່ງທີ່ຕ້ອງສົ່ງ:** ໂຄດ, ຜົນ `go test -v`, ຜົນຂອງໂປຣແກຣມເມື່ອ argument ຖືກ ແລະ ຜິດ, ແລະ ຄຳຕອບ.

### ວິທີແກ້

- [`charging.go`](../lesson02/internal/charging/charging.go):
  - `ConnectorType` + `iota` (`Type2` = 0, `CCS2` = 1, `CHAdeMO` = 2).
  - `PricePerKWh(c) (price int64, ok bool)` ໃຊ້ `switch`; `default` return `ok = false` ສຳລັບ `ConnectorType(42)`. ນີ້ແມ່ນ pattern "comma ok": Go ບໍ່ມີ `undefined`, ຈຶ່ງບອກວ່າ "ບໍ່ພົບ" ດ້ວຍ `bool` ຕົວທີສອງ.
  - `Cost` return `0` ຖ້າ connector ບໍ່ຮູ້ຈັກ ຫຼື `energyWh <= 0`, ເພາະ signature `Cost(...) int64` return error ບໍ່ໄດ້. ເມື່ອຮຽນເລື່ອງ error ແລ້ວ ຄວນປ່ຽນເປັນ `(int64, error)`.
- [`charging_test.go`](../lesson02/internal/charging/charging_test.go): table-driven test 8 ກໍລະນີ: ພໍດີ, ປັດຂຶ້ນ, 0 Wh, connector ທີ່ບໍ່ຮູ້ຈັກ, ແລະ ພະລັງງານຕິດລົບ.
- [`main.go`](../lesson02/cmd/cost/main.go): ກວດ `len(os.Args)`, `strconv.ParseInt(s, 10, 64)` + `if err != nil`, ປະຕິເສດຄ່າຕິດລົບ. Error ທັງໝົດພິມໄປ stderr ແລະ `os.Exit(1)`.

**ການປັດຂຶ້ນດ້ວຍ integer:** `(energyWh*price + 999) / 1000`

- ການຫານ integer ປັດລົງສະເໝີ (`2500 / 1000 = 2`).
- ບວກ `ຕົວຫານ − 1` (999) ກ່ອນຫານ: ຖ້າຫານລົງຕົວພໍດີ, 999 ບໍ່ພໍທີ່ຈະໄປເຖິງພັນຖັດໄປ; ຖ້າມີເສດ ແມ້ແຕ່ 1, ມັນຈະໄປເຖິງພັນຖັດໄປສະເໝີ.

| a | a + 999 | ÷ 1000 |
|---|---|---|
| 2000 (2.0 ພໍດີ) | 2999 | **2** |
| 2001 (2.001) | 3000 | **3** |
| 2500 (2.5) | 3499 | **3** |

- **ຄູນກ່ອນຫານ.** `energyWh / 1000 * price` ຈະໄດ້ 1 Wh / 1000 = 0, ລາຄາກາຍເປັນ 0.

### ຜົນລັບ

```text
$ go test -v ./...
?       example.com/lesson02/cmd/cost   [no test files]
=== RUN   TestCost
=== RUN   TestCost/1_kWh_exact
=== RUN   TestCost/1_Wh_rounds_up
=== RUN   TestCost/zero_energy
=== RUN   TestCost/Type2_half_kip
=== RUN   TestCost/Type2_999_Wh
=== RUN   TestCost/CHAdeMO_12345_Wh
=== RUN   TestCost/unknown_connector
=== RUN   TestCost/negative_energy
--- PASS: TestCost (0.00s)
    --- PASS: TestCost/1_kWh_exact (0.00s)
    --- PASS: TestCost/1_Wh_rounds_up (0.00s)
    --- PASS: TestCost/zero_energy (0.00s)
    --- PASS: TestCost/Type2_half_kip (0.00s)
    --- PASS: TestCost/Type2_999_Wh (0.00s)
    --- PASS: TestCost/CHAdeMO_12345_Wh (0.00s)
    --- PASS: TestCost/unknown_connector (0.00s)
    --- PASS: TestCost/negative_energy (0.00s)
PASS
ok      example.com/lesson02/internal/charging

$ go test -cover ./...
        example.com/lesson02/cmd/cost           coverage: 0.0% of statements
ok      example.com/lesson02/internal/charging  coverage: 100.0% of statements

$ go run ./cmd/cost 12345
CCS2, 12345 Wh: 30863 kip
$ go run ./cmd/cost 1
CCS2, 1 Wh: 3 kip
$ go run ./cmd/cost abc
error: energy must be a whole number of Wh: strconv.ParseInt: parsing "abc": invalid syntax
exit status 1
$ go run ./cmd/cost -5
error: energy cannot be negative
exit status 1
$ go run ./cmd/cost
usage: cost <energy-Wh>
exit status 1
```

- `12345` → 12345 × 2500 = 30,862,500 ÷ 1000 = 30862.5 → ປັດຂຶ້ນເປັນ **30863**.
- `1` → 2.5 → **3**. ຖ້າບໍ່ມີ `+ 999` ຈະໄດ້ 2.
- ຂໍ້ຄວາມ error ລວມ error ຂອງ `strconv` ໄວ້ນຳ (`parsing "abc": invalid syntax`), ຜູ້ໃຊ້ຈຶ່ງຮູ້ວ່າຜິດຫຍັງ.

### ຄຳຖາມ ແລະ ຄຳຕອບ

ຜົນຈາກ playground ([`cmd/play`](../lesson02/cmd/play/main.go), ເກັບໄວ້ເພື່ອທົດລອງຕໍ່):

```text
A: 21 7
B const: 0.3
B vars:  0.30000000000000004
C: A 65

$ node -e 'console.log("ສະບາຍດີ".length)'
7
```

**A. `len("ສະບາຍດີ")` ແລະ `utf8.RuneCountInString("ສະບາຍດີ")`**
`len` = **21**, ເພາະ `len` ນັບ byte ຂອງ UTF-8 ແລະ ຕົວອັກສອນລາວແຕ່ລະຕົວໃຊ້ 3 byte (7 × 3). `utf8.RuneCountInString` = **7** code point. `.length` ຂອງ JavaScript = **7**, ເພາະ JS ນັບເປັນ UTF-16 unit ແລະ ຕົວອັກສອນລາວໃຊ້ unit ດຽວຕໍ່ຕົວ. Rune ບໍ່ແມ່ນຕົວອັກສອນທີ່ຕາເຫັນສະເໝີ: `ດີ` ເບິ່ງຄືຕົວດຽວ ແຕ່ເປັນ 2 rune.

**B. ເປັນຫຍັງໃຊ້ integer ກັບເງິນ, ແລະ `0.1 + 0.2`**
`float64` ເກັບ 0.1 ແລະ 0.2 ແບບພໍດີບໍ່ໄດ້ (ໃນລະບົບເລກຖານສອງ), ແຕ່ລະຕົວຈຶ່ງຜິດໜ້ອຍໜຶ່ງ. ເມື່ອເອົາ variable ມາບວກກັນ ຄວາມຜິດກໍບວກກັນ: `a + b` = `0.30000000000000004`.
ສ່ວນ `0.1 + 0.2` ທີ່ຂຽນໂດຍກົງ ພິມ `0.3` ເພາະເປັນ **untyped constant**: Go ຄຳນວນ constant ແບບພໍດີຕອນ compile, ແລ້ວຈຶ່ງແປງເປັນ `float64`. ຜົນນີ້ຫຼອກຕາ, ເພາະເງິນຈິງມາຈາກ variable (database, meter, request) ສະເໝີ.
ເງິນເປັນ integer (ກີບ, ຫຼື ໜ່ວຍນ້ອຍສຸດຂອງສະກຸນເງິນ) ຈຶ່ງພໍດີສະເໝີ, ບໍ່ມີຄວາມຜິດສະສົມ.

**C. `string(65)`**
ໄດ້ `"A"`, ບໍ່ແມ່ນ `"65"`: ມັນແປງ **code point** 65 ເປັນຕົວອັກສອນ. `go run` ຍັງ compile ຜ່ານ, ແຕ່ `go vet` (ແລະ `go test`) ເຕືອນ:

```text
conversion from int to string yields a string of one rune, not a string of digits
```

ວິທີທີ່ຖືກ: `strconv.Itoa(65)` → `"65"`. ຖ້າຕ້ອງການຕົວອັກສອນແທ້ ໃຫ້ຂຽນໃຫ້ຊັດ: `string(rune(65))` → `"A"`.

**D. Shadowing: ເປັນຫຍັງ `cfg` ຕົວນອກຍັງຫວ່າງ?**
`err` ເປັນຊື່ໃໝ່ໃນ block, ສະນັ້ນ `:=` ໃຊ້ໄດ້, ແຕ່ `:=` ປະກາດ **ທັງສອງຊື່ໃໝ່** ໃນ scope ຂອງ block, ລວມທັງ `cfg` ຕົວໃໝ່ທີ່ບັງ `cfg` ຕົວນອກ. ຄ່າຖືກໃສ່ໃນ `cfg` ຕົວໃນ ເຊິ່ງຫາຍໄປເມື່ອ block ຈົບ. ວິທີແກ້: ປະກາດ `var err error` ກ່ອນ, ແລ້ວໃຊ້ `cfg, err = loadDebug()` (ໃຊ້ `=` ບໍ່ແມ່ນ `:=`). ເບິ່ງໂຄດໃນພາກ 1.

### ບັນທຶກການ review

- ✅ ເງິນເປັນ `int64` ກີບ, ບໍ່ມີ `float64` ຫຼື `math.Ceil`; ປັດຂຶ້ນດ້ວຍ `(a + b - 1) / b`, ແລະ ຄູນກ່ອນຫານ.
- ✅ `switch` ມີ `default` ສຳລັບ connector type ທີ່ບໍ່ຮູ້ຈັກ, ເພາະ enum ຂອງ Go ບໍ່ແມ່ນຊຸດປິດ.
- ✅ Test ຄອບຄຸມທຸກ branch (coverage 100%), ລວມທັງ input ທີ່ບໍ່ຖືກຕ້ອງ. ທົດລອງປ່ຽນ `+ 999` ເປັນ `+ 0` ເພື່ອເບິ່ງວ່າ test ກໍລະນີປັດຂຶ້ນ fail ແທ້.
- ✅ Error ໄປ stderr ແລະ exit code 1; return ອອກໄວ (early exit) ເຮັດໃຫ້ເສັ້ນທາງປົກກະຕິບໍ່ຕ້ອງຫຍໍ້ໜ້າ.
- ⚠️ `Cost` return `0` ສຳລັບ input ທີ່ບໍ່ຖືກຕ້ອງ, ເຊິ່ງແຍກບໍ່ອອກຈາກ "ສາກ 0 Wh". ຈະແກ້ເປັນ `(int64, error)` ເມື່ອຮຽນເລື່ອງ error.
- ⚠️ `energyWh*price` ອາດ overflow ຖ້າ `energyWh` ໃຫຍ່ກວ່າປະມານ 3.7 × 10¹⁵ Wh. ບໍ່ເກີດຂຶ້ນໃນການສາກຈິງ, ແຕ່ຄວນຮູ້ວ່າ integer overflow ວົນກັບແບບງຽບໆ.
