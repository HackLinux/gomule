/*                                                                              
 * Copyright (C) 2013 Deepin, Inc.                                                 
 *               2013 Leslie Zhai <zhaixiang@linuxdeepin.com>                   
 *                                                                              
 * This program is free software: you can redistribute it and/or modify         
 * it under the terms of the GNU General Public License as published by         
 * the Free Software Foundation, either version 3 of the License, or            
 * any later version.                                                           
 *                                                                              
 * This program is distributed in the hope that it will be useful,              
 * but WITHOUT ANY WARRANTY; without even the implied warranty of               
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the                
 * GNU General Public License for more details.                                 
 *                                                                              
 * You should have received a copy of the GNU General Public License            
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.        
 */

package emule

import (
    "fmt"
    "bytes"
    "encoding/binary"
    "math"
)

func byteToInt32(data []byte) (ret int32) {
    buf := bytes.NewBuffer(data)
    binary.Read(buf, binary.LittleEndian, &ret)
    return
}

func byteToInt16(data []byte) (ret int16) {
    buf := bytes.NewBuffer(data)
    binary.Read(buf, binary.LittleEndian, &ret)
    return
}

func int16ToByte(data int16) (ret []byte) {
    ret = []byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, data)
    ret = buf.Bytes()
    return
}

func int32ToByte(data int32) (ret []byte) {
    ret = []byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, data)
    ret = buf.Bytes()
    return
}

func highId(addr string) (ret int32) {
    ret = 0
    var ip1, ip2, ip3, ip4, port int32
    fmt.Sscanf(addr, "%d.%d.%d.%d:%d", &ip1, &ip2, &ip3, &ip4, &port)
    ret = ip1 + int32(math.Pow(2, 8)) * ip2 + int32(math.Pow(2, 16)) * ip3 +
        int32(math.Pow(2, 24)) * ip4
    return
}
