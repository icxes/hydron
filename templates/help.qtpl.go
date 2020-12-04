// Code generated by qtc from "help.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line help.qtpl:1
package templates

//line help.qtpl:1
import "github.com/bakape/hydron/common"

//line help.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line help.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line help.qtpl:3
func StreamHelpPage(qw422016 *qt422016.Writer) {
//line help.qtpl:4
	streamhead(qw422016, "Help")
//line help.qtpl:4
	qw422016.N().S(`<body><div><b>Drag&Drop</b><article>Files can be imported by dragging and dropping them into either the search page or the imports page.`)
//line help.qtpl:10
	qw422016.N().S(` `)
//line help.qtpl:10
	qw422016.N().S(`Files imported this way will fetch tags from Danbooru.</article></div><hr><div><b>Search</b><article>Tags can include an order:$x parameter where $x is one of:<br>size, width, height, duration, tag_count, random.<br>Prefixing - before $x will reverse the order.<br>Using an order: tag will override the order selected in the dropdown box.</article><article>Tags can be prefixed with - to match a subset that does not include that tag.</article><article>Tags can include prefixed system tags for searching by file metadata:<br>size, width, height, duration, tag_count,<br>followed by one of these comparison operators:<br>>, <, =, >=, <=<br>and a positive integer.<br>e.g. system:width>1920 or system:tag_count=0</article><article>Files can be filtered by the following ratings:<br>safe, questionable, explicit.<br>e.g. rating:safe</article><article>The number of results per page can be controlled with the limit tag. The default amount is`)
//line help.qtpl:50
	qw422016.N().S(` `)
//line help.qtpl:50
	qw422016.N().D(common.PageSize)
//line help.qtpl:50
	qw422016.N().S(`.<br>It takes an integer between 1 and`)
//line help.qtpl:52
	qw422016.N().S(` `)
//line help.qtpl:52
	qw422016.N().D(common.PageSize)
//line help.qtpl:52
	qw422016.N().S(`.<br>e.g. limit:50</article><article>Tags can be prefixed to match a specific tag category like artist (artist:$tag or author:$tag), series (series:$tag or copyright:$tag),`)
//line help.qtpl:58
	qw422016.N().S(` `)
//line help.qtpl:58
	qw422016.N().S(`character (character:$tag), and meta (meta:$tag), where $tag is the suffixing tag.<br>Example meta tags are meta:highres and meta:animated.</article></div><hr><div><b>Keyboard Shortcuts</b><article>The search page can be navigated via keyboard Shortcuts.</article><article>Ctrl+l brings focus to the search bar.</article><article>Ctrl+a toggles the value of all checkboxes.<br>Space toggles the highlighted result's checkbox.</article><article>The arrow keys can be used to move the highlight selection.<br>PgUp and PgDn move the highlight selection greatly up or down respectively.<br>Home moves the highlight selection to the first result in the page, and End moves it to the last result in the page.</article><article>Enter navigates to the highlighted result's image page.</article></div></body>`)
//line help.qtpl:90
}

//line help.qtpl:90
func WriteHelpPage(qq422016 qtio422016.Writer) {
//line help.qtpl:90
	qw422016 := qt422016.AcquireWriter(qq422016)
//line help.qtpl:90
	StreamHelpPage(qw422016)
//line help.qtpl:90
	qt422016.ReleaseWriter(qw422016)
//line help.qtpl:90
}

//line help.qtpl:90
func HelpPage() string {
//line help.qtpl:90
	qb422016 := qt422016.AcquireByteBuffer()
//line help.qtpl:90
	WriteHelpPage(qb422016)
//line help.qtpl:90
	qs422016 := string(qb422016.B)
//line help.qtpl:90
	qt422016.ReleaseByteBuffer(qb422016)
//line help.qtpl:90
	return qs422016
//line help.qtpl:90
}
