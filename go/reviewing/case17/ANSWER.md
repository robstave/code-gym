# Answer: Health check

**Difficulty: Medium**

`http.Client.Get` returns an error for problems making the request, not for
non-200 HTTP status codes. The test server returns 503, but the function
returns nil and the sample prints `Healthy: true`; it should report an error.

After checking the request error and arranging to close the response body,
check `response.StatusCode == http.StatusOK`. Return an informative error
containing the unexpected status otherwise. Keep closing the body on both
success and unexpected-status paths.
