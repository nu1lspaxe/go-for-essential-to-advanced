package threadpools

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

var (
	errDefault = errors.New("worng argument type.")
	desciptor  = JobDescriptor{
		ID:    JobID("1"),
		JType: JobType("anyType"),
		MetaData: JobMetadata{
			"foo": "foo",
			"bar": "bar",
		},
	}
	execFn = func(ctx context.Context, args interface{}) (interface{}, error) {
		argVal, ok := args.(int)
		if !ok {
			return nil, errDefault
		}
		return argVal * 2, nil
	}
)

func TestJobExecute(t *testing.T) {
	ctx := context.TODO()

	// Job for single test
	type fields struct {
		descriptor JobDescriptor
		execFn     ExecutionFn
		args       interface{}
	}

	// Anonymous struct 使用 (匿名結構目的: 僅使用一次)
	tests := []struct {
		name   string // Test case name
		fields fields // Job for this test
		want   Result // Expected result
	}{
		{
			name: "job execution success",
			fields: fields{
				descriptor: desciptor,
				execFn:     execFn,
				args:       10,
			},
			want: Result{
				Value:      20,
				Descriptor: desciptor,
			},
		},
		{
			name: "job execution failure",
			fields: fields{
				descriptor: desciptor,
				execFn:     execFn,
				args:       "10",
			},
			want: Result{
				Err:        errDefault,
				Descriptor: desciptor,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			j := Job{
				Descriptor: test.fields.descriptor,
				ExecFn:     test.fields.execFn,
				Args:       test.fields.args,
			}

			res := j.execute(ctx)
			if test.want.Err != nil {
				if !reflect.DeepEqual(res.Err, test.want.Err) {
					t.Errorf("execute() = %v, wantError = %v", res.Err, test.want.Err)
				}
				return
			}

			if !reflect.DeepEqual(res.Value, test.want.Value) {
				t.Errorf("execute() = %v, wantError = %v", res.Value, test.want.Value)
			}
		})
	}
}
