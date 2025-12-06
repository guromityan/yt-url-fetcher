<!--
Sync Impact Report:
- Version change: 0.0.0 -> 1.0.0 (Initial Ratification)
- Added Principles:
  - I. 単一責任の原則 (Single Responsibility)
  - II. 標準入出力とCLIの重視 (Standard I/O & CLI)
  - III. 日本語優先 (Japanese First)
  - IV. テスト駆動 (Test Driven)
  - V. シンプルさ (Simplicity)
- Templates requiring updates: ✅ Verified (Generic templates are compatible)
-->
# yt-url-fetcher Constitution

## Core Principles

### I. 単一責任の原則 (Single Responsibility)
YouTube等のURL取得機能に集中し、複雑な解析や変換（ダウンロード、フォーマット変換など）は行わない。一つのツールは一つのことをうまくやるべきである。

### II. 標準入出力とCLIの重視 (Standard I/O & CLI)
UNIX哲学に従い、テキストストリームとしてデータを扱い、パイプライン処理を可能にする。入力はstdin/引数から受け取り、結果はstdoutへ、エラーはstderrへ出力する。JSON形式の出力もサポートし、他のツールとの連携を容易にする。

### III. 日本語優先 (Japanese First)
ドキュメント、コード内のコメント、ユーザーへのメッセージ、コミットメッセージなど、あらゆるアウトプットは日本語を第一言語とする。これにより、主要な利用者である日本の開発者にとっての理解しやすさを最優先する。

### IV. テスト駆動 (Test Driven)
信頼性を担保するため、機能実装前にテストを作成する（TDD）。テストはドキュメントの一部としても機能し、変更に対する恐怖を取り除く。高いテストカバレッジを維持する。

### V. シンプルさ (Simplicity)
必要最小限の機能から始め、YAGNI (You Aren't Gonna Need It) 原則に従う。複雑さはバグの温床であり、メンテナンスコストを増大させるため、常にシンプルな解決策を選択する。

## 技術スタックと制約

### 実装言語と環境
CLIツールとして配布・実行が容易な言語（Goなど）を選択する。外部依存を最小限に抑え、シングルバイナリでの配布を目指す。

### セキュリティ
外部入力（URLなど）は常に信頼できないものとして扱い、適切なバリデーションを行う。

## 開発ワークフロー

### 品質管理
全てのコード変更において、既存のテストが通過すること、および新しい機能に対するテストが追加されていることを確認する。Linterによる静的解析を行い、コードスタイルを統一する。

## Governance

本憲章はプロジェクトの全ての慣習に優先する。憲章の修正は、明確な理由とドキュメント化を伴うPull Requestを通じて行われ、チームの合意形成を経て承認される。バージョン管理はセマンティックバージョニングに従う。

**Version**: 1.0.0 | **Ratified**: 2025-12-06 | **Last Amended**: 2025-12-06