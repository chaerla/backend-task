package cmd

import (
	"backend-task/bootstrap/worker"
	"backend-task/internal/usecase"
	"backend-task/pkg/logger"
	"fmt"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/net/context"
)

type GetPostsCmd *cobra.Command

func NewGetPostsCmd(postService usecase.PostService) GetPostsCmd {
	var (
		totalPages     int
		numWorkers     int
		defaultPages   = 10
		defaultWorkers = 5
	)

	cmd := &cobra.Command{
		Use:   "posts",
		Short: "Use this command to scrape posts data",
		Run: func(cmd *cobra.Command, args []string) {
			_, span := otel.Tracer("GetPosts").Start(context.Background(), "GetPosts")
			defer span.End()
			span.SetAttributes(
				attribute.Int("pages", totalPages),
				attribute.Int("workers", numWorkers),
			)

			wp := worker.NewWorkerPool(numWorkers)
			wp.Run()
			logger.Log.Info(fmt.Sprintf("Scraping %d pages of posts data with %d workers", totalPages, numWorkers))

			for i := 1; i <= totalPages; i++ {
				page := i
				wp.Enqueue(
					func() {
						postService.GetManyAndPrint(page)
					},
				)
			}

			wp.Close()
		},
	}

	cmd.Flags().IntVarP(&totalPages, "pages", "p", defaultPages, "Number of pages to scrape")
	cmd.Flags().IntVarP(&numWorkers, "workers", "w", defaultWorkers, "Number of workers in the worker pool")

	return cmd
}
