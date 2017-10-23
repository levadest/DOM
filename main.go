package crawler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"lib/config"

	"net/http"
	"net/url"
	"strconv"
	//"net/http/httputil"
)

type (
	CrawlParam struct {
		Key   string
		Value string
	}

	Crawler struct {
		config *config.Config
	}

	SedoDocument struct {
		Diadoc         *Documents
		Spheradoc      *Doc
		Compared       bool
		Id             string
		Function       string
		Step           int
		RabbitDocument *RabbitDocumentEntity
		Sedo           interface{}
	}
	RabbitDocumentEntity struct {
		ConnectorType                      string
		Success                            bool
		FailReason                         string
		DocumentId                         string
		EntityId                           string
		UniversalTransferDocument          string
		UniversalTransferDocumentSignature string
		Invoices                           []Invoice
	}

	Invoice struct {
		EntityId                string
		FileName                string
		Type                    string
		InvoiceContent          string
		InvoiceContentSignature string
	}

	StoredDocumentFlow struct {
		DocumentId string
		StepId     string
		RabbitDocumentEntity
	}
)

const StatusCodeSuccess int = 200

func NewCrawler() *Crawler {
	return &Crawler{
		config: config.NewConfig(config.Default),
	}

}

func (c *Crawler) Sign(body string) (string, error) {

	externalSignerUrl := c.config.GetString("externalSigner")

	signedContent, err := http.PostForm(externalSignerUrl, url.Values{"content": {body}})

	if err != nil {
		return "", err
	}
	signedContentBody, err := ioutil.ReadAll(signedContent.Body)
	if err != nil {
		return "", err
	}
	signedContent.Body.Close()
	signedContentBodyString, err := strconv.Unquote(string(signedContentBody))

	return signedContentBodyString, nil

}

func (c *Crawler) Crawl(method, urlString string, body []byte, headers *[]CrawlParam) (respBody []byte, err error) {

	req, err := http.NewRequest(method, urlString, bytes.NewBuffer(body))

	if *headers != nil && len(*headers) > 0 {
		for _, header := range *headers {

			req.Header.Add(header.Key, header.Value)
		}

	}

	//log, err := httputil.DumpRequest(req, true)
	//fmt.Print(string(log))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	responce, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != StatusCodeSuccess {

		return nil, fmt.Errorf("%s method responce code %d, error=%s", urlString, resp.StatusCode, string(responce))
	}

	if err != nil {
		return nil, err
	}

	return responce, nil

}
