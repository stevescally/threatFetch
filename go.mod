module github.com/stevescally/threatFetch

go 1.22.0

replace github.com/stevescally/threatFetch/threatview => ./threatview

replace github.com/stevescally/threatFetch/greensnow => ./greensnow

replace github.com/stevescally/threatFetch/rescure => ./rescure

require (
	github.com/pterm/pterm v0.12.79
	github.com/stevescally/threatFetch/greensnow v0.0.0-00010101000000-000000000000
	github.com/stevescally/threatFetch/rescure v0.0.0-00010101000000-000000000000
	github.com/stevescally/threatFetch/threatview v0.0.0-00010101000000-000000000000
)

require (
	atomicgo.dev/cursor v0.2.0 // indirect
	atomicgo.dev/keyboard v0.2.9 // indirect
	atomicgo.dev/schedule v0.1.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/lithammer/fuzzysearch v1.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/term v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
