# Tor Webscraper in Golang

## requirement

This program is written in Go so to compile and use it you need install Go language on you computer.

This program also require next Golang libraries:
 - github.com/PuerkitoBio/goquery
 - golang.org/x/net/proxy
 
You can install them automatically with:

	go get -u "github.com/PuerkitoBio/goquery"

Or manually with:

	git clone https://github.com/golang/net/ $GOPATH/src/golang.org/x/net
	
## compilation

	go build

## [TODO-List]

 - name file like `outputPrefix_id-[id].html` or `outputPrefix_pid-[pid].html` and move them into directory named with `outputPrefix`


 - check if now created page already exist and is different if so we can overwrite it or store it under a subdirectory named like `outputPrefix_id-[id]` or `outputPrefix_pid-[pid]` and each new version end by `(n)` where 'n' is the version number


----------------------------------------
Copyright (C) 2019  Guillaume GRONNIER

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.