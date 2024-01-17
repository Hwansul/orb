package internal

import (
	"github.com/hoehwa/gopkg/bubbletea/table"
)

// Hardcoded common header. Both shortcut command for obsidian and vscode use this data commonly.
// var TableHeader = []string{"Command", "Description", "shortcut"}
var Columns = table.Columns{
	{Title: "Command", Width: 4},
	{Title: "Description", Width: 12},
	{Title: "Keypress", Width: 6},
}

// Hardcoded table body data for vscode.
var TableBodyForVscode = table.Rows{
	{"Quick Fix", "Extract source codes into a new method or function.", "Ctrl+."},
	{"Refactor", "see refactorings without Quick Fixes.", "Ctrl+Shift+R"},
	{"Extract/Inline variable", "Being able to take magic values and give them a name lets you simplify your code quickly.", "Ctrl+."},
	{"Extract method/function", "It is vital to be able to take a section of code and extract functions/methods.", "Ctrl+."},
	{"Rename symbol", "You should be able to confidently rename symbols across files.", "F2"},
	{"IntelliSense", "You can see function definition in the editor.", "auto"},
	{"Go to definition", "You can see function definition in the editor.", "F12 |Ctrl+Click "},
	{"Go to type definition", "You can see function definition in the editor.", "custom "},
	{"Go to Implementation", "You can see type definition in the editor.", "Ctrl+F12"},
	{"Go to Symbol", "You can navigate symbols inside a file.", "Ctrl+Shift+O"},
	{"Open symbol by name", "You can jump to a symbol across files.", "Ctrl+T"},
	{"Peek", "You can navigate between different references in the peeked editor.", "Alt+F12"},
	{"Go to References", "You can navigate between different references.", "Shift+F12"},
}

// Hardcoded table body data for obsidian.
var TableBodyForObsidian = table.Rows{
	{"Quick Fix", "Extract source codes into a new method or function.", "Ctrl+."},
}
