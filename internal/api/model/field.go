/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package model

// Field represents a name/value pair to display on an account's profile.
//
// swagger:model field
type Field struct {
	// The key/name of this field.
	// example: pronouns
	Name string `json:"name"`
	// The value of this field.
	// example: they/them
	Value string `json:"value"`
	// If this field has been verified, when did this occur? (ISO 8601 Datetime).
	// example: 2021-07-30T09:20:25+00:00
	VerifiedAt string `json:"verified_at,omitempty"`
}
