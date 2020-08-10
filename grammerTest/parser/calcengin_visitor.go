// Code generated from .\CalcEngin.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // CalcEngin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalcEnginParser.
type CalcEnginVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcEnginParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#StrStat.
	VisitStrStat(ctx *StrStatContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#BoolExp.
	VisitBoolExp(ctx *BoolExpContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#CidrStat.
	VisitCidrStat(ctx *CidrStatContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#IpStat.
	VisitIpStat(ctx *IpStatContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#BoolStat.
	VisitBoolStat(ctx *BoolStatContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#CompareStat.
	VisitCompareStat(ctx *CompareStatContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#ListStat.
	VisitListStat(ctx *ListStatContext) interface{}

	// Visit a parse tree produced by CalcEnginParser#InStat.
	VisitInStat(ctx *InStatContext) interface{}
}
