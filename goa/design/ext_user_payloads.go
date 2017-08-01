package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// UserPayload defines the data structure used in the create user request body.
var UserPayload = Type("UserPayload", func() {
	//`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
	//`user_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'User Name',
	//`email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'E-Mail Address',
	//`password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Password',
	//`delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
	//`created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
	//`updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',

	Attribute("user_name", String, "First name", func() {
		MinLength(2)
		MaxLength(20)
		Example("Hiroki")
	})
	Attribute("email", String, "E-mail of user", fieldEmail)
	Attribute("password", String, "Password", fieldPassword)

	//Attribute("email", String, "E-mail of user", func() {
	//	Format("email")
	//	Example("hy@gmail.com")
	//})
	//Attribute("password", String, "Password", func() {
	//	MinLength(10)
	//	MinLength(20)
	//	Example("xxxxx")
	//})
})
