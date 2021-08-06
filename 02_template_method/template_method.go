package template_method

import "fmt"

// 父类定义流程
// 子类实现方法

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

type HTTPDownloader struct {
	*template
}

func newHTTPDownloader() *HTTPDownloader {
	downloader := &HTTPDownloader{}
	temp := newTemplate(downloader)
	downloader.template = temp

	return downloader
}

func (d HTTPDownloader) download() {
	fmt.Printf("HTTPDownloader download,uri: %s \n", d.uri)
}

func (d HTTPDownloader) save() {
	fmt.Print("HTTPDownloader saving\n")
}

type FTPDownloader struct {
	*template
}

func newFTPDownloader() *FTPDownloader {
	downloader := &FTPDownloader{}
	temp := newTemplate(downloader)
	downloader.template = temp

	return downloader
}

func (d FTPDownloader) download() {
	fmt.Printf("FTPDownloader download,uri: %s \n", d.uri)
}

func (d FTPDownloader) save() {
	fmt.Print("FTPDownloader saving\n")
}
