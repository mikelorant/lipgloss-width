# lipgloss-width

## Issue

Determining displayed width using `lipgloss.Width()` will report incorrect
results in some cases.

## Test case

The test cases take a string and add a 1 character margin to the right side.

## Regression

The width detection has always had problems however with the change from commit
`f754c404f6dd63e783b8fde8f245306581335932` the functionality has changed. This
change includes an update to the `muesil/termenv` package which modified
terminal detection. The last commit before this regression was
`14eeaa6ffac0d2dec3f6b1806fbb1f2f5b347613`.

## Outcome modifiers

### ANSI

Using bold adds ANSI codes which increases the string length which allows the
right margin to be added correctly.

If the terminal detection identified no ANSI capabilities, the characters will
be stripped causing the string to only contain whitespace characters.

### Termenv profile detection

While `termenv` attempts to detect the capabilities of the terminal, it can be
overriden which changes the outcome.

`export NO_COLOR=1` forces the terminal to be `ascii` only. ANSI codes such as
bold will be stripped. This will then match the results from `go test` which is
marked as having minimal terminal capabilities.

`export CLICOLOR_FORCE=1` will make `go test` match the capabilities of a fully
featured terminal. This will match the output of running the test application
directly.
