package rpc

import (
	"context"

	"connectrpc.com/connect"

	. "github.com/river-build/river/core/node/base"
	. "github.com/river-build/river/core/node/events"
	. "github.com/river-build/river/core/node/protocol"
	. "github.com/river-build/river/core/node/shared"
)

func (s *Service) localGetStream(
	ctx context.Context,
	req *connect.Request[GetStreamRequest],
) (*connect.Response[GetStreamResponse], error) {
	streamId, err := StreamIdFromBytes(req.Msg.StreamId)
	if err != nil {
		return nil, err
	}

	var streamView StreamView
	stream, err := s.cache.GetStream(ctx, streamId)
	if err == nil {
		streamView, err = stream.GetView(ctx)
	}

	if err != nil {
		if req.Msg.Optional && AsRiverError(err).Code == Err_NOT_FOUND {
			// aellis - this is actually an error, if the forwarder thinks the stream exists, but it doesn't exist in the cache
			// it's a real error, but currently (feb 2024) in single node this will reach here
			// If optional is set, empty response indicates that there is no stream.
			// This reduces log spam for the case where stream legitimately may not exist yet.
			return connect.NewResponse(&GetStreamResponse{}), nil
		} else {
			return nil, err
		}
	} else {
		_, _ = s.scrubTaskProcessor.TryScheduleScrub(ctx, stream, false)
		return connect.NewResponse(&GetStreamResponse{
			Stream: &StreamAndCookie{
				Events:         streamView.MinipoolEnvelopes(),
				NextSyncCookie: streamView.SyncCookie(s.wallet.Address),
				Miniblocks:     streamView.MiniblocksFromLastSnapshot(),
			},
		}), nil
	}
}
