package v2_by_list

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListStack_Push(t *testing.T) {
	type fields struct {
		stack *list.List
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case-1",
			fields: fields{
				stack: list.New(),
			},
			args: args{
				value: "Vladimir",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &listStack{
				stack: tt.fields.stack,
			}
			s.Push(tt.args.value)

			value, err := s.Front()
			assert.Empty(t, err, "Error must be nil")
			assert.Equal(t, tt.args.value, value, "Values must be equal")
		})
	}
}

func TestListStack_Pop(t *testing.T) {
	testList := list.New()
	testList.PushFront("Vladimir")
	testList.PushFront("Oly")
	testList.PushFront("Uly")

	type fields struct {
		stack *list.List
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult string
		wantErr    bool
	}{
		{
			name: "case-1",
			fields: fields{
				stack: testList,
			},
			wantResult: "Uly",
			wantErr:    false,
		},
		{
			name: "case-2",
			fields: fields{
				stack: testList,
			},
			wantResult: "Oly",
			wantErr:    false,
		},
		{
			name: "case-3",
			fields: fields{
				stack: testList,
			},
			wantResult: "Vladimir",
			wantErr:    false,
		},
		{
			name: "case-4",
			fields: fields{
				stack: testList,
			},
			wantResult: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &listStack{
				stack: tt.fields.stack,
			}
			gotResult, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("listStack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("listStack.Pop() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestListStack_Front(t *testing.T) {
	listForTest := list.New()
	listForTest.PushFront("Vladimir")

	type fields struct {
		stack *list.List
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult string
		wantErr    bool
	}{
		{
			name: "case-1",
			fields: fields{
				stack: listForTest,
			},
			wantResult: "Vladimir",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &listStack{
				stack: tt.fields.stack,
			}
			gotResult, err := s.Front()
			if (err != nil) != tt.wantErr {
				t.Errorf("listStack.Front() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("listStack.Front() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestListStack_Size(t *testing.T) {
	testList := list.New()
	testList.PushFront("Vladimir")

	testList2 := list.New()
	testList2.PushFront("Vladimir")
	testList2.PushFront("Oly")

	testList3 := list.New()
	testList3.PushFront("Vladimir")
	testList3.PushFront("Oly")
	testList3.PushFront("Uly")

	type fields struct {
		stack *list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "case-1",
			fields: fields{
				stack: testList,
			},
			want: 1,
		},
		{
			name: "case-2",
			fields: fields{
				stack: testList2,
			},
			want: 2,
		},
		{
			name: "case-3",
			fields: fields{
				stack: testList3,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &listStack{
				stack: tt.fields.stack,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("listStack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listStack_Empty(t *testing.T) {
	testList := list.New()
	testList.PushFront("Vladimir")

	type fields struct {
		stack *list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case-1",
			fields: fields{
				stack: list.New(),
			},
			want: true,
		},
		{
			name: "case-2",
			fields: fields{
				stack: testList,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &listStack{
				stack: tt.fields.stack,
			}
			if got := s.Empty(); got != tt.want {
				t.Errorf("listStack.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
