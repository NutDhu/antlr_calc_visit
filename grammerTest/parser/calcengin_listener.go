// Code generated from .\CalcEngin.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // CalcEngin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcEnginListener is a complete listener for a parse tree produced by CalcEnginParser.
type CalcEnginListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterStrStat is called when entering the StrStat production.
	EnterStrStat(c *StrStatContext)

	// EnterBoolExp is called when entering the BoolExp production.
	EnterBoolExp(c *BoolExpContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterCidrStat is called when entering the CidrStat production.
	EnterCidrStat(c *CidrStatContext)

	// EnterIpStat is called when entering the IpStat production.
	EnterIpStat(c *IpStatContext)

	// EnterBoolStat is called when entering the BoolStat production.
	EnterBoolStat(c *BoolStatContext)

	// EnterCompareStat is called when entering the CompareStat production.
	EnterCompareStat(c *CompareStatContext)

	// EnterListStat is called when entering the ListStat production.
	EnterListStat(c *ListStatContext)

	// EnterInStat is called when entering the InStat production.
	EnterInStat(c *InStatContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStrStat is called when exiting the StrStat production.
	ExitStrStat(c *StrStatContext)

	// ExitBoolExp is called when exiting the BoolExp production.
	ExitBoolExp(c *BoolExpContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitCidrStat is called when exiting the CidrStat production.
	ExitCidrStat(c *CidrStatContext)

	// ExitIpStat is called when exiting the IpStat production.
	ExitIpStat(c *IpStatContext)

	// ExitBoolStat is called when exiting the BoolStat production.
	ExitBoolStat(c *BoolStatContext)

	// ExitCompareStat is called when exiting the CompareStat production.
	ExitCompareStat(c *CompareStatContext)

	// ExitListStat is called when exiting the ListStat production.
	ExitListStat(c *ListStatContext)

	// ExitInStat is called when exiting the InStat production.
	ExitInStat(c *InStatContext)
}
