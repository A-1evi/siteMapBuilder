# Site Map Builder

## Why Generate XML?

### Purpose of this Program

You're building a site map generator — a tool that:

- Crawls all pages of a website up to a certain depth (using BFS)
- Gathers internal links
- Outputs them as a sitemap XML file

##  What is a Sitemap?

A sitemap is an XML file that lists URLs for a site, like this:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://gophercises.com</loc>
  </url>
  <url>
    <loc>https://gophercises.com/exercises</loc>
  </url>
</urlset>
```

##  Why XML?

- Google and other search engines use XML sitemaps to understand what pages to index.
- XML is a standard format for this — especially with specific tags like `<url>`, `<loc>`, and the correct XML namespace (xmlns).
- You're generating this so search engines can read it.

## What is xmlns?

### xmlns means XML Namespace

```go
const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
```

This tells search engines: "Hey, this XML follows the official sitemap schema."

It's added to your `<urlset>` tag via:

```go
Xmlns string `xml:"xmlns,attr"`
```

###  It ensures compatibility with tools like:

- Google Search Console
- Bing Webmaster Tools
- Other SEO tools

##  Struct Mapping: loc and urlset

You're using Go structs to encode data into XML, like so:

###  loc struct

```go
type loc struct {
    Value string `xml:"loc"`
}
```

This means:

- Each loc instance becomes a `<loc>` tag
- Example:
  ```go
  loc{"https://example.com"} → <loc>https://example.com</loc>
  ```

###  urlset struct

```go
type urlset struct {
    Urls  []loc  `xml:"url"`
    Xmlns string `xml:"xmlns,attr"`
}
```

- Urls: a list of `<url>` entries
- Each url tag wraps a `<loc>` tag
- xmlns attribute gets written into the `<urlset>` root tag

###  Final XML Output

Your Go code:

```go
urlset{
    Xmlns: xmlns,
    Urls: []loc{
        {Value: "https://site.com"},
        {Value: "https://site.com/about"},
    },
}
```

Will be serialized to:

```xml
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://site.com</loc>
  </url>
  <url>
    <loc>https://site.com/about</loc>
  </url>
</urlset>
```

## Summary

| Term | Meaning |
|------|---------|
| XML | Standard format for structured data; used for sitemaps |
| xmlns | XML namespace; required for sitemaps to be valid |
| loc struct | Represents a `<loc>` tag in sitemap XML |
| urlset | The root element of the sitemap (`<urlset>`) |
| xml:"..." | Struct tags in Go that map fields to XML elements or attributes |