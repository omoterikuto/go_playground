package main

import (
	"context"
	"fmt"
	"github.com/hasura/go-graphql-client"
	"net/http"
)

//var Mutation struct {
//	BulkOperationRunQuery struct {
//		BulkOperationRunQueryResponse
//	} `graphql:"bulkOperationRunQuery(query: $query)"`
//	BulkOperationRunMutation struct {
//		BulkOperationRunMutationResponse
//	} `graphql:"bulkOperationRunQuery(mutation: $mutation, stagedUploadPath, $path)"`
//}

var Query struct {
	CurrentBulkOperation struct {
		BulkOperation
	} `graphql:"currentBulkOperation(type: $type)"`
}

type BulkOperationType string

type BulkOperationMutationInput struct {
	Query            string `graphql:"query"`
	StagedUploadPath string `graphql:"stagedUploadPath"`
}

type BulkOperationRunQueryResponse struct {
	BulkOperation BulkOperation    `graphql:"bulkOperation"`
	UserErrors    []QueryUserError `graphql:"userErrors"`
}

type BulkOperationRunMutationResponse struct {
	BulkOperation BulkOperation       `graphql:"bulkOperation"`
	UserErrors    []MutationUserError `graphql:"userErrors"`
}

type CurrentBulkOperationResponse struct {
	BulkOperation BulkOperation `graphql:"bulkOperation"`
}

type BulkOperation struct {
	ID        string     `graphql:"id"`
	Status    Status     `graphql:"status"`
	ErrorCode *ErrorCode `graphql:"errorCode"`
	URL       *string    `graphql:"url"`
}

type Status string

const (
	StatusCompleted = Status("COMPLETED")
	StatusRunning   = Status("RUNNING")
	StatusFailed    = Status("FAILED")
)

type ErrorCode string

const (
	ErrorCodeAccessDenied        = ErrorCode("ACCESS_DENIED")
	ErrorCodeInternalServerError = ErrorCode("INTERNAL_SERVER_ERROR")
	ErrorCodeTimeout             = ErrorCode("TIMEOUT")
)

type QueryUserError struct {
	Field   []string `graphql:"field"`
	Message string   `graphql:"message"`
}

type MutationUserError struct {
	Code    UserErrorCode `graphql:"code"`
	Field   []string      `graphql:"field"`
	Message string        `graphql:"message"`
}

type UserErrorCode string

const (
	UserErrorCodeInternalFileServerError = UserErrorCode("INTERNAL_FILE_SERVER_ERROR")
	UserErrorCodeInvalidMutation         = UserErrorCode("INVALID_MUTATION")
	UserErrorCodeInvalidStagedUploadFile = UserErrorCode("INVALID_STAGED_UPLOAD_FILE")
	UserErrorCodeNoSuchFile              = UserErrorCode("NO_SUCH_FILE")
	UserErrorCodeOperationInProgress     = UserErrorCode("OPERATION_IN_PROGRESS")
)

func main() {
	err := CreateMutationBulkOperation(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	return

	cli := graphql.NewClient("hoge", new(http.Client))
	cli = cli.WithRequestModifier(func(r *http.Request) {
		r.Header.Add("X-Shopify-Access-Token", "hoge")
	})

	err = cli.Query(context.Background(), &Query, map[string]interface{}{
		"type": BulkOperationType("QUERY"),
	})
	fmt.Println(err)
	fmt.Printf("%+v\n", *Query.CurrentBulkOperation.BulkOperation.URL)
	fmt.Printf(Query.CurrentBulkOperation.BulkOperation.ID)
}

func CreateQueryBulkOperation(ctx context.Context) error {
	cli := graphql.NewClient("hoge", new(http.Client))
	cli = cli.WithRequestModifier(func(r *http.Request) {
		r.Header.Add("X-Shopify-Access-Token", "hoge")
	})
	var m struct {
		BulkOperationRunQuery struct {
			BulkOperationRunQueryResponse
		} `graphql:"bulkOperationRunQuery(query: $query)"`
	}

	if err := cli.Mutate(ctx, &m, map[string]interface{}{
		"query": `{ products(query: "created_at:>=2020-01-01 AND created_at:<2023-11-01") { edges { node { id createdAt updatedAt title handle descriptionHtml productType options { name position values } priceRange { minVariantPrice { amount currencyCode } maxVariantPrice { amount currencyCode } } } }  } }`,
	}); err != nil {
		return err
	}
	resp := m.BulkOperationRunQuery.BulkOperationRunQueryResponse

	if len(resp.UserErrors) != 0 {
		return fmt.Errorf("failed to run mutation. errors: %+v", resp.UserErrors)
	}

	fmt.Println(resp.BulkOperation.ID, resp.BulkOperation.Status)

	return nil
}

func CreateMutationBulkOperation(ctx context.Context) error {
	cli := graphql.NewClient("hoge", new(http.Client))
	cli = cli.WithRequestModifier(func(r *http.Request) {
		r.Header.Add("X-Shopify-Access-Token", "hoge")
	})
	var m struct {
		BulkOperationRunMutation struct {
			BulkOperationRunMutationResponse
		} `graphql:"bulkOperationRunMutation(mutation: $mutation, stagedUploadPath: $path)"`
	}
	err := cli.Mutate(ctx, &m, map[string]interface{}{
		"mutation": `mutation call($input: ProductInput!) { productCreate(input: $input) { product {id title } userErrors { message field } } }`,
		"path":     "tmp/76421038384/bulk/f6ecced6-8416-4f4b-bd60-dda5d65fdad1/input.jsonl",
	})
	if err != nil {
		return err
	}
	resp := m.BulkOperationRunMutation.BulkOperationRunMutationResponse

	if len(resp.UserErrors) != 0 {
		return fmt.Errorf("failed to run mutation. errors: %+v", resp.UserErrors)
	}
	fmt.Println(resp.BulkOperation.URL, resp.BulkOperation.ID, resp.BulkOperation.Status)

	return nil
}
