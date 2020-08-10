// Code generated from .\CalcEngin.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // CalcEngin

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCalcEnginVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcEnginVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitStrStat(ctx *StrStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitBoolExp(ctx *BoolExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitCidrStat(ctx *CidrStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitIpStat(ctx *IpStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitBoolStat(ctx *BoolStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitCompareStat(ctx *CompareStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitListStat(ctx *ListStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcEnginVisitor) VisitInStat(ctx *InStatContext) interface{} {
	return v.VisitChildren(ctx)
}
