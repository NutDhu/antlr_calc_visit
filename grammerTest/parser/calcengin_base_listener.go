// Code generated from .\CalcEngin.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // CalcEngin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcEnginListener is a complete listener for a parse tree produced by CalcEnginParser.
type BaseCalcEnginListener struct{}

var _ CalcEnginListener = &BaseCalcEnginListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcEnginListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcEnginListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcEnginListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcEnginListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcEnginListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcEnginListener) ExitStart(ctx *StartContext) {}

// EnterStrStat is called when production StrStat is entered.
func (s *BaseCalcEnginListener) EnterStrStat(ctx *StrStatContext) {}

// ExitStrStat is called when production StrStat is exited.
func (s *BaseCalcEnginListener) ExitStrStat(ctx *StrStatContext) {}

// EnterBoolExp is called when production BoolExp is entered.
func (s *BaseCalcEnginListener) EnterBoolExp(ctx *BoolExpContext) {}

// ExitBoolExp is called when production BoolExp is exited.
func (s *BaseCalcEnginListener) ExitBoolExp(ctx *BoolExpContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseCalcEnginListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseCalcEnginListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalcEnginListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalcEnginListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcEnginListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcEnginListener) ExitAddSub(ctx *AddSubContext) {}

// EnterCidrStat is called when production CidrStat is entered.
func (s *BaseCalcEnginListener) EnterCidrStat(ctx *CidrStatContext) {}

// ExitCidrStat is called when production CidrStat is exited.
func (s *BaseCalcEnginListener) ExitCidrStat(ctx *CidrStatContext) {}

// EnterIpStat is called when production IpStat is entered.
func (s *BaseCalcEnginListener) EnterIpStat(ctx *IpStatContext) {}

// ExitIpStat is called when production IpStat is exited.
func (s *BaseCalcEnginListener) ExitIpStat(ctx *IpStatContext) {}

// EnterBoolStat is called when production BoolStat is entered.
func (s *BaseCalcEnginListener) EnterBoolStat(ctx *BoolStatContext) {}

// ExitBoolStat is called when production BoolStat is exited.
func (s *BaseCalcEnginListener) ExitBoolStat(ctx *BoolStatContext) {}

// EnterCompareStat is called when production CompareStat is entered.
func (s *BaseCalcEnginListener) EnterCompareStat(ctx *CompareStatContext) {}

// ExitCompareStat is called when production CompareStat is exited.
func (s *BaseCalcEnginListener) ExitCompareStat(ctx *CompareStatContext) {}

// EnterListStat is called when production ListStat is entered.
func (s *BaseCalcEnginListener) EnterListStat(ctx *ListStatContext) {}

// ExitListStat is called when production ListStat is exited.
func (s *BaseCalcEnginListener) ExitListStat(ctx *ListStatContext) {}

// EnterInStat is called when production InStat is entered.
func (s *BaseCalcEnginListener) EnterInStat(ctx *InStatContext) {}

// ExitInStat is called when production InStat is exited.
func (s *BaseCalcEnginListener) ExitInStat(ctx *InStatContext) {}
