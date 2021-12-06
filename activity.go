// implement the activies such as
// calling /transcription /sentiment /tone /submission
// calling /recordings/:id/status /recordings/:id/delete
package app

import (
	"context"
	"fmt"
)

func ConfigService(ctx context.Context, protoDetails ProtobuffDetails) error {
	fmt.Printf(
		"connecting to Kafka Interfact (targetting config service) \nMessage %s, dub-point-id %s, recording_id %s\n",
		protoDetails.Message,
		protoDetails.DubPointID,
		protoDetails.RecordingID,
	)
	return nil
}

func TranscriptionService(ctx context.Context, protoDetails ProtobuffDetails) error {
	fmt.Printf(
		"connecting to Kafka Interfact  (targetting Transcription service) \nMessage %s, dub-point-id %s, recording_id %s\n",
		protoDetails.Message,
		protoDetails.DubPointID,
		protoDetails.RecordingID,
	)
	return nil
}

func DeletionService(ctx context.Context, protoDetails ProtobuffDetails) error {
	fmt.Printf(
		"connecting to Kafka Interfact  (targetting Deletion service) \nMessage %s, dub-point-id %s, recording_id %s\n",
		protoDetails.Message,
		protoDetails.DubPointID,
		protoDetails.RecordingID,
	)
	return nil
}
