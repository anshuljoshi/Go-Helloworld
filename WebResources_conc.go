// This will fetch concurrently various web resources
// It will wait for all of them to be fetched
// It will process them all at once after it is fetched

package main

import (
    "fmt"
    "time"
    "net/http"
)

// Define a struct
type ResponseFromURL struct {
    url      string
    response *http.Response
    err      error
}

// Array of URLs that we are going to fetch
var urls = []string{
    "https://xebia.com/",
    "http://golang.org/",
    "https://www.facebook.com/XebiaIndia",
    "http://blog.xebia.in/2015/09/23/concurrency-in-go/",
}

// Function asynchronousHttpResponse takes an argument the array of URLs

func asynchronousHttpResponse(urls []string) []*ResponseFromURL {
    //
    ch := make(chan *ResponseFromURL)
    // an empty instance of a slice containing pointers to HttpResponse objects.
    responses := []*ResponseFromURL{}
    // iterate through our urls
    for _, url := range urls {
        // define an anonymous function.
        // takes a string argument representing a url.
        go func(url string) {
            fmt.Printf("Wait. Currently fetching %s \n", url)
            // uses the net/http library to fetch the web resource.
            resp, err := http.Get(url)
            // returned data to create an instance of HttpResponse type
            // and send it to the channel.
            ch <- &ResponseFromURL{url, resp, err}
            }(url)
        }

    for {
        select {
        //  case statement checks if something is in the channel.
        //  allocate the data to the r variable
        //  print the resource’s url
        //  append the resource to the slice
        case r := <-ch:
            fmt.Printf("%s was fetched\n", r.url)
            responses = append(responses, r)
            // if all resources are fetched then return.
            if len(responses) == len(urls) {
                return responses
            }
        default:
            fmt.Printf(".")
            // print "." every 25ms.
            time.Sleep(25 * time.Millisecond)
        }
    }
    return responses
}

func main() {
    results := asynchronousHttpResponse(urls)
    for _, result := range results {
        fmt.Printf("\n%s \nstatus: %s\n", result.url, result.response.Status)
    }
}
