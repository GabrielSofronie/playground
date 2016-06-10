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

/*
*** This should be Entities accessed by Interactors and used by EntityGatewayImplementers ***
*/

package entities

import "time"

type User struct{
	//Id bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Name string //`json:"name"`
	Age byte
	Registered time.Time
	Updated time.Time
}
