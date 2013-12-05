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

package main

import (
    "fmt"
    "flag"
    "os"
    "github.com/xiangzhai/gomule/emule"
)

var (
    debug   bool
    host    string
    port    int
)

func init() {
    flag.BoolVar(&debug, "d", false, "Debug")
    flag.StringVar(&host, "h", "0.0.0.0", "Host address")
    flag.IntVar(&port, "p", 7111, "Port number")
}

func main() {
    if len(os.Args) == 1 {
        fmt.Println("Usage: gomule [options]")
        return
    }

    if os.Args[1] == "-v" {
        fmt.Println("GoMule server Version 1.0")
        fmt.Println("Copyright 2013 Leslie Zhai")
        return
    }

    flag.Parse()

    sock := emule.NewSockSrv(host, port, debug)
    sock.Start()
    defer sock.Stop()
}
