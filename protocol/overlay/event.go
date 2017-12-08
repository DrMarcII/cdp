// Code generated by cdpgen. DO NOT EDIT.

package overlay

import (
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

// InspectNodeRequestedClient is a client for InspectNodeRequested events. Fired when the node should be inspected. This happens after call to `setInspectMode` or when
// user manually inspects an element.
type InspectNodeRequestedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*InspectNodeRequestedReply, error)
	rpcc.Stream
}

// InspectNodeRequestedReply is the reply for InspectNodeRequested events.
type InspectNodeRequestedReply struct {
	BackendNodeID dom.BackendNodeID `json:"backendNodeId"` // Id of the node to inspect.
}

// NodeHighlightRequestedClient is a client for NodeHighlightRequested events. Fired when the node should be highlighted. This happens after call to `setInspectMode`.
type NodeHighlightRequestedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*NodeHighlightRequestedReply, error)
	rpcc.Stream
}

// NodeHighlightRequestedReply is the reply for NodeHighlightRequested events.
type NodeHighlightRequestedReply struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// ScreenshotRequestedClient is a client for ScreenshotRequested events. Fired when user asks to capture screenshot of some area on the page.
type ScreenshotRequestedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScreenshotRequestedReply, error)
	rpcc.Stream
}

// ScreenshotRequestedReply is the reply for ScreenshotRequested events.
type ScreenshotRequestedReply struct {
	Viewport page.Viewport `json:"viewport"` // Viewport to capture, in CSS.
}
