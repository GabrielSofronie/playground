/**
Content model
A content can be used in the context of a Page.
It can have the following structure/data:

V_1:
- Title
- File(s) / Image(s)
- Section(s)
- Author
- Creation date
- Update date
- Comments

V_2:
- Favorite
- Status (published/unpublished)
- Type
- Tags ??? (are they part of the content?)
- Teaser ???

*/

package models

import (
	"time"
)

type Content struct {
	Title string
	//Author User
	Created time.Time
	Updated time.Time
}
