package catchy

type DoableNode struct {
	This IDoable
	Next *DoableNode
}

func newDoableNode(this IDoable) *DoableNode {
	return &DoableNode{
		This: this,
	}
}

func (node *DoableNode) Do() error {
	if err := node.This.Do(); err != nil {
		return err
	}

	last := node.Next
	for last != nil {
		if err := last.This.Do(); err != nil {
			return err
		}
		last = last.Next
	}
	return nil
}
