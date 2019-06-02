package intgconvconf

import (
	
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

//SpfConvConf SpfConvConf
type SpfConvConf interface{
	GetSpfConversion(dir string) *map[string]SpfConfFile
}


//ConfFile ConfFile
type ConfFile struct {
	Name     string
	FullName string
}

//ConfDir ConfDir
type ConfDir struct {
	Name  string
	Files []ConfFile
}


//SpfConfFile SpfConfFile
type SpfConfFile struct{
	CfDirectory string `json:"cfDirectory"`
	Distributor string `json:"distributor"`
	SpfFields []SpfField `json:"spfFields"`
}


//SpfField SpfField
type SpfField struct{
	SpfKey string `json:"spfKey"`
	CartKey string `json:"cartKey"`
	Prefix string `json:"prefix"`
	SpfSubKeys []string `json:"spfSubKeys"`
}

