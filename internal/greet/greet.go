// Package greet builds friendly greeting messages.
//
// ENG: Provides Hello, which greets a person by name and falls back to
// "stranger" when no name is given.
//
// LAO: ສ້າງຂໍ້ຄວາມທັກທາຍ — ທັກທາຍຕາມຊື່, ຖ້າບໍ່ມີຊື່ຈະໃຊ້ "stranger".
package greet

import "strings"

// Hello returns a greeting message for the given name.
//
// ENG: Trims surrounding whitespace from name and returns "Hello, <name>!".
// Returns "Hello, stranger!" if name is empty or contains only whitespace.
//
// LAO: ສົ່ງຄືນຄຳທັກທາຍ "Hello, <name>!" — ຖ້າຊື່ຫວ່າງເປົ່າ ຈະສົ່ງຄືນ "Hello, stranger!".
func Hello(name string) string {
	return "Hello, " + displayName(name) + "!"
}

// displayName returns the name to show in a greeting.
//
// ENG: Trims surrounding whitespace from name. Returns "stranger" if nothing
// is left, otherwise the trimmed name.
//
// LAO: ຕັດຊ່ອງວ່າງຫົວທ້າຍອອກ — ຖ້າບໍ່ມີຫຍັງເຫຼືອ ຈະສົ່ງຄືນ "stranger".
func displayName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "stranger"
	}
	return name
}
