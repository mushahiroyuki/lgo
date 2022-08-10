// パッケージ money は金銭を管理するための様々なユーティリティを提供する。
package money

// Money は「金額」と「通貨」の組を表す。
type Money struct {
    Value decimal.Decimal
    Currency string
}

// Convert はひとつの通貨を別の通貨に変換する。
// 引数は2つ。ひとつはMoneyのインスタンスで変換元の値。もうひとつは文字列で変換後の通貨を表す。
// 戻り値として変換後の値とエラー（未知あるいは変換不可能な通貨が指定された場合）を返す。
//
// 対応通貨：
//        USD - US Dollar（米ドル）
//        CAD - Canadian Dollar（カナダドル）
//        EUR - Euro（ユーロ）
//        INR - Indian Rupee（インドルピー）
// 為替レートに関する詳しい情報は次のURL参照:
// https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
    // ...
}

