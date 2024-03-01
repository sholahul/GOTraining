// Code generated by usestdlibvars, DO NOT EDIT.

package http_test

import "net/http"

var (
	_ = "connect"
	_ = "delete"
	_ = "get"
	_ = "head"
	_ = "options"
	_ = "patch"
	_ = "post"
	_ = "put"
	_ = "trace"
)

const (
	_ = "connect"
	_ = "delete"
	_ = "get"
	_ = "head"
	_ = "options"
	_ = "patch"
	_ = "post"
	_ = "put"
	_ = "trace"
)

var (
	_ = http.Request{Method: "connect"} // want `"connect" can be replaced by http\.MethodConnect`
	_ = http.Request{Method: "delete"}  // want `"delete" can be replaced by http\.MethodDelete`
	_ = http.Request{Method: "get"}     // want `"get" can be replaced by http\.MethodGet`
	_ = http.Request{Method: "head"}    // want `"head" can be replaced by http\.MethodHead`
	_ = http.Request{Method: "options"} // want `"options" can be replaced by http\.MethodOptions`
	_ = http.Request{Method: "patch"}   // want `"patch" can be replaced by http\.MethodPatch`
	_ = http.Request{Method: "post"}    // want `"post" can be replaced by http\.MethodPost`
	_ = http.Request{Method: "put"}     // want `"put" can be replaced by http\.MethodPut`
	_ = http.Request{Method: "trace"}   // want `"trace" can be replaced by http\.MethodTrace`
)

var (
	_ = &http.Request{Method: "connect"} // want `"connect" can be replaced by http\.MethodConnect`
	_ = &http.Request{Method: "delete"}  // want `"delete" can be replaced by http\.MethodDelete`
	_ = &http.Request{Method: "get"}     // want `"get" can be replaced by http\.MethodGet`
	_ = &http.Request{Method: "head"}    // want `"head" can be replaced by http\.MethodHead`
	_ = &http.Request{Method: "options"} // want `"options" can be replaced by http\.MethodOptions`
	_ = &http.Request{Method: "patch"}   // want `"patch" can be replaced by http\.MethodPatch`
	_ = &http.Request{Method: "post"}    // want `"post" can be replaced by http\.MethodPost`
	_ = &http.Request{Method: "put"}     // want `"put" can be replaced by http\.MethodPut`
	_ = &http.Request{Method: "trace"}   // want `"trace" can be replaced by http\.MethodTrace`
)

var (
	_, _ = http.NewRequest("connect", "", http.NoBody) // want `"connect" can be replaced by http\.MethodConnect`
	_, _ = http.NewRequest("delete", "", http.NoBody)  // want `"delete" can be replaced by http\.MethodDelete`
	_, _ = http.NewRequest("get", "", http.NoBody)     // want `"get" can be replaced by http\.MethodGet`
	_, _ = http.NewRequest("head", "", http.NoBody)    // want `"head" can be replaced by http\.MethodHead`
	_, _ = http.NewRequest("options", "", http.NoBody) // want `"options" can be replaced by http\.MethodOptions`
	_, _ = http.NewRequest("patch", "", http.NoBody)   // want `"patch" can be replaced by http\.MethodPatch`
	_, _ = http.NewRequest("post", "", http.NoBody)    // want `"post" can be replaced by http\.MethodPost`
	_, _ = http.NewRequest("put", "", http.NoBody)     // want `"put" can be replaced by http\.MethodPut`
	_, _ = http.NewRequest("trace", "", http.NoBody)   // want `"trace" can be replaced by http\.MethodTrace`
)

var (
	_, _ = http.NewRequestWithContext(nil, "connect", "", http.NoBody) // want `"connect" can be replaced by http\.MethodConnect`
	_, _ = http.NewRequestWithContext(nil, "delete", "", http.NoBody)  // want `"delete" can be replaced by http\.MethodDelete`
	_, _ = http.NewRequestWithContext(nil, "get", "", http.NoBody)     // want `"get" can be replaced by http\.MethodGet`
	_, _ = http.NewRequestWithContext(nil, "head", "", http.NoBody)    // want `"head" can be replaced by http\.MethodHead`
	_, _ = http.NewRequestWithContext(nil, "options", "", http.NoBody) // want `"options" can be replaced by http\.MethodOptions`
	_, _ = http.NewRequestWithContext(nil, "patch", "", http.NoBody)   // want `"patch" can be replaced by http\.MethodPatch`
	_, _ = http.NewRequestWithContext(nil, "post", "", http.NoBody)    // want `"post" can be replaced by http\.MethodPost`
	_, _ = http.NewRequestWithContext(nil, "put", "", http.NoBody)     // want `"put" can be replaced by http\.MethodPut`
	_, _ = http.NewRequestWithContext(nil, "trace", "", http.NoBody)   // want `"trace" can be replaced by http\.MethodTrace`
)