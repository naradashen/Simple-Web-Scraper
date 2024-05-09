package main

import (
    "fmt"
    "net/http"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    var url string
    fmt.Println("Enter the website URL you want to scrape:")
    fmt.Scanln(&url)

    // Fetch the webpage
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error fetching webpage:", err)
        return
    }
    defer resp.Body.Close()

    // Parse HTML
    doc, err := html.Parse(resp.Body)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    }

    // Extract directories
    directories := extractDirectories(doc)
    fmt.Println("Directories found:")
    for _, directory := range directories {
        fmt.Println(directory)
    }
}

func extractDirectories(n *html.Node) []string {
    var directories []string
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                href := attr.Val
                // Check if it's a directory (you might need more sophisticated logic here)
                if strings.HasSuffix(href, "/") {
                    directories = append(directories, href)
                }
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        directories = append(directories, extractDirectories(c)...)
    }
    return directories
}
