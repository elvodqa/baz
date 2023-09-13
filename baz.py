from lark import Lark, Transformer

with open('grammer.ebnf') as f:
    grammer = f.read()

parser = Lark(grammer, start='start', parser='lalr', transformer=Transformer())

with open('test.baz') as f:
    code = f.read()
    l = parser.parse(code)
    print(l.pretty())

