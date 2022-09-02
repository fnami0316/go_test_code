package return_values

import "testing"

// テスト関数名はTestXxxという形でさえあればいいが、「Test+<テスト対象関数名>」という命名にするのが無難そう

// ReturnAddValue()のテスト
func TestReturnAddValue(t *testing.T) {
    // テストに際し、テスト対象関数に与える引数を定義
    x, y := 2, 3

    // 期待結果を定義
    expect := 5

    // テスト対象関数を実行
    got := ReturnAddValue(x, y)
    
    // 実行結果と期待結果を比較し、エラーだったら出力
    if got != expect {
      t.Errorf("exec ReturnAddValue() expect: %v, got: %v, x: %v, y: %v", expect, got, x, y)
    }
}

// ReturnSubValue()のテスト
func TestReturnSubValue(t *testing.T) {
	// テストに際し、テスト対象関数に与える引数を定義
	x, y := 8, 3

	// 期待結果を定義
	expect := 5

	// テスト対象関数を実行
	got := ReturnSubValue(x, y)
	
	// 実行結果と期待結果を比較し、エラーだったら出力
	if got != expect {
		t.Errorf("exec ReturnSubValue() expect: %v, got: %v, x: %v, y: %v", expect, got, x, y)
	}
}

