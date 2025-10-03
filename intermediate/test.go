package intermediate

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"
)

// ==================== 数据结构定义 ====================

// BitcoinTransaction 比特币交易结构
// 代表从比特币网络获取的交易数据，包含输入输出和元数据
type BitcoinTransaction struct {
	TxID      string                // 交易哈希（64字节十六进制字符串）
	BlockHash string                // 所在区块哈希
	BlockTime int64                 // 区块时间戳（Unix时间）
	Inputs    []TransactionInput    // 交易输入列表
	Outputs   []TransactionOutput   // 交易输出列表
}

// TransactionInput 交易输入结构
// 表示一个UTXO的引用，包含SegWit的Witness数据
type TransactionInput struct {
	TxID    string   // 引用的交易ID
	Vout    int      // 引用的输出索引
	Witness []string // Witness数据（SegWit），铭文通常存储在这里
}

// TransactionOutput 交易输出结构
// 表示交易创建的新UTXO，包含金额和脚本
type TransactionOutput struct {
	Value        int64  // 输出金额（satoshi）
	ScriptPubKey string // 锁定脚本（十六进制）
	Address      string // 接收地址（如果可解析）
}

// BRC20Inscription BRC-20铭文数据结构
// 符合BRC-20协议标准的JSON格式铭文内容
// 参考：https://domo-2.gitbook.io/brc-20-experiment/
type BRC20Inscription struct {
	Protocol  string `json:"p"`              // 协议标识，必须为 "brc-20"
	Operation string `json:"op"`             // 操作类型：deploy（部署）/mint（铸造）/transfer（转账）
	Tick      string `json:"tick"`           // 代币标识符（必须为4个字符）
	Max       string `json:"max,omitempty"`  // 最大供应量（仅在deploy操作中使用）
	Limit     string `json:"lim,omitempty"`  // 单次铸造限额（仅在deploy操作中使用）
	Amount    string `json:"amt,omitempty"`  // 操作数量（mint和transfer操作中使用）
	Decimals  string `json:"dec,omitempty"`  // 小数位数（仅在deploy操作中使用，可选）
}

// RuneInscription 符文（Runes）铭文结构
// Bitcoin Runes协议的数据结构，用于原生代币
// 参考：https://docs.ordinals.com/runes.html
type RuneInscription struct {
	RuneID       string // 符文唯一标识符（格式：block:tx）
	Operation    string // 操作类型：etch（刻蚀）/mint（铸造）/transfer（转账）
	RuneName     string // 符文名称（可包含特殊字符）
	Symbol       string // 符文符号（单个字符）
	Divisibility int    // 可分割性（小数位数，0-38）
	Amount       uint64 // 操作数量
	Premine      uint64 // 预铸造数量（仅在etch操作中）
}

// InscriptionResult 铭文解析结果
// 统一的解析结果结构，用于返回不同类型铭文的解析状态
type InscriptionResult struct {
	TxID      string      // 交易ID
	BlockHash string      // 区块哈希
	BlockTime time.Time   // 区块时间
	Type      string      // 铭文类型："brc-20", "rune", "ordinal", "unknown"
	Content   interface{} // 解析后的内容（具体类型取决于Type）
	RawData   string      // 原始数据（十六进制或字符串）
	IsValid   bool        // 是否通过格式验证
	ErrorMsg  string      // 错误信息（如果IsValid为false）
}

// BRC20Token BRC-20代币完整信息
// 维护一个BRC-20代币的完整状态，包括供应量、持有人和交易历史
type BRC20Token struct {
	Tick          string                    // 代币标识符
	MaxSupply     string                    // 最大供应量（字符串格式以避免精度问题）
	MintLimit     string                    // 单次铸造限额
	Decimals      string                    // 小数位数
	TotalMinted   string                    // 已铸造总量
	Holders       map[string]string         // 持有人余额映射：address -> balance
	DeployTxID    string                    // 部署交易ID
	DeployTime    time.Time                 // 部署时间
	Transactions  []BRC20Transaction        // 所有相关交易历史
}

// BRC20Transaction BRC-20交易记录
// 记录BRC-20代币的单笔操作历史
type BRC20Transaction struct {
	TxID      string    // 交易ID
	From      string    // 发送方地址（mint操作时为"mint"）
	To        string    // 接收方地址
	Amount    string    // 交易数量（字符串格式）
	Operation string    // 操作类型：deploy/mint/transfer
	Timestamp time.Time // 交易时间
}

// InscriptionParser 铭文解析器主结构
// 维护所有已解析的BRC-20代币和符文的状态
type InscriptionParser struct {
	BRC20Tokens map[string]*BRC20Token      // BRC-20代币映射：tick -> token info
	RuneTokens  map[string]*RuneInscription // 符文映射：runeID -> rune info
}

// ==================== 核心功能实现 ====================

// NewInscriptionParser 创建新的铭文解析器
// 初始化BRC-20代币和符文的存储映射
// 返回: *InscriptionParser - 新的解析器实例
func NewInscriptionParser() *InscriptionParser {
	return &InscriptionParser{
		BRC20Tokens: make(map[string]*BRC20Token),
		RuneTokens:  make(map[string]*RuneInscription),
	}
}

// ScanBlock 扫描比特币区块，寻找带有Ordinals或符文数据的交易
// 这是解析器的主要入口点，负责批量处理区块中的所有交易
// 参数:
//   - blockHash: 区块哈希值
//   - transactions: 区块中的所有交易列表
// 返回: []InscriptionResult - 解析结果列表，只包含找到铭文的交易
func (ip *InscriptionParser) ScanBlock(blockHash string, transactions []BitcoinTransaction) []InscriptionResult {
	results := make([]InscriptionResult, 0)

	fmt.Printf("正在扫描区块: %s\n", blockHash)
	fmt.Printf("交易数量: %d\n\n", len(transactions))

	// 遍历区块中的每笔交易
	for _, tx := range transactions {
		// 尝试解析交易中的铭文数据
		result := ip.ParseTransaction(tx)
		if result != nil {
			results = append(results, *result)

			// 如果解析成功且格式有效，则处理铭文（更新状态）
			if result.IsValid {
				ip.ProcessInscription(result)
			}
		}
	}

	return results
}

// ParseTransaction 解析单笔交易，提取Witness数据和OP_RETURN数据
// 寻找可能包含铭文的数据位置：
// 1. SegWit Witness数据（Ordinals铭文通常在这里）
// 2. OP_RETURN输出（符文数据可能在这里）
// 参数:
//   - tx: 要解析的比特币交易
// 返回: *InscriptionResult - 如果找到铭文则返回解析结果，否则返回nil
func (ip *InscriptionParser) ParseTransaction(tx BitcoinTransaction) *InscriptionResult {
	// 优先检查交易输入的Witness数据（Ordinals铭文）
	for _, input := range tx.Inputs {
		if len(input.Witness) > 0 {
			// 尝试从Witness数据中提取铭文
			inscriptionData := ip.ExtractInscriptionFromWitness(input.Witness)
			if inscriptionData != "" {
				// 找到铭文数据，进行格式解析
				return ip.ParseInscriptionData(tx, inscriptionData)
			}
		}
	}

	// 检查交易输出中的OP_RETURN脚本（符文协议）
	for _, output := range tx.Outputs {
		// OP_RETURN脚本以"6a"开头
		if strings.HasPrefix(output.ScriptPubKey, "6a") {
			runeData := ip.ExtractRuneFromOpReturn(output.ScriptPubKey)
			if runeData != "" {
				return ip.ParseRuneData(tx, runeData)
			}
		}
	}

	// 未找到任何铭文数据
	return nil
}

// ExtractInscriptionFromWitness 从Witness数据中提取铭文内容
// Ordinals铭文嵌入在SegWit交易的Witness字段中
// 标准格式: OP_FALSE OP_IF "ord" OP_1 "content-type" OP_0 "content" OP_ENDIF
// 参数:
//   - witness: Witness数据数组（十六进制字符串）
// 返回: string - 提取的铭文内容，如果未找到则返回空字符串
func (ip *InscriptionParser) ExtractInscriptionFromWitness(witness []string) string {
	// Ordinals铭文通常在Witness的最后一个元素
	if len(witness) < 2 {
		return ""
	}

	// 获取最后一个witness元素（通常包含脚本数据）
	lastWitness := witness[len(witness)-1]

	// 增强的输入验证
	if len(lastWitness)%2 != 0 || len(lastWitness) == 0 {
		return "" // 无效的十六进制字符串
	}

	// 将十六进制字符串解码为字节数组
	data, err := hex.DecodeString(lastWitness)
	if err != nil {
		fmt.Printf("Witness数据解码失败: %v\n", err)
		return ""
	}

	// 防止过大的数据导致性能问题
	if len(data) > 100*1024 { // 限制为100KB
		fmt.Printf("Witness数据过大: %d bytes\n", len(data))
		return ""
	}

	// 查找Ordinals协议标记"ord"
	ordMarker := []byte("ord")
	if idx := indexOf(data, ordMarker); idx != -1 {
		// 找到ord标记，提取后续的内容数据
		content := ip.extractContent(data[idx+len(ordMarker):])
		return content
	}

	return ""
}

// extractContent 从脚本数据中提取铭文内容
// 这是一个简化的实现，实际的Ordinals脚本解析更复杂
// 主要查找JSON格式的BRC-20数据
// 参数:
//   - data: 脚本字节数据
// 返回: string - 提取的内容字符串
func (ip *InscriptionParser) extractContent(data []byte) string {
	// 防止处理过大的数据
	if len(data) > 10*1024 { // 限制为10KB
		return ""
	}

	// 简化实现：查找JSON格式的内容（以{开始，以}结束）
	// 实际实现需要完整解析Bitcoin脚本操作码结构

	startIdx := -1
	endIdx := -1
	braceCount := 0

	// 查找JSON对象的边界，支持嵌套结构
	for i := 0; i < len(data); i++ {
		if data[i] == '{' {
			if startIdx == -1 {
				startIdx = i
			}
			braceCount++
		} else if data[i] == '}' {
			braceCount--
			if braceCount == 0 && startIdx != -1 {
				endIdx = i + 1
				break
			}
		}
	}

	// 如果找到完整的JSON对象，返回该部分
	if startIdx != -1 && endIdx != -1 && endIdx > startIdx {
		jsonStr := string(data[startIdx:endIdx])
		// 验证是否为有效的JSON
		var test json.RawMessage
		if json.Unmarshal([]byte(jsonStr), &test) == nil {
			return jsonStr
		}
	}

	// 否则尝试将整个数据作为字符串返回（如果是可打印字符）
	if ip.isPrintableString(data) {
		return string(data)
	}

	return ""
}

// isPrintableString 检查字节数组是否包含可打印字符
// 参数:
//   - data: 要检查的字节数组
// 返回: bool - 如果主要包含可打印字符返回true
func (ip *InscriptionParser) isPrintableString(data []byte) bool {
	if len(data) == 0 {
		return false
	}

	printableCount := 0
	for _, b := range data {
		if (b >= 32 && b <= 126) || b == '\n' || b == '\r' || b == '\t' {
			printableCount++
		}
	}

	// 如果80%以上是可打印字符，认为是文本
	return float64(printableCount)/float64(len(data)) >= 0.8
}

// ExtractRuneFromOpReturn 从OP_RETURN脚本中提取符文数据
// 符文协议使用OP_RETURN输出来存储协议数据
// OP_RETURN格式: 6a + 长度 + 数据
// 参数:
//   - scriptPubKey: 输出脚本的十六进制字符串
// 返回: string - 提取的符文数据，如果解析失败则返回空字符串
func (ip *InscriptionParser) ExtractRuneFromOpReturn(scriptPubKey string) string {
	// 验证是否为OP_RETURN脚本
	if !strings.HasPrefix(scriptPubKey, "6a") {
		return ""
	}

	// 去掉OP_RETURN操作码前缀（6a）
	dataHex := scriptPubKey[2:]

	// 解码十六进制数据
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		return ""
	}

	// 简化实现：直接返回解码后的数据
	// 实际的符文协议解析需要处理更复杂的数据结构
	return string(data)
}

// ParseInscriptionData 解析铭文数据，识别BRC-20或其他类型
// 尝试将原始数据解析为已知的铭文格式
// 参数:
//   - tx: 包含铭文的交易
//   - data: 提取的铭文数据字符串
// 返回: *InscriptionResult - 解析结果，包含类型识别和验证信息
func (ip *InscriptionParser) ParseInscriptionData(tx BitcoinTransaction, data string) *InscriptionResult {
	result := &InscriptionResult{
		TxID:      tx.TxID,
		BlockHash: tx.BlockHash,
		BlockTime: time.Unix(tx.BlockTime, 0),
		RawData:   data,
		IsValid:   false,
	}

	// 尝试解析为BRC-20格式
	var brc20 BRC20Inscription
	if err := json.Unmarshal([]byte(data), &brc20); err == nil {
		if brc20.Protocol == "brc-20" {
			// 验证BRC-20格式的完整性和有效性
			if ip.ValidateBRC20(brc20) {
				result.Type = "brc-20"
				result.Content = brc20
				result.IsValid = true
				return result
			} else {
				result.ErrorMsg = "BRC-20格式验证失败"
			}
		}
	}

	// 如果不是BRC-20，标记为普通Ordinals铭文
	result.Type = "ordinal"
	result.Content = data
	result.IsValid = true

	return result
}

// ParseRuneData 解析符文协议数据
// 将OP_RETURN中的数据解析为符文格式
// 参数:
//   - tx: 包含符文的交易
//   - data: 提取的符文数据
// 返回: *InscriptionResult - 符文解析结果
func (ip *InscriptionParser) ParseRuneData(tx BitcoinTransaction, data string) *InscriptionResult {
	result := &InscriptionResult{
		TxID:      tx.TxID,
		BlockHash: tx.BlockHash,
		BlockTime: time.Unix(tx.BlockTime, 0),
		Type:      "rune",
		RawData:   data,
		IsValid:   false,
	}

	// 简化的符文解析（实际需要根据符文协议规范进行复杂解析）
	rune := &RuneInscription{
		RuneID:    fmt.Sprintf("%s:0", tx.TxID), // 简化的ID格式
		Operation: "unknown",                    // 需要从数据中解析
		RuneName:  data,                        // 简化处理
	}

	result.Content = rune
	result.IsValid = true

	return result
}

// ValidateBRC20 验证BRC-20铭文格式的有效性
// 检查必需字段和字段格式是否符合BRC-20协议规范
// 参数:
//   - brc20: 要验证的BRC-20铭文结构
// 返回: bool - 如果格式有效返回true，否则返回false
func (ip *InscriptionParser) ValidateBRC20(brc20 BRC20Inscription) bool {
	// 检查协议标识
	if brc20.Protocol != "brc-20" {
		return false
	}

	// 检查代币标识符（必须为4个字符）
	if brc20.Tick == "" || len(brc20.Tick) != 4 {
		return false
	}

	// 根据操作类型验证必需字段
	switch brc20.Operation {
	case "deploy":
		// 部署操作必须包含最大供应量和铸造限额
		if brc20.Max == "" || brc20.Limit == "" {
			return false
		}
		// 验证数值格式
		return ip.validateNumericString(brc20.Max) && ip.validateNumericString(brc20.Limit)
	case "mint":
		// 铸造操作必须包含数量
		if brc20.Amount == "" {
			return false
		}
		return ip.validateNumericString(brc20.Amount)
	case "transfer":
		// 转账操作必须包含数量
		if brc20.Amount == "" {
			return false
		}
		return ip.validateNumericString(brc20.Amount)
	default:
		// 未知操作类型
		return false
	}
}

// validateNumericString 验证数值字符串的有效性
// 检查字符串是否表示有效的正数
// 参数:
//   - numStr: 要验证的数值字符串
// 返回: bool - 如果是有效数值返回true
func (ip *InscriptionParser) validateNumericString(numStr string) bool {
	// 检查空字符串
	if numStr == "" {
		return false
	}

	// 尝试解析为大整数
	num, ok := new(big.Int).SetString(numStr, 10)
	if !ok {
		return false
	}

	// 必须为正数
	return num.Sign() > 0
}

// ProcessInscription 处理有效的铭文，更新解析器状态
// 根据铭文类型调用相应的处理函数来更新代币状态
// 参数:
//   - result: 已验证的铭文解析结果
func (ip *InscriptionParser) ProcessInscription(result *InscriptionResult) {
	if result.Type == "brc-20" {
		brc20, ok := result.Content.(BRC20Inscription)
		if !ok {
			return
		}

		// 根据BRC-20操作类型进行相应处理
		switch brc20.Operation {
		case "deploy":
			ip.ProcessBRC20Deploy(result.TxID, result.BlockTime, brc20)
		case "mint":
			ip.ProcessBRC20Mint(result.TxID, result.BlockTime, brc20)
		case "transfer":
			ip.ProcessBRC20Transfer(result.TxID, result.BlockTime, brc20)
		}
	}
}

// ProcessBRC20Deploy 处理BRC-20代币部署操作
// 创建新的代币记录，如果代币已存在则忽略
// 参数:
//   - txID: 部署交易ID
//   - blockTime: 部署时间
//   - brc20: BRC-20部署数据
func (ip *InscriptionParser) ProcessBRC20Deploy(txID string, blockTime time.Time, brc20 BRC20Inscription) {
	// 检查代币是否已经部署（BRC-20协议：首次部署有效原则）
	if _, exists := ip.BRC20Tokens[brc20.Tick]; exists {
		fmt.Printf("代币 %s 已经部署，忽略重复部署\n", brc20.Tick)
		return
	}

	// 创建新代币记录
	token := &BRC20Token{
		Tick:         brc20.Tick,
		MaxSupply:    brc20.Max,
		MintLimit:    brc20.Limit,
		Decimals:     brc20.Decimals,
		TotalMinted:  "0", // 初始铸造量为0
		Holders:      make(map[string]string),
		DeployTxID:   txID,
		DeployTime:   blockTime,
		Transactions: make([]BRC20Transaction, 0),
	}

	ip.BRC20Tokens[brc20.Tick] = token
	fmt.Printf("✓ 部署代币: %s, 最大供应量: %s\n", brc20.Tick, brc20.Max)
}

// ProcessBRC20Mint 处理BRC-20代币铸造操作
// 记录铸造交易，验证铸造限额和总供应量
// 参数:
//   - txID: 铸造交易ID
//   - blockTime: 铸造时间
//   - brc20: BRC-20铸造数据
func (ip *InscriptionParser) ProcessBRC20Mint(txID string, blockTime time.Time, brc20 BRC20Inscription) {
	token, exists := ip.BRC20Tokens[brc20.Tick]
	if !exists {
		fmt.Printf("代币 %s 未部署，无法铸造\n", brc20.Tick)
		return
	}

	// 验证铸造数量
	mintAmount, ok := new(big.Int).SetString(brc20.Amount, 10)
	if !ok || mintAmount.Sign() <= 0 {
		fmt.Printf("铸造数量无效: %s\n", brc20.Amount)
		return
	}

	// 验证单次铸造限额
	mintLimit, ok := new(big.Int).SetString(token.MintLimit, 10)
	if !ok {
		fmt.Printf("代币 %s 铸造限额配置错误\n", brc20.Tick)
		return
	}

	if mintAmount.Cmp(mintLimit) > 0 {
		fmt.Printf("铸造数量 %s 超过单次限额 %s\n", brc20.Amount, token.MintLimit)
		return
	}

	// 验证总供应量限制
	maxSupply, ok := new(big.Int).SetString(token.MaxSupply, 10)
	if !ok {
		fmt.Printf("代币 %s 最大供应量配置错误\n", brc20.Tick)
		return
	}

	totalMinted, ok := new(big.Int).SetString(token.TotalMinted, 10)
	if !ok {
		totalMinted = big.NewInt(0) // 如果解析失败，默认为0
	}

	newTotal := new(big.Int).Add(totalMinted, mintAmount)
	if newTotal.Cmp(maxSupply) > 0 {
		fmt.Printf("铸造后总量 %s 将超过最大供应量 %s\n", newTotal.String(), token.MaxSupply)
		return
	}

	// 更新总铸造量
	token.TotalMinted = newTotal.String()

	// 记录铸造交易
	tx := BRC20Transaction{
		TxID:      txID,
		From:      "mint",                // 铸造操作的发送方标记为"mint"
		To:        "minter_address",      // 实际应该从交易输出中提取接收地址
		Amount:    brc20.Amount,
		Operation: "mint",
		Timestamp: blockTime,
	}

	token.Transactions = append(token.Transactions, tx)
	fmt.Printf("✓ 铸造代币: %s, 数量: %s, 总铸造量: %s\n", brc20.Tick, brc20.Amount, token.TotalMinted)
}

// ProcessBRC20Transfer 处理BRC-20代币转账操作
// 记录转账交易，实际应用中需要验证余额和更新持有人状态
// 参数:
//   - txID: 转账交易ID
//   - blockTime: 转账时间
//   - brc20: BRC-20转账数据
func (ip *InscriptionParser) ProcessBRC20Transfer(txID string, blockTime time.Time, brc20 BRC20Inscription) {
	token, exists := ip.BRC20Tokens[brc20.Tick]
	if !exists {
		fmt.Printf("代币 %s 未部署，无法转账\n", brc20.Tick)
		return
	}

	// 记录转账交易
	// 注意：这里缺少实际的余额验证和更新逻辑
	tx := BRC20Transaction{
		TxID:      txID,
		From:      "sender_address",      // 实际应该从交易输入中提取发送地址
		To:        "receiver_address",    // 实际应该从交易输出中提取接收地址
		Amount:    brc20.Amount,
		Operation: "transfer",
		Timestamp: blockTime,
	}

	token.Transactions = append(token.Transactions, tx)
	fmt.Printf("✓ 转账代币: %s, 数量: %s\n", brc20.Tick, brc20.Amount)
}

// ==================== 输出和查询功能 ====================

// PrintBRC20TokenInfo 打印指定BRC-20代币的详细信息
// 包含供应量、持有人、交易历史等完整信息
// 参数:
//   - tick: 要查询的代币标识符
func (ip *InscriptionParser) PrintBRC20TokenInfo(tick string) {
	token, exists := ip.BRC20Tokens[tick]
	if !exists {
		fmt.Printf("代币 %s 不存在\n", tick)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("BRC-20 代币信息: %s\n", tick)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("最大供应量:     %s\n", token.MaxSupply)
	fmt.Printf("铸造限额:       %s\n", token.MintLimit)
	fmt.Printf("小数位数:       %s\n", token.Decimals)
	fmt.Printf("已铸造总量:     %s\n", token.TotalMinted)
	fmt.Printf("部署交易ID:     %s\n", token.DeployTxID)
	fmt.Printf("部署时间:       %s\n", token.DeployTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("持有人数量:     %d\n", len(token.Holders))
	fmt.Printf("交易历史数量:   %d\n", len(token.Transactions))

	// 显示最近的交易历史
	if len(token.Transactions) > 0 {
		fmt.Println("\n最近的交易:")
		for i, tx := range token.Transactions {
			if i >= 5 { // 只显示最近5条
				break
			}
			fmt.Printf("  %d. [%s] %s -> %s: %s (TxID: %s...)\n",
				i+1, tx.Operation, tx.From, tx.To, tx.Amount, tx.TxID[:16])
		}
	}

	fmt.Println(strings.Repeat("=", 60))
}

// PrintRuneInfo 打印指定符文的详细信息
// 参数:
//   - runeID: 要查询的符文ID
func (ip *InscriptionParser) PrintRuneInfo(runeID string) {
	rune, exists := ip.RuneTokens[runeID]
	if !exists {
		fmt.Printf("符文 %s 不存在\n", runeID)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("符文信息: %s\n", runeID)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("符文名称:       %s\n", rune.RuneName)
	fmt.Printf("符文符号:       %s\n", rune.Symbol)
	fmt.Printf("可分割性:       %d\n", rune.Divisibility)
	fmt.Printf("铸造数量:       %d\n", rune.Amount)
	fmt.Printf("预铸造量:       %d\n", rune.Premine)
	fmt.Println(strings.Repeat("=", 60))
}

// PrintAllTokens 打印所有已解析的BRC-20代币列表
// 提供代币概览信息
func (ip *InscriptionParser) PrintAllTokens() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("所有BRC-20代币列表")
	fmt.Println(strings.Repeat("=", 60))

	if len(ip.BRC20Tokens) == 0 {
		fmt.Println("暂无代币")
		return
	}

	// 遍历并显示所有代币的基本信息
	for tick, token := range ip.BRC20Tokens {
		fmt.Printf("%-6s | 供应量: %-15s | 交易数: %d\n",
			tick, token.MaxSupply, len(token.Transactions))
	}

	fmt.Println(strings.Repeat("=", 60))
}

// ==================== 工具函数 ====================

// indexOf 在字节数组中查找子切片的位置
// 简单的字节模式匹配算法，用于在铭文数据中查找特定标记
// 参数:
//   - data: 要搜索的字节数组
//   - pattern: 要查找的子切片模式
// 返回: int - 如果找到返回第一个匹配位置的索引，否则返回-1
func indexOf(data []byte, pattern []byte) int {
	// 遍历数据，寻找模式匹配
	for i := 0; i <= len(data)-len(pattern); i++ {
		match := true
		// 检查当前位置是否匹配整个模式
		for j := 0; j < len(pattern); j++ {
			if data[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

// ==================== 测试和演示 ====================

// main 主函数，演示BRC-20铭文解析器的完整功能
// 包含多个场景：区块扫描、代币部署、铸造、转账等操作的模拟
func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          BRC-20 铭文解析器 (Inscription Parser)           ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// 创建解析器实例
	parser := NewInscriptionParser()

	// 场景1：模拟比特币区块扫描，包含BRC-20铭文的交易
	fmt.Println("\n【场景1: 扫描区块并解析BRC-20铭文】")

	// 构造测试交易数据（模拟真实的比特币交易结构）
	transactions := []BitcoinTransaction{
		// 交易1: 部署ORDI代币
		{
			TxID:      "abc123def456789012345678901234567890123456789012345678901234",
			BlockHash: "000000000000000000001234567890abcdef",
			BlockTime: time.Now().Unix(),
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id",
					Vout: 0,
					// 模拟包含BRC-20部署铭文的Witness数据
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"deploy","tick":"ordi","max":"21000000","lim":"1000","dec":"8"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
		// 交易2: 铸造ORDI代币
		{
			TxID:      "def456abc789012345678901234567890123456789012345678901234567",
			BlockHash: "000000000000000000001234567890abcdef",
			BlockTime: time.Now().Unix() + 600, // 10分钟后
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id_2",
					Vout: 0,
					// 模拟包含BRC-20铸造铭文的Witness数据
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"mint","tick":"ordi","amt":"1000"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
		// 交易3: 转账ORDI代币
		{
			TxID:      "ghi789jkl012345678901234567890123456789012345678901234567890",
			BlockHash: "000000000000000000001234567890abcdef",
			BlockTime: time.Now().Unix() + 1200, // 20分钟后
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id_3",
					Vout: 0,
					// 模拟包含BRC-20转账铭文的Witness数据
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"transfer","tick":"ordi","amt":"500"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
	}

	// 执行区块扫描
	results := parser.ScanBlock("000000000000000000001234567890abcdef", transactions)

	// 场景2：显示解析结果汇总
	fmt.Println("\n【解析结果汇总】")
	for i, result := range results {
		fmt.Printf("\n交易 %d:\n", i+1)
		fmt.Printf("  TxID:     %s...\n", result.TxID[:32])
		fmt.Printf("  类型:     %s\n", result.Type)
		fmt.Printf("  是否有效: %t\n", result.IsValid)

		// 如果是有效的BRC-20铭文，显示详细信息
		if result.IsValid && result.Type == "brc-20" {
			brc20 := result.Content.(BRC20Inscription)
			fmt.Printf("  操作:     %s\n", brc20.Operation)
			fmt.Printf("  代币:     %s\n", brc20.Tick)
		}
	}

	// 场景3：查询特定代币的详细信息
	fmt.Println("\n【场景2: 查询BRC-20代币详情】")
	parser.PrintBRC20TokenInfo("ordi")

	// 场景4：显示所有代币概览
	fmt.Println("\n【场景3: 列出所有代币】")
	parser.PrintAllTokens()

	// 场景5：演示部署更多代币
	fmt.Println("\n【场景4: 部署更多代币】")

	// 创建新的交易来部署SATS代币
	moreTxs := []BitcoinTransaction{
		{
			TxID:      "sats001234567890123456789012345678901234567890123456789012345",
			BlockHash: "000000000000000000002345678901bcdefg",
			BlockTime: time.Now().Unix() + 1800, // 30分钟后
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id_4",
					Vout: 0,
					// 部署SATS代币的铭文
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"deploy","tick":"sats","max":"2100000000000000","lim":"100000","dec":"8"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
	}

	// 扫描新区块
	parser.ScanBlock("000000000000000000002345678901bcdefg", moreTxs)
	// 再次显示所有代币（现在应该包含SATS）
	parser.PrintAllTokens()

	fmt.Println("\n✓ 铭文解析器演示完成")
}

// createBRC20WitnessData 创建BRC-20 Witness数据（模拟函数）
// 将JSON格式的BRC-20数据编码为模拟的Witness格式
// 实际的Ordinals铭文格式更复杂，包含完整的Bitcoin脚本结构
// 参数:
//   - jsonData: JSON格式的BRC-20铭文内容
// 返回: string - 模拟的十六进制Witness数据
func createBRC20WitnessData(jsonData string) string {
	// 模拟Ordinals铭文的Witness结构
	// 实际格式：OP_FALSE OP_IF "ord" OP_1 "content-type" OP_0 "content" OP_ENDIF
	// 这里简化为：前缀"ord" + JSON内容的十六进制编码
	prefix := "6f7264"                             // "ord"的十六进制编码
	content := hex.EncodeToString([]byte(jsonData)) // JSON内容的十六进制编码
	return prefix + content
}
