// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: template.proto

/*
Package template is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package template

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_Template_Get_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_Template_Get_0(ctx context.Context, marshaler runtime.Marshaler, client TemplateClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TemplateData
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_Template_Get_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Get(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Template_Get_0(ctx context.Context, marshaler runtime.Marshaler, server TemplateServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TemplateData
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_Template_Get_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Get(ctx, &protoReq)
	return msg, metadata, err

}

func request_Template_Create_0(ctx context.Context, marshaler runtime.Marshaler, client TemplateClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TemplateData
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Create(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Template_Create_0(ctx context.Context, marshaler runtime.Marshaler, server TemplateServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TemplateData
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Create(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterTemplateHandlerServer registers the http handlers for service Template to "mux".
// UnaryRPC     :call TemplateServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterTemplateHandlerFromEndpoint instead.
func RegisterTemplateHandlerServer(ctx context.Context, mux *runtime.ServeMux, server TemplateServer) error {

	mux.Handle("GET", pattern_Template_Get_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/template.Template/Get", runtime.WithHTTPPathPattern("/v1/template_golang"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Template_Get_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Template_Get_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Template_Create_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/template.Template/Create", runtime.WithHTTPPathPattern("/v1/template_golang"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Template_Create_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Template_Create_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterTemplateHandlerFromEndpoint is same as RegisterTemplateHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTemplateHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTemplateHandler(ctx, mux, conn)
}

// RegisterTemplateHandler registers the http handlers for service Template to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTemplateHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTemplateHandlerClient(ctx, mux, NewTemplateClient(conn))
}

// RegisterTemplateHandlerClient registers the http handlers for service Template
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TemplateClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TemplateClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TemplateClient" to call the correct interceptors.
func RegisterTemplateHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TemplateClient) error {

	mux.Handle("GET", pattern_Template_Get_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/template.Template/Get", runtime.WithHTTPPathPattern("/v1/template_golang"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Template_Get_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Template_Get_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Template_Create_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/template.Template/Create", runtime.WithHTTPPathPattern("/v1/template_golang"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Template_Create_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Template_Create_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Template_Get_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "template_golang"}, ""))

	pattern_Template_Create_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "template_golang"}, ""))
)

var (
	forward_Template_Get_0 = runtime.ForwardResponseMessage

	forward_Template_Create_0 = runtime.ForwardResponseMessage
)
