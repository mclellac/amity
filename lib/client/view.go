package client

import (
	"fmt"
	"strconv"

	"github.com/mclellac/amity/lib/ui"
)

// ListView is the view used by the amity cli client when listing
// posts.
func (client *Client) ListView(id int32, title string) {
	idStr := strconv.FormatInt(int64(id), 10)

	fmt.Printf("%s[%s%s%s]\t%s%s%s\n",
		ui.LightBlue,
		ui.LightCyan,
		idStr,
		ui.LightBlue,
		ui.LightCyan,
		title,
		ui.Reset)
}

// ReadView is the view used for the amity cli client when reading posts.
func (client *Client) ReadView(id int32, title string, article string) {
	idStr := strconv.FormatInt(int64(id), 10)

	fmt.Printf("%s[%s]%s\t%+v%s\n", ui.Blue, idStr, ui.LightCyan, title, ui.Reset)
	ui.DrawDivider()
	fmt.Printf("%s%+v%s\n", ui.LightGreen, article, ui.Reset)
}
