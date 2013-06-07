github.com/bjdean/canAccessInternet
=====================================

go executable - check if general internet access is possible

See also: http://bjdean.id.au/wiki/Programming_Notes/GoLang

This program is a simple wrapper around the gonetcheck
package: https://github.com/bjdean/gonetcheck

Synopsis
--------

	#!/bin/bash
	
	canAccessInternet
	if [ $? -eq 0 ]
	then
		# Do something requiring internet access
	else
		# Report internet access down
	fi

License
-------

Copyright 2013 Bradley Dean

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http:www.gnu.org/licenses/>.

