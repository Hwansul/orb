package internal

import (
	"github.com/hoehwa/gopkg/bubbletea/table"
)

// Hardcoded common header. Both shortcut command for obsidian and vscode use this data commonly.
// var TableHeader = []string{"Command", "Description", "shortcut"}
var Columns = table.Columns{
	{Title: "Command", Width: 20},
	{Title: "Description", Width: 60},
	{Title: "Keypress", Width: 6},
}

// Hardcoded table body data for vscode.
var TableBodyForVim = table.Rows{
	{"Command Mode", "Get back to command mode", "ESC"},
	{"Insert Mode", "Enter to insert mode", "i"},
	{"Append Mode", "Enter to insert mode at the next to the current cursor", "a"},
	{"Insert Mode at start", "Enter insert mode and put cursor to start of sentence.", "I"},
	{"Append Mode at start", "Enter insert mode and put cursor to end of sentence.", "A"},
	{"Scroll Left", "In command mode, move screen to left", "h"},
	{"Scroll Down", "In command mode, move screen to bottom", "j"},
	{"Scroll Up", "In command mode, move screen to top", "k"},
	{"Scroll Right", "In command mode, move screen to right", "l"},
	{"jump to home", "jump to the start of the line", "0"},
	{"jump to end", "jump to the end of the line", "$"},
	{"jump forwards to", "jump forwards to the start of a word. it is also used with number in front of command.", "w"},
	{"jump backwards to", "jump forwards to the start of a word. e.g. 3b", "b"},
	{"Move Top", "move to top of screen", "H"},
	{"Move Middle", "move to middle of screen", "M"},
	{"Move Buttom", "move to bottom of screen", "L"},
	{"Go First", "Go to the first line of the document", "gg"},
	{"Go Last", "Go to the last line of the document. e.g. 20G", "G"},
	{"Scroll Up Half", "move cursor and screen up 1/2 page", "ctrl+u"},
	{"Scroll Down Half", "move cursor and screen down 1/2 page", "ctrl+d"},
	{"Previous Paragraph", "jump to previous paragraph (or function/block, when editing code)", "{"},
	{"Next Paragraph", "jump to next paragraph (or function/block, when editing code)", "}"},
	{"Cut a character", "delete (cut) character", "x"},
	{"Cut a line", "delete (cut) a line", "dd"},
	{"Copy a line", "yank (copy) a line", "yy"},
	{"Paste", "put (paste) the clipboard after cursor", "p"},
	{"Paster Clipboard", "put (paste) the clipboard after cursor from clipboard", "*p"},
}

var TableBodyForVimCmd = table.Rows{
	{"Repeat last command", "repeat last command", "."},
	{"Undo", "undo", "u"},
	{"Redo", "redo", "Ctrl + r"},
	{"Cutting Example", "Delete a word", "daw"},
	{"Cutting Example", "Delete three word", "d3w"},
	{"Cutting Example", "Delete two lines after cursor", "d2j"},
	{"Cutting Example", "Delete two lines before cursor", "d3k"},
	{"Cutting Example", "Delete contents in curly brace", "di{"},
	{"Cutting Example", "Delete contents in brace", "di("},
	{"Cutting Example", "Delete contents in brace and brace itself.", "da("},
	{"Cutting Example", "Delete contents in single quote.", "di'"},
	{"Cutting Example", "Delete contents in single quote and quote itself.", "da'"},
	{"Cutting Example", "Delete contents before bracket", "dt("},
	{"Cutting Example", "Delete contents before bracket and bracket itself.", "df("},
	{"Cutting Example", "Delete contents from current cursor to matched contents", "d/("},
	{"Change Example", "Change contnets in square bracket. After deleting, vim'll be set to insert mode.", "ci["},
	{"Search forward", "Search some keyword to forward", "/"},
	{"Search backward", "Search some keyword to backward", "?"},
	{"Keep searching", "Repeat search to direction before you did at last", "n"},
	{"Visual Mode Example", "Visual select a word", "vaw"},
	{"Visual Mode Example", "Multiple block selection.", "ctrl + v"},
}

// Hardcoded table body data for vscode.
var TableBodyForSurfingKeys = table.Rows{
	{"Scroll half page up", "Scroll screen to top as half of current page", "e"},
	{"Scroll half page down", "Scroll screen to bottom as half of current page", "d"},
	{"Open a URL", "Open up search console in current page", "t"},
	{"Scroll left", "Scroll screen to left", "h"},
	{"Scroll down", "Scroll screen to bottom", "j"},
	{"Scroll up", "Scroll screen to top", "k"},
	{"Scroll right", "Scroll screen to right", "l"},
	{"Copy URL", "Copy current page's URL", "yy"},
	{"Copy elements", "Yank text of an element", "yv"},
	{"Go to last tab", "Go to last used tab", "Ctrl + 6"},
	{"Last tab", "Go to last used tab", "Ctrl + 6"},
	{"Last tab", "Go to edit box", "i"},
	{"Select element", "Click a element in current page", "q"},
	{"Copy form data", "Copy form data in JSON on current page", "yf"},
}

// Hardcoded table body data for vscode.
var TableBodyForVscode = table.Rows{
	{"Toggle Line Comment", "In javascript, add dobule slash(//) in source code.", "Ctrl+/"},
	{"Toggle Block Comment", "In javascript, wrap with slack and asterisk to block.", "Shift+Alt+A"},
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
