package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T){
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T){
		_, got := dictionary.Search("123")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	
	t.Run("new word", func(t *testing.T){
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
	
	t.Run("new word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new definition")
		assertError(t, err, ErrWordExist)

		assertDefinition(t, dictionary, word, definition)	
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		updateDefinition := "the content after update"
		dictionary.Update(word, updateDefinition)

		assertDefinition(t, dictionary, word, updateDefinition)
	})

	t.Run("new word", func(t *testing.T){
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error.")
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find a world:", err)
	}

	assertString(t, got, definition)	
}