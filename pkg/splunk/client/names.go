package client

const (
	// Azure token fetch URL
	azureTokenFetchURL = "http://169.254.169.254/metadata/identity/oauth2/token"

	// Azure http header XMS version
	azureHTTPHeaderXmsVersion = "2017-11-09"

	// Azure URL for listing apps
	azureBlobListAppFetchURL = "%s/%s?prefix=%s&restype=container&comp=list&include=snapshots&include=metadata"

	// Header strings
	headerAuthorization      = "Authorization"
	headerCacheControl       = "Cache-Control"
	headerContentEncoding    = "Content-Encoding"
	headerContentDisposition = "Content-Disposition"
	headerContentLanguage    = "Content-Language"
	headerContentLength      = "Content-Length"
	headerContentMD5         = "Content-MD5"
	headerContentType        = "Content-Type"
	headerDate               = "Date"
	headerIfMatch            = "If-Match"
	headerIfModifiedSince    = "If-Modified-Since"
	headerIfNoneMatch        = "If-None-Match"
	headerIfUnmodifiedSince  = "If-Unmodified-Since"
	headerRange              = "Range"
	headerUserAgent          = "User-Agent"
	headerXmsDate            = "x-ms-date"
	headerXmsVersion         = "x-ms-version"
)
