# GO Interpreter GOQ
> go语言实现解释器
## 使用示例
```aiignore
let x = 5
let y = 10
le add = f(x,y) {
    x+y;
}
let result = add(x,y)
print(result)
```

## 组成
- 词法分析器
  - 分词器lexer
- 语法分析器
  - 抽象语法树ast
# TODO
- go语言怎么添加默认值
- interface不需要用指针吗？