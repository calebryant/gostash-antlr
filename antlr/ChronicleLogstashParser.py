# Generated from /Users/caleb.bryant/Library/CloudStorage/OneDrive-Cyderes/TelEng/github/chronicle-parser-language-server/antlr/ChronicleLogstashParser.g4 by ANTLR 4.13.0
# encoding: utf-8
from antlr4 import *
from io import StringIO
import sys
if sys.version_info[1] > 5:
	from typing import TextIO
else:
	from typing.io import TextIO

def serializedATN():
    return [
        4,1,39,220,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,
        2,14,7,14,2,15,7,15,2,16,7,16,2,17,7,17,2,18,7,18,1,0,1,0,1,0,1,
        0,5,0,43,8,0,10,0,12,0,46,9,0,1,0,1,0,1,0,1,1,1,1,1,1,1,1,1,1,5,
        1,56,8,1,10,1,12,1,59,9,1,1,1,1,1,1,1,1,1,1,1,1,1,5,1,67,8,1,10,
        1,12,1,70,9,1,1,1,1,1,1,1,1,1,3,1,76,8,1,1,1,1,1,1,1,1,1,1,1,1,1,
        5,1,84,8,1,10,1,12,1,87,9,1,1,1,3,1,90,8,1,1,2,1,2,1,2,1,2,1,2,1,
        2,1,2,1,2,1,2,1,2,1,2,1,2,3,2,104,8,2,1,2,1,2,1,2,1,2,5,2,110,8,
        2,10,2,12,2,113,9,2,1,3,1,3,3,3,117,8,3,1,4,1,4,1,5,1,5,1,5,1,5,
        1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,3,6,133,8,6,1,7,1,7,1,7,1,7,1,7,
        1,7,3,7,141,8,7,1,7,1,7,1,7,5,7,146,8,7,10,7,12,7,149,9,7,1,8,1,
        8,1,8,1,8,1,9,1,9,1,10,1,10,1,11,1,11,1,11,5,11,162,8,11,10,11,12,
        11,165,9,11,1,11,1,11,1,12,1,12,1,12,1,12,3,12,173,8,12,1,13,1,13,
        1,14,1,14,1,14,1,14,1,14,1,14,3,14,183,8,14,1,15,1,15,5,15,187,8,
        15,10,15,12,15,190,9,15,1,15,1,15,1,16,1,16,1,16,5,16,197,8,16,10,
        16,12,16,200,9,16,3,16,202,8,16,1,16,1,16,1,17,1,17,1,17,5,17,209,
        8,17,10,17,12,17,212,9,17,3,17,214,8,17,1,17,1,17,1,18,1,18,1,18,
        0,2,4,14,19,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,
        0,8,1,0,3,4,2,0,29,29,32,32,2,0,32,32,34,34,1,0,27,28,2,0,7,7,19,
        26,2,0,30,30,33,33,2,0,29,30,34,34,3,0,16,16,29,30,33,34,237,0,38,
        1,0,0,0,2,89,1,0,0,0,4,103,1,0,0,0,6,116,1,0,0,0,8,118,1,0,0,0,10,
        120,1,0,0,0,12,132,1,0,0,0,14,140,1,0,0,0,16,150,1,0,0,0,18,154,
        1,0,0,0,20,156,1,0,0,0,22,158,1,0,0,0,24,168,1,0,0,0,26,174,1,0,
        0,0,28,182,1,0,0,0,30,184,1,0,0,0,32,193,1,0,0,0,34,205,1,0,0,0,
        36,217,1,0,0,0,38,39,5,33,0,0,39,44,5,9,0,0,40,43,3,22,11,0,41,43,
        3,2,1,0,42,40,1,0,0,0,42,41,1,0,0,0,43,46,1,0,0,0,44,42,1,0,0,0,
        44,45,1,0,0,0,45,47,1,0,0,0,46,44,1,0,0,0,47,48,5,10,0,0,48,49,5,
        0,0,1,49,1,1,0,0,0,50,51,7,0,0,0,51,52,3,4,2,0,52,57,5,9,0,0,53,
        56,3,22,11,0,54,56,3,2,1,0,55,53,1,0,0,0,55,54,1,0,0,0,56,59,1,0,
        0,0,57,55,1,0,0,0,57,58,1,0,0,0,58,60,1,0,0,0,59,57,1,0,0,0,60,61,
        5,10,0,0,61,90,1,0,0,0,62,63,5,5,0,0,63,68,5,9,0,0,64,67,3,22,11,
        0,65,67,3,2,1,0,66,64,1,0,0,0,66,65,1,0,0,0,67,70,1,0,0,0,68,66,
        1,0,0,0,68,69,1,0,0,0,69,71,1,0,0,0,70,68,1,0,0,0,71,90,5,10,0,0,
        72,75,5,6,0,0,73,74,5,38,0,0,74,76,5,36,0,0,75,73,1,0,0,0,75,76,
        1,0,0,0,76,77,1,0,0,0,77,78,5,38,0,0,78,79,5,37,0,0,79,80,5,38,0,
        0,80,85,5,39,0,0,81,84,3,22,11,0,82,84,3,2,1,0,83,81,1,0,0,0,83,
        82,1,0,0,0,84,87,1,0,0,0,85,83,1,0,0,0,85,86,1,0,0,0,86,88,1,0,0,
        0,87,85,1,0,0,0,88,90,5,10,0,0,89,50,1,0,0,0,89,62,1,0,0,0,89,72,
        1,0,0,0,90,3,1,0,0,0,91,92,6,2,-1,0,92,93,5,13,0,0,93,94,3,4,2,0,
        94,95,5,14,0,0,95,104,1,0,0,0,96,97,5,11,0,0,97,98,3,4,2,0,98,99,
        5,12,0,0,99,104,1,0,0,0,100,101,5,17,0,0,101,104,3,4,2,2,102,104,
        3,6,3,0,103,91,1,0,0,0,103,96,1,0,0,0,103,100,1,0,0,0,103,102,1,
        0,0,0,104,111,1,0,0,0,105,106,10,3,0,0,106,107,3,18,9,0,107,108,
        3,4,2,4,108,110,1,0,0,0,109,105,1,0,0,0,110,113,1,0,0,0,111,109,
        1,0,0,0,111,112,1,0,0,0,112,5,1,0,0,0,113,111,1,0,0,0,114,117,3,
        10,5,0,115,117,3,8,4,0,116,114,1,0,0,0,116,115,1,0,0,0,117,7,1,0,
        0,0,118,119,7,1,0,0,119,9,1,0,0,0,120,121,3,12,6,0,121,122,3,20,
        10,0,122,123,3,12,6,0,123,11,1,0,0,0,124,133,3,14,7,0,125,133,5,
        34,0,0,126,133,3,34,17,0,127,133,5,32,0,0,128,133,5,30,0,0,129,133,
        5,31,0,0,130,133,5,29,0,0,131,133,5,33,0,0,132,124,1,0,0,0,132,125,
        1,0,0,0,132,126,1,0,0,0,132,127,1,0,0,0,132,128,1,0,0,0,132,129,
        1,0,0,0,132,130,1,0,0,0,132,131,1,0,0,0,133,13,1,0,0,0,134,135,6,
        7,-1,0,135,136,5,13,0,0,136,137,3,14,7,0,137,138,5,14,0,0,138,141,
        1,0,0,0,139,141,3,16,8,0,140,134,1,0,0,0,140,139,1,0,0,0,141,147,
        1,0,0,0,142,143,10,2,0,0,143,144,5,18,0,0,144,146,3,14,7,3,145,142,
        1,0,0,0,146,149,1,0,0,0,147,145,1,0,0,0,147,148,1,0,0,0,148,15,1,
        0,0,0,149,147,1,0,0,0,150,151,7,2,0,0,151,152,5,18,0,0,152,153,7,
        2,0,0,153,17,1,0,0,0,154,155,7,3,0,0,155,19,1,0,0,0,156,157,7,4,
        0,0,157,21,1,0,0,0,158,159,5,33,0,0,159,163,5,9,0,0,160,162,3,24,
        12,0,161,160,1,0,0,0,162,165,1,0,0,0,163,161,1,0,0,0,163,164,1,0,
        0,0,164,166,1,0,0,0,165,163,1,0,0,0,166,167,5,10,0,0,167,23,1,0,
        0,0,168,169,3,26,13,0,169,170,5,15,0,0,170,172,3,28,14,0,171,173,
        5,16,0,0,172,171,1,0,0,0,172,173,1,0,0,0,173,25,1,0,0,0,174,175,
        7,5,0,0,175,27,1,0,0,0,176,183,5,34,0,0,177,183,3,32,16,0,178,183,
        3,30,15,0,179,183,5,30,0,0,180,183,5,29,0,0,181,183,5,33,0,0,182,
        176,1,0,0,0,182,177,1,0,0,0,182,178,1,0,0,0,182,179,1,0,0,0,182,
        180,1,0,0,0,182,181,1,0,0,0,183,29,1,0,0,0,184,188,5,9,0,0,185,187,
        3,24,12,0,186,185,1,0,0,0,187,190,1,0,0,0,188,186,1,0,0,0,188,189,
        1,0,0,0,189,191,1,0,0,0,190,188,1,0,0,0,191,192,5,10,0,0,192,31,
        1,0,0,0,193,201,5,11,0,0,194,198,3,36,18,0,195,197,3,36,18,0,196,
        195,1,0,0,0,197,200,1,0,0,0,198,196,1,0,0,0,198,199,1,0,0,0,199,
        202,1,0,0,0,200,198,1,0,0,0,201,194,1,0,0,0,201,202,1,0,0,0,202,
        203,1,0,0,0,203,204,5,12,0,0,204,33,1,0,0,0,205,213,5,11,0,0,206,
        210,7,6,0,0,207,209,3,36,18,0,208,207,1,0,0,0,209,212,1,0,0,0,210,
        208,1,0,0,0,210,211,1,0,0,0,211,214,1,0,0,0,212,210,1,0,0,0,213,
        206,1,0,0,0,213,214,1,0,0,0,214,215,1,0,0,0,215,216,5,12,0,0,216,
        35,1,0,0,0,217,218,7,7,0,0,218,37,1,0,0,0,24,42,44,55,57,66,68,75,
        83,85,89,103,111,116,132,140,147,163,172,182,188,198,201,210,213
    ]

class ChronicleLogstashParser ( Parser ):

    grammarFileName = "ChronicleLogstashParser.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "<INVALID>", "<INVALID>", "'if'", "'else if'", 
                     "'else'", "'for'", "'in'", "'not'", "<INVALID>", "'}'", 
                     "'['", "']'", "'('", "')'", "<INVALID>", "','", "'!'", 
                     "<INVALID>", "'=='", "'!='", "'<'", "'>'", "'<='", 
                     "'>='", "'=~'", "'!~'" ]

    symbolicNames = [ "<INVALID>", "WS", "COMMENT", "IF", "ELSEIF", "ELSE", 
                      "FOR", "IN", "NOT", "LBRACE", "RBRACE", "LBRACKET", 
                      "RBRACKET", "LPAREN", "RPAREN", "KVSEPARATOR", "COMMA", 
                      "BOOLNOT", "MATHOP", "EQUAL", "NOTEQUAL", "LESSTHAN", 
                      "GREATERTHAN", "LTEQUAL", "GTEQUAL", "MATCH", "NOTMATCH", 
                      "AND", "OR", "BOOLEAN", "STRING", "REGEX", "IFSTATEMENTID", 
                      "ID", "NUMBER", "FORWS", "FORCOMMA", "FORIN", "FORID", 
                      "FOROPENER" ]

    RULE_filterblock = 0
    RULE_conditionalblock = 1
    RULE_statement = 2
    RULE_expression = 3
    RULE_unary_expression = 4
    RULE_binary_expression = 5
    RULE_expression_val = 6
    RULE_math_statement = 7
    RULE_math_expression = 8
    RULE_boolean_op = 9
    RULE_boolean_eval = 10
    RULE_plugin = 11
    RULE_keyvalue = 12
    RULE_kv_lvalue = 13
    RULE_kv_rvalue = 14
    RULE_hash = 15
    RULE_list = 16
    RULE_if_list = 17
    RULE_list_value = 18

    ruleNames =  [ "filterblock", "conditionalblock", "statement", "expression", 
                   "unary_expression", "binary_expression", "expression_val", 
                   "math_statement", "math_expression", "boolean_op", "boolean_eval", 
                   "plugin", "keyvalue", "kv_lvalue", "kv_rvalue", "hash", 
                   "list", "if_list", "list_value" ]

    EOF = Token.EOF
    WS=1
    COMMENT=2
    IF=3
    ELSEIF=4
    ELSE=5
    FOR=6
    IN=7
    NOT=8
    LBRACE=9
    RBRACE=10
    LBRACKET=11
    RBRACKET=12
    LPAREN=13
    RPAREN=14
    KVSEPARATOR=15
    COMMA=16
    BOOLNOT=17
    MATHOP=18
    EQUAL=19
    NOTEQUAL=20
    LESSTHAN=21
    GREATERTHAN=22
    LTEQUAL=23
    GTEQUAL=24
    MATCH=25
    NOTMATCH=26
    AND=27
    OR=28
    BOOLEAN=29
    STRING=30
    REGEX=31
    IFSTATEMENTID=32
    ID=33
    NUMBER=34
    FORWS=35
    FORCOMMA=36
    FORIN=37
    FORID=38
    FOROPENER=39

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.13.0")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class FilterblockContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(ChronicleLogstashParser.ID, 0)

        def LBRACE(self):
            return self.getToken(ChronicleLogstashParser.LBRACE, 0)

        def RBRACE(self):
            return self.getToken(ChronicleLogstashParser.RBRACE, 0)

        def EOF(self):
            return self.getToken(ChronicleLogstashParser.EOF, 0)

        def plugin(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.PluginContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.PluginContext,i)


        def conditionalblock(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.ConditionalblockContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.ConditionalblockContext,i)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_filterblock

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFilterblock" ):
                listener.enterFilterblock(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFilterblock" ):
                listener.exitFilterblock(self)




    def filterblock(self):

        localctx = ChronicleLogstashParser.FilterblockContext(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_filterblock)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 38
            self.match(ChronicleLogstashParser.ID)
            self.state = 39
            self.match(ChronicleLogstashParser.LBRACE)
            self.state = 44
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while (((_la) & ~0x3f) == 0 and ((1 << _la) & 8589934712) != 0):
                self.state = 42
                self._errHandler.sync(self)
                token = self._input.LA(1)
                if token in [33]:
                    self.state = 40
                    self.plugin()
                    pass
                elif token in [3, 4, 5, 6]:
                    self.state = 41
                    self.conditionalblock()
                    pass
                else:
                    raise NoViableAltException(self)

                self.state = 46
                self._errHandler.sync(self)
                _la = self._input.LA(1)

            self.state = 47
            self.match(ChronicleLogstashParser.RBRACE)
            self.state = 48
            self.match(ChronicleLogstashParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ConditionalblockContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def statement(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.StatementContext,0)


        def LBRACE(self):
            return self.getToken(ChronicleLogstashParser.LBRACE, 0)

        def RBRACE(self):
            return self.getToken(ChronicleLogstashParser.RBRACE, 0)

        def IF(self):
            return self.getToken(ChronicleLogstashParser.IF, 0)

        def ELSEIF(self):
            return self.getToken(ChronicleLogstashParser.ELSEIF, 0)

        def plugin(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.PluginContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.PluginContext,i)


        def conditionalblock(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.ConditionalblockContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.ConditionalblockContext,i)


        def ELSE(self):
            return self.getToken(ChronicleLogstashParser.ELSE, 0)

        def FOR(self):
            return self.getToken(ChronicleLogstashParser.FOR, 0)

        def FORID(self, i:int=None):
            if i is None:
                return self.getTokens(ChronicleLogstashParser.FORID)
            else:
                return self.getToken(ChronicleLogstashParser.FORID, i)

        def FORIN(self):
            return self.getToken(ChronicleLogstashParser.FORIN, 0)

        def FOROPENER(self):
            return self.getToken(ChronicleLogstashParser.FOROPENER, 0)

        def FORCOMMA(self):
            return self.getToken(ChronicleLogstashParser.FORCOMMA, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_conditionalblock

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterConditionalblock" ):
                listener.enterConditionalblock(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitConditionalblock" ):
                listener.exitConditionalblock(self)




    def conditionalblock(self):

        localctx = ChronicleLogstashParser.ConditionalblockContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_conditionalblock)
        self._la = 0 # Token type
        try:
            self.state = 89
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [3, 4]:
                self.enterOuterAlt(localctx, 1)
                self.state = 50
                _la = self._input.LA(1)
                if not(_la==3 or _la==4):
                    self._errHandler.recoverInline(self)
                else:
                    self._errHandler.reportMatch(self)
                    self.consume()
                self.state = 51
                self.statement(0)
                self.state = 52
                self.match(ChronicleLogstashParser.LBRACE)
                self.state = 57
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while (((_la) & ~0x3f) == 0 and ((1 << _la) & 8589934712) != 0):
                    self.state = 55
                    self._errHandler.sync(self)
                    token = self._input.LA(1)
                    if token in [33]:
                        self.state = 53
                        self.plugin()
                        pass
                    elif token in [3, 4, 5, 6]:
                        self.state = 54
                        self.conditionalblock()
                        pass
                    else:
                        raise NoViableAltException(self)

                    self.state = 59
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)

                self.state = 60
                self.match(ChronicleLogstashParser.RBRACE)
                pass
            elif token in [5]:
                self.enterOuterAlt(localctx, 2)
                self.state = 62
                self.match(ChronicleLogstashParser.ELSE)
                self.state = 63
                self.match(ChronicleLogstashParser.LBRACE)
                self.state = 68
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while (((_la) & ~0x3f) == 0 and ((1 << _la) & 8589934712) != 0):
                    self.state = 66
                    self._errHandler.sync(self)
                    token = self._input.LA(1)
                    if token in [33]:
                        self.state = 64
                        self.plugin()
                        pass
                    elif token in [3, 4, 5, 6]:
                        self.state = 65
                        self.conditionalblock()
                        pass
                    else:
                        raise NoViableAltException(self)

                    self.state = 70
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)

                self.state = 71
                self.match(ChronicleLogstashParser.RBRACE)
                pass
            elif token in [6]:
                self.enterOuterAlt(localctx, 3)
                self.state = 72
                self.match(ChronicleLogstashParser.FOR)
                self.state = 75
                self._errHandler.sync(self)
                la_ = self._interp.adaptivePredict(self._input,6,self._ctx)
                if la_ == 1:
                    self.state = 73
                    self.match(ChronicleLogstashParser.FORID)
                    self.state = 74
                    self.match(ChronicleLogstashParser.FORCOMMA)


                self.state = 77
                self.match(ChronicleLogstashParser.FORID)
                self.state = 78
                self.match(ChronicleLogstashParser.FORIN)
                self.state = 79
                self.match(ChronicleLogstashParser.FORID)
                self.state = 80
                self.match(ChronicleLogstashParser.FOROPENER)
                self.state = 85
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while (((_la) & ~0x3f) == 0 and ((1 << _la) & 8589934712) != 0):
                    self.state = 83
                    self._errHandler.sync(self)
                    token = self._input.LA(1)
                    if token in [33]:
                        self.state = 81
                        self.plugin()
                        pass
                    elif token in [3, 4, 5, 6]:
                        self.state = 82
                        self.conditionalblock()
                        pass
                    else:
                        raise NoViableAltException(self)

                    self.state = 87
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)

                self.state = 88
                self.match(ChronicleLogstashParser.RBRACE)
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def LPAREN(self):
            return self.getToken(ChronicleLogstashParser.LPAREN, 0)

        def statement(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.StatementContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.StatementContext,i)


        def RPAREN(self):
            return self.getToken(ChronicleLogstashParser.RPAREN, 0)

        def LBRACKET(self):
            return self.getToken(ChronicleLogstashParser.LBRACKET, 0)

        def RBRACKET(self):
            return self.getToken(ChronicleLogstashParser.RBRACKET, 0)

        def BOOLNOT(self):
            return self.getToken(ChronicleLogstashParser.BOOLNOT, 0)

        def expression(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.ExpressionContext,0)


        def boolean_op(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Boolean_opContext,0)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_statement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterStatement" ):
                listener.enterStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitStatement" ):
                listener.exitStatement(self)



    def statement(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = ChronicleLogstashParser.StatementContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 4
        self.enterRecursionRule(localctx, 4, self.RULE_statement, _p)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 103
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,10,self._ctx)
            if la_ == 1:
                self.state = 92
                self.match(ChronicleLogstashParser.LPAREN)
                self.state = 93
                self.statement(0)
                self.state = 94
                self.match(ChronicleLogstashParser.RPAREN)
                pass

            elif la_ == 2:
                self.state = 96
                self.match(ChronicleLogstashParser.LBRACKET)
                self.state = 97
                self.statement(0)
                self.state = 98
                self.match(ChronicleLogstashParser.RBRACKET)
                pass

            elif la_ == 3:
                self.state = 100
                self.match(ChronicleLogstashParser.BOOLNOT)
                self.state = 101
                self.statement(2)
                pass

            elif la_ == 4:
                self.state = 102
                self.expression()
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 111
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,11,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    localctx = ChronicleLogstashParser.StatementContext(self, _parentctx, _parentState)
                    self.pushNewRecursionContext(localctx, _startState, self.RULE_statement)
                    self.state = 105
                    if not self.precpred(self._ctx, 3):
                        from antlr4.error.Errors import FailedPredicateException
                        raise FailedPredicateException(self, "self.precpred(self._ctx, 3)")
                    self.state = 106
                    self.boolean_op()
                    self.state = 107
                    self.statement(4) 
                self.state = 113
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,11,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class ExpressionContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def binary_expression(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Binary_expressionContext,0)


        def unary_expression(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Unary_expressionContext,0)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_expression

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterExpression" ):
                listener.enterExpression(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitExpression" ):
                listener.exitExpression(self)




    def expression(self):

        localctx = ChronicleLogstashParser.ExpressionContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_expression)
        try:
            self.state = 116
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,12,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 114
                self.binary_expression()
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 115
                self.unary_expression()
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Unary_expressionContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def IFSTATEMENTID(self):
            return self.getToken(ChronicleLogstashParser.IFSTATEMENTID, 0)

        def BOOLEAN(self):
            return self.getToken(ChronicleLogstashParser.BOOLEAN, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_unary_expression

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterUnary_expression" ):
                listener.enterUnary_expression(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitUnary_expression" ):
                listener.exitUnary_expression(self)




    def unary_expression(self):

        localctx = ChronicleLogstashParser.Unary_expressionContext(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_unary_expression)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 118
            _la = self._input.LA(1)
            if not(_la==29 or _la==32):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Binary_expressionContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expression_val(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.Expression_valContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.Expression_valContext,i)


        def boolean_eval(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Boolean_evalContext,0)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_binary_expression

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBinary_expression" ):
                listener.enterBinary_expression(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBinary_expression" ):
                listener.exitBinary_expression(self)




    def binary_expression(self):

        localctx = ChronicleLogstashParser.Binary_expressionContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_binary_expression)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 120
            self.expression_val()
            self.state = 121
            self.boolean_eval()
            self.state = 122
            self.expression_val()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Expression_valContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def math_statement(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Math_statementContext,0)


        def NUMBER(self):
            return self.getToken(ChronicleLogstashParser.NUMBER, 0)

        def if_list(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.If_listContext,0)


        def IFSTATEMENTID(self):
            return self.getToken(ChronicleLogstashParser.IFSTATEMENTID, 0)

        def STRING(self):
            return self.getToken(ChronicleLogstashParser.STRING, 0)

        def REGEX(self):
            return self.getToken(ChronicleLogstashParser.REGEX, 0)

        def BOOLEAN(self):
            return self.getToken(ChronicleLogstashParser.BOOLEAN, 0)

        def ID(self):
            return self.getToken(ChronicleLogstashParser.ID, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_expression_val

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterExpression_val" ):
                listener.enterExpression_val(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitExpression_val" ):
                listener.exitExpression_val(self)




    def expression_val(self):

        localctx = ChronicleLogstashParser.Expression_valContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_expression_val)
        try:
            self.state = 132
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,13,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 124
                self.math_statement(0)
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 125
                self.match(ChronicleLogstashParser.NUMBER)
                pass

            elif la_ == 3:
                self.enterOuterAlt(localctx, 3)
                self.state = 126
                self.if_list()
                pass

            elif la_ == 4:
                self.enterOuterAlt(localctx, 4)
                self.state = 127
                self.match(ChronicleLogstashParser.IFSTATEMENTID)
                pass

            elif la_ == 5:
                self.enterOuterAlt(localctx, 5)
                self.state = 128
                self.match(ChronicleLogstashParser.STRING)
                pass

            elif la_ == 6:
                self.enterOuterAlt(localctx, 6)
                self.state = 129
                self.match(ChronicleLogstashParser.REGEX)
                pass

            elif la_ == 7:
                self.enterOuterAlt(localctx, 7)
                self.state = 130
                self.match(ChronicleLogstashParser.BOOLEAN)
                pass

            elif la_ == 8:
                self.enterOuterAlt(localctx, 8)
                self.state = 131
                self.match(ChronicleLogstashParser.ID)
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Math_statementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def LPAREN(self):
            return self.getToken(ChronicleLogstashParser.LPAREN, 0)

        def math_statement(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.Math_statementContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.Math_statementContext,i)


        def RPAREN(self):
            return self.getToken(ChronicleLogstashParser.RPAREN, 0)

        def math_expression(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Math_expressionContext,0)


        def MATHOP(self):
            return self.getToken(ChronicleLogstashParser.MATHOP, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_math_statement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMath_statement" ):
                listener.enterMath_statement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMath_statement" ):
                listener.exitMath_statement(self)



    def math_statement(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = ChronicleLogstashParser.Math_statementContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 14
        self.enterRecursionRule(localctx, 14, self.RULE_math_statement, _p)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 140
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [13]:
                self.state = 135
                self.match(ChronicleLogstashParser.LPAREN)
                self.state = 136
                self.math_statement(0)
                self.state = 137
                self.match(ChronicleLogstashParser.RPAREN)
                pass
            elif token in [32, 34]:
                self.state = 139
                self.math_expression()
                pass
            else:
                raise NoViableAltException(self)

            self._ctx.stop = self._input.LT(-1)
            self.state = 147
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,15,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    localctx = ChronicleLogstashParser.Math_statementContext(self, _parentctx, _parentState)
                    self.pushNewRecursionContext(localctx, _startState, self.RULE_math_statement)
                    self.state = 142
                    if not self.precpred(self._ctx, 2):
                        from antlr4.error.Errors import FailedPredicateException
                        raise FailedPredicateException(self, "self.precpred(self._ctx, 2)")
                    self.state = 143
                    self.match(ChronicleLogstashParser.MATHOP)
                    self.state = 144
                    self.math_statement(3) 
                self.state = 149
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,15,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class Math_expressionContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def MATHOP(self):
            return self.getToken(ChronicleLogstashParser.MATHOP, 0)

        def IFSTATEMENTID(self, i:int=None):
            if i is None:
                return self.getTokens(ChronicleLogstashParser.IFSTATEMENTID)
            else:
                return self.getToken(ChronicleLogstashParser.IFSTATEMENTID, i)

        def NUMBER(self, i:int=None):
            if i is None:
                return self.getTokens(ChronicleLogstashParser.NUMBER)
            else:
                return self.getToken(ChronicleLogstashParser.NUMBER, i)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_math_expression

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMath_expression" ):
                listener.enterMath_expression(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMath_expression" ):
                listener.exitMath_expression(self)




    def math_expression(self):

        localctx = ChronicleLogstashParser.Math_expressionContext(self, self._ctx, self.state)
        self.enterRule(localctx, 16, self.RULE_math_expression)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 150
            _la = self._input.LA(1)
            if not(_la==32 or _la==34):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
            self.state = 151
            self.match(ChronicleLogstashParser.MATHOP)
            self.state = 152
            _la = self._input.LA(1)
            if not(_la==32 or _la==34):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Boolean_opContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def AND(self):
            return self.getToken(ChronicleLogstashParser.AND, 0)

        def OR(self):
            return self.getToken(ChronicleLogstashParser.OR, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_boolean_op

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBoolean_op" ):
                listener.enterBoolean_op(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBoolean_op" ):
                listener.exitBoolean_op(self)




    def boolean_op(self):

        localctx = ChronicleLogstashParser.Boolean_opContext(self, self._ctx, self.state)
        self.enterRule(localctx, 18, self.RULE_boolean_op)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 154
            _la = self._input.LA(1)
            if not(_la==27 or _la==28):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Boolean_evalContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def EQUAL(self):
            return self.getToken(ChronicleLogstashParser.EQUAL, 0)

        def NOTEQUAL(self):
            return self.getToken(ChronicleLogstashParser.NOTEQUAL, 0)

        def LESSTHAN(self):
            return self.getToken(ChronicleLogstashParser.LESSTHAN, 0)

        def GREATERTHAN(self):
            return self.getToken(ChronicleLogstashParser.GREATERTHAN, 0)

        def LTEQUAL(self):
            return self.getToken(ChronicleLogstashParser.LTEQUAL, 0)

        def GTEQUAL(self):
            return self.getToken(ChronicleLogstashParser.GTEQUAL, 0)

        def MATCH(self):
            return self.getToken(ChronicleLogstashParser.MATCH, 0)

        def NOTMATCH(self):
            return self.getToken(ChronicleLogstashParser.NOTMATCH, 0)

        def IN(self):
            return self.getToken(ChronicleLogstashParser.IN, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_boolean_eval

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBoolean_eval" ):
                listener.enterBoolean_eval(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBoolean_eval" ):
                listener.exitBoolean_eval(self)




    def boolean_eval(self):

        localctx = ChronicleLogstashParser.Boolean_evalContext(self, self._ctx, self.state)
        self.enterRule(localctx, 20, self.RULE_boolean_eval)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 156
            _la = self._input.LA(1)
            if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 133693568) != 0)):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class PluginContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(ChronicleLogstashParser.ID, 0)

        def LBRACE(self):
            return self.getToken(ChronicleLogstashParser.LBRACE, 0)

        def RBRACE(self):
            return self.getToken(ChronicleLogstashParser.RBRACE, 0)

        def keyvalue(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.KeyvalueContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.KeyvalueContext,i)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_plugin

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterPlugin" ):
                listener.enterPlugin(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitPlugin" ):
                listener.exitPlugin(self)




    def plugin(self):

        localctx = ChronicleLogstashParser.PluginContext(self, self._ctx, self.state)
        self.enterRule(localctx, 22, self.RULE_plugin)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 158
            self.match(ChronicleLogstashParser.ID)
            self.state = 159
            self.match(ChronicleLogstashParser.LBRACE)
            self.state = 163
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==30 or _la==33:
                self.state = 160
                self.keyvalue()
                self.state = 165
                self._errHandler.sync(self)
                _la = self._input.LA(1)

            self.state = 166
            self.match(ChronicleLogstashParser.RBRACE)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class KeyvalueContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def kv_lvalue(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Kv_lvalueContext,0)


        def KVSEPARATOR(self):
            return self.getToken(ChronicleLogstashParser.KVSEPARATOR, 0)

        def kv_rvalue(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.Kv_rvalueContext,0)


        def COMMA(self):
            return self.getToken(ChronicleLogstashParser.COMMA, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_keyvalue

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterKeyvalue" ):
                listener.enterKeyvalue(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitKeyvalue" ):
                listener.exitKeyvalue(self)




    def keyvalue(self):

        localctx = ChronicleLogstashParser.KeyvalueContext(self, self._ctx, self.state)
        self.enterRule(localctx, 24, self.RULE_keyvalue)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 168
            self.kv_lvalue()
            self.state = 169
            self.match(ChronicleLogstashParser.KVSEPARATOR)
            self.state = 170
            self.kv_rvalue()
            self.state = 172
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if _la==16:
                self.state = 171
                self.match(ChronicleLogstashParser.COMMA)


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Kv_lvalueContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(ChronicleLogstashParser.ID, 0)

        def STRING(self):
            return self.getToken(ChronicleLogstashParser.STRING, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_kv_lvalue

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterKv_lvalue" ):
                listener.enterKv_lvalue(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitKv_lvalue" ):
                listener.exitKv_lvalue(self)




    def kv_lvalue(self):

        localctx = ChronicleLogstashParser.Kv_lvalueContext(self, self._ctx, self.state)
        self.enterRule(localctx, 26, self.RULE_kv_lvalue)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 174
            _la = self._input.LA(1)
            if not(_la==30 or _la==33):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Kv_rvalueContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def NUMBER(self):
            return self.getToken(ChronicleLogstashParser.NUMBER, 0)

        def list_(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.ListContext,0)


        def hash_(self):
            return self.getTypedRuleContext(ChronicleLogstashParser.HashContext,0)


        def STRING(self):
            return self.getToken(ChronicleLogstashParser.STRING, 0)

        def BOOLEAN(self):
            return self.getToken(ChronicleLogstashParser.BOOLEAN, 0)

        def ID(self):
            return self.getToken(ChronicleLogstashParser.ID, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_kv_rvalue

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterKv_rvalue" ):
                listener.enterKv_rvalue(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitKv_rvalue" ):
                listener.exitKv_rvalue(self)




    def kv_rvalue(self):

        localctx = ChronicleLogstashParser.Kv_rvalueContext(self, self._ctx, self.state)
        self.enterRule(localctx, 28, self.RULE_kv_rvalue)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 182
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [34]:
                self.state = 176
                self.match(ChronicleLogstashParser.NUMBER)
                pass
            elif token in [11]:
                self.state = 177
                self.list_()
                pass
            elif token in [9]:
                self.state = 178
                self.hash_()
                pass
            elif token in [30]:
                self.state = 179
                self.match(ChronicleLogstashParser.STRING)
                pass
            elif token in [29]:
                self.state = 180
                self.match(ChronicleLogstashParser.BOOLEAN)
                pass
            elif token in [33]:
                self.state = 181
                self.match(ChronicleLogstashParser.ID)
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class HashContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def LBRACE(self):
            return self.getToken(ChronicleLogstashParser.LBRACE, 0)

        def RBRACE(self):
            return self.getToken(ChronicleLogstashParser.RBRACE, 0)

        def keyvalue(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.KeyvalueContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.KeyvalueContext,i)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_hash

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterHash" ):
                listener.enterHash(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitHash" ):
                listener.exitHash(self)




    def hash_(self):

        localctx = ChronicleLogstashParser.HashContext(self, self._ctx, self.state)
        self.enterRule(localctx, 30, self.RULE_hash)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 184
            self.match(ChronicleLogstashParser.LBRACE)
            self.state = 188
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==30 or _la==33:
                self.state = 185
                self.keyvalue()
                self.state = 190
                self._errHandler.sync(self)
                _la = self._input.LA(1)

            self.state = 191
            self.match(ChronicleLogstashParser.RBRACE)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ListContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def LBRACKET(self):
            return self.getToken(ChronicleLogstashParser.LBRACKET, 0)

        def RBRACKET(self):
            return self.getToken(ChronicleLogstashParser.RBRACKET, 0)

        def list_value(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.List_valueContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.List_valueContext,i)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_list

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterList" ):
                listener.enterList(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitList" ):
                listener.exitList(self)




    def list_(self):

        localctx = ChronicleLogstashParser.ListContext(self, self._ctx, self.state)
        self.enterRule(localctx, 32, self.RULE_list)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 193
            self.match(ChronicleLogstashParser.LBRACKET)
            self.state = 201
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if (((_la) & ~0x3f) == 0 and ((1 << _la) & 27380482048) != 0):
                self.state = 194
                self.list_value()
                self.state = 198
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while (((_la) & ~0x3f) == 0 and ((1 << _la) & 27380482048) != 0):
                    self.state = 195
                    self.list_value()
                    self.state = 200
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)



            self.state = 203
            self.match(ChronicleLogstashParser.RBRACKET)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class If_listContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def LBRACKET(self):
            return self.getToken(ChronicleLogstashParser.LBRACKET, 0)

        def RBRACKET(self):
            return self.getToken(ChronicleLogstashParser.RBRACKET, 0)

        def STRING(self):
            return self.getToken(ChronicleLogstashParser.STRING, 0)

        def BOOLEAN(self):
            return self.getToken(ChronicleLogstashParser.BOOLEAN, 0)

        def NUMBER(self):
            return self.getToken(ChronicleLogstashParser.NUMBER, 0)

        def list_value(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ChronicleLogstashParser.List_valueContext)
            else:
                return self.getTypedRuleContext(ChronicleLogstashParser.List_valueContext,i)


        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_if_list

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterIf_list" ):
                listener.enterIf_list(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitIf_list" ):
                listener.exitIf_list(self)




    def if_list(self):

        localctx = ChronicleLogstashParser.If_listContext(self, self._ctx, self.state)
        self.enterRule(localctx, 34, self.RULE_if_list)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 205
            self.match(ChronicleLogstashParser.LBRACKET)
            self.state = 213
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if (((_la) & ~0x3f) == 0 and ((1 << _la) & 18790481920) != 0):
                self.state = 206
                _la = self._input.LA(1)
                if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 18790481920) != 0)):
                    self._errHandler.recoverInline(self)
                else:
                    self._errHandler.reportMatch(self)
                    self.consume()
                self.state = 210
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while (((_la) & ~0x3f) == 0 and ((1 << _la) & 27380482048) != 0):
                    self.state = 207
                    self.list_value()
                    self.state = 212
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)



            self.state = 215
            self.match(ChronicleLogstashParser.RBRACKET)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class List_valueContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def STRING(self):
            return self.getToken(ChronicleLogstashParser.STRING, 0)

        def ID(self):
            return self.getToken(ChronicleLogstashParser.ID, 0)

        def BOOLEAN(self):
            return self.getToken(ChronicleLogstashParser.BOOLEAN, 0)

        def NUMBER(self):
            return self.getToken(ChronicleLogstashParser.NUMBER, 0)

        def COMMA(self):
            return self.getToken(ChronicleLogstashParser.COMMA, 0)

        def getRuleIndex(self):
            return ChronicleLogstashParser.RULE_list_value

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterList_value" ):
                listener.enterList_value(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitList_value" ):
                listener.exitList_value(self)




    def list_value(self):

        localctx = ChronicleLogstashParser.List_valueContext(self, self._ctx, self.state)
        self.enterRule(localctx, 36, self.RULE_list_value)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 217
            _la = self._input.LA(1)
            if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 27380482048) != 0)):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx



    def sempred(self, localctx:RuleContext, ruleIndex:int, predIndex:int):
        if self._predicates == None:
            self._predicates = dict()
        self._predicates[2] = self.statement_sempred
        self._predicates[7] = self.math_statement_sempred
        pred = self._predicates.get(ruleIndex, None)
        if pred is None:
            raise Exception("No predicate with index:" + str(ruleIndex))
        else:
            return pred(localctx, predIndex)

    def statement_sempred(self, localctx:StatementContext, predIndex:int):
            if predIndex == 0:
                return self.precpred(self._ctx, 3)
         

    def math_statement_sempred(self, localctx:Math_statementContext, predIndex:int):
            if predIndex == 1:
                return self.precpred(self._ctx, 2)
         




