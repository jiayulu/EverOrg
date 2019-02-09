//
//  evernote.go
//  EverOrg
//
//  Created by Mario Martelli on 26.02.17.
//  Copyright © 2017 Mario Martelli. All rights reserved.
//
//  This file is part of EverOrg.
//
//  EverOrg is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  EverOrg is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with EverOrg.  If not, see <http://www.gnu.org/licenses/>.

package main

import "golang.org/x/net/html"

//
// Data Structures
//

// Query ...
type Query struct {
	Notes []Note `xml:"note"`
}

// Node ...
type Node struct {
	Token html.Token
	Text  string
}

// Nodes ...
type Nodes []Node

// Note ...
type Note struct {
	Title      string   `xml:"title"`
	Content    string   `xml:"content"`
	Created    string   `xml:"created"`
	Updated    string   `xml:"updated"`
	Tags       []string `xml:"tag"`
	Attributes struct {
		Author    string  `xml:"author"`
		Latitude  float64 `xml:"latitude"`
		Longitude float64 `xml:"longitude"`
		// Altitude  float64 `xml:"altitude"`
		Source    string `xml:"source"`
		SourceUrl string `xml:"source-url"`
	} `xml:"note-attributes"`

	Resources []Resource `xml:"resource"`
}

// Resource ...
type Resource struct {
	Mime string `xml:"mime"`
	Data struct {
		Content  string `xml:",chardata"`
		Encoding string `xml:"encoding,attr"`
	} `xml:"data"`
}
