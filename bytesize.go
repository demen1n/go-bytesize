// Package bytesize provides functionality for measuring and formatting  byte
// sizes.
//
// You can also perform mathematical operation with ByteSize's and the result
// will be a valid ByteSize with the correct size suffix.
package bytesize

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// This code was originally based on http://golang.org/doc/progs/eff_bytesize.go
//
// Since then many improvements have been made. The following is the original
// copyright notice:

// Copyright 2009 The Go Authors. All rights reserved. Use of this source code
// is governed by a BSD-style license that can be found in the LICENSE file.

// ByteSize represents a number of bytes
type ByteSize uint64

// Byte size suffixes
const (
	B  ByteSize = 1
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

// Locale represents a supported locale
type Locale string

const (
	LocaleEN Locale = "en"
	LocaleRU Locale = "ru"
)

// unitDefinitions is a struct for unit definitions for different locales
type unitDefinitions struct {
	// longUnits used for returning long unit form of string representation.
	longUnits map[ByteSize]string
	// shortUnits used for returning string representation.
	shortUnits map[ByteSize]string
	// parseMap used to convert user input to ByteSize
	parseMap map[string]ByteSize
}

// Localized unit definitions
var localizedUnits = map[Locale]unitDefinitions{
	LocaleEN: {
		longUnits: map[ByteSize]string{
			B:  "byte",
			KB: "kilobyte",
			MB: "megabyte",
			GB: "gigabyte",
			TB: "terabyte",
			PB: "petabyte",
			EB: "exabyte",
		},
		shortUnits: map[ByteSize]string{
			B:  "B",
			KB: "KB",
			MB: "MB",
			GB: "GB",
			TB: "TB",
			PB: "PB",
			EB: "EB",
		},
		parseMap: map[string]ByteSize{
			"B": B, "BYTE": B, "BYTES": B,
			"KB": KB, "KILOBYTE": KB, "KILOBYTES": KB,
			"MB": MB, "MEGABYTE": MB, "MEGABYTES": MB,
			"GB": GB, "GIGABYTE": GB, "GIGABYTES": GB,
			"TB": TB, "TERABYTE": TB, "TERABYTES": TB,
			"PB": PB, "PETABYTE": PB, "PETABYTES": PB,
			"EB": EB, "EXABYTE": EB, "EXABYTES": EB,
		},
	},
	LocaleRU: {
		longUnits: map[ByteSize]string{
			B:  "байт",
			KB: "килобайт",
			MB: "мегабайт",
			GB: "гигабайт",
			TB: "терабайт",
			PB: "петабайт",
			EB: "эксабайт",
		},
		shortUnits: map[ByteSize]string{
			B:  "Б",
			KB: "КБ",
			MB: "МБ",
			GB: "ГБ",
			TB: "ТБ",
			PB: "ПБ",
			EB: "ЭБ",
		},
		parseMap: map[string]ByteSize{
			"Б": B, "БАЙТ": B, "БАЙТЫ": B, "БАЙТОВ": B,
			"КБ": KB, "КИЛОБАЙТ": KB, "КИЛОБАЙТЫ": KB, "КИЛОБАЙТОВ": KB,
			"МБ": MB, "МЕГАБАЙТ": MB, "МЕГАБАЙТЫ": MB, "МЕГАБАЙТОВ": MB,
			"ГБ": GB, "ГИГАБАЙТ": GB, "ГИГАБАЙТЫ": GB, "ГИГАБАЙТОВ": GB,
			"ТБ": TB, "ТЕРАБАЙТ": TB, "ТЕРАБАЙТЫ": TB, "ТЕРАБАЙТОВ": TB,
			"ПБ": PB, "ПЕТАБАЙТ": PB, "ПЕТАБАЙТЫ": PB, "ПЕТАБАЙТОВ": PB,
			"ЭБ": EB, "ЭКСАБАЙТ": EB, "ЭКСАБАЙТЫ": EB, "ЭКСАБАЙТОВ": EB,
		},
	},
}

func init() {
	for k, v := range localizedUnits[LocaleEN].parseMap {
		if _, exists := localizedUnits[LocaleRU].parseMap[k]; !exists {
			localizedUnits[LocaleRU].parseMap[k] = v
		}
	}
}

var (
	// Current locale
	CurrentLocale Locale = LocaleEN

	// Use long units, such as "megabytes" instead of "MB".
	LongUnits bool = false

	// Format var is a string format of bytesize output. The unit of measure will be appended
	// to the end. Uses the same formatting options as the fmt package.
	Format string = "%.2f"
)

// SetLocale sets the current locale for formatting and parsing
func SetLocale(locale Locale) {
	if _, exists := localizedUnits[locale]; exists {
		CurrentLocale = locale
	}
}

// parseWithLocale parses a byte size string using the specified locale
func parseWithLocale(s string, locale Locale) (ByteSize, error) {
	units, ok := localizedUnits[locale]
	if !ok {
		return 0, errors.New("unsupported locale: " + string(locale))
	}

	// Remove leading and trailing whitespace
	s = strings.TrimSpace(s)

	split := make([]string, 0)
	for i, r := range s {
		if !unicode.IsDigit(r) && r != '.' {
			// Split the string by digit and size designator, remove whitespace
			split = append(split, strings.TrimSpace(string(s[:i])))
			split = append(split, strings.TrimSpace(string(s[i:])))
			break
		}
	}

	// Check to see if we split successfully
	if len(split) != 2 {
		return 0, errors.New("unrecognized size suffix")
	}

	// Check for unit in the parse map
	unit, ok := units.parseMap[strings.ToUpper(split[1])]
	if !ok {
		return 0, errors.New("unrecognized size suffix: " + split[1])
	}

	value, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return 0, err
	}

	bytesize := ByteSize(value * float64(unit))
	return bytesize, nil
}

// Parse parses a byte size string. A byte size string is a number followed by
// a unit suffix, such as "1024B" or "1 MB". Valid byte units are "B", "KB",
// "MB", "GB", "TB", "PB" and "EB". You can also use the long
// format of units, such as "kilobyte" or "kilobytes".
func Parse(s string) (ByteSize, error) {
	bs, err := parseWithLocale(s, CurrentLocale)
	return bs, err
}

// Satisfy the flag package Value interface.
func (b *ByteSize) Set(s string) error {
	bs, err := Parse(s)
	if err != nil {
		return err
	}
	*b = bs
	return nil
}

// Satisfy the pflag package Value interface.
func (b *ByteSize) Type() string { return "byte_size" }

// Satisfy the encoding.TextUnmarshaler interface.
func (b *ByteSize) UnmarshalText(text []byte) error {
	return b.Set(string(text))
}

// Satisfy the flag package Getter interface.
func (b *ByteSize) Get() interface{} { return ByteSize(*b) }

// New returns a new ByteSize type set to s.
func New(s float64) ByteSize {
	return ByteSize(s)
}

// Returns a string representation of b with the specified formatting and units.
func (b ByteSize) Format(format string, unit string, longUnits bool) string {
	return b.formatWithLocale(format, unit, longUnits, CurrentLocale)
}

// formatWithLocale returns a string representation using the specified locale
func (b ByteSize) formatWithLocale(format string, unit string, longUnits bool, locale Locale) string {
	units, ok := localizedUnits[locale]
	if !ok {
		locale = LocaleEN
		units = localizedUnits[LocaleEN]
	}

	return b.formatWithUnits(format, unit, longUnits, units)
}

// String returns the string form of b using the package global options
func (b ByteSize) String() string {
	return b.stringWithLocale(CurrentLocale)
}

// stringWithLocale returns the string form using the specified locale
func (b ByteSize) stringWithLocale(locale Locale) string {
	return b.formatWithLocale(Format, "", LongUnits, locale)
}

func (b ByteSize) formatWithUnits(format string, unit string, longUnits bool, units unitDefinitions) string {
	var unitSize ByteSize
	if unit != "" {
		var ok bool
		unitSize, ok = units.parseMap[strings.ToUpper(unit)]
		if !ok {
			return "Unrecognized unit: " + unit
		}
	} else {
		switch {
		case b >= EB:
			unitSize = EB
		case b >= PB:
			unitSize = PB
		case b >= TB:
			unitSize = TB
		case b >= GB:
			unitSize = GB
		case b >= MB:
			unitSize = MB
		case b >= KB:
			unitSize = KB
		default:
			unitSize = B
		}
	}

	value := float64(b) / float64(unitSize)

	if longUnits {
		unitStr := units.longUnits[unitSize]
		// russian plural form based on the number
		if CurrentLocale == LocaleRU {
			unitStr = getRussianPlural(value, unitSize)
		} else if CurrentLocale == LocaleEN {
			if value > 0 && value != 1 {
				unitStr = unitStr + "s"
			}
		}
		return fmt.Sprintf(format+" %s", value, unitStr)
	}

	return fmt.Sprintf(format+"%s", value, units.shortUnits[unitSize])
}

// getRussianPlural returns the correct Russian plural form based on the number
func getRussianPlural(value float64, unit ByteSize) string {
	intValue := int(value)

	var forms []string
	switch unit {
	case B:
		forms = []string{"байт", "байта", "байтов"}
	case KB:
		forms = []string{"килобайт", "килобайта", "килобайтов"}
	case MB:
		forms = []string{"мегабайт", "мегабайта", "мегабайтов"}
	case GB:
		forms = []string{"гигабайт", "гигабайта", "гигабайтов"}
	case TB:
		forms = []string{"терабайт", "терабайта", "терабайтов"}
	case PB:
		forms = []string{"петабайт", "петабайта", "петабайтов"}
	case EB:
		forms = []string{"эксабайт", "эксабайта", "эксабайтов"}
	}

	if intValue%100 >= 11 && intValue%100 <= 19 {
		return forms[2] // много (11-19)
	}

	switch intValue % 10 {
	case 1:
		return forms[0] // один
	case 2, 3, 4:
		return forms[1] // несколько (2-4)
	default:
		return forms[2] // много (0, 5-9)
	}
}
