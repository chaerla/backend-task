package cmd

import (
	"backend-task/bootstrap/config"
	"backend-task/bootstrap/kafka"
	"github.com/spf13/cobra"
)

type KafkaRunnerCmd *cobra.Command

func NewKafkaRunnerCmd(config *config.Config) KafkaRunnerCmd {
	var topic string
	var message string

	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "Use this command to connect to kafka broker",
		Run: func(cmd *cobra.Command, args []string) {
			producer := kafka.NewKafkaProducer(config)
			if topic != "" && message != "" {
				producer.SendMessage(topic, message)
			} else {
				cmd.Help()
			}
		},
	}

	cmd.Flags().StringVarP(&topic, "topic", "t", "", "Kafka topic")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Message to send")

	cmd.MarkFlagRequired("topic")
	cmd.MarkFlagRequired("message")

	return cmd
}
