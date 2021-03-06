package main

import (
	"fmt"
	"time"
	"strings"
	"web"
)

var rss_head = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">

<channel>
<title>fefemama</title>
<link>http://fettemama.org</link>
<description>THE BEST BLOG IN THE UNIVERSE WRITTEN IN Go :]</description>
<language>de</language>
`

var rss_item = `
<item>
<description><![CDATA[
$descriptioncontent$
]]>
</description>
<title>
$titlecontent$...
</title>
<link>$linkcontent$</link>
<guid>$guidcontent$</guid>
<pubDate>$datecontent$</pubDate>
</item>
`
var rss_footer =`
</channel>
</rss>
`

func renderRSSHeader() string {
	return rss_head
}

func renderRSSItem(post *BlogPost) string {
	s := rss_item
	s = strings.Replace(s, "$descriptioncontent$", post.Content, -1)

	title := htmlstrip(post.Content)
	l := len(title)
	if (l > 64) {
		l = 64
	}

	s = strings.Replace(s, "$titlecontent$", title[0:l], -1)

	link := fmt.Sprintf("http://fettemama.org/post?id=%d", post.Id)
	s = strings.Replace(s, "$linkcontent$", link, -1)
	s = strings.Replace(s, "$guidcontent$", link, -1)
	
	post_date := time.SecondsToLocalTime(post.Timestamp)
	date := post_date.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	s = strings.Replace(s, "$datecontent$", date, -1)
	
	return s
}

func renderRSSFooter() string {
	return rss_footer
}

func RenderRSS(posts *[]BlogPost) string {
	s := renderRSSHeader()
	for _, p := range *posts {
		s += renderRSSItem(&p)
	}
	s += renderRSSFooter()

	return s
}

func rss(ctx *web.Context) string {
    posts,_ := Db.GetLastNPosts(20) //postsForMonth(time.LocalTime())//
    s := RenderRSS(&posts)
	return s
}
