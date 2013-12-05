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
    "net"
)

func login(buf []byte, protocol byte, conn net.Conn, debug bool) {
    high_id := highId(conn.RemoteAddr().String())
    uuid := fmt.Sprintf("%x-%x-%x-%x-%x-%x-%x-%x",
        buf[1:3], buf[3:5], buf[5:7], buf[7:9], buf[9:11], buf[11:13],
        buf[13:15], buf[15:17])
    port := byteToInt16(buf[21:23])
    if debug {
        fmt.Println("DEBUG:", high_id)
        fmt.Println("DEBUG:", uuid)
        fmt.Println("DEBUG:", port)
    }

    data := []byte{protocol,
                   8, 0, 0, 0,
                   0x38,
                   5, 0,
                   'h', 'e', 'l', 'l', 'o'}
    if debug { fmt.Println("DEBUG:", data) }
    conn.Write(data)

    data = []byte{protocol,
                  9, 0, 0, 0,
                  0x40,
                  0, 0, 0, 0,
                  1, 0, 0, 0}
    high_id_b := int32ToByte(high_id)
    for i := 0; i < len(high_id_b); i++ { data[i + 6] = high_id_b[i] }
    if debug { fmt.Println("DEBUG:", data) }
    conn.Write(data)
}
