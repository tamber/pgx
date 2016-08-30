package pgx

// import (
// 	"bytes"
// 	"database/sql/driver"
// 	"encoding/binary"
// 	"encoding/hex"
// 	"errors"
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/lib/pq/oid"
// )

// func binaryEncode(parameterStatus *parameterStatus, x interface{}) []byte {
// 	switch v := x.(type) {
// 	case []byte:
// 		return v
// 	default:
// 		return encode(parameterStatus, x, oid.T_unknown)
// 	}
// }

// func encode(parameterStatus *parameterStatus, x interface{}, pgtypOid oid.Oid) []byte {
// 	switch v := x.(type) {
// 	case int64:
// 		return strconv.AppendInt(nil, v, 10)
// 	case float64:
// 		return strconv.AppendFloat(nil, v, 'f', -1, 64)
// 	case []byte:
// 		if pgtypOid == oid.T_bytea {
// 			return encodeBytea(parameterStatus.serverVersion, v)
// 		}

// 		return v
// 	case string:
// 		if pgtypOid == oid.T_bytea {
// 			return encodeBytea(parameterStatus.serverVersion, []byte(v))
// 		}

// 		return []byte(v)
// 	case bool:
// 		return strconv.AppendBool(nil, v)
// 	case time.Time:
// 		return formatTs(v)

// 	default:
// 		errorf("encode: unknown type for %T", v)
// 	}

// 	panic("not reached")
// }
