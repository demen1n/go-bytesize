ByteSize
========

A localized fork of [go-bytesize](https://github.com/inhies/go-bytesize) with internationalization support for working with byte size measurements.

Using this package allows you to easily add 100KB to 4928MB and get a nicely formatted string representation of the result in multiple languages with proper grammar rules.

[![Go Reference](https://pkg.go.dev/badge/github.com/demen1n/go-bytesize.svg)](https://pkg.go.dev/github.com/demen1n/go-bytesize)
[![Go Report Card](https://goreportcard.com/badge/github.com/demen1n/go-bytesize)](https://goreportcard.com/report/github.com/demen1n/go-bytesize)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

## ✨ New Features

This fork adds the following features while maintaining **100% backward compatibility**:

- 🌍 **Multi-language support** (English, Russian)
- 📝 **Proper plural forms** and grammar rules for each language
- 🔄 **Enhanced parsing** - supports both localized and English units
- 🎯 **Flexible formatting** - per-locale customization
- ⚡ **Improved input handling** - comma decimal separator support
- 🧪 **Extensive test coverage** - all localization features tested

## 📦 Installation

```bash
go get github.com/demen1n/go-bytesize
```

## 🚀 Quick Start

### Basic Usage (Same as Original)

```go
package main

import (
    "fmt"
    "github.com/demen1n/go-bytesize"
)

func main() {
    // Works exactly like the original library
    size := bytesize.New(1024 * 1024 * 1.5) // 1.5 MB
    fmt.Println(size) // "1.50MB"
    
    // Parse sizes
    parsed, _ := bytesize.Parse("2.5 GB")
    fmt.Println(parsed) // "2.50GB"
}
```

### New Localization Features

```go
package main

import (
    "fmt"
    "github.com/demen1n/go-bytesize"
)

func main() {
    // Switch to Russian locale
    bytesize.SetLocale(bytesize.LocaleRU)
    bytesize.LongUnits = true
    
    size := bytesize.New(1024 * 1024 * 1.5) // 1.5 MB
    fmt.Println(size) // "1.50 мегабайта"
    
    // Parse Russian units
    parsed, _ := bytesize.Parse("2 ГБ")
    fmt.Println(parsed) // "2.00 гигабайта"
    
    // English units still work in Russian locale
    english, _ := bytesize.Parse("500 MB")
    fmt.Println(english) // "500.00 мегабайтов"
}
```

### Russian Pluralization

Russian language has complex pluralization rules that are properly handled:

```go
bytesize.SetLocale(bytesize.LocaleRU)
bytesize.LongUnits = true
bytesize.Format = "%.0f"

fmt.Println(bytesize.New(1 * bytesize.MB))   // "1 мегабайт"
fmt.Println(bytesize.New(2 * bytesize.MB))   // "2 мегабайта" 
fmt.Println(bytesize.New(5 * bytesize.MB))   // "5 мегабайтов"
fmt.Println(bytesize.New(11 * bytesize.MB))  // "11 мегабайтов"
fmt.Println(bytesize.New(21 * bytesize.MB))  // "21 мегабайт"
```

## 📖 Usage Examples

### Parsing Different Formats

```go
// English formats
sizes := []string{"1 MB", "2.5 GB", "1024B", "1 megabyte", "500 kilobytes"}

// Russian formats  
русскиеРазмеры := []string{"1 МБ", "2,5 ГБ", "1024Б", "1 мегабайт", "500 килобайтов"}

bytesize.SetLocale(bytesize.LocaleRU)
for _, size := range русскиеРазмеры {
    parsed, err := bytesize.Parse(size)
    if err == nil {
        fmt.Printf("%s = %s\n", size, parsed.String())
    }
}
```

### Per-Operation Locale

```go
size := bytesize.New(1024 * 1024 * 2.5)

// Format in different locales without changing global settings
english := size.StringWithLocale(bytesize.LocaleEN)   // "2.50MB"
russian := size.StringWithLocale(bytesize.LocaleRU)   // "2.50МБ"

fmt.Printf("English: %s, Russian: %s\n", english, russian)
```

### Custom Formatting

```go
size := bytesize.New(1536) // 1.5 KB

// Custom format with specific locale
formatted := size.FormatWithLocale("%.1f", "KB", true, bytesize.LocaleRU)
fmt.Println(formatted) // "1.5 килобайта"
```

## 🌐 Supported Locales

| Locale | Code | Short Units | Long Units | Parse Support |
|--------|------|-------------|------------|---------------|
| English | `en` | B, KB, MB, GB, TB, PB, EB | byte, kilobyte, megabyte, ... | ✅ |
| Russian | `ru` | Б, КБ, МБ, ГБ, ТБ, ПБ, ЭБ | байт, килобайт, мегабайт, ... | ✅ |

**Note**: Russian locale also supports parsing English units for maximum compatibility.

## 🔧 Configuration

### Global Settings

```go
// Set active locale
bytesize.SetLocale(bytesize.LocaleRU)

// Use long unit names
bytesize.LongUnits = true

// Custom number format
bytesize.Format = "%.1f"

// Get supported locales
locales := bytesize.GetSupportedLocales()
fmt.Println(locales) // [en ru]
```

### Thread Safety

All parsing and formatting operations are thread-safe. Global settings (`CurrentLocale`, `LongUnits`, `Format`) should be set once during initialization or protected with your own synchronization.

## 🔄 Migration from Original

**Zero changes needed!** Your existing code will work unchanged:

```go
// This code works exactly the same
size := bytesize.New(1024)
parsed, err := bytesize.Parse("1 MB")
formatted := size.Format("%.2f", "KB", false)
```

New features are opt-in through new functions and locale settings.

## 🎯 API Reference

### New Functions

- `SetLocale(locale Locale)` - Set global locale
- `GetSupportedLocales() []Locale` - Get available locales
- `ParseWithLocale(s string, locale Locale) (ByteSize, error)` - Parse with specific locale
- `(b ByteSize) StringWithLocale(locale Locale) string` - Format with specific locale
- `(b ByteSize) FormatWithLocale(format, unit string, longUnits bool, locale Locale) string` - Custom format with locale

### New Types

- `Locale` - Locale identifier (`LocaleEN`, `LocaleRU`)

### Compatible API

All original functions work unchanged:
- `New(float64) ByteSize`
- `Parse(string) (ByteSize, error)`
- `(b ByteSize) String() string`
- `(b ByteSize) Format(format, unit string, longUnits bool) string`
- Plus all `flag.Value` and `encoding.TextUnmarshaler` interfaces

## 🧪 Testing

```bash
# Run all tests
go test -v

# Run specific locale tests
go test -run TestRuLocale -v

# Run benchmarks
go test -bench=. -benchmem

# Check test coverage
go test -cover
```
