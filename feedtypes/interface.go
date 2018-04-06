package feedtypes

import "time"

type feedtype int

const (
	feedtypeRSS feedtype = iota
	feedtypeGopher
)

type newsfeedentry interface {
	getTitle() string
	getAbstract() string
	getDate() time.Time
	getAuthor() string
	getName() string
	getLink() string
}

type newsfeed interface {
	GetEntries(count int) []newsfeedentry
	getLang() language
	getType() feedtype
	getName() string
}

type outletentry struct {
	Title    string
	Type     feedtype
	Language string
	URL      string
}
