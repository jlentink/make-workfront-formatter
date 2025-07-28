package main

import "strings"

type parseState struct {
	char   rune
	runes  []rune
	input  string
	output string
	pos    int
	indent int
	state  []string
}

func indentString(indent int) string {
	if indent < 0 {
		indent = 0
	}
	return strings.Repeat("  ", indent)

}

func openingsTag(state parseState) parseState {
	if state.char == '{' && state.input[state.pos+1] == '{' {
		state.pos += 1
		state.indent += 1
		state.output += "{{\n" + indentString(state.indent)
		state.state = append(state.state, "OPEN_BLOCK")
	}
	return state
}

func closingTag(state parseState) parseState {
	if state.input[state.pos] == '}' && state.input[state.pos+1] == '}' {
		state.pos += 1
		state.indent -= 1
		state.output += "}}"
		state.state = []string{}
	}
	return state
}

func countLeadingSpaces(s string) int {
	count := 0
	for _, ch := range s {
		if ch == ' ' {
			count++
		} else {
			break
		}
	}
	return count
}

func dotCommaTag(state parseState) parseState {
	if state.char == ';' && len(state.state) > 0 && state.state[0] == "OPEN_BLOCK" {
		state.output += ";\n" + indentString(state.indent)
		s := countLeadingSpaces(state.input[state.pos+1:])
		state.pos += s
	} else {
		state.output += string(state.char)
	}
	return state
}

func NewlineTag(state parseState) parseState {
	if state.char == '\n' && len(state.state) > 0 && state.state[0] == "OPEN_BLOCK" {

	} else {
		state.output += string(state.char)
	}
	return state
}
func closeHookTag(state parseState) parseState {
	if state.input[state.pos] == ')' && len(state.state) > 0 && state.state[0] == "OPEN_BLOCK" {
		state.indent -= 1
		state.output += "\n" + indentString(state.indent) + ")"

		if state.pos < len(state.input)-1 && state.input[state.pos+1] != ';' {
			state.output += "\n"
		}
		state.state = state.state[:len(state.state)-1]
	} else {
		state.output += string(state.char)
	}
	return state
}

func GenericFunction(fnName string, state parseState) parseState {
	if state.state[0] == "OPEN_BLOCK" && state.input[state.pos:state.pos+len(fnName)] == fnName {

		fragment := state.input[state.pos:]
		offset := strings.Index(fragment, "(")
		state.output += fnName + "(\n" + indentString(state.indent+1)
		state.indent += 1
		state.state = append(state.state, strings.ToUpper(fnName)+"_BLOCK")
		state.pos += offset
	} else {
		state.output += string(state.char)
	}
	return state
}

func mergeArrays(arrs ...[]string) []string {
	output := []string{}
	for _, arr := range arrs {
		output = append(output, arr...)
	}
	return output
}

// IndentMakeScript formats and indents makescript content
func IndentMakeScript(input string) string {
	genericFunctions := []string{"get", "if", "ifempty", "switch", "omit", "pick", "mergeCollections"}
	dateFunctions := []string{"addSeconds", "addMinutes", "addHours", "addDays", "setSecond", "setMinute", "setHour", "setDay", "setDate", "setMonth", "setYear", "formatDate", "parseDate"}
	stringFunctions := []string{"length", "lower", "capitalize", "startcase", "ascii", "replace", "trim", "upper", "substring", "indexOf", "toBinary", "toString", "encodeURL", "decodeURL", "escapeHTML", "escapeMarkdown", "stripHTML", "contains", "split", "md5", "sha256", "sha512", "bas64"}
	mathFunctions := []string{"average", "ceil", "floor", "max", "min", "round", "sum", "parseNumber", "formatNumber"}
	arrayFunctions := []string{"join", "length", "keys", "slice", "merge", "contains", "remove", "add", "map", "shuffle", "reverse", "flatten", "distinct", "toCollection", "toArray", "arrayDifference", "deDuplicate"}
	allFunctions := mergeArrays(genericFunctions, dateFunctions, stringFunctions, mathFunctions, arrayFunctions)
	lineState := parseState{
		char:   0,
		runes:  []rune(input),
		input:  input,
		output: "",
		pos:    0,
		indent: 0,
		state:  []string{},
	}
	for lineState.pos = 0; lineState.pos < len(lineState.runes); lineState.pos++ {
		lineState.char = lineState.runes[lineState.pos]
		switch lineState.char {
		case '{':
			lineState = openingsTag(lineState)
		case '}':
			lineState = closingTag(lineState)
		case ')':
			lineState = closeHookTag(lineState)
		case ';':
			lineState = dotCommaTag(lineState)
		case '\n':
			lineState = NewlineTag(lineState)
		default:
			executed := false
			if len(lineState.state) > 0 && lineState.state[0] == "OPEN_BLOCK" {
				for _, fn := range allFunctions {
					if strings.HasPrefix(lineState.input[lineState.pos:], fn) {
						lineState = GenericFunction(fn, lineState)
						executed = true
						break
					}
				}
			}
			if !executed {
				lineState.output += string(lineState.char)
			}
		}
	}
	return lineState.output
}
