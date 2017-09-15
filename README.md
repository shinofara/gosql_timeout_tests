## 発生・検証する問題

PR
https://github.com/go-sql-driver/mysql/pull/302

上記PRに対する問題提起Issue
https://github.com/go-sql-driver/mysql/issues/657

## テストケース

1. 検証する `go-sql-driver` はPRの前後バージョンで実行する
2. MySQL側でwait_timeを3秒、10秒と設定してそれぞれで同じGoTestを実行する
3. GO言語側では、`SetConnMaxLifetime` を設定したテストと設定していない時のテストを実行する
4. テストコード内部で、wait_timeの短い時間より多めSleepを発生させて、コネクションが切られる様にする
5. もちろんwait_timeが長い場合は問題無いことも検証

## 実行手順

Goの依存解決

```
$ make
```

MySQL起動

```
$ docker-compose up
```

テスト実行

```
$ make test
```

## 実行履歴

```
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| wait_timeout  | 10    |
+---------------+-------+
=== RUN   TestUnsetLifetime
--- PASS: TestUnsetLifetime (5.01s)
=== RUN   TestSetLifetime
--- PASS: TestSetLifetime (5.02s)
PASS
ok  	gosql/21d7e97c9f760ca685a01ecea202e1c84276daa1	10.053s
=== RUN   TestUnsetLifetime
--- PASS: TestUnsetLifetime (5.01s)
=== RUN   TestSetLifetime
--- PASS: TestSetLifetime (5.02s)
PASS
ok  	gosql/26471af196a17ee75a22e6481b5a5897fb16b081	10.046s
```
```
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| wait_timeout  | 3     |
+---------------+-------+
=== RUN   TestUnsetLifetime
[mysql] 2017/09/15 11:09:24 packets.go:36: unexpected EOF
--- PASS: TestUnsetLifetime (5.02s)
=== RUN   TestSetLifetime
--- PASS: TestSetLifetime (5.01s)
PASS
ok  	gosql/21d7e97c9f760ca685a01ecea202e1c84276daa1	10.042s
=== RUN   TestUnsetLifetime
[mysql] 2017/09/15 11:09:35 packets.go:36: unexpected EOF
--- FAIL: TestUnsetLifetime (5.01s)
	main_test.go:60: invalid connection
	main_test.go:15: runtime error: invalid memory address or nil pointer dereference
=== RUN   TestSetLifetime
--- PASS: TestSetLifetime (5.01s)
FAIL
exit status 1
FAIL	gosql/26471af196a17ee75a22e6481b5a5897fb16b081	10.038s
```

## まとめ

|                               | wait_timeを3秒 | wait_timewを10秒                                         |
| ------------------------------| -------------- | ------------------------------------------------------- |
| SetConnMaxLifetimeを2秒で設定 | (/)            | (/)                                                      |
| SetConnMaxLifetimeを未設定    | Panic発生      | コネクションロストするが、再接続が実行され正常に終了する | 
