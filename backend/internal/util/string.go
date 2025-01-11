package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
)

func IsStringEmpty(s string) bool { return len(strings.TrimSpace(s)) == 0 }

const (
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	idLength = 22
)

// GenerateNanoID generates a random string using the custom alphabet
func GenerateNanoID(length int) (string, error) {

	// Default length is 22
	if length == 0 {
		length = idLength
	}

	bytes := make([]byte, length)
	alphabetLen := big.NewInt(int64(len(alphabet)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, alphabetLen)
		if err != nil {
			return "", err
		}
		bytes[i] = alphabet[n.Int64()]
	}

	return string(bytes), nil
}

// GeneratePrefixedID generates an ID with the given prefix
func GeneratePrefixedID(prefix string, separator string, length int) (string, error) {
	id, err := GenerateNanoID(length)
	if err != nil {
		return "", err
	}

	if separator == "" {
		return fmt.Sprintf("%s_%s", prefix, id), nil
	}

	return fmt.Sprintf("%s%s%s", prefix, separator, id), nil
}

// Slugify converts a string to a URL-friendly slug
func Slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Normalize unicode characters
	s = norm.NFKD.String(s)

	// Remove accents
	s = regexp.MustCompile(`[^\w\s-]`).ReplaceAllString(s, "")

	// Convert spaces and underscores to hyphens
	s = regexp.MustCompile(`[\s_]+`).ReplaceAllString(s, "-")

	// Remove leading/trailing hyphens
	s = strings.Trim(s, "-")

	// Replace multiple consecutive hyphens with single hyphen
	s = regexp.MustCompile(`-+`).ReplaceAllString(s, "-")

	return s
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
