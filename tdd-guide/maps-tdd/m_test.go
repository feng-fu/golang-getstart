package m

import "testing"

var testValue = "this is just a test"

func TestSearch(t *testing.T) {
	dictonary := Dictonary{"test": testValue}

	t.Run("Known word", func(t *testing.T) {
		got, err := dictonary.Search("test")
		assertNoError(t, err)
		assertStrings(t, got, testValue)
	})

	t.Run("Unkonwn word", func(t *testing.T) {
		_, err := dictonary.Search("blalba")
		assertError(t, err, ErrorNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictonary := Dictonary{"test": testValue}
	t.Run("Add word", func(t *testing.T) {
		err := dictonary.Add("hello", "moto")
		assertNoError(t, err)
		got, e := dictonary.Search("hello")
		assertNoError(t, e)
		assertStrings(t, got, "moto")
	})

	t.Run("Repeat add word", func(t *testing.T) {
		err := dictonary.Add("test", "this is a repeat add")
		assertError(t, err, ErrorRepeatDefine)
		got, err := dictonary.Search("test")
		assertNoError(t, err)
		assertStrings(t, got, testValue)
	})
}

func TestUpdate(t *testing.T) {
	dictonary := Dictonary{"test": "this is just a test"}

	t.Run("Update word", func(t *testing.T) {
		err := dictonary.Update("test", "Jack")
		assertNoError(t, err)
		got, e := dictonary.Search("test")
		assertNoError(t, e)
		assertStrings(t, got, "Jack")
	})

	t.Run("Update undefined word", func(t *testing.T) {
		err := dictonary.Update("test1", "Jack")
		assertError(t, err, ErrorUpdateNotFound)
		_, e := dictonary.Search("test1")
		assertError(t, e, ErrorNotFound)
	})
}

func TestRemove(t *testing.T) {
	dictonary := Dictonary{"test": "this is just a test"}

	t.Run("remove word", func(t *testing.T) {
		err := dictonary.Remove("test")
		assertNoError(t, err)
		_, e := dictonary.Search("test")
		assertError(t, e, ErrorNotFound)
	})

	t.Run("remove undefined word", func(t *testing.T) {
		err := dictonary.Remove("test1")
		assertError(t, err, ErrorRemoveNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("Should find addded word:", got)
	}
}

func assertError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("want an error but got nil")
	}

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
