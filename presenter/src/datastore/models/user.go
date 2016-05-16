/**
User starts as the Author of the Content.

It holds the following fields:

V_1:
- Name
- Age
- Registered
- Updated

V_2:
- Email
- Password
- Type (Admin/Editor etc.)
++ should have role-based permissions
++ should have a "prestige" system (check SO and badges)
++ should have OAuth login
++ what about "follow" model? (see Twitter, Medium etc.)

*/

package models

import (
	"time"
)

type User struct{
	Name string
	Age byte
	Registered time.Time
	Updated time.Time
}