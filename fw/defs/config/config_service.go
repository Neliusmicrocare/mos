// Code generated by clubbygen.
// GENERATED FILE DO NOT EDIT
// +build !clubby_strict

package config

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"cesanta.com/clubby"
	"cesanta.com/clubby/endpoint"
	"cesanta.com/clubby/frame"
	"cesanta.com/common/go/ourjson"
	"cesanta.com/common/go/ourtrace"
	"github.com/cesanta/errors"
	"golang.org/x/net/trace"
)

var _ = bytes.MinRead
var _ = fmt.Errorf
var emptyMessage = ourjson.RawMessage{}
var _ = ourtrace.New
var _ = trace.New

const ServiceID = "http://mongoose-iot.com/fwConfig"

type SaveArgs struct {
	Reboot *bool `json:"reboot,omitempty"`
}

type SetArgs struct {
	Config ourjson.RawMessage `json:"config,omitempty"`
}

type Service interface {
	Get(ctx context.Context) (ourjson.RawMessage, error)
	Save(ctx context.Context, args *SaveArgs) error
	Set(ctx context.Context, args *SetArgs) error
}

type Instance interface {
	Call(context.Context, string, *frame.Command) (*frame.Response, error)
	TraceCall(context.Context, string, *frame.Command) (context.Context, trace.Trace, func(*error))
}

func NewClient(i Instance, addr string) Service {
	return &_Client{i: i, addr: addr}
}

type _Client struct {
	i    Instance
	addr string
}

func (c *_Client) Get(pctx context.Context) (res ourjson.RawMessage, err error) {
	cmd := &frame.Command{
		Cmd: "Config.Get",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return ourjson.RawMessage{}, errors.Trace(err)
	}
	if resp.Status != 0 {
		return ourjson.RawMessage{}, errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	var r ourjson.RawMessage
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return ourjson.RawMessage{}, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Save(pctx context.Context, args *SaveArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "Config.Save",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func (c *_Client) Set(pctx context.Context, args *SetArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "Config.Set",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func RegisterService(i *clubby.Instance, impl Service) error {
	s := &_Server{impl}
	i.RegisterCommandHandler("Config.Get", s.Get)
	i.RegisterCommandHandler("Config.Save", s.Save)
	i.RegisterCommandHandler("Config.Set", s.Set)
	i.RegisterService(ServiceID, _ServiceDefinition)
	return nil
}

type _Server struct {
	impl Service
}

func (s *_Server) Get(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	return s.impl.Get(ctx)
}

func (s *_Server) Save(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args SaveArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.Save(ctx, &args)
}

func (s *_Server) Set(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args SetArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.Set(ctx, &args)
}

var _ServiceDefinition = json.RawMessage([]byte(`{
  "methods": {
    "Get": {
      "doc": "Get device config",
      "result": {
        "keep_as_json": true
      }
    },
    "Save": {
      "args": {
        "reboot": {
          "doc": "If set to ` + "`" + `true` + "`" + `, the device will be rebooted after saving config. It\nis often desirable because it's the only way to apply saved config.\n",
          "type": "boolean"
        }
      },
      "doc": "Save device config"
    },
    "Set": {
      "args": {
        "config": {
          "keep_as_json": true
        }
      },
      "doc": "Set device config"
    }
  },
  "name": "Config",
  "namespace": "http://mongoose-iot.com/fw"
}`))