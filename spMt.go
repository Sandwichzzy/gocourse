package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// SparseMerkleTree 稀疏默克尔树
type SparseMerkleTree struct {
	root  *Node
	depth int // 树的深度
}

// Node 树节点
type Node struct {
	hash  []byte
	left  *Node
	right *Node
	key   []byte // 叶子节点的键
	value []byte // 叶子节点的值
}

// NewSparseMerkleTree 创建新的稀疏默克尔树
// depth: 树的深度，支持 2^depth 个叶子节点
func NewSparseMerkleTree(depth int) *SparseMerkleTree {
	return &SparseMerkleTree{
		root:  newEmptyNode(),
		depth: depth,
	}
}

// newEmptyNode 创建空节点
func newEmptyNode() *Node {
	emptyHash := sha256.Sum256([]byte{})
	return &Node{
		hash: emptyHash[:],
	}
}

// hashData 对数据进行哈希
func hashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// hashNodes 合并两个节点的哈希
func hashNodes(left, right []byte) []byte {
	combined := append(left, right...)
	return hashData(combined)
}

// getBit 获取字节数组在指定位置的比特位
func getBit(data []byte, position int) bool {
	byteIndex := position / 8
	bitIndex := 7 - (position % 8)
	if byteIndex >= len(data) {
		return false
	}
	return (data[byteIndex]>>bitIndex)&1 == 1
}

// Update 更新或插入键值对
func (smt *SparseMerkleTree) Update(key, value []byte) {
	keyHash := hashData(key)
	valueHash := hashData(value)
	smt.root = smt.update(smt.root, keyHash, value, valueHash, 0)
}

// update 递归更新节点
func (smt *SparseMerkleTree) update(node *Node, keyHash, value, valueHash []byte, depth int) *Node {
	// 到达叶子节点
	if depth == smt.depth {
		return &Node{
			hash:  valueHash,
			key:   keyHash,
			value: value,
		}
	}

	// 如果当前节点为空，创建新节点
	if node == nil {
		node = newEmptyNode()
	}

	// 根据 key 的比特位决定往左还是右
	bit := getBit(keyHash, depth)

	if bit { // bit = 1，往右
		if node.right == nil {
			node.right = newEmptyNode()
		}
		node.right = smt.update(node.right, keyHash, value, valueHash, depth+1)
	} else { // bit = 0，往左
		if node.left == nil {
			node.left = newEmptyNode()
		}
		node.left = smt.update(node.left, keyHash, value, valueHash, depth+1)
	}

	// 更新当前节点的哈希
	var leftHash, rightHash []byte
	if node.left == nil {
		leftHash = newEmptyNode().hash
	} else {
		leftHash = node.left.hash
	}
	if node.right == nil {
		rightHash = newEmptyNode().hash
	} else {
		rightHash = node.right.hash
	}
	node.hash = hashNodes(leftHash, rightHash)

	return node
}

// Get 获取键对应的值
func (smt *SparseMerkleTree) Get(key []byte) ([]byte, bool) {
	keyHash := hashData(key)
	return smt.get(smt.root, keyHash, 0)
}

// get 递归获取值
func (smt *SparseMerkleTree) get(node *Node, keyHash []byte, depth int) ([]byte, bool) {
	if node == nil {
		return nil, false
	}

	// 到达叶子节点
	if depth == smt.depth {
		if string(node.key) == string(keyHash) {
			return node.value, true
		}
		return nil, false
	}

	// 根据 key 的比特位决定往左还是右
	bit := getBit(keyHash, depth)
	if bit {
		return smt.get(node.right, keyHash, depth+1)
	}
	return smt.get(node.left, keyHash, depth+1)
}

// Proof Merkle 证明
type Proof struct {
	Siblings [][]byte // 兄弟节点的哈希值
	Path     []bool   // 路径（false=左，true=右）
}

// GenerateProof 生成 Merkle 证明
func (smt *SparseMerkleTree) GenerateProof(key []byte) *Proof {
	keyHash := hashData(key)
	proof := &Proof{
		Siblings: make([][]byte, 0, smt.depth),
		Path:     make([]bool, 0, smt.depth),
	}
	smt.generateProof(smt.root, keyHash, 0, proof)
	return proof
}

// generateProof 递归生成证明
func (smt *SparseMerkleTree) generateProof(node *Node, keyHash []byte, depth int, proof *Proof) {
	if depth == smt.depth || node == nil {
		return
	}

	bit := getBit(keyHash, depth)
	proof.Path = append(proof.Path, bit)

	if bit { // 往右走，记录左兄弟
		if node.left != nil {
			proof.Siblings = append(proof.Siblings, node.left.hash)
		} else {
			proof.Siblings = append(proof.Siblings, newEmptyNode().hash)
		}
		smt.generateProof(node.right, keyHash, depth+1, proof)
	} else { // 往左走，记录右兄弟
		if node.right != nil {
			proof.Siblings = append(proof.Siblings, node.right.hash)
		} else {
			proof.Siblings = append(proof.Siblings, newEmptyNode().hash)
		}
		smt.generateProof(node.left, keyHash, depth+1, proof)
	}
}

// VerifyProof 验证 Merkle 证明
func (smt *SparseMerkleTree) VerifyProof(key, value []byte, proof *Proof) bool {
	valueHash := hashData(value)

	// 从叶子节点开始计算哈希
	currentHash := valueHash

	// 从底部往上计算
	for i := len(proof.Siblings) - 1; i >= 0; i-- {
		sibling := proof.Siblings[i]
		path := proof.Path[i]

		if path { // 当前节点在右边
			currentHash = hashNodes(sibling, currentHash)
		} else { // 当前节点在左边
			currentHash = hashNodes(currentHash, sibling)
		}
	}

	// 比较计算出的根哈希与树的根哈希
	return string(currentHash) == string(smt.root.hash)
}

// GetRoot 获取根节点哈希
func (smt *SparseMerkleTree) GetRoot() []byte {
	return smt.root.hash
}

// PrintTree 打印树结构（用于调试）
func (smt *SparseMerkleTree) PrintTree() {
	fmt.Println("稀疏默克尔树结构:")
	smt.printNode(smt.root, 0, "Root")
}

func (smt *SparseMerkleTree) printNode(node *Node, depth int, prefix string) {
	if node == nil {
		return
	}

	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	hashStr := hex.EncodeToString(node.hash[:8]) // 只显示前8字节
	fmt.Printf("%s%s: %s...\n", indent, prefix, hashStr)

	if node.key != nil {
		fmt.Printf("%s  Key: %s\n", indent, hex.EncodeToString(node.key[:4]))
		fmt.Printf("%s  Value: %s\n", indent, string(node.value))
	}

	if depth < smt.depth {
		if node.left != nil {
			smt.printNode(node.left, depth+1, "L")
		}
		if node.right != nil {
			smt.printNode(node.right, depth+1, "R")
		}
	}
}

func main() {
	// 创建深度为 8 的稀疏默克尔树
	smt := NewSparseMerkleTree(8)

	fmt.Println("=== 稀疏默克尔树示例 ===")

	// 插入一些键值对
	fmt.Println("1. 插入键值对:")
	data := map[string]string{
		"alice": "100",
		"bob":   "200",
		"carol": "300",
	}

	for key, value := range data {
		smt.Update([]byte(key), []byte(value))
		fmt.Printf("   插入: %s = %s\n", key, value)
	}

	// 显示根哈希
	fmt.Printf("\n2. 根哈希: %s\n", hex.EncodeToString(smt.GetRoot()))

	// 查询值
	fmt.Println("\n3. 查询值:")
	if value, found := smt.Get([]byte("alice")); found {
		fmt.Printf("   alice = %s\n", string(value))
	}
	if value, found := smt.Get([]byte("bob")); found {
		fmt.Printf("   bob = %s\n", string(value))
	}
	if _, found := smt.Get([]byte("dave")); !found {
		fmt.Println("   dave 不存在")
	}

	// 生成证明
	fmt.Println("\n4. 生成 Merkle 证明:")
	proof := smt.GenerateProof([]byte("alice"))
	fmt.Printf("   alice 的证明包含 %d 个兄弟节点\n", len(proof.Siblings))

	// 验证证明
	fmt.Println("\n5. 验证 Merkle 证明:")
	valid := smt.VerifyProof([]byte("alice"), []byte("100"), proof)
	fmt.Printf("   alice=100 的证明验证: %v\n", valid)

	// 验证错误的值
	invalid := smt.VerifyProof([]byte("alice"), []byte("999"), proof)
	fmt.Printf("   alice=999 的证明验证: %v\n", invalid)

	// 更新值
	fmt.Println("\n6. 更新值:")
	smt.Update([]byte("alice"), []byte("150"))
	fmt.Println("   更新 alice = 150")
	if value, found := smt.Get([]byte("alice")); found {
		fmt.Printf("   新值: alice = %s\n", string(value))
	}
	fmt.Printf("   新根哈希: %s\n", hex.EncodeToString(smt.GetRoot()))

	// 旧证明应该失效
	fmt.Println("\n7. 旧证明验证:")
	stillValid := smt.VerifyProof([]byte("alice"), []byte("100"), proof)
	fmt.Printf("   旧证明(alice=100)验证: %v (应该为 false)\n", stillValid)
}
