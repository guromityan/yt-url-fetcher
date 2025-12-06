# Research: Add Short Flags

## 決定事項

### 1. フラグエイリアスの実装方法
**Decision**: `flag.StringVar` を使用して、同じ変数に対してロングフラグとショートフラグの両方をバインドする。
**Rationale**: 
- Goの標準 `flag` パッケージにはエイリアス機能そのものはないが、同じ変数を複数のフラグ定義の宛先として指定することで実質的なエイリアスを実現できる。
- これにより、後でどのフラグ（ロングかショートか）が使われたかを気にする必要がなくなり、変数の中身だけを見ればよくなる。

**Alternatives Considered**: 
- 別の変数を定義して、パース後にマージする: 複雑になるだけでメリットがない。

```go
// Example Implementation
var channelID string
flag.StringVar(&channelID, "channel", "", "Channel ID")
flag.StringVar(&channelID, "c", "", "Channel ID (short)")
```

### 2. ヘルプメッセージの表示
**Decision**: デフォルトの `flag.Usage` は各フラグを個別に表示するため、同じ説明が2回出てくることになるが、今回は許容する。
**Rationale**: カスタムUsageを実装するほどの規模ではない。シンプルさを優先。
