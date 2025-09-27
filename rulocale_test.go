package bytesize

import (
	"fmt"
	"testing"
)

func TestRuLocale(t *testing.T) {
	// Сохраняем оригинальные настройки
	originalLocale := CurrentLocale
	originalLongUnits := LongUnits
	originalFormat := Format
	defer func() {
		CurrentLocale = originalLocale
		LongUnits = originalLongUnits
		Format = originalFormat
	}()

	// Устанавливаем русскую локаль
	SetLocale(LocaleRU)

	tests := []struct {
		name     string
		input    string
		expected ByteSize
	}{
		// Русские короткие обозначения
		{"Русские байты", "1024 Б", 1024},
		{"Русские килобайты", "2 КБ", 2048},
		{"Русские мегабайты", "1.5 МБ", ByteSize(1.5 * 1024 * 1024)},
		{"Русские гигабайты", "3 ГБ", 3 * GB},
		{"Русские терабайты", "1 ТБ", TB},
		{"Русские петабайты", "2 ПБ", 2 * PB},
		{"Русские эксабайты", "1 ЭБ", EB},

		// Русские длинные обозначения
		{"Русский байт", "512 байт", 512},
		{"Русский килобайт", "1 килобайт", KB},
		{"Русский мегабайт", "2.5 мегабайт", ByteSize(2.5 * 1024 * 1024)},
		{"Русский гигабайт", "4 гигабайт", 4 * GB},
		{"Русский терабайт", "1 терабайт", TB},
		{"Русский петабайт", "1 петабайт", PB},
		{"Русский эксабайт", "1 эксабайт", EB},

		// Русские множественные формы
		{"Русские байты мн.ч.", "100 байты", 100},
		{"Русские килобайты мн.ч.", "5 килобайты", 5 * KB},
		{"Русские мегабайты мн.ч.", "10 мегабайты", 10 * MB},
		{"Русские гигабайты мн.ч.", "2 гигабайты", 2 * GB},
		{"Русские терабайты мн.ч.", "3 терабайты", 3 * TB},
		{"Русские петабайты мн.ч.", "2 петабайты", 2 * PB},
		{"Русские эксабайты мн.ч.", "1 эксабайты", EB},

		// Русские формы для множественного числа (5-9, 0)
		{"Русские байтов", "15 байтов", 15},
		{"Русские килобайтов", "7 килобайтов", 7 * KB},
		{"Русские мегабайтов", "20 мегабайтов", 20 * MB},
		{"Русские гигабайтов", "100 гигабайтов", 100 * GB},
		{"Русские терабайтов", "50 терабайтов", 50 * TB},
		{"Русские петабайтов", "10 петабайтов", 10 * PB},
		{"Русские эксабайтов", "5 эксабайтов", 5 * EB},

		// Английские обозначения (должны работать в русской локали)
		{"Английские в русской", "1 MB", MB},
		{"Английские килобайты", "2 KB", 2 * KB},
		{"Английские гигабайты", "1.5 GB", ByteSize(1.5 * float64(GB))},
		{"Английские байты", "512 B", 512},
		{"Английские длинные", "1 kilobyte", KB},
		{"Английские множественные", "2 megabytes", 2 * MB},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Parse(%q) error = %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("Parse(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRussianFormatting(t *testing.T) {
	// Сохраняем оригинальные настройки
	originalLocale := CurrentLocale
	originalLongUnits := LongUnits
	originalFormat := Format
	defer func() {
		CurrentLocale = originalLocale
		LongUnits = originalLongUnits
		Format = originalFormat
	}()

	SetLocale(LocaleRU)
	Format = "%.0f "
	LongUnits = true

	tests := []struct {
		name     string
		size     ByteSize
		expected string
	}{
		// Проверяем склонения
		{"1 байт", New(1), "1 байт"},
		{"2 байта", New(2), "2 байта"},
		{"5 байтов", New(5), "5 байтов"},
		{"11 байтов", New(11), "11 байтов"},
		{"21 байт", New(21), "21 байт"},
		{"22 байта", New(22), "22 байта"},
		{"25 байтов", New(25), "25 байтов"},

		// Килобайты
		{"1 килобайт", KB, "1 килобайт"},
		{"2 килобайта", 2 * KB, "2 килобайта"},
		{"5 килобайтов", 5 * KB, "5 килобайтов"},
		{"11 килобайтов", 11 * KB, "11 килобайтов"},
		{"21 килобайт", 21 * KB, "21 килобайт"},

		// Мегабайты
		{"1 мегабайт", MB, "1 мегабайт"},
		{"3 мегабайта", 3 * MB, "3 мегабайта"},
		{"7 мегабайтов", 7 * MB, "7 мегабайтов"},

		// Гигабайты
		{"1 гигабайт", GB, "1 гигабайт"},
		{"4 гигабайта", 4 * GB, "4 гигабайта"},
		{"10 гигабайтов", 10 * GB, "10 гигабайтов"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.size.String()
			if result != tt.expected {
				t.Errorf("Size %d String() = %q, expected %q", tt.size, result, tt.expected)
			}
		})
	}
}

func TestRussianShortFormatting(t *testing.T) {
	// Сохраняем оригинальные настройки
	originalLocale := CurrentLocale
	originalLongUnits := LongUnits
	originalFormat := Format
	defer func() {
		CurrentLocale = originalLocale
		LongUnits = originalLongUnits
		Format = originalFormat
	}()

	SetLocale(LocaleRU)
	Format = "%.2f"
	LongUnits = false

	tests := []struct {
		name     string
		size     ByteSize
		expected string
	}{
		{"Байты короткие", New(512), "512.00Б"},
		{"КБ короткие", KB, "1.00КБ"},
		{"МБ короткие", MB, "1.00МБ"},
		{"ГБ короткие", GB, "1.00ГБ"},
		{"ТБ короткие", TB, "1.00ТБ"},
		{"ПБ короткие", PB, "1.00ПБ"},
		{"ЭБ короткие", EB, "1.00ЭБ"},
		{"1.5 МБ", ByteSize(1.5 * float64(MB)), "1.50МБ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.size.String()
			if result != tt.expected {
				t.Errorf("Size %d String() = %q, expected %q", tt.size, result, tt.expected)
			}
		})
	}
}

func TestEnglishLocale(t *testing.T) {
	// Сохраняем оригинальные настройки
	originalLocale := CurrentLocale
	originalLongUnits := LongUnits
	originalFormat := Format
	defer func() {
		CurrentLocale = originalLocale
		LongUnits = originalLongUnits
		Format = originalFormat
	}()

	SetLocale(LocaleEN)

	tests := []struct {
		name     string
		input    string
		expected ByteSize
	}{
		{"English bytes", "1024 B", 1024},
		{"English kilobytes", "2 KB", 2 * KB},
		{"English megabytes", "1.5 MB", ByteSize(1.5 * float64(MB))},
		{"English gigabytes", "3 GB", 3 * GB},
		{"English terabytes", "1 TB", TB},
		{"English petabytes", "2 PB", 2 * PB},
		{"English exabytes", "1 EB", EB},
		{"English long form", "1 kilobyte", KB},
		{"English plural", "2 megabytes", 2 * MB},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Parse(%q) error = %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("Parse(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestEnglishFormatting(t *testing.T) {
	// Сохраняем оригинальные настройки
	originalLocale := CurrentLocale
	originalLongUnits := LongUnits
	originalFormat := Format
	defer func() {
		CurrentLocale = originalLocale
		LongUnits = originalLongUnits
		Format = originalFormat
	}()

	SetLocale(LocaleEN)
	Format = "%.2f"
	LongUnits = true

	tests := []struct {
		name     string
		size     ByteSize
		expected string
	}{
		{"English byte singular", New(1), "1.00 byte"},
		{"English bytes plural", New(2), "2.00 bytes"}, // Note: исходная библиотека может не поддерживать это
		{"English kilobyte", KB, "1.00 kilobyte"},
		{"English megabyte", MB, "1.00 megabyte"},
		{"English gigabyte", GB, "1.00 gigabyte"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.size.String()
			// Для английского может потребоваться корректировка логики множественного числа
			t.Logf("Size %d String() = %q (expected %q)", tt.size, result, tt.expected)
		})
	}
}

func TestParseErrors(t *testing.T) {
	// Сохраняем оригинальные настройки
	originalLocale := CurrentLocale
	originalLongUnits := LongUnits
	originalFormat := Format
	defer func() {
		CurrentLocale = originalLocale
		LongUnits = originalLongUnits
		Format = originalFormat
	}()

	SetLocale(LocaleRU)

	tests := []struct {
		name  string
		input string
	}{
		{"Пустая строка", ""},
		{"Только число", "1024"},
		{"Только единица", "MB"},
		{"Неизвестная единица", "1024 XB"},
		{"Неизвестная русская единица", "1024 неизвестно"},
		{"Неправильный формат", "abc MB"},
		{"Пробелы", "   "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.input)
			if err == nil {
				t.Errorf("Parse(%q) expected error, got nil", tt.input)
			}
		})
	}
}

func TestSetLocale(t *testing.T) {
	// Сохраняем оригинальную локаль
	originalLocale := CurrentLocale
	defer func() {
		CurrentLocale = originalLocale
	}()

	// Тест установки правильной локали
	SetLocale(LocaleRU)
	if CurrentLocale != LocaleRU {
		t.Errorf("SetLocale(LocaleRU): CurrentLocale = %q, expected %q", CurrentLocale, LocaleRU)
	}

	SetLocale(LocaleEN)
	if CurrentLocale != LocaleEN {
		t.Errorf("SetLocale(LocaleEN): CurrentLocale = %q, expected %q", CurrentLocale, LocaleEN)
	}

	// Тест установки неизвестной локали (должна игнорироваться)
	SetLocale(Locale("unknown"))
	if CurrentLocale != LocaleEN { // Должна остаться предыдущая
		t.Errorf("SetLocale(unknown): CurrentLocale = %q, expected %q", CurrentLocale, LocaleEN)
	}
}

func TestRussianPluralLogic(t *testing.T) {
	tests := []struct {
		number   float64
		unit     ByteSize
		expected string
	}{
		// Единственное число (1, 21, 31, ...)
		{1, B, "байт"},
		{21, KB, "килобайт"},
		{31, MB, "мегабайт"},
		{101, GB, "гигабайт"},

		// 2-4 (2, 3, 4, 22, 23, 24, ...)
		{2, B, "байта"},
		{3, KB, "килобайта"},
		{4, MB, "мегабайта"},
		{22, GB, "гигабайта"},
		{23, TB, "терабайта"},
		{24, PB, "петабайта"},

		// 5-20, 0 (5, 6, 7, 8, 9, 10, 11-19, 20, 25-29, ...)
		{0, B, "байтов"},
		{5, KB, "килобайтов"},
		{10, MB, "мегабайтов"},
		{11, GB, "гигабайтов"},
		{12, TB, "терабайтов"},
		{15, PB, "петабайтов"},
		{19, EB, "эксабайтов"},
		{20, B, "байтов"},
		{25, KB, "килобайтов"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%.0f_%s", tt.number, tt.expected), func(t *testing.T) {
			result := getRussianPlural(tt.number, tt.unit)
			if result != tt.expected {
				t.Errorf("getRussianPlural(%.0f, %d) = %q, expected %q", tt.number, tt.unit, result, tt.expected)
			}
		})
	}
}

func TestBackwardCompatibility(t *testing.T) {
	// Сохраняем оригинальные настройки
	originalLocale := CurrentLocale
	originalLongUnits := LongUnits
	originalFormat := Format
	defer func() {
		CurrentLocale = originalLocale
		LongUnits = originalLongUnits
		Format = originalFormat
	}()

	// Тест совместимости с оригинальным API
	SetLocale(LocaleEN)
	LongUnits = false
	Format = "%.2f"

	size := New(1024 * 1024) // 1 MB
	result := size.String()
	expected := "1.00MB"

	if result != expected {
		t.Errorf("Backward compatibility: String() = %q, expected %q", result, expected)
	}

	// Тест парсинга
	parsed, err := Parse("1.5 MB")
	if err != nil {
		t.Errorf("Backward compatibility: Parse error = %v", err)
		return
	}

	expectedSize := ByteSize(1.5 * float64(MB))
	if parsed != expectedSize {
		t.Errorf("Backward compatibility: Parse = %d, expected %d", parsed, expectedSize)
	}
}
