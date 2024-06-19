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

type GetUsersCmd *cobra.Command

func NewGetUsersCmd(userService usecase.UserService) GetUsersCmd {
	var (
		totalPages     int
		numWorkers     int
		defaultPages   = 10
		defaultWorkers = 5
	)

	cmd := &cobra.Command{
		Use:   "users",
		Short: "Use this command to scrape users data",
		Run: func(cmd *cobra.Command, args []string) {
			_, span := otel.Tracer("GetUsers").Start(context.Background(), "GetUsers")
			defer span.End()
			span.SetAttributes(
				attribute.Int("pages", totalPages),
				attribute.Int("workers", numWorkers),
			)

			wp := worker.NewWorkerPool(numWorkers)
			wp.Run()
			defer wp.Close()
			logger.Log.Info(fmt.Sprintf("Scraping %d pages of users data with %d workers", totalPages, numWorkers))

			for i := 1; i <= totalPages; i++ {
				page := i
				wp.Enqueue(
					func() {
						userService.GetManyAndPrint(page)
					},
				)
			}

		},
	}

	cmd.Flags().IntVarP(&totalPages, "pages", "p", defaultPages, "Number of pages to scrape")
	cmd.Flags().IntVarP(&numWorkers, "workers", "w", defaultWorkers, "Number of workers in the worker pool")

	return cmd
}
