// Code generated from .\CalcEngin.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // CalcEngin

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 39, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 17, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 34, 10, 3, 12, 3, 14,
	3, 37, 11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 6, 3, 2, 3, 4, 3, 2, 5, 10, 3,
	2, 12, 13, 3, 2, 14, 15, 2, 46, 2, 6, 3, 2, 2, 2, 4, 16, 3, 2, 2, 2, 6,
	7, 5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8, 3, 3, 2, 2, 2, 9, 10, 8, 3, 1, 2, 10,
	17, 7, 17, 2, 2, 11, 17, 7, 20, 2, 2, 12, 17, 7, 16, 2, 2, 13, 17, 7, 24,
	2, 2, 14, 17, 7, 23, 2, 2, 15, 17, 7, 26, 2, 2, 16, 9, 3, 2, 2, 2, 16,
	11, 3, 2, 2, 2, 16, 12, 3, 2, 2, 2, 16, 13, 3, 2, 2, 2, 16, 14, 3, 2, 2,
	2, 16, 15, 3, 2, 2, 2, 17, 35, 3, 2, 2, 2, 18, 19, 12, 13, 2, 2, 19, 20,
	9, 2, 2, 2, 20, 34, 5, 4, 3, 14, 21, 22, 12, 12, 2, 2, 22, 23, 7, 11, 2,
	2, 23, 34, 5, 4, 3, 13, 24, 25, 12, 11, 2, 2, 25, 26, 9, 3, 2, 2, 26, 34,
	5, 4, 3, 12, 27, 28, 12, 10, 2, 2, 28, 29, 9, 4, 2, 2, 29, 34, 5, 4, 3,
	11, 30, 31, 12, 9, 2, 2, 31, 32, 9, 5, 2, 2, 32, 34, 5, 4, 3, 10, 33, 18,
	3, 2, 2, 2, 33, 21, 3, 2, 2, 2, 33, 24, 3, 2, 2, 2, 33, 27, 3, 2, 2, 2,
	33, 30, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3,
	2, 2, 2, 36, 5, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 5, 16, 33, 35,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'&&'", "'||'", "'>'", "'>='", "'<'", "'<='", "'='", "'!='", "'in'",
	"'*'", "'/'", "'+'", "'-'", "", "", "", "", "", "'true'", "'false'", "",
	"", "", "", "'['", "", "", "", "']'",
}
var symbolicNames = []string{
	"", "AND", "OR", "GT", "GE", "LT", "LE", "EQ", "NE", "IN", "MUL", "DIV",
	"ADD", "SUB", "BOOL", "NUMBER", "WHITESPACE", "ID", "STRING", "TRUE", "FALSE",
	"CIDR", "IP", "SET", "LIST", "LK", "PARAM", "P1", "P2", "RK",
}

var ruleNames = []string{
	"start", "expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CalcEnginParser struct {
	*antlr.BaseParser
}

func NewCalcEnginParser(input antlr.TokenStream) *CalcEnginParser {
	this := new(CalcEnginParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CalcEngin.g4"

	return this
}

// CalcEnginParser tokens.
const (
	CalcEnginParserEOF        = antlr.TokenEOF
	CalcEnginParserAND        = 1
	CalcEnginParserOR         = 2
	CalcEnginParserGT         = 3
	CalcEnginParserGE         = 4
	CalcEnginParserLT         = 5
	CalcEnginParserLE         = 6
	CalcEnginParserEQ         = 7
	CalcEnginParserNE         = 8
	CalcEnginParserIN         = 9
	CalcEnginParserMUL        = 10
	CalcEnginParserDIV        = 11
	CalcEnginParserADD        = 12
	CalcEnginParserSUB        = 13
	CalcEnginParserBOOL       = 14
	CalcEnginParserNUMBER     = 15
	CalcEnginParserWHITESPACE = 16
	CalcEnginParserID         = 17
	CalcEnginParserSTRING     = 18
	CalcEnginParserTRUE       = 19
	CalcEnginParserFALSE      = 20
	CalcEnginParserCIDR       = 21
	CalcEnginParserIP         = 22
	CalcEnginParserSET        = 23
	CalcEnginParserLIST       = 24
	CalcEnginParserLK         = 25
	CalcEnginParserPARAM      = 26
	CalcEnginParserP1         = 27
	CalcEnginParserP2         = 28
	CalcEnginParserRK         = 29
)

// CalcEnginParser rules.
const (
	CalcEnginParserRULE_start      = 0
	CalcEnginParserRULE_expression = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcEnginParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcEnginParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcEnginParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcEnginParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.expression(0)
	}
	{
		p.SetState(5)
		p.Match(CalcEnginParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcEnginParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcEnginParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StrStatContext struct {
	*ExpressionContext
}

func NewStrStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrStatContext {
	var p = new(StrStatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StrStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrStatContext) STRING() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserSTRING, 0)
}

func (s *StrStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterStrStat(s)
	}
}

func (s *StrStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitStrStat(s)
	}
}

func (s *StrStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitStrStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolExpContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewBoolExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExpContext {
	var p = new(BoolExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExpContext) GetOp() antlr.Token { return s.op }

func (s *BoolExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BoolExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BoolExpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BoolExpContext) AND() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserAND, 0)
}

func (s *BoolExpContext) OR() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserOR, 0)
}

func (s *BoolExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterBoolExp(s)
	}
}

func (s *BoolExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitBoolExp(s)
	}
}

func (s *BoolExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitBoolExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	*ExpressionContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type CidrStatContext struct {
	*ExpressionContext
}

func NewCidrStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CidrStatContext {
	var p = new(CidrStatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CidrStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CidrStatContext) CIDR() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserCIDR, 0)
}

func (s *CidrStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterCidrStat(s)
	}
}

func (s *CidrStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitCidrStat(s)
	}
}

func (s *CidrStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitCidrStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type IpStatContext struct {
	*ExpressionContext
}

func NewIpStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IpStatContext {
	var p = new(IpStatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IpStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpStatContext) IP() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserIP, 0)
}

func (s *IpStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterIpStat(s)
	}
}

func (s *IpStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitIpStat(s)
	}
}

func (s *IpStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitIpStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolStatContext struct {
	*ExpressionContext
}

func NewBoolStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolStatContext {
	var p = new(BoolStatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolStatContext) BOOL() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserBOOL, 0)
}

func (s *BoolStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterBoolStat(s)
	}
}

func (s *BoolStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitBoolStat(s)
	}
}

func (s *BoolStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitBoolStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareStatContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewCompareStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareStatContext {
	var p = new(CompareStatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompareStatContext) GetOp() antlr.Token { return s.op }

func (s *CompareStatContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareStatContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompareStatContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareStatContext) GT() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserGT, 0)
}

func (s *CompareStatContext) GE() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserGE, 0)
}

func (s *CompareStatContext) EQ() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserEQ, 0)
}

func (s *CompareStatContext) NE() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserNE, 0)
}

func (s *CompareStatContext) LT() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserLT, 0)
}

func (s *CompareStatContext) LE() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserLE, 0)
}

func (s *CompareStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterCompareStat(s)
	}
}

func (s *CompareStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitCompareStat(s)
	}
}

func (s *CompareStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitCompareStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListStatContext struct {
	*ExpressionContext
}

func NewListStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListStatContext {
	var p = new(ListStatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ListStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStatContext) LIST() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserLIST, 0)
}

func (s *ListStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterListStat(s)
	}
}

func (s *ListStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitListStat(s)
	}
}

func (s *ListStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitListStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type InStatContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewInStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InStatContext {
	var p = new(InStatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InStatContext) GetOp() antlr.Token { return s.op }

func (s *InStatContext) SetOp(v antlr.Token) { s.op = v }

func (s *InStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InStatContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InStatContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InStatContext) IN() antlr.TerminalNode {
	return s.GetToken(CalcEnginParserIN, 0)
}

func (s *InStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.EnterInStat(s)
	}
}

func (s *InStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcEnginListener); ok {
		listenerT.ExitInStat(s)
	}
}

func (s *InStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcEnginVisitor:
		return t.VisitInStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcEnginParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CalcEnginParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcEnginParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcEnginParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(8)
			p.Match(CalcEnginParserNUMBER)
		}

	case CalcEnginParserSTRING:
		localctx = NewStrStatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Match(CalcEnginParserSTRING)
		}

	case CalcEnginParserBOOL:
		localctx = NewBoolStatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(10)
			p.Match(CalcEnginParserBOOL)
		}

	case CalcEnginParserIP:
		localctx = NewIpStatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(CalcEnginParserIP)
		}

	case CalcEnginParserCIDR:
		localctx = NewCidrStatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)
			p.Match(CalcEnginParserCIDR)
		}

	case CalcEnginParserLIST:
		localctx = NewListStatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.Match(CalcEnginParserLIST)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBoolExpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcEnginParserRULE_expression)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BoolExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcEnginParserAND || _la == CalcEnginParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BoolExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)
					p.expression(12)
				}

			case 2:
				localctx = NewInStatContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcEnginParserRULE_expression)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(20)

					var _m = p.Match(CalcEnginParserIN)

					localctx.(*InStatContext).op = _m
				}
				{
					p.SetState(21)
					p.expression(11)
				}

			case 3:
				localctx = NewCompareStatContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcEnginParserRULE_expression)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(23)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareStatContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalcEnginParserGT)|(1<<CalcEnginParserGE)|(1<<CalcEnginParserLT)|(1<<CalcEnginParserLE)|(1<<CalcEnginParserEQ)|(1<<CalcEnginParserNE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareStatContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(24)
					p.expression(10)
				}

			case 4:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcEnginParserRULE_expression)
				p.SetState(25)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(26)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcEnginParserMUL || _la == CalcEnginParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(27)
					p.expression(9)
				}

			case 5:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcEnginParserRULE_expression)
				p.SetState(28)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(29)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcEnginParserADD || _la == CalcEnginParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(30)
					p.expression(8)
				}

			}

		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *CalcEnginParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcEnginParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
