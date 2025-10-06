package exercise

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// SparseMerkleTree 稀疏默克尔树
// 稀疏默克尔树是一种优化的默克尔树，专门用于处理大量可能的键值对，但实际只存储少量数据的场景
// 与传统默克尔树不同，稀疏默克尔树不需要为所有可能的叶子节点分配内存
// 它通过使用默认的"空节点"来表示未使用的位置，从而节省大量空间
type SparseMerkleTree struct {
	root  *Node // 树的根节点
	depth int   // 树的深度，决定了树可以容纳的最大键数量 (2^depth)
}

// Node 树节点
// 在稀疏默克尔树中，节点可以是内部节点或叶子节点
// 内部节点包含左右子节点的引用，叶子节点包含实际的键值对
type Node struct {
	hash  []byte // 节点的哈希值，对于内部节点是左右子节点哈希的组合，对于叶子节点是值的哈希
	left  *Node  // 左子节点指针
	right *Node  // 右子节点指针
	key   []byte // 叶子节点的键（已哈希），用于在到达叶子节点时验证是否找到了正确的键
	value []byte // 叶子节点的值（原始数据）
}

// NewSparseMerkleTree 创建新的稀疏默克尔树
// 参数:
//   depth: 树的深度，支持 2^depth 个叶子节点
//          例如: depth=8 可以支持 256 个不同的键
//          depth=256 可以支持 2^256 个键（接近无限）
// 返回:
//   初始化的稀疏默克尔树实例，初始根节点为空节点
func NewSparseMerkleTree(depth int) *SparseMerkleTree {
	return &SparseMerkleTree{
		root:  newEmptyNode(),
		depth: depth,
	}
}

// newEmptyNode 创建空节点
// 空节点代表树中未被使用的位置
// 所有空节点具有相同的哈希值（空字节数组的SHA256哈希）
// 这是稀疏默克尔树的关键优化：不需要为每个空位置创建新对象
func newEmptyNode() *Node {
	emptyHash := sha256.Sum256([]byte{})
	return &Node{
		hash: emptyHash[:],
	}
}

// hashData 对数据进行哈希
// 使用 SHA256 算法对输入数据进行哈希
// 参数:
//   data: 待哈希的字节数据
// 返回:
//   32字节的哈希值
func hashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// hashNodes 合并两个节点的哈希
// 在默克尔树中，父节点的哈希是通过合并子节点的哈希计算得出的
// 参数:
//   left: 左子节点的哈希值
//   right: 右子节点的哈希值
// 返回:
//   合并后的哈希值 (Hash(left || right))
func hashNodes(left, right []byte) []byte {
	combined := append(left, right...)
	return hashData(combined)
}

// getBit 获取字节数组在指定位置的比特位
// 该函数用于确定键的路径：在树的每一层，根据键的对应比特位决定向左(0)还是向右(1)
// 参数:
//   data: 字节数组（通常是键的哈希值）
//   position: 比特位的位置（0表示最高位root）
// 返回:
//   true 表示该位为1，false 表示该位为0
// 示例:
//   对于字节 0b10110000，getBit(data, 0)=true, getBit(data, 1)=false
func getBit(data []byte, position int) bool {
	byteIndex := position / 8          // 确定比特位所在的字节索引
	bitIndex := 7 - (position % 8)     // 确定比特位在该字节中的位置（从高位到低位）
	if byteIndex >= len(data) {        // 如果超出数据范围，返回false（视为0）
		return false
	}
	return (data[byteIndex]>>bitIndex)&1 == 1  // 提取指定位置的比特位
}

// Update 更新或插入键值对
// 这是稀疏默克尔树的核心操作，支持插入新键或更新现有键的值
// 参数:
//   key: 原始键（任意字节数组）
//   value: 要存储的值（任意字节数组）
// 工作流程:
//   1. 对键进行哈希，得到固定长度的键哈希（用于确定路径）
//   2. 对值进行哈希，得到叶子节点的哈希值
//   3. 从根节点开始，递归更新树结构
//   4. 更新路径上所有节点的哈希值
func (smt *SparseMerkleTree) Update(key, value []byte) {
	keyHash := hashData(key)
	valueHash := hashData(value)
	smt.root = smt.update(smt.root, keyHash, value, valueHash, 0)
}

// update 递归更新节点
// 这是 Update 方法的内部递归实现
// 参数:
//   node: 当前处理的节点
//   keyHash: 键的哈希值（用于确定路径）
//   value: 原始值（存储在叶子节点）
//   valueHash: 值的哈希（存储在叶子节点的hash字段）
//   depth: 当前深度（0表示根节点）
// 返回:
//   更新后的节点（可能是新创建的节点）
// 工作原理:
//   - 如果到达叶子层(depth == smt.depth)，创建新的叶子节点
//   - 否则，根据keyHash的当前比特位决定向左或向右递归
//   - 递归返回后，重新计算当前节点的哈希值
func (smt *SparseMerkleTree) update(node *Node, keyHash, value, valueHash []byte, depth int) *Node {
	// 到达叶子节点层，创建新的叶子节点存储键值对
	if depth == smt.depth {
		return &Node{
			hash:  valueHash,  // 叶子节点的哈希是值的哈希
			key:   keyHash,    // 存储键的哈希用于后续验证
			value: value,      // 存储原始值
		}
	}

	// 如果当前节点为空（第一次访问该路径），创建新的空节点
	if node == nil {
		node = newEmptyNode()
	}

	// 根据 keyHash 的第 depth 个比特位决定往左还是右
	// 这确保了相同的键总是走相同的路径
	bit := getBit(keyHash, depth)

	if bit { // bit = 1，往右子树递归
		if node.right == nil {
			node.right = newEmptyNode()
		}
		node.right = smt.update(node.right, keyHash, value, valueHash, depth+1)
	} else { // bit = 0，往左子树递归
		if node.left == nil {
			node.left = newEmptyNode()
		}
		node.left = smt.update(node.left, keyHash, value, valueHash, depth+1)
	}

	// 更新当前节点的哈希值
	// 父节点的哈希 = Hash(左子节点哈希 || 右子节点哈希)
	var leftHash, rightHash []byte
	if node.left == nil {
		leftHash = newEmptyNode().hash  // 如果左子节点不存在，使用空节点的哈希
	} else {
		leftHash = node.left.hash
	}
	if node.right == nil {
		rightHash = newEmptyNode().hash  // 如果右子节点不存在，使用空节点的哈希
	} else {
		rightHash = node.right.hash
	}
	node.hash = hashNodes(leftHash, rightHash)

	return node
}

// Get 获取键对应的值
// 从稀疏默克尔树中查询指定键的值
// 参数:
//   key: 要查询的键（原始字节数组）
// 返回:
//   value: 键对应的值（如果存在）
//   found: 布尔值，表示是否找到该键
func (smt *SparseMerkleTree) Get(key []byte) ([]byte, bool) {
	keyHash := hashData(key)
	return smt.get(smt.root, keyHash, 0)
}

// get 递归获取值
// 这是 Get 方法的内部递归实现
// 参数:
//   node: 当前处理的节点
//   keyHash: 键的哈希值（用于确定路径）
//   depth: 当前深度
// 返回:
//   value: 找到的值
//   found: 是否找到
// 工作原理:
//   按照与 Update 相同的路径规则向下遍历，直到到达叶子节点
//   在叶子节点处验证键是否匹配
func (smt *SparseMerkleTree) get(node *Node, keyHash []byte, depth int) ([]byte, bool) {
	// 如果节点为空，说明该键不存在
	if node == nil {
		return nil, false
	}

	// 到达叶子节点层，检查键是否匹配
	if depth == smt.depth {
		// 比较存储的键哈希与查询的键哈希
		if string(node.key) == string(keyHash) {
			return node.value, true  // 键匹配，返回值
		}
		return nil, false  // 键不匹配，该位置存储的是其他键
	}

	// 根据 keyHash 的第 depth 个比特位决定往左还是右
	bit := getBit(keyHash, depth)
	if bit {
		return smt.get(node.right, keyHash, depth+1)  // 往右子树查找
	}
	return smt.get(node.left, keyHash, depth+1)  // 往左子树查找
}

// Proof Merkle 证明
// Merkle 证明是一个紧凑的数据结构，用于证明某个键值对存在于树中
// 无需提供整棵树，只需提供从叶子节点到根节点路径上的"兄弟节点"
// 验证者可以使用这些兄弟节点重新计算根哈希，如果与已知的根哈希匹配，则证明有效
type Proof struct {
	Siblings [][]byte // 从叶子到根路径上所有兄弟节点的哈希值（按从根到叶的顺序）
	Path     []bool   // 路径信息（false=该层向左，true=该层向右）
}

// GenerateProof 生成 Merkle 证明
// 为指定的键生成一个证明，证明该键值对存在于树中
// 参数:
//   key: 要生成证明的键
// 返回:
//   包含兄弟节点哈希和路径信息的 Proof 结构
// 工作原理:
//   沿着键对应的路径向下遍历，记录每一层的兄弟节点哈希
func (smt *SparseMerkleTree) GenerateProof(key []byte) *Proof {
	keyHash := hashData(key)
	proof := &Proof{
		Siblings: make([][]byte, 0, smt.depth),  // 预分配容量以提高效率
		Path:     make([]bool, 0, smt.depth),
	}
	smt.generateProof(smt.root, keyHash, 0, proof)
	return proof
}

// generateProof 递归生成证明
// 这是 GenerateProof 方法的内部递归实现
// 参数:
//   node: 当前处理的节点
//   keyHash: 键的哈希值
//   depth: 当前深度
//   proof: 正在构建的证明对象（通过指针修改）
// 工作原理:
//   在每一层，记录兄弟节点的哈希和当前的路径方向
//   如果向左走，记录右兄弟；如果向右走，记录左兄弟
func (smt *SparseMerkleTree) generateProof(node *Node, keyHash []byte, depth int, proof *Proof) {
	// 到达叶子节点层或遇到空节点，停止递归
	if depth == smt.depth || node == nil {
		return
	}

	bit := getBit(keyHash, depth)
	proof.Path = append(proof.Path, bit)  // 记录路径方向

	if bit { // 往右走，记录左兄弟节点的哈希
		if node.left != nil {
			proof.Siblings = append(proof.Siblings, node.left.hash)
		} else {
			// 如果左兄弟不存在，使用空节点的哈希
			proof.Siblings = append(proof.Siblings, newEmptyNode().hash)
		}
		smt.generateProof(node.right, keyHash, depth+1, proof)
	} else { // 往左走，记录右兄弟节点的哈希
		if node.right != nil {
			proof.Siblings = append(proof.Siblings, node.right.hash)
		} else {
			// 如果右兄弟不存在，使用空节点的哈希
			proof.Siblings = append(proof.Siblings, newEmptyNode().hash)
		}
		smt.generateProof(node.left, keyHash, depth+1, proof)
	}
}

// VerifyProof 验证 Merkle 证明
// 验证给定的键值对是否存在于树中，无需访问整棵树
// 参数:
//   key: 要验证的键
//   value: 要验证的值
//   proof: 之前生成的 Merkle 证明
// 返回:
//   true 表示证明有效（键值对存在于树中），false 表示证明无效
// 工作原理:
//   1. 从叶子节点的值哈希开始
//   2. 使用证明中的兄弟节点哈希，逐层向上计算父节点哈希
//   3. 最终得到计算出的根哈希
//   4. 将计算出的根哈希与树的实际根哈希比较
// 注意：这个验证过程是从叶子向根进行的，所以需要从 Siblings 数组的末尾开始遍历
func (smt *SparseMerkleTree) VerifyProof(key, value []byte, proof *Proof) bool {
	valueHash := hashData(value)

	// 从叶子节点的值哈希开始
	currentHash := valueHash

	// 从底部往上计算，重建根哈希
	// 遍历 Siblings 数组（从后往前，因为生成时是从上往下记录的）
	for i := len(proof.Siblings) - 1; i >= 0; i-- {
		sibling := proof.Siblings[i]  // 当前层的兄弟节点哈希
		path := proof.Path[i]         // 当前层的路径方向

		if path { // 当前节点在右边，兄弟节点在左边
			currentHash = hashNodes(sibling, currentHash)  // Hash(左 || 右)
		} else { // 当前节点在左边，兄弟节点在右边
			currentHash = hashNodes(currentHash, sibling)  // Hash(左 || 右)
		}
	}

	// 比较计算出的根哈希与树的实际根哈希
	// 如果匹配，说明该键值对确实存在于树中
	return string(currentHash) == string(smt.root.hash)
}

// GetRoot 获取根节点哈希
// 返回树的根哈希，可用于验证整棵树的完整性
// 任何对树的修改都会导致根哈希的变化
func (smt *SparseMerkleTree) GetRoot() []byte {
	return smt.root.hash
}

// PrintTree 打印树结构（用于调试）
// 以层次结构的形式打印整棵树，便于理解树的结构
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
