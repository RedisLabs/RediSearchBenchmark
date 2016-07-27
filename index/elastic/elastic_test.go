package elastic

import (
	"fmt"
	"strings"
	"testing"

	"github.com/RedisLabs/RediSearchBenchmark/index"
	"github.com/RedisLabs/RediSearchBenchmark/query"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	// todo: run redisearch automatically
	md := index.NewMetadata().AddField(index.NewTextField("title", 1.0)).
		AddField(index.NewNumericField("score"))

	idx, err := NewIndex("http://localhost:9200", "testung", md)
	assert.NoError(t, err)

	docs := []index.Document{}
	for i := 0; i < 100; i++ {
		docs = append(docs, index.NewDocument(fmt.Sprintf("doc%d", i), 0.1).Set("title", "hello world").Set("body", "lorem ipsum foo bar"))

		//index.NewDocument("doc2", 1.0).Set("title", "foo bar hello").Set("score", 2),
	}

	//	assert.NoError(t, idx.Drop())
	//	assert.NoError(t, idx.Create())

	//	assert.NoError(t, idx.Index(docs, nil))

	q := query.NewQuery("doc", "hello world")
	docs, total, err := idx.Search(*q)

	assert.NoError(t, err)
	assert.True(t, total == 100)
	assert.Len(t, docs, 10)
	assert.True(t, strings.HasPrefix(docs[0].Id, "doc"))
	assert.Equal(t, docs[0].Properties["title"], "hello world")

}