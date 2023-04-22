package main

import "testing"

func TestHello(t *testing.T) { //I learned that test functions need to start with a capital T!!
	t.Run("testing with a person's name", func(t *testing.T) {
		got := hello("tom", "")
		want := "Hello, tom!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("testing with no input", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("spanish input", func(t *testing.T) {
		got := hello("Elon", "Spanish")
		want := "Hola, Elon!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("french input", func(t *testing.T) {
		got := hello("Bill", "French")
		want := "Bonjour, Bill!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("dutch input", func(t *testing.T) {
		got := hello("Jack", "Dutch")
		want := "Hallo, Jack!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
