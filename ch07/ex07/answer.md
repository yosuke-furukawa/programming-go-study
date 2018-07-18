Q. 20.0 のデフォルト値は゜を含んでいないのに、ヘルプが゜を含んでいる理由
A. flag.CommandLine.Var でusageを登録する際にDefValue を flag の struct として保持しており、 DefValue は Celsius のString()メソッドを呼んだものになっているため。
