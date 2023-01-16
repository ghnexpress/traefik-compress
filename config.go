package traefik_compress

type Compress struct {
	// ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing.
	// `application/grpc` is always excluded.
	ExcludedContentTypes []string `json:"excludedContentTypes,omitempty" toml:"excludedContentTypes,omitempty" yaml:"excludedContentTypes,omitempty" export:"true"`
	// MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed.
	// Default: 1024.
	MinResponseBodyBytes int `json:"minResponseBodyBytes,omitempty" toml:"minResponseBodyBytes,omitempty" yaml:"minResponseBodyBytes,omitempty" export:"true"`
}
