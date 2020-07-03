package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Known Word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test."}
		got, _ := dictionary.Search("test")
		want := "this is a test."
		assertStrings(t, got, want)
	})
	t.Run("Unknow Word", func(t *testing.T) {
		dictionary := Dictionary{}
		_, got := dictionary.Search("unknown")
		want := ErrNotFound
		assertStrings(t, got.Error(), want.Error())
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s while want %s ", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("Test Add : ", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Test Update : ", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		err := dictionary.Update(word, definition)
		want := ErrDoesNotExists
		if err == want {
			t.Errorf("got %v while want %v", err, want)
		}
	})
}

func assertDefinition(t *testing.T, dictinary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictinary.Search(word)
	if err != nil {
		t.Fatal("Should have found the word : ", err)
	}
	if got != definition {
		t.Errorf("got %s while want %s", got, definition)
	}
}
