package configs_test

import (
	"errors"
	"testing"

	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppError"
	configs "github.com/F1zm0n/enoki/enoki/utils/pkg/config"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		key   string
		value string
	}{
		{"country", "KZ"},
		{"url", "https://mirror.hoster.kz"},
		{"learn", "bim"},
	}
	conf := make(map[string]string)
	var mirrors []string

	c := &configs.ConfigApp{Config: conf, MirrorLinks: mirrors}

	err := c.ReadConfig("mock.conf")
	if err != nil {
		if errors.Is(err, apperror.ErrReadingConfig) {
			t.Fatalf("%v", err)
		}
		t.Fatalf("error: %v", err)
	}
	t.Log(c.Config)

	if l := len(c.Config); l != len(tests) {
		t.Fatalf("wrong config map length expected=%d got=%d",
			l,
			len(tests),
		)
	}

	for i, test := range tests {
		val, ok := c.Config[test.key]
		if !ok {
			t.Fatalf("test[%d]: key=%s value=%s expected=%s got nothing",
				i,
				test.key,
				test.value,
				test.value,
			)
		}
		if val != test.value {
			t.Fatalf("test [%d]: wrong value key=%s expected=%s got=%s",
				i,
				test.key,
				test.value,
				val,
			)
		}
	}
}

func TestMirrorLink(t *testing.T) {
	tests := []string{"https://mirror.hoster.kz", "bim", "zim", "fim"}

	conf := make(map[string]string)
	var mirrors []string

	c := &configs.ConfigApp{Config: conf, MirrorLinks: mirrors}

	err := c.ReadConfig("mock.conf")
	if err != nil {
		if errors.Is(err, apperror.ErrReadingConfig) {
			t.Fatalf("%v", err)
		}
		t.Fatalf("error: %v", err)
	}

	for i, l := range tests {
		if l != c.MirrorLinks[i] {
			t.Fatalf("test[%d]: wrong value of link expected=%s got=%s",
				i,
				l,
				c.MirrorLinks[i],
			)
		}
	}
}
