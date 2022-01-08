package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/deckarep/gosx-notifier"
	"io"
	"os"
	"strings"
)

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// FileHash create sha256 file hash for later comparison
func FileHash(file string) (string, error) {
	if _, err := os.Stat(file); err != nil {
		return "", err
	}
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	hashString := hex.EncodeToString(h.Sum(nil))
	return hashString, err
}

func LessString(v interface{}) func(i, j int) bool {
	s := *v.(*[]string)
	return func(i, j int) bool { return s[i] < s[j] }
}

// Notification macOS Notification using github.com/deckarep/gosx-notifier
func Notification(message string) error {
	note := gosxnotifier.NewNotification(message)
	note.Title = "Jnana"
	note.Sound = gosxnotifier.Default

	if err := note.Push(); err != nil {
		return err
	}
	return nil
}
func ArrayRemoveEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func StringContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func StringSplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

// StringForSQLite  prepare string for SQLite FTS query
// make all terms wildcard
func StringForSQLite(query string) *string {
	var querySlice []string
	queryOperators := []string{"and", "or", "not", "AND", "OR", "NOT"}

	slc := strings.Split(query, " ")
	for i := range slc {
		term := slc[i]
		if strings.HasPrefix(term, "-") {
			// exclude terms beginning with '-', change to 'NOT [term]'
			querySlice = append(querySlice, "NOT "+term[1:]+"*")
		} else if StringInSlice(term, queryOperators) {
			// auto capitalize operators 'and', 'or', 'not'
			querySlice = append(querySlice, strings.ToUpper(term))
		} else if strings.Contains(term, ".") || strings.Contains(term, "-") {
			// quote terms containing dot
			querySlice = append(querySlice, "\""+term+"*\"")
		} else {
			querySlice = append(querySlice, term+"*")
		}
	}
	s := strings.Join(querySlice, " ")
	return &s
}

// StringInSlice Test if string is included in slice
func StringInSlice(a string, list []string) bool {
	for i := range list {
		if list[i] == a {
			return true
		}
	}
	return false
}
