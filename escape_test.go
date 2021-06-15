package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscapeHTMLTag(t *testing.T) {
	out := EscapeHTMLTag("<script>alert('hello')</script>")
	assert.Equal(t, "&lt;script&gt;alert(&#39;hello&#39;)&lt;/script&gt;", out, "The html string escape test case is not successful.")
}
