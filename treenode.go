// Package treenode provides a set of functions for working with the
// https://leetcode.com/ binary tree structure TreeNode
package treenode

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	pad       = " "
	line      = "\u2500"
	leftBend  = "\u256d"
	rightBend = "\u256e"
)

var (
	r *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Depth returns depth of the TreeNode
func (t *TreeNode) Depth() int {
	if t == nil {
		return -1
	}

	return 1 + max(t.Left.Depth(), t.Right.Depth())
}

// String returns pretty-formatted TreeNode
func (t *TreeNode) String() string {
	p := placeholder(t.Depth())

	values(t, p, 0, 0, false)
	lines(p)

	s := ""
	for _, x := range p {
		s += fmt.Sprintf("%s\n", strings.Join(x, ""))
	}

	return s
}

// Inorder returns inorder traversal of TreeNode
func (t *TreeNode) Inorder() []int {
	if t == nil {
		return nil
	}

	return append(
		append(
			t.Left.Inorder(),
			t.Val,
		),
		t.Right.Inorder()...,
	)
}

// Preorder returns preorder traversal of TreeNode
func (t *TreeNode) Preorder() []int {
	if t == nil {
		return nil
	}

	return append(
		append(
			[]int{t.Val},
			t.Left.Preorder()...,
		),
		t.Right.Preorder()...,
	)
}

// Postorder returns postorder traversal of TreeNode
func (t *TreeNode) Postorder() []int {
	if t == nil {
		return nil
	}

	return append(
		append(
			t.Left.Postorder(),
			t.Right.Postorder()...,
		),
		t.Val,
	)
}

// Invert returns inverted TreeNode
func (t *TreeNode) Invert() *TreeNode {
	if t == nil {
		return nil
	}

	return &TreeNode{
		Val:   t.Val,
		Left:  t.Right.Invert(),
		Right: t.Left.Invert(),
	}
}

// IsEqual checks if trees are equal
func (t *TreeNode) IsEqual(other *TreeNode) bool {
	if t == nil && other == nil {
		return true
	}

	if (t != nil && other == nil) || (t == nil && other != nil) {
		return false
	}

	if t.Val != other.Val {
		return false
	}

	return t.Left.IsEqual(other.Left) && t.Right.IsEqual(other.Right)
}

// New returns *TreeNode using simple method
func New(xs []int) *TreeNode {
	if len(xs) == 0 {
		return nil
	}

	m := len(xs) / 2
	return &TreeNode{
		Val:   xs[m],
		Left:  New(xs[:m]),
		Right: New(xs[m+1:]),
	}
}

// NewFromInPre returns *TreeNode using inorder and preorder traversals of binary tree
func NewFromInPre(inorder, preorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}

	i := index(inorder, preorder[0])

	if i < 0 {
		return nil
	}

	lin := inorder[:i]
	lpre := preorder[1 : len(lin)+1]

	rin := inorder[i+1:]
	rpre := preorder[len(lpre)+1:]

	return &TreeNode{
		Val:   preorder[0],
		Left:  NewFromInPre(lin, lpre),
		Right: NewFromInPre(rin, rpre),
	}
}

// NewFromInPost returns *TreeNode using inorder and postorder traversals of binary tree
func NewFromInPost(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	i := index(inorder, postorder[len(postorder)-1])
	if i < 0 {
		return nil
	}

	lin := inorder[:i]
	lpost := postorder[:len(lin)]
	rin := inorder[i+1:]
	rpost := postorder[len(lpost) : len(postorder)-1]

	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  NewFromInPost(lin, lpost),
		Right: NewFromInPost(rin, rpost),
	}
}

// NewFromPrePost returns *TreeNode using preorder and postorder traversals of binary tree
func NewFromPrePost(preorder, postorder []int) *TreeNode {
	if len(preorder) == 0 && len(postorder) == 0 {
		return nil
	}

	if preorder[0] != postorder[len(postorder)-1] || len(preorder) != len(postorder) {
		return nil
	}

	if len(preorder) == 1 && len(postorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}

	i := index(postorder, preorder[1])
	if i < 0 {
		return nil
	}

	lpost := postorder[:i+1]
	lpre := preorder[1 : len(lpost)+1]

	rpre := preorder[1+len(lpre):]
	rpost := postorder[len(lpost) : len(postorder)-1]

	return &TreeNode{
		Val:   preorder[0],
		Left:  NewFromPrePost(lpre, lpost),
		Right: NewFromPrePost(rpre, rpost),
	}
}

func newRandom(depth, min, max, p int, bst bool) *TreeNode {
	// depth: depth of binary tree
	// min, max: etc
	// p: nil-node probability
	// bst: generate binary search tree
	if depth == 0 {
		return nil
	}

	if r.Int()%100 < p || max-min < 1 {
		return nil
	}

	v := r.Int()%(max-min+1) + min

	leftMax := max
	rightMin := min
	if bst {
		leftMax = v - 1
		rightMin = v + 1
	}

	return &TreeNode{
		Val:   v,
		Left:  newRandom(depth-1, min, leftMax, p, bst),
		Right: newRandom(depth-1, rightMin, max, p, bst),
	}
}

// NewRandom returns random generated not full binary tree *TreeNode
// with depth, min and max values of tree and probability of nil-node p (0..100)
func NewRandom(depth, min, max, p int) *TreeNode {
	return newRandom(depth, min, max, p, false)
}

// NewRandomFull returns random generated full binary tree *TreeNode
// with depth, min and max values of tree
func NewRandomFull(depth, min, max int) *TreeNode {
	return newRandom(depth, min, max, 0, false)
}

// NewRandomBST returns random generated not-full binary search tree *TreeNode
// with depth, min and max values of tree and probability of nil-node p (0..100)
func NewRandomBST(depth, min, max, p int) *TreeNode {
	return newRandom(depth, min, max, p, true)
}

// NewRandomFullBST returns random generated full binary search tree *TreeNode
// with depth, min and max values of tree
func NewRandomFullBST(depth, min, max int) *TreeNode {
	return newRandom(depth, min, max, 0, true)
}

// max
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// index
func index(xs []int, x int) int {
	for i := range xs {
		if x == xs[i] {
			return i
		}
	}

	return -1
}

// lines
func lines(p [][]string) {
	for i := 0; i < len(p)-1; i++ {
		step := len(p[i]) / (1 << (i + 1))
		for j := step; j < len(p[i]); j += 2 * (step + 1) {
			edge := len(p[i+1]) / (1 << (i + 2))

			for n := j - 1; n > j-edge-1; n-- {
				p[i][n] = strings.Repeat(line, len(p[i][n]))
			}

			for n := j + 1; n < j+edge+1; n++ {
				p[i][n] = strings.Repeat(line, len(p[i][n]))
			}
		}
	}
}

// placeholder
func placeholder(depth int) [][]string {
	p := make([][]string, depth+1)
	for i := range p {
		p[i] = make([]string, 1<<(depth+1)-1)
	}

	return p
}

// values
func values(root *TreeNode, p [][]string, offset int, depth int, left bool) {
	if root == nil {
		return
	}

	center := offset + len(p[depth])/(1<<(depth+1))
	p[depth][center] = fmt.Sprintf(" %d ", root.Val)

	for i := 0; i < len(p); i++ {
		if i != depth {
			p[i][center] = strings.Repeat(pad, len(p[depth][center]))
		}
	}

	if depth > 0 {
		n := len(p[depth-1][center])
		s := strings.Repeat(line, n/2-1)
		if n%2 != 0 {
			s += line
		}
		if left {
			s = strings.Repeat(pad, n/2) + leftBend + s
		} else {
			s = s + rightBend + strings.Repeat(pad, n/2)
		}

		p[depth-1][center] = s
	}

	values(root.Left, p, offset, depth+1, true)
	values(root.Right, p, center+1, depth+1, false)
}
