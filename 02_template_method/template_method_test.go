package template_method

import "testing"

func TestHTTP(t *testing.T) {
	var downloader Downloader
	downloader = newHTTPDownloader()
	downloader.Download("http.uri")
}

func TestFTP(t *testing.T) {
	var downloader Downloader
	downloader = newFTPDownloader()
	downloader.Download("ftp.uri")
}
