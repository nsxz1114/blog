package search_ser

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"go.uber.org/zap"
)

type ArticleSearchResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func SearchDocumentTerm(field, key string) (result []models.Article) {
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			Term: map[string]types.TermQuery{
				field: {Value: key},
			},
		}).
		Do(context.Background())
	if err != nil {
		global.Log.Error("search document failed, err", zap.Error(err))
		return result
	}
	for _, hit := range resp.Hits.Hits {
		var item models.Article
		err := json.Unmarshal(hit.Source_, &item)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}
		result = append(result, item)
	}
	return result
}

func SearchDocumentTerms(field string, key []string) (result []models.Article) {
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			Terms: &types.TermsQuery{
				TermsQuery: map[string]types.TermsQueryField{
					field: key,
				},
			},
		}).
		Do(context.Background())
	if err != nil {
		global.Log.Error("search document failed, err", zap.Error(err))
		return result
	}
	for _, hit := range resp.Hits.Hits {
		var item models.Article
		err := json.Unmarshal(hit.Source_, &item)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}

		result = append(result, item)
	}
	return result
}

func SearchDocumentMultiMatch(fields []string, key string, pageInfo models.PageInfo) (result []models.Article) {
	form := (pageInfo.Page - 1) * pageInfo.PageSize
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Fields: fields,
				Query:  key,
			},
		}).From(form).Size(pageInfo.PageSize).
		Do(context.Background())
	if err != nil {
		global.Log.Error("search document failed", zap.Error(err))
		return result
	}
	for _, hit := range resp.Hits.Hits {
		var item models.Article
		err := json.Unmarshal(hit.Source_, &item)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}
		result = append(result, item)
	}
	return result
}

func SearchDocumentMultiMatchByTitle(fields []string, key string) (result []ArticleSearchResponse) {
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				"title": {
					Query: key,
				},
			},
		}).Source_(fields).
		Do(context.Background())
	if err != nil {
		global.Log.Error("search document failed, err", zap.Error(err))
		return result
	}
	for _, hit := range resp.Hits.Hits {
		var item ArticleSearchResponse
		err := json.Unmarshal(hit.Source_, &item)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}
		result = append(result, item)
	}
	return result
}

func GetDocumentById(id string) (result models.Article, err error) {
	resp, err := global.Es.Get(models.Article{}.Index(), id).
		Do(context.Background())
	if err != nil {
		global.Log.Error("get document by id failed", zap.Error(err))
		return result, err
	}
	err = json.Unmarshal(resp.Source_, &result)
	if err != nil {
		global.Log.Error("unmarshal json failed", zap.Error(err))
		return result, err
	}
	return result, nil
}

func SearchAllDocuments(pageInfo models.PageInfo) (result []models.Article) {
	form := (pageInfo.Page - 1) * pageInfo.PageSize
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			MatchAll: &types.MatchAllQuery{},
		}).From(form).Size(pageInfo.PageSize).Do(context.Background())
	if err != nil {
		global.Log.Error("search all documents failed", zap.Error(err))
		return result
	}
	for _, hit := range resp.Hits.Hits {
		var item models.Article
		err := json.Unmarshal(hit.Source_, &item)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}
		result = append(result, item)
	}
	return result
}

func DocIsExistByTitle(title string) bool {
	res := SearchDocumentTerm("title.keyword", title)
	if len(res) == 0 {
		return false
	} else {
		return true
	}
}

func DocIsExistById(id string) bool {
	res, _ := global.Es.Exists(models.Article{}.Index(), id).Do(context.Background())
	return res
}
