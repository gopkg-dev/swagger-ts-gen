package generator

import (
	"regexp"
	"strings"
	"unicode"
)

var nonAlnumRegexp = regexp.MustCompile(`[^a-zA-Z0-9]+`)
var versionSegmentRegexp = regexp.MustCompile(`^v\d+([a-zA-Z0-9_-]+)?$`)
var validIdentRegexp = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

func groupFromPath(path string) string {
	segments := splitPath(path)
	if len(segments) == 0 {
		return "default"
	}
	if segments[0] == "api" && len(segments) >= 3 && versionSegmentRegexp.MatchString(segments[1]) {
		return sanitizePathSegment(segments[2])
	}
	return sanitizePathSegment(segments[0])
}

func splitPath(path string) []string {
	raw := strings.Split(path, "/")
	var segments []string
	for _, seg := range raw {
		if seg == "" {
			continue
		}
		segments = append(segments, seg)
	}
	return segments
}

func sanitizePathSegment(segment string) string {
	if segment == "" {
		return "default"
	}
	cleaned := make([]rune, 0, len(segment))
	for _, r := range segment {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' {
			cleaned = append(cleaned, unicode.ToLower(r))
		} else {
			cleaned = append(cleaned, '-')
		}
	}
	normalized := strings.Trim(strings.Trim(string(cleaned), "-"), "_")
	if normalized == "" {
		return "default"
	}
	words := splitWords(normalized)
	if len(words) == 0 {
		return "default"
	}
	var b strings.Builder
	for idx, word := range words {
		if word == "" {
			continue
		}
		if idx == 0 {
			b.WriteString(strings.ToLower(word))
			continue
		}
		b.WriteString(upperFirst(word))
	}
	result := b.String()
	if result == "" {
		return "default"
	}
	if unicode.IsDigit(rune(result[0])) {
		return "group" + result
	}
	return result
}

func sanitizeOperationID(opID string) string {
	trimmed := strings.TrimSpace(opID)
	if trimmed == "" {
		return ""
	}
	if !nonAlnumRegexp.MatchString(trimmed) {
		return lowerFirst(trimmed)
	}
	words := splitWords(trimmed)
	if len(words) == 0 {
		return ""
	}
	return lowerCamel(words)
}

func buildOperationName(method string, path string) string {
	segments := splitPath(path)
	var words []string
	words = append(words, strings.ToLower(method))
	for _, seg := range segments {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			param := strings.TrimSuffix(strings.TrimPrefix(seg, "{"), "}")
			if param != "" {
				words = append(words, "by", param)
			}
			continue
		}
		words = append(words, seg)
	}
	return lowerCamel(words)
}

func lowerCamel(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var b strings.Builder
	for i, w := range words {
		if w == "" {
			continue
		}
		if i == 0 {
			b.WriteString(strings.ToLower(w))
		} else {
			b.WriteString(upperFirst(w))
		}
	}
	result := b.String()
	if result == "" {
		return ""
	}
	if unicode.IsDigit(rune(result[0])) {
		return "op" + result
	}
	return result
}

func upperFirst(word string) string {
	if word == "" {
		return ""
	}
	r := []rune(word)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func lowerFirst(word string) string {
	if word == "" {
		return ""
	}
	r := []rune(word)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func splitWords(input string) []string {
	cleaned := nonAlnumRegexp.ReplaceAllString(input, " ")
	parts := strings.Fields(cleaned)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

func sanitizeTypeName(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "Model"
	}
	trimmed = strings.ReplaceAll(trimmed, ".", " ")
	words := splitWords(trimmed)
	if len(words) == 0 {
		return "Model"
	}
	var b strings.Builder
	for _, w := range words {
		if w == "" {
			continue
		}
		b.WriteString(upperFirst(w))
	}
	result := b.String()
	if result == "" {
		result = "Model"
	}
	if unicode.IsDigit(rune(result[0])) {
		result = "Model" + result
	}
	return result
}

func sanitizeIdentifier(name string) string {
	if validIdentRegexp.MatchString(name) {
		return name
	}
	cleaned := nonAlnumRegexp.ReplaceAllString(name, "_")
	cleaned = strings.Trim(cleaned, "_")
	if cleaned == "" {
		return "value"
	}
	if unicode.IsDigit(rune(cleaned[0])) {
		cleaned = "value_" + cleaned
	}
	if !validIdentRegexp.MatchString(cleaned) {
		return "value"
	}
	return cleaned
}

func isValidIdentifier(name string) bool {
	return validIdentRegexp.MatchString(name)
}
