package cache

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/koki-develop/clive/internal/util"
)

type Store struct {
	rootPath string
	ttl      time.Duration
}

func NewStore(ttl time.Duration) (*Store, error) {
	p, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	return &Store{
		rootPath: filepath.Join(p, "clive"),
		ttl:      ttl,
	}, nil
}

func (s *Store) Get(key string) (*Cache, error) {
	p := s.buildPath(key)
	exists, err := util.FileExists(p)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}

	c, err := s.load(p)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Store) Set(key string, data any) error {
	p := s.buildPath(key)

	f, err := util.CreateFile(p)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	c := New(s.ttl, data)
	if err := c.Write(f); err != nil {
		return err
	}

	return nil
}

func (s *Store) buildPath(key string) string {
	return filepath.Join(s.rootPath, fmt.Sprintf("%s.json", key))
}

func (s *Store) load(p string) (*Cache, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	c, err := s.decode(f)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Store) decode(r io.Reader) (*Cache, error) {
	var c Cache
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
