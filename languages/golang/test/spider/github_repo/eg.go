package main

/*
Packages must be imported:
    "core/common/page"
    "core/spider"
Pckages may be imported:
    "core/pipeline": scawler result persistent;
    "github.com/PuerkitoBio/goquery": html dom parser.
*/
import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"github.com/sirupsen/logrus"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to Page.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()

	if strings.Contains(p.GetPageItems().GetRequest().Url, "tab=repositories") {
		var urls []string
		query.Find("h3[class=wb-break-all] > a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			urls = append(urls, "https://github.com"+href)
		})
		logrus.Info("urls: ", urls)
		// these urls will be saved and crawed by other coroutines.
		p.AddTargetRequests(urls, "html")
	} else {
		// s := query.Find(".AppHeader-context-compact-parentItem").First().Text()
		// fmt.Println("first text: ", s)
		q := query.Find(".Truncate-text")
		fmt.Println("q: ", q.Text(), q.Length())
		fmt.Println("node: ", q.Nodes)
		fmt.Println("contents: ", q.Contents())
		fmt.Println(q.Contents().Children())
		fmt.Println(q.Html())
		// fmt.Println("contents: )
		// nv, exist := name.Attr("innerHTML")
		// fmt.Println(nv, exist)
		// name = strings.Trim(name, " \t\n")
		repository := query.Find(".AppHeader-context-compact-mainItem").Text()
		repository = strings.Trim(repository, " \t\n")
		//readme, _ := query.Find("#readme").Html()
		// if name == "" {
		// 	p.SetSkip(true)
		// }
		// the entity we want to save by Pipeline
		// p.AddField("author", name)
		p.AddField("project", repository)
		// logrus.Infof("author: %s, project: %s\r\n", name, repository)
		panic("")
	}

	//p.AddField("readme", readme)
}

func (this *MyPageProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

func main() {
	// Spider input:
	//  PageProcesser ;
	//  Task name used in Pipeline for record;
	req := request.NewRequest("https://github.com/feng6917?tab=repositories", "html", "", "GET", "", nil, nil, nil, nil)
	req.AddProxyHost("http://127.0.0.1:11809")
	spider.NewSpider(NewMyPageProcesser(), "TaskName").
		// AddUrl("https://github.com/hu17889?tab=repositories", "html"). // Start url, html is the responce type ("html" or "json" or "jsonp" or "text")
		AddRequest(req).
		AddPipeline(pipeline.NewPipelineConsole()). // Print result on screen
		SetThreadnum(3).                            // Crawl request by three Coroutines
		Run()
}
