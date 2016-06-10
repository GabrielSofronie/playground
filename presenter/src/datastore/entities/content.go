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

/*
Use cases:

Create <single> Content
--------------
Data:
<Title>
<Actual-content>
<Author-information>
<Create-information>
<Update-information>

Primary Course:
1. Author issues command to save new content using above data
   >> POST /content
   >> http.Body
2. System validates data
   >> fields have to match type (?)
   >> check required fields (e.g. Title, Author)
3. System creates the content and determines <Content-id>
   >>
4. System delivers <Content-id> to visitor
   >> <Content-id> can be used for further operations

Exception Course: Validation Error
1. System delivers error message to visitor
   >> POST fields validation fails
2. System redirects the visitor to Page-Not-Found/Unauthorized
   >> Author DOES NOT have the rights to create content

View <single> Content
--------------
Data:
<Content-id>
<Content-meta>
<Content-sections>
<Author-information>

Primary Course:
1. Visitor issues command to view content using <Content-id>
   >> GET /content/<Content-id>
2. System validates data
   >> /content MUST have a route and an action
   >> <Content-id> MUST be provided
3. System determines <Content-id>
   >> Parse and get value
4. System delivers content to visitor
   >> Query storage for data using <Content-id>

Exception Course: Validation Error
1. System delivers error message to visitor
   >> /content DOES NOT have an associated route
   >> <Content-id> parsing fails
2. System redirects the visitor to Page-Not-Found
   >> Query on <Content-id> returns NO result

*/

package entities

import "time"

type Content struct {
	//Id string
	Title string
	Meta string
	Author User
	Created time.Time
	Updated time.Time
}
