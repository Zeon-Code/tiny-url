package cache

import (
	"fmt"
	"strings"
)

type CacheKey struct {
	contexts []string
	parts    []string
}

func NewCacheKey(contexts ...string) CacheKey {
	return CacheKey{contexts: contexts}
}

func (k CacheKey) With(parts ...any) CacheKey {
	next := make([]string, 0, len(k.parts)+len(parts))
	next = append(next, k.parts...)

	for _, p := range parts {
		next = append(next, fmt.Sprint(p))
	}

	return CacheKey{contexts: k.contexts, parts: next}
}

func (k CacheKey) String() string {
	parts := strings.Join(k.parts, ":")
	contexts := strings.Join(k.contexts, "-")

	if len(contexts) > 0 && len(parts) > 0 {
		return fmt.Sprintf("%s:%s", contexts, parts)
	}

	return fmt.Sprintf("%s%s", contexts, parts)
}
