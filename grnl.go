package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type Entry struct {
	date  time.Time
	title string
	body  string
}

func parse(text string) (error, []*Entry) {
	return nil, []*Entry{}
}

func main() {
	// Comments are flags I should implement

	// parser.add_argument('-v', '--version', dest='version', action="store_true", help="prints version information and exits")
	versionFlag := flag.Bool("version", false, "Prints version information and exits.")
	// parser.add_argument('-ls', dest='ls', action="store_true", help="displays accessible journals")
	// parser.add_argument('-d', '--debug', dest='debug', action='store_true', help='execute in debug mode')

	// reading = parser.add_argument_group('Reading', 'Specifying either of these parameters will display posts of your journal')
	// reading.add_argument('-from', dest='start_date', metavar="DATE", help='View entries after this date')
	// reading.add_argument('-until', '-to', dest='end_date', metavar="DATE", help='View entries before this date')
	// reading.add_argument('-on', dest='on_date', metavar="DATE", help='View entries on this date')
	// reading.add_argument('-and', dest='strict', action="store_true", help='Filter by tags using AND (default: OR)')
	// reading.add_argument('-starred', dest='starred', action="store_true", help='Show only starred entries')
	// reading.add_argument('-n', dest='limit', default=None, metavar="N", help="Shows the last n entries matching the filter. '-n 3' and '-3' have the same effect.", nargs="?", type=int)

	// exporting = parser.add_argument_group('Export / Import', 'Options for transmogrifying your journal')
	// exporting.add_argument('--short', dest='short', action="store_true", help='Show only titles or line containing the search tags')
	// exporting.add_argument('--tags', dest='tags', action="store_true", help='Returns a list of all tags and number of occurences')
	// exporting.add_argument('--export', metavar='TYPE', dest='export', choices=['text', 'txt', 'markdown', 'md', 'json'], help='Export your journal. TYPE can be json, markdown, or text.', default=False, const=None)
	// exporting.add_argument('-o', metavar='OUTPUT', dest='output', help='Optionally specifies output file when using --export. If OUTPUT is a directory, exports each entry into an individual file instead.', default=False, const=None)
	// exporting.add_argument('--encrypt', metavar='FILENAME', dest='encrypt', help='Encrypts your existing journal with a new password', nargs='?', default=False, const=None)
	// exporting.add_argument('--decrypt', metavar='FILENAME', dest='decrypt', help='Decrypts your journal and stores it in plain text', nargs='?', default=False, const=None)
	// exporting.add_argument('--edit', dest='edit', help='Opens your editor to edit the selected entries.', action="store_true")

	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the points with e.g. `*wordPtr`
	// to get the actual option values.
	fmt.Println("Version:", *versionFlag)
	fmt.Println("tail:", len(flag.Args()))

	parse(strings.Join(flag.Args(), " "))
}
