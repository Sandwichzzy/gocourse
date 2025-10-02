package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ==================== 数据结构定义 ====================

// 比特币交易结构
type BitcoinTransaction struct {
	TxID      string
	BlockHash string
	BlockTime int64
	Inputs    []TransactionInput
	Outputs   []TransactionOutput
}

// 交易输入
type TransactionInput struct {
	TxID    string
	Vout    int
	Witness []string // Witness数据（SegWit）
}

// 交易输出
type TransactionOutput struct {
	Value        int64
	ScriptPubKey string
	Address      string
}

// BRC-20铭文结构
type BRC20Inscription struct {
	Protocol  string `json:"p"`   // 协议标识，通常为 "brc-20"
	Operation string `json:"op"`  // 操作类型：deploy/mint/transfer
	Tick      string `json:"tick"` // 代币标识（4个字符）
	Max       string `json:"max,omitempty"`  // 最大供应量（仅deploy）
	Limit     string `json:"lim,omitempty"`  // 每次铸造限额（仅deploy）
	Amount    string `json:"amt,omitempty"`  // 数量（mint/transfer）
	Decimals  string `json:"dec,omitempty"`  // 小数位数（仅deploy）
}

// 符文（Runes）结构
type RuneInscription struct {
	RuneID      string // 符文ID
	Operation   string // 操作类型：etch/mint/transfer
	RuneName    string // 符文名称
	Symbol      string // 符文符号
	Divisibility int   // 可分割性
	Amount      uint64 // 数量
	Premine     uint64 // 预铸造量
}

// 铭文解析结果
type InscriptionResult struct {
	TxID        string
	BlockHash   string
	BlockTime   time.Time
	Type        string // "brc-20", "rune", "ordinal", "unknown"
	Content     interface{}
	RawData     string
	IsValid     bool
	ErrorMsg    string
}

// BRC-20代币信息
type BRC20Token struct {
	Tick          string
	MaxSupply     string
	MintLimit     string
	Decimals      string
	TotalMinted   string
	Holders       map[string]string // address -> balance
	DeployTxID    string
	DeployTime    time.Time
	Transactions  []BRC20Transaction
}

// BRC-20交易记录
type BRC20Transaction struct {
	TxID      string
	From      string
	To        string
	Amount    string
	Operation string
	Timestamp time.Time
}

// 铭文解析器
type InscriptionParser struct {
	BRC20Tokens map[string]*BRC20Token // tick -> token info
	RuneTokens  map[string]*RuneInscription
}

// ==================== 核心功能实现 ====================

// 创建新的铭文解析器
func NewInscriptionParser() *InscriptionParser {
	return &InscriptionParser{
		BRC20Tokens: make(map[string]*BRC20Token),
		RuneTokens:  make(map[string]*RuneInscription),
	}
}

// 1. 扫描比特币区块，寻找带有 Ordinals 或符文数据的交易
func (ip *InscriptionParser) ScanBlock(blockHash string, transactions []BitcoinTransaction) []InscriptionResult {
	results := make([]InscriptionResult, 0)
	
	fmt.Printf("正在扫描区块: %s\n", blockHash)
	fmt.Printf("交易数量: %d\n\n", len(transactions))
	
	for _, tx := range transactions {
		result := ip.ParseTransaction(tx)
		if result != nil {
			results = append(results, *result)
			
			// 处理解析结果
			if result.IsValid {
				ip.ProcessInscription(result)
			}
		}
	}
	
	return results
}

// 2. 解析交易，提取 Witness 数据
func (ip *InscriptionParser) ParseTransaction(tx BitcoinTransaction) *InscriptionResult {
	// 遍历所有输入，查找Witness数据
	for _, input := range tx.Inputs {
		if len(input.Witness) > 0 {
			// 尝试解析Witness数据
			inscriptionData := ip.ExtractInscriptionFromWitness(input.Witness)
			if inscriptionData != "" {
				// 找到铭文数据，进行解析
				return ip.ParseInscriptionData(tx, inscriptionData)
			}
		}
	}
	
	// 检查OP_RETURN输出（符文可能在这里）
	for _, output := range tx.Outputs {
		if strings.HasPrefix(output.ScriptPubKey, "6a") { // OP_RETURN
			runeData := ip.ExtractRuneFromOpReturn(output.ScriptPubKey)
			if runeData != "" {
				return ip.ParseRuneData(tx, runeData)
			}
		}
	}
	
	return nil
}

// 3. 从Witness数据中提取铭文
func (ip *InscriptionParser) ExtractInscriptionFromWitness(witness []string) string {
	// Ordinals铭文通常在Witness的最后一个元素
	// 格式: OP_FALSE OP_IF "ord" OP_1 "content-type" OP_0 "content" OP_ENDIF
	
	if len(witness) < 2 {
		return ""
	}
	
	// 获取最后一个witness元素
	lastWitness := witness[len(witness)-1]
	
	// 解码hex数据
	data, err := hex.DecodeString(lastWitness)
	if err != nil {
		return ""
	}
	
	// 查找"ord"标记
	ordMarker := []byte("ord")
	if idx := indexOf(data, ordMarker); idx != -1 {
		// 找到ord标记，提取内容
		content := ip.extractContent(data[idx+len(ordMarker):])
		return content
	}
	
	return ""
}

// 提取铭文内容
func (ip *InscriptionParser) extractContent(data []byte) string {
	// 简化实现：查找JSON格式的内容
	// 实际实现需要解析完整的脚本结构
	
	// 查找 { 和 } 之间的内容
	startIdx := -1
	endIdx := -1
	
	for i := 0; i < len(data); i++ {
		if data[i] == '{' && startIdx == -1 {
			startIdx = i
		}
		if data[i] == '}' {
			endIdx = i + 1
		}
	}
	
	if startIdx != -1 && endIdx != -1 && endIdx > startIdx {
		return string(data[startIdx:endIdx])
	}
	
	// 尝试将整个数据作为字符串返回
	return string(data)
}

// 从OP_RETURN中提取符文数据
func (ip *InscriptionParser) ExtractRuneFromOpReturn(scriptPubKey string) string {
	// OP_RETURN脚本格式: 6a + 数据
	if !strings.HasPrefix(scriptPubKey, "6a") {
		return ""
	}
	
	// 去掉6a前缀
	dataHex := scriptPubKey[2:]
	
	// 解码hex数据
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		return ""
	}
	
	// 查找符文标记（简化实现）
	return string(data)
}

// 4. 解析铭文数据，识别不同类型
func (ip *InscriptionParser) ParseInscriptionData(tx BitcoinTransaction, data string) *InscriptionResult {
	result := &InscriptionResult{
		TxID:      tx.TxID,
		BlockHash: tx.BlockHash,
		BlockTime: time.Unix(tx.BlockTime, 0),
		RawData:   data,
		IsValid:   false,
	}
	
	// 尝试解析为BRC-20
	var brc20 BRC20Inscription
	if err := json.Unmarshal([]byte(data), &brc20); err == nil {
		if brc20.Protocol == "brc-20" {
			// 验证BRC-20格式
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
	
	// 尝试解析为其他类型的铭文
	result.Type = "ordinal"
	result.Content = data
	result.IsValid = true
	
	return result
}

// 解析符文数据
func (ip *InscriptionParser) ParseRuneData(tx BitcoinTransaction, data string) *InscriptionResult {
	result := &InscriptionResult{
		TxID:      tx.TxID,
		BlockHash: tx.BlockHash,
		BlockTime: time.Unix(tx.BlockTime, 0),
		Type:      "rune",
		RawData:   data,
		IsValid:   false,
	}
	
	// 简化的符文解析（实际需要更复杂的协议解析）
	rune := &RuneInscription{
		RuneID:    fmt.Sprintf("%s:0", tx.TxID),
		Operation: "unknown",
		RuneName:  data,
	}
	
	result.Content = rune
	result.IsValid = true
	
	return result
}

// 验证BRC-20格式
func (ip *InscriptionParser) ValidateBRC20(brc20 BRC20Inscription) bool {
	// 检查必需字段
	if brc20.Protocol != "brc-20" {
		return false
	}
	
	if brc20.Tick == "" || len(brc20.Tick) != 4 {
		return false
	}
	
	// 根据操作类型验证
	switch brc20.Operation {
	case "deploy":
		return brc20.Max != "" && brc20.Limit != ""
	case "mint":
		return brc20.Amount != ""
	case "transfer":
		return brc20.Amount != ""
	default:
		return false
	}
}

// 5. 处理铭文，更新状态
func (ip *InscriptionParser) ProcessInscription(result *InscriptionResult) {
	if result.Type == "brc-20" {
		brc20, ok := result.Content.(BRC20Inscription)
		if !ok {
			return
		}
		
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

// 处理BRC-20部署
func (ip *InscriptionParser) ProcessBRC20Deploy(txID string, blockTime time.Time, brc20 BRC20Inscription) {
	// 检查代币是否已经部署
	if _, exists := ip.BRC20Tokens[brc20.Tick]; exists {
		fmt.Printf("代币 %s 已经部署，忽略重复部署\n", brc20.Tick)
		return
	}
	
	// 创建新代币
	token := &BRC20Token{
		Tick:         brc20.Tick,
		MaxSupply:    brc20.Max,
		MintLimit:    brc20.Limit,
		Decimals:     brc20.Decimals,
		TotalMinted:  "0",
		Holders:      make(map[string]string),
		DeployTxID:   txID,
		DeployTime:   blockTime,
		Transactions: make([]BRC20Transaction, 0),
	}
	
	ip.BRC20Tokens[brc20.Tick] = token
	fmt.Printf("✓ 部署代币: %s, 最大供应量: %s\n", brc20.Tick, brc20.Max)
}

// 处理BRC-20铸造
func (ip *InscriptionParser) ProcessBRC20Mint(txID string, blockTime time.Time, brc20 BRC20Inscription) {
	token, exists := ip.BRC20Tokens[brc20.Tick]
	if !exists {
		fmt.Printf("代币 %s 未部署，无法铸造\n", brc20.Tick)
		return
	}
	
	// 记录交易
	tx := BRC20Transaction{
		TxID:      txID,
		From:      "mint",
		To:        "minter_address", // 实际应该从交易中提取
		Amount:    brc20.Amount,
		Operation: "mint",
		Timestamp: blockTime,
	}
	
	token.Transactions = append(token.Transactions, tx)
	fmt.Printf("✓ 铸造代币: %s, 数量: %s\n", brc20.Tick, brc20.Amount)
}

// 处理BRC-20转账
func (ip *InscriptionParser) ProcessBRC20Transfer(txID string, blockTime time.Time, brc20 BRC20Inscription) {
	token, exists := ip.BRC20Tokens[brc20.Tick]
	if !exists {
		fmt.Printf("代币 %s 未部署，无法转账\n", brc20.Tick)
		return
	}
	
	// 记录交易
	tx := BRC20Transaction{
		TxID:      txID,
		From:      "sender_address",   // 实际应该从交易中提取
		To:        "receiver_address", // 实际应该从交易中提取
		Amount:    brc20.Amount,
		Operation: "transfer",
		Timestamp: blockTime,
	}
	
	token.Transactions = append(token.Transactions, tx)
	fmt.Printf("✓ 转账代币: %s, 数量: %s\n", brc20.Tick, brc20.Amount)
}

// 6. 输出解析结果
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

// 打印符文信息
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

// 打印所有代币
func (ip *InscriptionParser) PrintAllTokens() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("所有BRC-20代币列表")
	fmt.Println(strings.Repeat("=", 60))
	
	if len(ip.BRC20Tokens) == 0 {
		fmt.Println("暂无代币")
		return
	}
	
	for tick, token := range ip.BRC20Tokens {
		fmt.Printf("%-6s | 供应量: %-15s | 交易数: %d\n",
			tick, token.MaxSupply, len(token.Transactions))
	}
	
	fmt.Println(strings.Repeat("=", 60))
}

// ==================== 工具函数 ====================

// 查找子切片
func indexOf(data []byte, pattern []byte) int {
	for i := 0; i <= len(data)-len(pattern); i++ {
		match := true
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

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          BRC-20 铭文解析器 (Inscription Parser)           ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	
	// 创建解析器
	parser := NewInscriptionParser()
	
	// 模拟比特币区块数据
	fmt.Println("\n【场景1: 扫描区块并解析BRC-20铭文】")
	
	// 构造测试交易
	transactions := []BitcoinTransaction{
		// 1. 部署代币
		{
			TxID:      "abc123def456789012345678901234567890123456789012345678901234",
			BlockHash: "000000000000000000001234567890abcdef",
			BlockTime: time.Now().Unix(),
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id",
					Vout: 0,
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"deploy","tick":"ordi","max":"21000000","lim":"1000","dec":"8"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
		// 2. 铸造代币
		{
			TxID:      "def456abc789012345678901234567890123456789012345678901234567",
			BlockHash: "000000000000000000001234567890abcdef",
			BlockTime: time.Now().Unix() + 600,
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id_2",
					Vout: 0,
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"mint","tick":"ordi","amt":"1000"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
		// 3. 转账代币
		{
			TxID:      "ghi789jkl012345678901234567890123456789012345678901234567890",
			BlockHash: "000000000000000000001234567890abcdef",
			BlockTime: time.Now().Unix() + 1200,
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id_3",
					Vout: 0,
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"transfer","tick":"ordi","amt":"500"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
	}
	
	// 扫描区块
	results := parser.ScanBlock("000000000000000000001234567890abcdef", transactions)
	
	// 显示解析结果
	fmt.Println("\n【解析结果汇总】")
	for i, result := range results {
		fmt.Printf("\n交易 %d:\n", i+1)
		fmt.Printf("  TxID:     %s...\n", result.TxID[:32])
		fmt.Printf("  类型:     %s\n", result.Type)
		fmt.Printf("  是否有效: %t\n", result.IsValid)
		
		if result.IsValid && result.Type == "brc-20" {
			brc20 := result.Content.(BRC20Inscription)
			fmt.Printf("  操作:     %s\n", brc20.Operation)
			fmt.Printf("  代币:     %s\n", brc20.Tick)
		}
	}
	
	// 显示代币信息
	fmt.Println("\n【场景2: 查询BRC-20代币详情】")
	parser.PrintBRC20TokenInfo("ordi")
	
	// 显示所有代币
	fmt.Println("\n【场景3: 列出所有代币】")
	parser.PrintAllTokens()
	
	// 演示更多代币
	fmt.Println("\n【场景4: 部署更多代币】")
	
	moreTxs := []BitcoinTransaction{
		{
			TxID:      "sats001234567890123456789012345678901234567890123456789012345",
			BlockHash: "000000000000000000002345678901bcdefg",
			BlockTime: time.Now().Unix() + 1800,
			Inputs: []TransactionInput{
				{
					TxID: "prev_tx_id_4",
					Vout: 0,
					Witness: []string{
						"",
						createBRC20WitnessData(`{"p":"brc-20","op":"deploy","tick":"sats","max":"2100000000000000","lim":"100000","dec":"8"}`),
					},
				},
			},
			Outputs: []TransactionOutput{},
		},
	}
	
	parser.ScanBlock("000000000000000000002345678901bcdefg", moreTxs)
	parser.PrintAllTokens()
	
	fmt.Println("\n✓ 铭文解析器演示完成")
}

// 创建BRC-20 Witness数据（模拟）
func createBRC20WitnessData(jsonData string) string {
	// 模拟Ordinals铭文的Witness结构
	// 实际格式更复杂，这里简化处理
	prefix := "6f7264" // "ord" in hex
	content := hex.EncodeToString([]byte(jsonData))
	return prefix + content
}
