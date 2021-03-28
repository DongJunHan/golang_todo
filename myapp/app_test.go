package myapp

import(
	"testing"
	"net/http"
	"WEB-INF/golang_todo/myapp"
	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T){
	assert := assert.New(t)
	ts := httptest.NewServer(MakeHandler())
	defer ts.Close()


	res, err := http.PostForm(ts.URL+"/todos",url.Values{"name":{"Test todo"}} )
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
}
