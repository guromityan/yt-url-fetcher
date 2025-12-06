# Feature Specification: Fetch YouTube URLs

**Feature Branch**: `001-fetch-youtube-urls`  
**Created**: 2025-12-06  
**Status**: Draft  
**Input**: User description: "NotebookLM を有効活用するため、YouTube ソースを充実にインポートできるようにしたいです。YouTube のチャンネルや、特定のプレイリストを指定したらその動画の URL の一覧を取得できるようにしたいです。"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - チャンネルからの動画一覧取得 (Priority: P1)

ユーザーは特定のYouTubeチャンネルの全ての公開動画のURLリストを取得したい。これにより、チャンネル全体をNotebookLMのソースとして一括インポートできる。

**Why this priority**: チャンネル単位でのインポートは最も基本的かつ需要が高いユースケースであるため。

**Independent Test**: 公開されているYouTubeチャンネルのURLを入力し、そのチャンネルに含まれる動画のURLリストが標準出力に出力されることを確認する。

**Acceptance Scenarios**:

1. **Given** 有効なYouTubeチャンネルのURLまたはID, **When** ツールを実行する, **Then** そのチャンネルの公開動画のURL一覧が標準出力に表示される
2. **Given** 動画が存在しないチャンネル, **When** ツールを実行する, **Then** 空のリストが表示され、エラーにはならない
3. **Given** 存在しないチャンネルID, **When** ツールを実行する, **Then** 適切なエラーメッセージが標準エラー出力に表示される

---

### User Story 2 - プレイリストからの動画一覧取得 (Priority: P2)

ユーザーは特定のYouTubeプレイリストに含まれる動画のURLリストを取得したい。これにより、トピックごとに整理された動画群をNotebookLMにインポートできる。

**Why this priority**: チャンネル全体ではなく、特定のテーマに絞ったインポートを可能にするため重要だが、チャンネル全体の次に来るニーズである。

**Independent Test**: 公開されているYouTubeプレイリストのURLを入力し、そのリストに含まれる動画のURL一覧が出力されることを確認する。

**Acceptance Scenarios**:

1. **Given** 有効なYouTubeプレイリストのURLまたはID, **When** ツールを実行する, **Then** そのプレイリスト内の動画URL一覧が標準出力に表示される
2. **Given** 非公開または存在しないプレイリスト, **When** ツールを実行する, **Then** 適切なエラーメッセージが表示される

### Edge Cases

- チャンネルやプレイリストの動画数が非常に多い場合（ページネーションの処理が必要）
- ライブ配信アーカイブやショート動画が含まれる場合の扱い（現在は区別せずURLを取得する）
- API制限に達した場合の挙動

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: システムは、YouTubeチャンネルのURLまたはIDを入力として受け入れなければならない (MUST)。
- **FR-002**: システムは、YouTubeプレイリストのURLまたはIDを入力として受け入れなければならない (MUST)。
- **FR-003**: システムは、指定されたソース（チャンネル/プレイリスト）に関連付けられた公開動画のURLを抽出できなければならない (MUST)。
- **FR-004**: 出力は、1行につき1つの動画URLという形式のプレーンテキストでなければならない (MUST)。
- **FR-005**: システムは、無効な入力やアクセス不可能なリソースに対して、標準エラー出力にわかりやすいエラーメッセージを表示し、終了コード1で終了しなければならない (MUST)。
- **FR-006**: システムは、YouTube Data API (または同等の手段) を使用して情報を取得する (SHOULD)。APIキーの入力方法は環境変数 `YOUTUBE_API_KEY` を想定する (SHOULD)。

### Key Entities *(include if feature involves data)*

- **Source**: 入力となるYouTubeチャンネルまたはプレイリスト。
- **Video URL**: 出力となる個々の動画のリンク。

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 指定されたチャンネル/プレイリストの公開動画URLを、API制限の範囲内で100%取得できること。
- **SC-002**: 出力されたURLリストをそのままコピー＆ペーストすることで、NotebookLMのソースとして認識されること（余計な装飾がないこと）。
- **SC-003**: 100件程度の動画を持つリストの取得が、APIレイテンシを除き数秒以内に完了すること。