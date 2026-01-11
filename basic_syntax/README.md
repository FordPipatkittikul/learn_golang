# basic syntax

## types

🔹 Value Types (❌ cannot be nil)

These always have a concrete value (zero value if not set).

| Type                      | Example                | Zero Value                       |
| ------------------------- | ---------------------- | -------------------------------- |
| `bool`                    | `true`                 | `false`                          |
| All integers              | `int`, `int64`, `uint` | `0`                              |
| Floating point            | `float32`, `float64`   | `0.0`                            |
| `complex64`, `complex128` |                        | `0+0i`                           |
| `string`                  | `"abc"`                | `""` (empty string, **not nil**) |
| `array`                   | `[3]int`               | `[0,0,0]`                        |
| `struct`                  | `struct{}`             | fields’ zero values              |
| `uintptr`                 |                        | `0`                              |

🔹 Reference / Nil-capable Types (✅ can be nil)

These are often called reference types (even though Go doesn’t officially use that term).

| Type                 | Can be `nil` | Notes                              |
| -------------------- | ------------ | ---------------------------------- |
| `pointer` (`*T`)     | ✅            | Most common for optional values    |
| `slice` (`[]T`)      | ✅            | `nil` slice ≠ empty slice          |
| `map` (`map[K]V`)    | ✅            | Writing to nil map panics          |
| `channel` (`chan T`) | ✅            | Send/receive blocks forever        |
| `function` (`func`)  | ✅            | Calling nil func panics            |
| `interface`          | ✅            | Subtle: typed nil vs nil interface |

##