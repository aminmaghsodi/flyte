package v1alpha1

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
)

type BooleanExpression struct {
	*core.BooleanExpression
}

func (in BooleanExpression) MarshalJSON() ([]byte, error) {
	if in.BooleanExpression == nil {
		return nilJSON, nil
	}

	var buf bytes.Buffer
	if err := marshaler.Marshal(&buf, in.BooleanExpression); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (in *BooleanExpression) UnmarshalJSON(b []byte) error {
	in.BooleanExpression = &core.BooleanExpression{}
	return jsonpb.Unmarshal(bytes.NewReader(b), in.BooleanExpression)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BooleanExpression) DeepCopyInto(out *BooleanExpression) {
	*out = *in
	// We do not manipulate the object, so its ok
	// Once we figure out the autogenerate story we can replace this
}

type IfBlock struct {
	Condition BooleanExpression `json:"condition"`
	ThenNode  *NodeID           `json:"then"`
}

func (in IfBlock) GetCondition() *core.BooleanExpression {
	return in.Condition.BooleanExpression
}

func (in *IfBlock) GetThenNode() *NodeID {
	return in.ThenNode
}

type BranchNodeSpec struct {
	If       IfBlock     `json:"if"`
	ElseIf   []*IfBlock  `json:"elseIf,omitempty"`
	Else     *NodeID     `json:"else,omitempty"`
	ElseFail *core.Error `json:"elseFail,omitempty"`
}

// BranchNodeSpec contains at least one proto message directly, which means that we need to
// implement DeepCopyInto since the generated proto messages do not implement that function.
func (in *BranchNodeSpec) DeepCopyInto(out *BranchNodeSpec) {
	*out = *in
	in.If.DeepCopyInto(&out.If)
	if in.ElseIf != nil {
		in, out := &in.ElseIf, &out.ElseIf
		*out = make([]*IfBlock, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IfBlock)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Else != nil {
		in, out := &in.Else, &out.Else
		*out = new(string)
		**out = **in
	}
	if in.ElseFail != nil {
		out.ElseFail = proto.Clone(in.ElseFail).(*core.Error)
	}
}

func (in *BranchNodeSpec) GetIf() ExecutableIfBlock {
	return &in.If
}

func (in *BranchNodeSpec) GetElse() *NodeID {
	return in.Else
}

func (in *BranchNodeSpec) GetElseIf() []ExecutableIfBlock {
	elifs := make([]ExecutableIfBlock, 0, len(in.ElseIf))
	for _, b := range in.ElseIf {
		elifs = append(elifs, b)
	}
	return elifs
}

func (in *BranchNodeSpec) GetElseFail() *core.Error {
	return in.ElseFail
}
