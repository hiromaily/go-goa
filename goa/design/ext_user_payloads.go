package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// UserPayload defines the data structure used in the create user request body.
var UserPayload = Type("UserPayload", func() {
	//`user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
	//`first_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'First Name',
	//`last_name` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Last Name',
	//`email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'E-Mail Address',
	//`password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Password',
	//`oauth2_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'oauth_flg flg',
	//`delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
	//`create_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
	//`update_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',

	Attribute("first_name", String, "First name", func() {
		MinLength(2)
		MinLength(10)
		Example("Hiroki")
	})
	Attribute("last_name", String, "Last name", func() {
		MinLength(2)
		MinLength(10)
		Example("Yasui")
	})
	Attribute("email", String, "E-mail of user", func() {
		Format("email")
		Example("hy@gmail.com")
	})
	Attribute("password", String, "Password", func() {
		MinLength(10)
		MinLength(20)
		Example("xxxxx")
	})
})
