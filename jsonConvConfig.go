package intgconvconf

import (
	"encoding/json"
	"os"
	//"fmt"
	"path/filepath"
	"io/ioutil"
)


/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//JSONConvConf JSONConvConf
type JSONConvConf struct{

}

//GetSpfConversion GetSpfConversion
func (c *JSONConvConf)GetSpfConversion(dir string) *map[string]SpfConfFile{
	var rtn = make(map[string]SpfConfFile) 
	var conFigDirs []ConfDir
	files, err := ioutil.ReadDir(dir)
	if err == nil{
		for _, file := range files {
			//fmt.Println("file name: ", file.Name())
			if file.IsDir() {
				var cfd ConfDir
				cfd.Name = file.Name()
				conFigDirs = append(conFigDirs, cfd)
			}
		}
		for _, cfgd := range conFigDirs {
			var sdirname = dir + string(filepath.Separator) + cfgd.Name
			//fmt.Println("sdirname: ", sdirname)
			jsonFile, err := os.Open(sdirname + string(filepath.Separator) + "config.json")
			defer jsonFile.Close()
			if err == nil{
				var cf SpfConfFile
				//fmt.Println("jsonFile binary: ", jsonFile)
				byteValue, _ := ioutil.ReadAll(jsonFile)
				json.Unmarshal(byteValue, &cf)				
				rtn[cfgd.Name] = cf // = append(rtnFiles, cf)
				// fmt.Println("jsonFile: ", cf)
				// fmt.Println("CfDirectory: ", cf.CfDirectory)
				// fmt.Println("Distributor: ", cf.Distributor)
				// fmt.Println("rtn: ", rtn)
			}
		}
	}
	return &rtn
}