package print_values

// PrintAddValue()のテスト
func ExamplePrintAddValue() {
	// テストに際し、テスト対象関数に与える引数を定義
	x, y := 2, 3

	// テスト対象関数を実行
  PrintAddValue(x,y)

	// 期待結果

  // Output:
  // 5
}

// PrintSubValue()のテスト
func ExamplePrintSubValue() {
	// テストに際し、テスト対象関数に与える引数を定義
	x, y := 2, 3

	// テスト対象関数を実行
  PrintSubValue(x,y)

	// 期待結果
	
  // Output:
  // -1
}