# gopherbadgeimg

Simple, educational tool to transform images to the bitmap format supported by
the 2024 gophercon badger-w (and therefore many TinyGo devices!).

All of the most interesting bits in the code are heavily documented in-line.

Build the tool using go build:

`go build . -o gopherbadgeimg`

Run it as follows to generate a profile image:

`./gopherbadgeimage profile image.jpg`

Or, for the splash screen:

`./gopherbadgeimage splash image.jpg`

This will result in 3 different outputs:

1. In the [Makefile](https://github.com/hybridgroup/badger2040/blob/main/Makefile)
from the badge repo, you're allowed to replace the profile image by dropping in
base64-encoded pixel data.
This is printed to STDOUT by default, and can be pasted directly into the Makefile.
1. The generator will also create a .bin file, which can be embedded into your
code using `go:embed`. An example of this can be found in `main_test.go`.
1. Finally, the generator will also create a `*-generated.go` file similar to the
[tainigo.go](https://github.com/hybridgroup/badger2040/blob/main/tainigo.go) file,
to demonstrate an alternative to go embed.
