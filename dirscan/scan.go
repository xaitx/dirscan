package dirscan

import (
	"net/http"
	"sync"

	"github.com/xaitx/dirscan/config"
)

type Scan struct {
	Cof      *config.Config      // config
	urlChan  chan string         // url channel
	request  *Request            // request
	w        *sync.WaitGroup     // wait group
	response chan *http.Response // response channel
}

func (s *Scan) Run() {
	// read dictionary
	dict, err := NewDict(s.Cof.DictPath, true)
	if err != nil {
		println(err.Error())
		return
	}
	defer dict.Close()

	// initialize all variables
	s.urlChan = make(chan string)
	s.request = &Request{Method: s.Cof.Method, Proxy: s.Cof.Proxy}
	s.w = &sync.WaitGroup{}
	s.response = make(chan *http.Response)

	s.w.Add(s.Cof.Threads)
	go s.ScanDirThread()

	// read dictionary
	for {
		line, err := dict.ReadLine()
		if err != nil {
			// eof close channel
			close(s.urlChan)
			break
		}
		s.urlChan <- line
	}
	s.w.Wait()
}

// scan directory thread
func (s *Scan) ScanDirThread() {
	// create thread
	for i := 0; i < s.Cof.Threads; i++ {
		go func() {
			for {
				url, ok := <-s.urlChan
				if !ok {
					break
				}
				// scan url
				response, err := s.request.Request(url)
				if err != nil {
					s.response <- response
				}
			}
			s.w.Done()
		}()
	}
}

// print
func (s *Scan) Print() {
	for {
		response, ok := <-s.response
		if !ok {
			break
		}
		println(response.Status)
	}
}
