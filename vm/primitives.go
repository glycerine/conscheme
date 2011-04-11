// This file is part of conscheme
// Automatically generated by compiler/primitives.scm
package conscheme
import "fmt"
import "os"
func evprim(primop string, args []Obj) Obj {
	switch primop {
	case "$write/2":
		return write(args[0], args[1])
	case "$display/2":
		return display(args[0], args[1])
	case "lookahead-u8/1":
		return lookahead_u8(args[0])
	case "put-u8/2":
		return put_u8(args[0], args[1])
	case "get-u8/1":
		return get_u8(args[0])
	case "$write-char/2":
		return _write_char(args[0], args[1])
	case "$read-char/1":
		return _read_char(args[0])
	case "open-input-file/1":
		return open_input_file(args[0])
	case "current-output-port/0":
		return current_output_port()
	case "current-input-port/0":
		return current_input_port()
	case "output-port?/1":
		return output_port_p(args[0])
	case "input-port?/1":
		return input_port_p(args[0])
	case "port?/1":
		return port_p(args[0])
	case "bytevector?/1":
		return bytevector_p(args[0])
	case "$cell-set!/2":
		v := args[0]
		vv := (*v).(*[1]Obj)
		vv[0] = args[1]
		return Void
	case "$cell-ref/1":
		v := args[0]
		vv := (*v).(*[1]Obj)
		return vv[0]
	case "$make-cell/1":
		var v [1]Obj
		v[0] = args[0]
		var vv interface{} = &v
		return Obj(&vv)
	case "command-line/0":
		return Command_line()
	case "exit/1":
		os.Exit(number_to_int(args[0]))
	case "eq?/2":
		if args[0] == args[1] {
			return True
		} else {
			return False
		}
	case "eof-object/0":
		return Eof
	case "unspecified/0":
		return Void
	case "apply/n":
		return apply(args)
	case "make-string/2":
		return Make_string(args[0], args[1])
	case "make-string/1":
		return Make_string(args[0],Make_char(32))
	case "string-set!/3":
		return String_set_ex(args[0], args[1], args[2])
	case "string-ref/2":
		return String_ref(args[0], args[1])
	case "string-length/1":
		return String_length(args[0])
	case "string?/1":
		return string_p(args[0])
	case "greatest-fixnum/0":
		return Make_fixnum(fixnum_max)
	case "least-fixnum/0":
		return Make_fixnum(fixnum_min)
	case "$cmp/2":
		return number_cmp(args[0], args[1])
	case "$-/2":
		return number_subtract(args[0], args[1])
	case "$//2":
		return number_divide(args[0], args[1])
	case "$+/2":
		return number_add(args[0], args[1])
	case "$number->string/2":
		return _number_to_string(args[0], args[1])
	case "number?/1":
		return number_p(args[0])
	case "vector-set!/3":
		return Vector_set_ex(args[0], args[1], args[2])
	case "vector-ref/2":
		return Vector_ref(args[0], args[1])
	case "vector-length/1":
		return vector_length(args[0])
	case "make-vector/2":
		return Make_vector(args[0], args[1])
	case "vector?/1":
		return vector_p(args[0])
	case "integer->char/1":
		return integer_to_char(args[0])
	case "char->integer/1":
		return char_to_integer(args[0])
	case "char?/1":
		return char_p(args[0])
	case "string->symbol/1":
		return String_to_symbol(args[0])
	case "symbol->string/1":
		return Symbol_to_string(args[0])
	case "symbol?/1":
		return symbol_p(args[0])
	case "set-cdr!/2":
		return set_cdr_ex(args[0], args[1])
	case "set-car!/2":
		return set_car_ex(args[0], args[1])
	case "length/1":
		return Length(args[0])
	case "floyd/1":
		return Floyd(args[0])
	case "cdr/1":
		return cdr(args[0])
	case "car/1":
		return car(args[0])
	case "cons/2":
		return Cons(args[0], args[1])
	case "pair?/1":
		return pair_p(args[0])
	case "not/1":
		return not(args[0])
	case "boolean?/1":
		return boolean_p(args[0])
	default:
		fmt.Fprintf(os.Stderr, "Please regenerate primitives.go\n")
		panic(fmt.Sprintf("Unimplemented primitive: %s",primop))
	}
	panic(fmt.Sprintf("Fell off the edge in evprim(): %s",primop))
}
