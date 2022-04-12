package blogposts_test

import (
	"errors"
	"io/fs"
	blogposts "reading-files-demo"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestBlogposts(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fs := fstest.MapFS{
			"hello-world.md": {Data: []byte(`Title: Hello, TDD world!
Description: lol
Tags: tdd, go
---
Hello
world!`)},
			// " hello-twitch.md": {Data: []byte("hello, twitch")},
		}
		//When
		posts, err := blogposts.PostsFromFS(fs)

		//Then
		if err != nil {
			t.Fatal(err)
		}
		if len(posts) != len(fs) {
			t.Errorf("Got %d, want %d", len(posts), len(fs))
		}

		expectedFirstPost := blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "lol",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
world!`,
		}
		if !reflect.DeepEqual(posts[0], expectedFirstPost) {
			t.Errorf("got %#v, want %#v", posts[0], expectedFirstPost)
		}
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})
		if err == nil {
			t.Error("expected an error, didnt get one")
		}
	})

}

type FailingFS struct {
}

func (f FailingFS) Open(string) (fs.File, error) {
	return nil, errors.New("always fail")
}

// func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
// 	t.Helper()
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %+v, want %+v", got, want)
// 	}
// }
