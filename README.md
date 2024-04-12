# EMA — Enhanced Markup for Authors

- utility and library written in pure [Go](https://go.dev/)
- `.ema` file format

## File Format
EMA file is text-based format in UTF-8 encoding.

File starts with header, which is four bytes: `{EMA` (values 123, 69, 77, 65)

There are two modes:

### Text mode
All bytes are interpreted as plain text.
Only special character is `{` (U+007B), which has the following meaning:
- If there is another `{` character immediately after `{`, text mode
  continues and both characters are interpreted as one occurrence of `{` character.
- If there is anything else, content is switcher to **data mode**.
  Character `{` is included as part of **data mode**.

### Data mode
All bytes are interpreted as JSON object (see [json.org](https://json.org)).
After end `}` character of object continues **text mode**.

There is also special form called **command** where starting `{` is followed by
letter `A`-`Z` or `a`-`z` (first command name letter).
Spaces between `{` and command name are allowed.
Command name match regular expression:
```regexp
^[A-Za-z][-0-9A-Za-z]*$
```
After command name, there can be optional spaces followed by:
- end `}` character
- JSON array followed by optional spaces and end `}` character
- JSON object followed by optional spaces and end `}` character

Spaces are all characters in Unicode category *Space Separator*.

Following commands and JSON object are equivalent:
```
{}             = {}
{bold}         = {"use": ["bold"]}
{bold[1]}      = {"use": ["bold", 1]}
{bold{"w": 1}} = {"use": ["bold", {"w": 1}]}
```
