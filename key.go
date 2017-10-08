package ini

import (
    "fmt"
    "strconv"
)

type Key interface {
    // get name of the key
    Name() string

    // get value of the key
    Value() (string, error )

    //get the value of key and return defValue if
    //the value does not exist
    ValueWithDefault( defValue string) string

    // get the value as int
    Int() (int, error )

    // get value as int and return defValue if the
    // value of the key does not exist
    IntWithDefault( defValue int) int

    // get the value as int64
    Int64() (int64, error )

    // get the value as int64 and return defValue
    // if the value of the key does not exist
    Int64WithDefault( defValue int64)int64

    // get the value as uint64
    Uint64() (uint64, error )

    // get the value as uint64 and return defValue
    // if the value of the key does not exist
    Uint64WithDefault( defValue uint64) uint64

    // get the value as float32
    Float32()(float32, error)

    // get the value as float32 and return defValue
    // if the value of the key does not exist
    Float32WithDefault(defValue float32)float32

    // get the value as float64
    Float64()(float64, error)

    // get the value as the float64 and return defValue
    // if the value of the key does not exist
    Float64WithDefault( defValue float64)float64

    // return a string as "key=value" format
    // and if no value return empty string
    String() string
}


type NonExistKey struct {
    name string
}

func NewNonExistKey( name string ) *NonExistKey {
    return &NonExistKey{ name: name }
}

func (nek *NonExistKey)Name() string {
    return nek.name
}

func (nek *NonExistKey)Value() (string, error) {
    return "", fmt.Errorf( "no such key:%s", nek.name )
}

func (nek *NonExistKey) ValueWithDefault( defValue string ) string {
    return defValue
}
func (nek *NonExistKey)Int()(int, error ) {
    return 0, fmt.Errorf( "no such key:%s", nek.name )
}

func (nek *NonExistKey)IntWithDefault(defValue int)int {
    return defValue
}

func (nek *NonExistKey)Int64()(int64, error ) {
    return 0, fmt.Errorf( "no such key:%s", nek.name )
}

func (nek *NonExistKey)Int64WithDefault( defValue int64) int64 {
    return defValue
}

func (nek *NonExistKey)Uint64()(uint64, error ) {
    return 0, fmt.Errorf( "no such key:%s", nek.name )
}

func (nek *NonExistKey)Uint64WithDefault(defValue uint64) uint64 {
    return defValue
}

func (nek *NonExistKey)Float32()(float32, error ) {
    return .0, fmt.Errorf( "no such key:%s", nek.name )
}

func (nek *NonExistKey)Float32WithDefault( defValue float32) float32 {
    return defValue
}

func (nek *NonExistKey)Float64() (float64, error ) {
    return .0, fmt.Errorf( "no such key:%s", nek.name )
}

func (nek *NonExistKey)Float64WithDefault( defValue float64) float64 {
    return defValue
}

func (nek *NonExistKey)String() string {
    return ""
}

type NormalKey struct {
    name string
    value string
}

func NewNormalKey( name, value string )*NormalKey {
    return &NormalKey{ name: name, value: value }
}

func (k *NormalKey) Name() string {
    return k.name
}

func (k *NormalKey) Value() (string, error) {
    return k.value, nil
}

func (k *NormalKey)ValueWithDefault(defValue string)string {
    return k.value
}

func (k *NormalKey)Int() (int, error ) {
    return strconv.Atoi( k.value )
}

func (k *NormalKey)IntWithDefault(defValue int ) int {
    i, err := strconv.Atoi( k.value )
    if err == nil {
        return i
    }
    return defValue
}

func (k *NormalKey)Int64()(int64, error ) {
    return strconv.ParseInt(k.value, 0, 64)
}

func (k *NormalKey)Int64WithDefault( defValue int64 )int64 {
    i, err := strconv.ParseInt(k.value, 0, 64)
    if err == nil {
        return i
    }
    return defValue
}

func (k *NormalKey)Uint64()(uint64, error ) {
    return strconv.ParseUint(k.value, 0, 64)
}

func (k *NormalKey)Uint64WithDefault( defValue uint64) uint64 {
    i, err := strconv.ParseUint(k.value, 0, 64)
    if err == nil {
        return i
    }
    return defValue
}

func (k *NormalKey)Float32()(float32, error ) {
    f, err := strconv.ParseFloat( k.value, 32 )
    return float32(f), err
}

func (k *NormalKey)Float32WithDefault( defValue float32) float32 {
    f, err := strconv.ParseFloat( k.value, 32 )
    if err == nil {
        return float32(f)
    }
    return defValue
}

func (k *NormalKey)Float64()(float64, error ) {
    return strconv.ParseFloat( k.value, 64 )
}

func (k *NormalKey)Float64WithDefault( defValue float64) float64 {
    f, err := strconv.ParseFloat( k.value, 64 )
    if err == nil {
        return f
    }
    return defValue
}


func (k* NormalKey)String() string {
    return fmt.Sprintf( "%s=%s", k.name, k.value )
}



