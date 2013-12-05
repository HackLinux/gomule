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

type SockSrv struct {
    Host        string
    Port        int
    Debug       bool
    listener    net.Listener
}

func NewSockSrv(host string, port int, debug bool) *SockSrv {
    return &SockSrv {
        Host: host,
        Port: port,
        Debug: debug}
}

func (this *SockSrv) respConn(conn net.Conn) {
    var protocol byte = 0xE3
    if this.Debug { fmt.Printf("DEBUG: %v connected\n", conn.RemoteAddr()) }
    buf := make([]byte, 5)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("ERROR:", err.Error())
        return
    }
    if buf[0] == 0xE3 {
        protocol = 0xE3
    }
    if this.Debug { fmt.Printf("DEBUG: 0x%02x\n", protocol) }
    size := byteToInt32(buf[1:n])
    if this.Debug { fmt.Printf("DEBUG: %v -> %d\n", buf[1:n], size) }
    buf = make([]byte, size)
    n, err = conn.Read(buf)
    if buf[0] == 0x01 {
        if this.Debug { fmt.Println("DEBUG: login") }
        login(buf, protocol, conn, this.Debug)
    }
}

func (this *SockSrv) Start() {
    ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Host, this.Port))
    if err != nil {
        fmt.Println("ERROR:", err.Error())
        return
    }
    this.listener = ln
    fmt.Printf("Staring server %s:%d\n", this.Host, this.Port)

    for {
        conn, err := this.listener.Accept()
        if err != nil {
            fmt.Println("ERROR:", err.Error())
            continue
        }
        go this.respConn(conn)
    }
}

func (this *SockSrv) Stop() {
    defer this.listener.Close()
    return
}
