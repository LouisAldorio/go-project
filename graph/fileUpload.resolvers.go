package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
	"github.com/davidyap2002/user-go/tools"
)

func (r *fileUploadResolver) ViewLink(ctx context.Context, obj *model.FileUpload) (string, error) {
	return service.FileUploadViewLink(ctx, obj)
}

func (r *fileUploadResolver) DownloadLink(ctx context.Context, obj *model.FileUpload) (string, error) {
	return service.FileUploadDownloadLink(ctx, obj)
}

func (r *fileUploadOpsResolver) UploadSingle(ctx context.Context, obj *model.FileUploadOps, file graphql.Upload) (*model.FileUpload, error) {
	return service.UploadFile(ctx, file)
}

func (r *fileUploadOpsResolver) UploadBatch(ctx context.Context, obj *model.FileUploadOps, files []*graphql.Upload) ([]*model.FileUpload, error) {
	return service.UploadFileBatch(ctx, files)
}

func (r *fileUploadPaginationResolver) TotalData(ctx context.Context, obj *model.FileUploadPagination) (int, error) {
	limit, page, asc, sortBy, filter := tools.PaginationVariableGenerator(obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Filter)
	return service.FileUploadTotalDataPagination(ctx, limit, page, asc, sortBy, filter, obj.Scopes, obj.UserID)
}

func (r *fileUploadPaginationResolver) Nodes(ctx context.Context, obj *model.FileUploadPagination) ([]*model.FileUpload, error) {
	limit, page, asc, sortBy, filter := tools.PaginationVariableGenerator(obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Filter)
	return service.FileUploadPagination(ctx, limit, page, asc, sortBy, filter, obj.Scopes, obj.UserID)
}

// FileUpload returns generated.FileUploadResolver implementation.
func (r *Resolver) FileUpload() generated.FileUploadResolver { return &fileUploadResolver{r} }

// FileUploadOps returns generated.FileUploadOpsResolver implementation.
func (r *Resolver) FileUploadOps() generated.FileUploadOpsResolver { return &fileUploadOpsResolver{r} }

// FileUploadPagination returns generated.FileUploadPaginationResolver implementation.
func (r *Resolver) FileUploadPagination() generated.FileUploadPaginationResolver {
	return &fileUploadPaginationResolver{r}
}

type fileUploadResolver struct{ *Resolver }
type fileUploadOpsResolver struct{ *Resolver }
type fileUploadPaginationResolver struct{ *Resolver }
