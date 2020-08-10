package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"grammerTest/parser"
	"strconv"
	"strings"
)
type Visitor struct {
	parser.BaseCalcEnginVisitor
	stack []interface{}
}


func NewVisitor() *Visitor {
	return &Visitor{}
}
func (l *Visitor) push(i interface{}) {
	l.stack = append(l.stack, i)
}


func (l *Visitor) pop() interface{} {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}
func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
	node.Accept(v)
	return nil
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.visitRule(ctx.Expression())
}
func (v *Visitor) VisitBoolStat(ctx *parser.BoolStatContext) interface{} {
	i := ctx.BOOL().GetText()
	v.push(i)
	return nil
}
func (v *Visitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	i, err := strconv.Atoi(ctx.NUMBER().GetText())
	if err != nil {
		panic(err.Error())
	}
	v.push(i)
	return nil
}

func (v *Visitor) VisitListStat(ctx *parser.ListStatContext) interface{} {
	i:= ctx.LIST().GetText()
	v.push(i)
	return nil
}

func (v *Visitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	//push expression result to stack
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))

	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	r,_:=right.(int)
	l,_:=left.(int)

	switch t.GetTokenType() {
	case parser.CalcEnginParserMUL:
		v.push(l * r)
	case parser.CalcEnginParserDIV:
		v.push(l / r)
	default:
		panic("should not happen")

	}
	return nil
}

func (v *Visitor) VisitCompareStat(ctx *parser.CompareStatContext) interface{} {
	//push expression result to stack
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))

	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()
	left := v.pop()

	r,_:=right.(int)
	l,_:=left.(int)
	switch t.GetTokenType() {
	case parser.CalcEnginParserGT:
		if l>r{
			v.push(true)
		}else {
			v.push(false)
		}
	case parser.CalcEnginParserGE:
		if l>=r{
			v.push(true)
		}else {
			v.push(false)
		}
	case parser.CalcEnginParserLT:
		if l<r{
			v.push(true)
		}else {
			v.push(false)
		}
	case parser.CalcEnginParserLE:
		if l<=r{
			v.push(true)
		}else {
			v.push(false)
		}
	case parser.CalcEnginParserEQ:
		if l==r{
			v.push(true)
		}else {
			v.push(false)
		}
	case parser.CalcEnginLexerNE:
		if l!=r{
			v.push(true)
		}else {
			v.push(false)
		}
	default:
		panic("should not happen")
	}
	return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	//push expression result to stack
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))
	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	r,_:=right.(int)
	l,_:=left.(int)
	switch t.GetTokenType() {
	case parser.CalcEnginParserADD:
		v.push(l + r)
	case parser.CalcEnginParserSUB:
		v.push(l - r)
	default:
		panic("should not happen")
	}
	return nil
}

func (v *Visitor) VisitBoolExp(ctx *parser.BoolExpContext) interface{} {
	//push expression result to stack
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))
	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop().(string)
	left := v.pop().(string)
	rb,_:=strconv.ParseBool(right)
	lb,_:=strconv.ParseBool(left)

	switch t.GetTokenType() {
	case parser.CalcEnginParserAND:
		v.push(strconv.FormatBool(rb && lb))
	case parser.CalcEnginParserOR:
		v.push(strconv.FormatBool(rb || lb))
	default:
		panic("should not happen")
	}
	return nil
}



func (v *Visitor) VisitInStat(ctx *parser.InStatContext) interface{} {
	//push expression result to stack
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))
	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()//list
	left := v.pop()
	var l string
	switch left.(type) {
	case int:
		l=strconv.Itoa(left.(int))
	case string:
		l=left.(string)
	}

	r,_:=right.(string)
	r=r[1:len(r)-1]
	rlist:=strings.Split(r,",")
	switch t.GetTokenType() {
	case parser.CalcEnginLexerIN:
		for _,i :=range rlist {
			if i==l{
				v.push(true)
				return nil
			}
		}
		v.push(false)
	}
	return nil
}



func main() {
	//is := antlr.NewInputStream("1 in [3,2,3]")
	//is := antlr.NewInputStream("1+2*3")
	is := antlr.NewInputStream("false && false || false")
	//is := antlr.NewInputStream("1>3 && true")
	lexer:=parser.NewCalcEnginLexer(is)
	//++++++++++++++++++++++++++++
	tokens:=antlr.NewCommonTokenStream(lexer,antlr.TokenDefaultChannel)
	p:=parser.NewCalcEnginParser(tokens)
	v:=NewVisitor()
	p.Start().Accept(v)
	fmt.Println(v.pop())
	//++++++++++++++++++++++++++++++++
	//for {
	//	t := lexer.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	//}
}
