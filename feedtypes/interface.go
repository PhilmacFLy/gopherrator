package feedtypes

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/mmcdole/gofeed"
)

type feedtype int

const (
	feedtypeRSS feedtype = iota
	feedtypeGopher
)

type newsfeedentry struct {
	title    string
	abstract string
	date     time.Time
	author   string
	link     string
	lang     language
	name     string
}

func (n *newsfeedentry) getTitle() string {
	return n.title
}

func (n *newsfeedentry) getAbstract() string {
	return n.title
}

func (n *newsfeedentry) getDate() time.Time {
	return n.date
}
func (n *newsfeedentry) getAuthor() string {
	return n.author
}
func (n *newsfeedentry) getLink() string {
	return n.link
}

func (n *newsfeedentry) getLang() language {
	return n.lang
}

func (n *newsfeedentry) getName() string {
	return n.name
}

type Newsfeed interface {
	getEntries(count int) ([]newsfeedentry, error)
	getType() feedtype
	getName() string
	getLang() language
}

type rssfeed struct {
	name string
	lang language
	typ  feedtype
	url  string
	feed gofeed.Feed
}

func (r *rssfeed) getInfofromFeed(url string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return err
	}
	r.feed = &feed
	r.name = feed.Title
	r.url = url

	return nil
}

func (r *rssfeed) getName() string {
	return r.name
}

func (r *rssfeed) getType() feedtype {
	return r.typ
}

func (r *rssfeed) getLang() language {
	return r.lang
}

func (r *rssfeed) getEntries(count int) ([]newsfeedentry, error) {
	var result []newsfeedentry
	i := 0
	for (i < len(r.feed.Items)) && (i < count) {
		it := r.feed.Items[i]
		nfe := newsfeedentry{}
		nfe.title = it.Title
		nfe.date = *it.PublishedParsed
		nfe.author = it.Author.Name
		nfe.abstract = it.Description
		nfe.link = it.Link
		nfe.lang = r.lang
		result = append(result, nfe)
	}
	fmt.Println(result)
	return result, nil
}

//Outletentry holds information about an Outlet Source
type Outletentry struct {
	Title    string
	Type     feedtype
	Language language
	URL      string
}

//LoadOutlets reads a outletentry array from an io.Reader
func LoadOutlets(input io.Reader) ([]Outletentry, error) {
	var result []Outletentry
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&result)
	return result, err
}

//ParseOutlets parses all Outlets and writes the gopher files
func ParseOutlets(outlets []Outletentry) ([]Newsfeed, error) {
	var nfs []Newsfeed
	for _, o := range outlets {
		switch o.Type {
		case feedtypeRSS:
			r := rssfeed{}
			r.typ = o.Type
			r.lang = o.Language
			err := r.getInfofromFeed(o.URL)
			if err != nil {
				return nfs, fmt.Errorf("Error with Outlet %s:%s", o.Title, err)
			}
			nfs = append(nfs, &r)
		}
	}
	return nfs, nil
}

func WriteEntriestoGopher(nfs []Newsfeed, entrycount int, gpath string) error {
	var nfes []newsfeedentry
	for _, n := range nfs {
		nfe, err := n.getEntries(entrycount)
		if err != nil {
			return fmt.Errorf("Error getting Entries from Outlet %s: %s", n.getName(), err)
		}
		nfes = append(nfes, nfe...)
	}
	return nil
}
